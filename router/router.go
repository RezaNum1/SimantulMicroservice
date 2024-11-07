package router

import (
	"Expire/config/middleware"
	"Expire/controller"
	"Expire/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewRouter(
	Db *gorm.DB,
	userController *controller.UserController,
	authController *controller.AuthController,
	tokenController *controller.TokenController,
	reportController *controller.ReportController,
	supervisorController *controller.SupervisorController,
	bankController *controller.BankController,
	leaderController *controller.LeaderController,
	reasonController *controller.ReasonController,
	externalController *controller.ExternalController,
) *gin.Engine {
	service := gin.Default()
	service.Use(CORS())

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Router has initialized")
	})

	service.NoRoute(func(c *gin.Context) {
		helper.ResponseError(c, helper.CustomError{
			Code:    404,
			Message: "Not Found.",
		})
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("/login", authController.SignInUser)
	authRouter.POST("/register", authController.SignUpUser)
	authRouter.POST("/forgetPassword", authController.VerifyForgetPassword)
	authRouter.PUT("/resetPassword", authController.ResetPassword)

	authRouter.GET("/refresh", tokenController.RefreshAccessToken)

	authRouter.Use(middleware.DeserializeUser(Db))
	authRouter.GET("/logout", authController.Logout)

	protectedUserRouter := router.Group("/user")
	protectedUserRouter.Use(middleware.AuthMiddleware())
	// protectedUserRouter.GET("", userController.GetUser)
	protectedUserRouter.POST("", userController.Create)
	protectedUserRouter.GET("", userController.GetAllUser)
	protectedUserRouter.DELETE("/:id", userController.Delete)

	protectedReportRouter := router.Group("/report")
	protectedReportRouter.Use(middleware.DeserializeUser(Db))
	protectedReportRouter.POST("", reportController.Create)
	protectedReportRouter.PUT("", reportController.Update)
	protectedReportRouter.PUT("/status", reportController.UpdateStatus)
	protectedReportRouter.PUT("/bank", reportController.UpdateBankReport)
	protectedReportRouter.PUT("/dokumen-temuan", reportController.UpdateDokumenTemuan)
	protectedReportRouter.PUT("/dokumen-tindak-lanjut", reportController.UpdateDokumenTindakLanjut)
	protectedReportRouter.GET("", reportController.GetAllReport)
	protectedReportRouter.GET("/:id", reportController.GetReport)
	protectedReportRouter.GET("/supervisor/:id", reportController.GetAllSupervisorReport)
	protectedReportRouter.GET("/leader/:id", reportController.GetAllLeaderReport)
	protectedReportRouter.GET("/bank/:id", reportController.GetAllBankReport)

	protectedSupervisorRouter := router.Group("/supervisor")
	protectedSupervisorRouter.Use(middleware.DeserializeUser(Db))
	protectedSupervisorRouter.POST("", supervisorController.Create)
	protectedSupervisorRouter.GET("", supervisorController.GetAllSupervisor)
	protectedSupervisorRouter.GET("/:id", supervisorController.GetSupervisor)

	protectedBankRouter := router.Group("/bank")
	protectedBankRouter.Use(middleware.DeserializeUser(Db))
	protectedBankRouter.POST("", bankController.Create)
	protectedBankRouter.GET("", bankController.GetAllBank)
	protectedBankRouter.GET("/:id", bankController.GetBank)

	protectedLeaderRouter := router.Group("/leader")
	protectedLeaderRouter.Use(middleware.DeserializeUser(Db))
	protectedLeaderRouter.POST("", leaderController.Create)
	protectedLeaderRouter.GET("", leaderController.GetAllLeader)
	protectedLeaderRouter.GET("/:id", leaderController.GetLeader)

	protectedReasonRouter := router.Group("/reason")
	protectedReasonRouter.Use(middleware.DeserializeUser(Db))
	protectedReasonRouter.POST("", reasonController.Create)
	protectedReasonRouter.GET("", reasonController.GetAllReason)
	protectedReasonRouter.GET("/:id", reasonController.GetReason)
	protectedReasonRouter.GET("/all/:id", reasonController.FindReasonsByReportID)

	protectedExternalRouter := router.Group("/external")
	protectedExternalRouter.Use(middleware.DeserializeUser(Db))
	protectedExternalRouter.POST("", externalController.Create)
	protectedExternalRouter.GET("", externalController.GetAllExternal)
	protectedExternalRouter.GET("/:id", externalController.GetExternal)
	protectedExternalRouter.PUT("", externalController.Update)

	return service
}
