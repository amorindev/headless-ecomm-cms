package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/internal/minio"
	mongoClient "github.com/amorindev/headless-ecomm-cms/internal/mongo"
	"github.com/amorindev/headless-ecomm-cms/internal/mongo/constants"
	resendClient "github.com/amorindev/headless-ecomm-cms/internal/resend"
	twilioClient "github.com/amorindev/headless-ecomm-cms/internal/twilio"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/admin/api"
	adminHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/admin/handler"
	authMethodHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/handler"
	authMethodService "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service"
	addressHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/handler"
	addressRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/repository/mongo"
	addressService "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/service"
	categoryHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/handler"
	categoryInitializer "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/initializer"
	categoryRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/repository/mongo"
	categoryService "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/service"
	productHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/handler"
	productRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/repository/mongo"
	productService "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/service"
	varOptInitializer "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/initializer"
	varOptRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/repository/mongo"
	variationHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/handler"
	variationInitializer "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/initializer"
	variationRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/repository/mongo"
	variationService "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/service"
	onboardingHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/handler"
	onboardingRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/repository/mongo"
	onboardingService "github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/service"
	otpCodeRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/repository/mongo"
	phoneHandler "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/handler"
	phoneRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/repository/mongo"
	phoneService "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/service"
	roleInitializer "github.com/amorindev/headless-ecomm-cms/pkg/app/roles/initializer"
	roleRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/roles/repository/mongo"
	sessionRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/repository/mongo"
	sessionService "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/service"
	userRepository "github.com/amorindev/headless-ecomm-cms/pkg/app/users/repository/mongo"
	resendAdapter "github.com/amorindev/headless-ecomm-cms/pkg/email/adapter/resend"
	emailService "github.com/amorindev/headless-ecomm-cms/pkg/email/service"
	minioAdapter "github.com/amorindev/headless-ecomm-cms/pkg/file-storage/minio/adapter"
	fileStgService "github.com/amorindev/headless-ecomm-cms/pkg/file-storage/service"
	twilioAdapter "github.com/amorindev/headless-ecomm-cms/pkg/sms/adapter/twilio"
	smsService "github.com/amorindev/headless-ecomm-cms/pkg/sms/service"
)

