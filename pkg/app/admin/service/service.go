package service

import "github.com/amorindev/headless-ecomm-cms/pkg/app/admin/port"

var _ port.AdminSrv = &Service{}

type Service struct {
	AdminRepo port.AdminRepo
}

func NewAdminSrv(adminRepo port.AdminRepo) *Service {
	return &Service{
		AdminRepo: adminRepo,
	}
}