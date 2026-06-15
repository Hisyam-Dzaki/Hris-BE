package main

import (
	"log"

	"hris-BE/internal/config"
	"hris-BE/internal/database"
	"hris-BE/internal/handlers"
	"hris-BE/internal/repositories"
	"hris-BE/internal/router"
	"hris-BE/internal/services"
	jwtpkg "hris-BE/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := config.Load()

	// Database
	database.Connect(cfg)
	database.Migrate()

	// JWT
	jwtMgr := jwtpkg.NewManager(cfg.JWTSecret, cfg.JWTExpiryHour)

	// Repositories
	userRepo := repositories.NewUserRepository(database.GetDB())
	companyRepo := repositories.NewCompanyRepository(database.GetDB())
	roleRepo := repositories.NewRoleRepository(database.GetDB())
	employeeRepo := repositories.NewEmployeeRepository(database.GetDB())
	deptRepo := repositories.NewDepartmentRepository(database.GetDB())
	locRepo := repositories.NewLocationRepository(database.GetDB())

	// Services
	authService := services.NewAuthService(userRepo, jwtMgr)
	companyService := services.NewCompanyService(companyRepo)
	roleService := services.NewRoleService(roleRepo)
	employeeService := services.NewEmployeeService(employeeRepo)
	deptService := services.NewDepartmentService(deptRepo)
	locService := services.NewLocationService(locRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	companyHandler := handlers.NewCompanyHandler(companyService)
	roleHandler := handlers.NewRoleHandler(roleService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	deptHandler := handlers.NewDepartmentHandler(deptService)
	locHandler := handlers.NewLocationHandler(locService)

	// Router
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	r := router.New(
		engine,
		jwtMgr,
		authHandler,
		companyHandler,
		roleHandler,
		employeeHandler,
		deptHandler,
		locHandler,
	)
	r.Setup()

	log.Printf("HRIS Backend running on port %s", cfg.AppPort)
	if err := engine.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