func New() http.Handler {

	mux := http.NewServeMux()

	// * Api version
	v1 := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	accessSecret := os.Getenv("")
	refreshSecret := os.Getenv("")
	issuer := os.Getenv("")

	// * Externals Clients

	// ** MongoDB
	mongoConn := mongoClient.New()

	mongoDB := mongoConn.DB.Database(os.Getenv("MONGO_INITDB_DATABASE"))
	mongoConn.Ping()

	// ** Email - Resend
	resendClient := resendClient.NewClient()

	// ** File Storage - Minio
	minioClient := minio.NewClient()
	minioClient.CreateStorage()

	// ** Sms - Twilio
	twilioClient := twilioClient.NewClient()

	// * Adapters layer
	minioAdp := minioAdapter.NewAdapter(minioClient.Client)
	resendAdapter := resendAdapter.NewAdapter(resendClient, os.Getenv("EMAIL_FROM"))
	twilioAdp := twilioAdapter.NewTwilioAdapter(twilioClient, os.Getenv("TWILIO_SMS_FROM"))

	// * Services layer
	fileStgSrv := fileStgService.NewFileStgSrv(minioAdp)
	emailSrv := emailService.NewEmailSrv(resendAdapter)
	twilioSrv := smsService.NewSmsSrv(twilioAdp)

	// * Collections
	roleColl := mongoDB.Collection(constants.CollRoles)
	varOptColl := mongoDB.Collection(constants.CollVarOptions)
	variationColl := mongoDB.Collection(constants.CollVariations)
	categoryColl := mongoDB.Collection(constants.CollCategories)
	productsColl := mongoDB.Collection(constants.CollProducts)
	otpColl := mongoDB.Collection(constants.CollOtpCodes)
	userColl := mongoDB.Collection(constants.CollUsers)
	mfaFaSmsColl := mongoDB.Collection(constants.CollMfaFaSms)
	phoneColl := mongoDB.Collection(constants.CollPhones)
	sessionColl := mongoDB.Collection(constants.CollSessions)
	addressColl := mongoDB.Collection(constants.CollAddress)
	onboardingColl := mongoDB.Collection(constants.CollOnboardings)

	// * Repositories
	variationRepo := variationRepository.NewVariationRepo(mongoConn.DB, variationColl)
	varOptRepo := varOptRepository.NewVarOptRepo(mongoConn.DB, varOptColl)
	categoryRepo := categoryRepository.NewCategoryRepo(mongoConn.DB, categoryColl)
	productRepo := productRepository.NewProductRepo(mongoConn.DB, productsColl)
	userMongoRepo := userRepository.NewUserRepo(mongoConn.DB, userColl, mfaFaSmsColl)
	otpCodeRepo := otpCodeRepository.NewOtpCodeRepo(mongoConn.DB, otpColl)
	phoneRepo := phoneRepository.NewPhoneRepository(mongoConn.DB, phoneColl)
	sessionMongoRepo := sessionRepository.NewSessionRepo(mongoConn.DB, sessionColl)
	roleMongoRepo := roleRepository.NewRoleRepo(mongoConn.DB, roleColl)
	addressRepo := addressRepository.NewAddressRepo(mongoConn.DB, addressColl)
	onboardingRepo := onboardingRepository.NewOnboardingRepo(mongoConn.DB, onboardingColl)

	// * Services
	authSrv := auth.NewTokenSrv(accessSecret, refreshSecret, issuer)
	variationSrv := variationService.NewVariationSrv(variationRepo, varOptRepo)
	productSrv := productService.NewProductSrv(productRepo, categoryRepo, fileStgSrv)
	categorySrv := categoryService.NewCategorySrv(categoryRepo, productRepo)
	sessionSrv := sessionService.NewSessionSrv(sessionMongoRepo, authSrv)
	authMethodSrv := authMethodService.NewAuthMethodSrv(userMongoRepo, roleMongoRepo, otpCodeRepo, sessionMongoRepo, phoneRepo, sessionSrv, emailSrv, twilioSrv)
	phoneSrv := phoneService.NewPhoneService(phoneRepo)
	addressSrv := addressService.NewAddressSrv(addressRepo)
	onboardingSrv := onboardingService.NewOnboardingSrv(onboardingRepo, fileStgSrv)

	// * Middlewares
	authMdw := middlewares.NewAuthMdw(authSrv)

	// * Handlers
	variationHandler.NewHandler(v1, variationSrv)
	productHandler.NewProductHdl(v1, productSrv, categoryRepo, varOptRepo, authMdw)
	categoryHandler.NewCategoryHdl(v1, categorySrv)
	authMethodHandler.NewAuthMethodHdl(v1, authMethodSrv, authMdw)
	phoneHandler.NewPhoneHdl(v1, phoneSrv, authMdw)
	addressHandler.NewAddressHdl(v1, addressSrv, authMdw)
	onboardingHandler.NewOnboardingHdl(v1, onboardingSrv, onboardingRepo, authMdw)

	// * Initializers
	variationItz := variationInitializer.NewVariationItz(variationRepo)
	if err := variationItz.SeedEssential(context.Background()); err != nil {
		log.Fatal(err)
	}
	varOptItz := varOptInitializer.NewVariationOptionItz(varOptRepo, variationRepo)
	if err := varOptItz.SeedEssential(context.Background()); err != nil {
		log.Fatal(err)
	}
	categoryItz := categoryInitializer.NewCategoryItz(categoryRepo)
	if err := categoryItz.SeedEssential(context.Background()); err != nil {
		log.Fatal(err)
	}
	roleItz := roleInitializer.NewRoleItz(roleMongoRepo)
	if err := roleItz.SeedEssentialRoles(context.Background()); err != nil {
		log.Fatal(err)
	}

	// * Ping Pong

	v1.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		type resp struct {
			Msg string `json:"pong"`
		}
		json.NewEncoder(w).Encode(resp{Msg: "pong"})
	})

	// * Templates
	// * redirect correct url
	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v1/admin/categories", http.StatusFound)
	})
	renderer := api.NewTemplateRenderer()
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		log.Fatal("missing required environment variable: API_BASE_URL")
	}
	adminHandler.NewAdminHandler(v1, renderer, apiBaseURL)

	return mux
}
