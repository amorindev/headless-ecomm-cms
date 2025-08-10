package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/minio"
	mongoClient "com.fernando/internal/mongo"
	config "com.fernando/internal/mongo/constants"
	resendClient "com.fernando/internal/resend"

	"com.fernando/internal/twilio"

	"com.fernando/internal/auth"
	"com.fernando/pkg/app/admin/api"
	adminHandler "com.fernando/pkg/app/admin/handler"
	addressHandler "com.fernando/pkg/app/ecomm/addresses/handler"
	addressRepository "com.fernando/pkg/app/ecomm/addresses/repository/mongo"
	addressService "com.fernando/pkg/app/ecomm/addresses/service"
	categoryHandler "com.fernando/pkg/app/ecomm/categories/handler"
	categoryInitializer "com.fernando/pkg/app/ecomm/categories/initializer"
	categoryRepository "com.fernando/pkg/app/ecomm/categories/repository/mongo"
	categoryService "com.fernando/pkg/app/ecomm/categories/service"
	productHandler "com.fernando/pkg/app/ecomm/products/handler"
	productService "com.fernando/pkg/app/ecomm/products/service"
	varOptInitializer "com.fernando/pkg/app/ecomm/var-options/initializer"
	"com.fernando/pkg/app/ecomm/var-options/repository/mongo"
	variationHandler "com.fernando/pkg/app/ecomm/variations/handler"
	variationInitializer "com.fernando/pkg/app/ecomm/variations/initializer"
	variationRepository "com.fernando/pkg/app/ecomm/variations/repository/mongo"
	variationService "com.fernando/pkg/app/ecomm/variations/service"
	onboardingHandler "com.fernando/pkg/app/onboardings/handler"
	onboardingRepository "com.fernando/pkg/app/onboardings/repository/mongo"
	onboardingService "com.fernando/pkg/app/onboardings/service"
	otpMongo "com.fernando/pkg/app/otp-codes/repository/mongo"
	phoneHandler "com.fernando/pkg/app/phones/handler"
	phoneRepository "com.fernando/pkg/app/phones/repository/mongo"
	phoneService "com.fernando/pkg/app/phones/service"

	twilioAdapter "com.fernando/pkg/sms/adapter/twilio"
	twilioService "com.fernando/pkg/sms/service"

	authMethodHandler "com.fernando/pkg/app/auth-methods/handler"
	userMongo "com.fernando/pkg/app/users/repository/mongo"

	authMethodSrv "com.fernando/pkg/app/auth-methods/service"

	productRepository "com.fernando/pkg/app/ecomm/products/repository/mongo"

	roleInitializer "com.fernando/pkg/app/roles/initializer"
	roleMongo "com.fernando/pkg/app/roles/repository/mongo"
	sessionMongo "com.fernando/pkg/app/sessions/repository/mongo"

	sessionSrv "com.fernando/pkg/app/sessions/service"
	resendAdapter "com.fernando/pkg/email/adapter/resend"
	emailService "com.fernando/pkg/email/service"
	minioAdapter "com.fernando/pkg/file-storage/minio/adapter"
	fileStgService "com.fernando/pkg/file-storage/service"
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
	twilioClient := twilio.NewClient()

	// * Adapters layer
	minioAdp := minioAdapter.NewAdapter(minioClient.Client)
	resendAdapter := resendAdapter.NewAdapter(resendClient, os.Getenv("EMAIL_FROM"))
	twilioAdp := twilioAdapter.NewTwilioAdapter(twilioClient, os.Getenv("TWILIO_SMS_FROM"))

	// * Services layer
	fileStgSrv := fileStgService.NewFileStgSrv(minioAdp)
	emailSrv := emailService.NewEmailSrv(resendAdapter)
	twilioSrv := twilioService.NewSmsSrv(twilioAdp)

	// * Collections
	roleColl := mongoDB.Collection(config.CollRoles)
	varOptColl := mongoDB.Collection(config.CollVarOptions)
	variationColl := mongoDB.Collection(config.CollVariations)
	categoryColl := mongoDB.Collection(config.CollCategories)
	productsColl := mongoDB.Collection(config.CollProducts)
	otpColl := mongoDB.Collection(config.CollOtpCodes)
	userColl := mongoDB.Collection(config.CollUsers)
	mfaFaSmsColl := mongoDB.Collection(config.CollMfaFaSms)
	phoneColl := mongoDB.Collection(config.CollPhones)
	sessionColl := mongoDB.Collection(config.CollSessions)
	addressColl := mongoDB.Collection(config.CollAddress)
	onboardingColl := mongoDB.Collection(config.CollOnboardings)

	// * Repositories
	variationRepo := variationRepository.NewRepository(mongoConn.DB, variationColl)
	varOptRepo := mongo.NewRepository(mongoConn.DB, varOptColl)
	categoryRepo := categoryRepository.NewCategoryRepo(mongoConn.DB, categoryColl)
	productRepo := productRepository.NewRepository(mongoConn.DB, productsColl)
	userMongoRepo := userMongo.NewUserRepo(mongoConn.DB, userColl, mfaFaSmsColl)
	otpRepository := otpMongo.NewRepository(mongoConn.DB, otpColl)
	phoneRepo := phoneRepository.NewRepository(mongoConn.DB, phoneColl)
	sessionMongoRepo := sessionMongo.NewRepository(mongoConn.DB, sessionColl)
	roleMongoRepo := roleMongo.NewRepository(mongoConn.DB, roleColl)
	addressRepo := addressRepository.NewAddressRepo(mongoConn.DB, addressColl)
	onboardingRepo := onboardingRepository.NewRepository(mongoConn.DB, onboardingColl)

	// * Services
	authSrv := auth.NewTokenSrv(accessSecret, refreshSecret, issuer)
	variationSrv := variationService.NewService(variationRepo, varOptRepo)
	productSrv := productService.NewService(productRepo, categoryRepo, fileStgSrv)
	categorySrv := categoryService.NewService(categoryRepo, productRepo)
	sessionSrv := sessionSrv.NewSessionSrv(sessionMongoRepo, authSrv)
	authMethodSrv := authMethodSrv.NewService(userMongoRepo, roleMongoRepo, otpRepository, sessionMongoRepo, phoneRepo, sessionSrv, emailSrv, twilioSrv)
	phoneSrv := phoneService.NewService(phoneRepo)
	addressSrv := addressService.NewService(addressRepo)
	onboardingSrv := onboardingService.NewService(onboardingRepo, fileStgSrv)

	// * Middlewares
	authMdw := middlewares.NewAuthMdw(authSrv)

	// * Handlers
	variationHandler.NewHandler(v1, variationSrv)
	productHandler.NewHandler(v1, productSrv, categoryRepo, varOptRepo, authMdw)
	categoryHandler.NewHandler(v1, categorySrv)
	authMethodHandler.NewHandler(v1, authMethodSrv, authMdw)
	phoneHandler.NewHandler(v1, phoneSrv, authMdw)
	addressHandler.NewHandler(v1, addressSrv, authMdw)
	onboardingHandler.NewHandler(v1, onboardingSrv, onboardingRepo, authMdw)

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
