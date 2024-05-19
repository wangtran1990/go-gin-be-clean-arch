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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
