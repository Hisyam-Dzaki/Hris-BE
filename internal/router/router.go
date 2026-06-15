package router

import (
	"hris-BE/internal/handlers"
	"hris-BE/internal/middleware"
	"hris-BE/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine           *gin.Engine
	jwtMgr           *jwt.JWTManager
	authHandler      *handlers.AuthHandler
	companyHandler   *handlers.CompanyHandler
	roleHandler      *handlers.RoleHandler
	employeeHandler  *handlers.EmployeeHandler
	deptHandler      *handlers.DepartmentHandler
	locHandler       *handlers.LocationHandler
}

func New(
	engine *gin.Engine,
	jwtMgr *jwt.JWTManager,
	authHandler *handlers.AuthHandler,
	companyHandler *handlers.CompanyHandler,
	roleHandler *handlers.RoleHandler,
	employeeHandler *handlers.EmployeeHandler,
	deptHandler *handlers.DepartmentHandler,
	locHandler *handlers.LocationHandler,
) *Router {
	return &Router{
		engine:          engine,
		jwtMgr:          jwtMgr,
		authHandler:     authHandler,
		companyHandler:  companyHandler,
		roleHandler:     roleHandler,
		employeeHandler: employeeHandler,
		deptHandler:     deptHandler,
		locHandler:      locHandler,
	}
}

func (r *Router) Setup() {
	r.engine.Use(middleware.CORSMiddleware())

	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "hris-BE"})
	})

	api := r.engine.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", r.authHandler.Register)
		auth.POST("/login", r.authHandler.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(r.jwtMgr))
	{
		// Auth
		protected.GET("/auth/me", r.authHandler.GetMe)

		// Company
		protected.POST("/companies", r.companyHandler.Create)
		protected.GET("/companies", r.companyHandler.GetAll)
		protected.GET("/companies/:id", r.companyHandler.GetByID)
		protected.PUT("/companies/:id", r.companyHandler.Update)
		protected.DELETE("/companies/:id", r.companyHandler.Delete)

		// Roles
		protected.POST("/roles", r.roleHandler.Create)
		protected.GET("/roles", r.roleHandler.GetByCompanyID)
		protected.GET("/roles/:id", r.roleHandler.GetByID)
		protected.PUT("/roles/:id", r.roleHandler.Update)
		protected.DELETE("/roles/:id", r.roleHandler.Delete)

		// Employees
		protected.POST("/employees", r.employeeHandler.Create)
		protected.GET("/employees", r.employeeHandler.GetAll)
		protected.GET("/employees/:id", r.employeeHandler.GetByID)
		protected.PUT("/employees/:id", r.employeeHandler.Update)
		protected.DELETE("/employees/:id", r.employeeHandler.Delete)

		// Departments
		protected.POST("/departments", r.deptHandler.Create)
		protected.GET("/departments", r.deptHandler.GetByCompanyID)
		protected.GET("/departments/:id", r.deptHandler.GetByID)
		protected.PUT("/departments/:id", r.deptHandler.Update)
		protected.DELETE("/departments/:id", r.deptHandler.Delete)

		// Locations
		protected.POST("/locations", r.locHandler.Create)
		protected.GET("/locations", r.locHandler.GetByCompanyID)
		protected.GET("/locations/:id", r.locHandler.GetByID)
		protected.PUT("/locations/:id", r.locHandler.Update)
		protected.DELETE("/locations/:id", r.locHandler.Delete)
	}
}
