package handler

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/helpers"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/errors"
)

func (h Handler) LoadFromZip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(20 << 20) // * max 20MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Error parsing form"})
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Error reading file"})
		return
	}
	defer file.Close()

	tempZipPath := filepath.Join(os.TempDir(), header.Filename)
	tempFile, err := os.Create(tempZipPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messages := fmt.Sprintf("Cannot create temp file, %s", err.Error())
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempZipPath)

	_, err = io.Copy(tempFile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messages := fmt.Sprintf("Error copying file, %s", err.Error())
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
		return
	}

	// * Open zip file
	zipReader, err := zip.OpenReader(tempZipPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messages := fmt.Sprintf("Invalid zip, %s", err.Error())
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
		return
	}
	defer zipReader.Close()

	var onboardings []*domain.Onboarding
	imageMap := make(map[string][]byte)

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		if path.Base(f.Name) == "data.json" {
			rc, err := f.Open()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				messages := fmt.Sprintf("failed to open data.json: %s", err.Error())
				json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
				return
			}
			content, err := io.ReadAll(rc)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				messages := fmt.Sprintf("failed to read data.json: %s", err.Error())
				json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
				return
			}
			rc.Close()
			err = json.Unmarshal(content, &onboardings)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				messages := fmt.Sprintf("failed to unmarshal data.json: %s", err.Error())
				json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
				return
			}
		} else {
			rc, err := f.Open()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(message.ErrorMessage{Message: "failed to open file in zip"})
				return
			}
			content, _ := io.ReadAll(rc)
			rc.Close()
			imageMap[path.Base(f.Name)] = content
		}
	}

	for i, obd := range onboardings {
		o, err := h.OnboardingRepo.GetByTitle(context.Background(), obd.Title)
		if err != nil {
			if err != errors.ErrOnboardingNotFound {
				w.WriteHeader(http.StatusInternalServerError)
				messages := fmt.Sprintf("error, %s", err.Error())
				json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
				return
			}
		}
		if o != nil {
			continue
		}

		if imgData, ok := imageMap[obd.FilePath]; ok {
			onboardings[i].File = imgData
		}

	}

	const maxTotalSize = 20 * 1024 * 1024 // 20MB
	totalSize := int64(0)
	for _, p := range onboardings {
		if len(p.File) > 0 {
			totalSize += int64(len(p.File))
		}
	}
	if totalSize > maxTotalSize {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Size: %s", helpers.FormatFileSize(totalSize))
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	for _, o := range onboardings {
		err = h.OnboardingSrv.Create(context.Background(), o)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}
	}

	count := 0
	for _, o := range onboardings {
		if len(o.File) > 0 {
			count++
		}

	}

	w.WriteHeader(http.StatusOK)
}
