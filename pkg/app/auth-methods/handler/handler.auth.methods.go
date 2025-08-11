package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/port"
)

type Handler struct {
	AuthMethodSrv port.AuthMethodSrv
}

func NewAuthMethodHdl(server *http.ServeMux, authMethodSrv port.AuthMethodSrv, authMdw *middlewares.AuthMiddleware) *Handler {
	h := &Handler{
		AuthMethodSrv: authMethodSrv,
	}

	// * Authentication handlers

	signUpH := logger.LoggerMdw(h.SignUp)
	server.HandleFunc("POST /auth/sign-up", signUpH)

	resendVEOH := logger.LoggerMdw(h.ResendVerifyEmailOtp)
	server.HandleFunc("POST /auth/resend-verify-email-otp", resendVEOH)

	verifyEmailOtpH := logger.LoggerMdw(h.VerifyEmailOtp)
	server.HandleFunc("POST /auth/verify-email-otp", verifyEmailOtpH)

	signInH := logger.LoggerMdw(h.SignIn)
	server.HandleFunc("POST /auth/sign-in", signInH)

	signOutH := logger.LoggerMdw(authMdw.RefreshTokenMdw(h.SignOut))
	server.HandleFunc("POST /auth/sign-out", signOutH)

	// * MFA handlers

	enableMfaSmsH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.EnableMfaSms))
	server.HandleFunc("POST /auth/enable-mfa-sms", enableMfaSmsH)

	resendVEMSOH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.ResendVerifyEnableMfaSmsOtp))
	server.HandleFunc("POST /auth/resend-verify-enable-mfa-sms-otp", resendVEMSOH)

	resendVMSO := logger.LoggerMdw(h.ResendVerifyMfaSmsOtp)
	server.HandleFunc("POST /auth/resend-verify-mfa-sms-otp", resendVMSO)

	verifyEMSOH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.VerifyEnableMfaSmsOtp))
	server.HandleFunc("POST /auth/verify-enable-mfa-sms-otp", verifyEMSOH)

	verifyMfaSmsOtp := logger.LoggerMdw(h.VerifyMfaSmsOtp)
	server.HandleFunc("POST /auth/verify-mfa-sms-otp", verifyMfaSmsOtp)

	return h
}
