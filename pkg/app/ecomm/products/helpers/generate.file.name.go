package helpers

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateUniqueFileName(imgPath string) string {
	fileName := imgPath[:len(imgPath)-len(filepath.Ext(imgPath))]
	uniqueFileName := fmt.Sprintf("%s-%d%s", fileName, time.Now().UnixNano(), filepath.Ext(imgPath))
	return uniqueFileName
}
