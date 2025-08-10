package handler

import "net/http"

func (h *Handler) SignInAdminPage(w http.ResponseWriter, r *http.Request) {
	h.Renderer.Render(w, "sign-in-admin.html", nil)
}
