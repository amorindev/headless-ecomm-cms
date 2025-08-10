package service

import "github.com/amorindev/headless-ecomm-cms/pkg/file-storage/port"

var _ port.FileStorageSrv = &Service{}

type Service struct {
	FileStgAdp port.FileStorageAdapter
}

func NewFileStgSrv(fileStgAdp port.FileStorageAdapter) *Service{
	return &Service{
		FileStgAdp: fileStgAdp,
	}
}