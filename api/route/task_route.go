package route

import (
	"time"

	"go-gin-be-clean-arch/api/controller"
	"go-gin-be-clean-arch/bootstrap"
	"go-gin-be-clean-arch/domain"
	"go-gin-be-clean-arch/mongo"
	"go-gin-be-clean-arch/repository"
	"go-gin-be-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
