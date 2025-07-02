package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"misconfig-integrator/app/controller"
	"misconfig-integrator/app/dbo/cli"
	"misconfig-integrator/app/dbo/repository"
	"misconfig-integrator/app/dbo/repository/project"
	"misconfig-integrator/app/service"
)

func ProjectRoute(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) *gin.RouterGroup {
	projectRepository := repository.NewProjectRepository()
	exclusionRepository := project.NewExclusionRepository()
	optionRepository := project.NewFilterOptionRepository()
	resultRepository := repository.NewResultRepository()
	projectAuthRepository := project.NewAuthRepository()
	trivyCli := cli.NewTrivyCli()
	gitCli := cli.NewGitCli()

	gitCliService := service.NewGitCliService(gitCli)
	projectService := service.NewProjectService(projectRepository,
		optionRepository,
		exclusionRepository,
		resultRepository,
		projectAuthRepository,
		trivyCli,
		gitCliService,
		validate,
		db)

	projectController := controller.NewProjectController(projectService)

	router.GET("projects", projectController.GetAllHandler)
	router.GET("project/:id", projectController.GetDetailByIdHandler)
	router.POST("project", projectController.CreateProjectHandler)
	router.POST("project/scan", projectController.ScanProjectHandler)
	router.POST("sonar-callback", projectController.SonarCallbackHandler)
	router.PUT("project/:id", projectController.UpdateProjectHandler)
	router.DELETE("project/:id", projectController.DeleteProjectHandler)

	return router
}
