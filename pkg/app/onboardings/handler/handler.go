package handler

import (
	"net/http"

	mdw "github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/port"
)

type Handler struct {
	OnboardingSrv  port.OnboardingSrv
	OnboardingRepo port.OnboardingRepo
}

func NewOnboardingHdl(server *http.ServeMux, onboardingSrv port.OnboardingSrv, onboardingRepo port.OnboardingRepo, authMdw *mdw.AuthMiddleware) *Handler {
	h := &Handler{
		OnboardingSrv:  onboardingSrv,
		OnboardingRepo: onboardingRepo,
	}

	server.HandleFunc("GET /onboardings", h.GetAll)
	server.HandleFunc("POST /onboardings/load-from-zip", h.LoadFromZip)
	return h
}
