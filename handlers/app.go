package handlers

import (
	"final-project-3/database"
	_ "final-project-3/docs"
	"final-project-3/handlers/http_handlers"
	"final-project-3/middlewares"
	"final-project-3/repositories/category_repository/category_pg"
	"final-project-3/repositories/task_repository/task_pg"
	"final-project-3/repositories/user_repository/user_pg"
	"final-project-3/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title MyGram API
// @version 1.0
// @description This is a sample server for ... .
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	database.StartDB()
	db := database.GetPostgresInstance()

	router := gin.Default()

	userRepo := user_pg.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := http_handlers.NewUserHandler(userService)

	// seeding admin with email: admin@gmail.com,
	// password: 123456
	userRepo.SeedingAdmin()

	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/register", userHandler.RegisterUser)
		usersRouter.POST("/login", userHandler.LoginUser)
		usersRouter.PUT("/update-account", middlewares.Authentication(), userHandler.UpdateUser)
		usersRouter.DELETE("/delete-account", middlewares.Authentication(), userHandler.DeleteUser)
	}

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := http_handlers.NewCategoryHandler(categoryService)

	categoryRouter := router.Group("/category")
	{
		categoryRouter.POST("/", middlewares.Authentication(), middlewares.CategoryAuthorization(), categoryHandler.CreateCategory)
		categoryRouter.GET("/", categoryHandler.GetAllCategory)
		categoryRouter.PATCH("/:categoryId", middlewares.Authentication(), middlewares.CategoryAuthorization(), categoryHandler.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", middlewares.Authentication(), middlewares.CategoryAuthorization(), categoryHandler.DeleteCategory)
	}

	taskRepo := task_pg.NewTaskPG(db)
	taskService := services.NewTaskService(taskRepo, categoryRepo)
	taskHandler := http_handlers.NewTaskHandler(taskService)

	tasksRouter := router.Group("/tasks")
	tasksRouter.Use(middlewares.Authentication())
	{
		tasksRouter.POST("/", taskHandler.CreateTask)
		// tasksRouter.POST("/", taskHandler.ViewAllTasks)
		// tasksRouter.PUT("/:taskId", middlewares.TaskAuthorization(), taskHandler.UpdateTitleAndDesc)
		tasksRouter.PATCH("/update-status/:taskId", middlewares.TaskAuthorization(), taskHandler.UpdateStatus)
		tasksRouter.PATCH("/update-category/:taskId", middlewares.TaskAuthorization(), taskHandler.UpdateCategoryIdOfTask)
		// tasksRouter.DELETE("/:taskId", middlewares.TaskAuthorization(), taskHandler.Deletetask)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}
