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
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/helpers"
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

	var products []*domain.Product
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
			err = json.Unmarshal(content, &products)
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

	for i, product := range products {

		if imgData, ok := imageMap[product.FilePath]; ok {
			products[i].File = imgData
		}

		for j, pItem := range product.ProductItems {
			var optionIDs []interface{}

			for _, option := range pItem.Options {
				vOpt, err := h.VarOptionRepo.FindByName(context.Background(), option.VarOptName)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					messages := fmt.Sprintf("error, %s", err.Error())
					json.NewEncoder(w).Encode(message.ErrorMessage{Message: messages})
					return

				}
				optionIDs = append(optionIDs, vOpt.ID.(string))
			}

			products[i].ProductItems[j].VarOptionIDs = optionIDs

			if imgData, ok := imageMap[pItem.FilePath]; ok {
				products[i].ProductItems[j].File = imgData
			}
		}

	}

	const maxTotalSize = 20 * 1024 * 1024 // 20MB
	totalSize := int64(0)
	for _, p := range products {
		if len(p.File) > 0 {
			totalSize += int64(len(p.File))
		}
		for _, pItem := range p.ProductItems {
			if len(pItem.File) > 0 {
				totalSize += int64(len(pItem.File))
			}
		}
	}
	if totalSize > maxTotalSize {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Size: %s", helpers.FormatFileSize(totalSize))
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}


	for _, p := range products {
		err = h.ProductSrv.CreateFromZip(context.TODO(), p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}
	}

	count := 0
	for _, p := range products {
		if len(p.File) > 0 {
			count++
		}
		for _, pItem := range p.ProductItems {
			if len(pItem.File) > 0 {
				count++
			}
		}
	}


	w.WriteHeader(http.StatusOK)
}

