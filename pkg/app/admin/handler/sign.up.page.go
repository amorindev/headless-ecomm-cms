package handler

import "net/http"

func (h *Handler) SignUpAdminPage(w http.ResponseWriter, r *http.Request) {
	h.Renderer.Render(w, "sign-up-admin.html", nil)
}
