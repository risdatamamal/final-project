package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/risdatamamal/final-project/config/gin"
	"github.com/risdatamamal/final-project/pkg/domain/user"
	"github.com/risdatamamal/final-project/pkg/server/http/middleware"
	"github.com/risdatamamal/final-project/pkg/server/http/router"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {

	// setiap yang /v1/user
	// harus melakukan pengecheckan auth
	// sehingga kita bisa meletakkan middleware di dalam group kita

	routerGroup := ginEngine.GetGin().Group("/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	// all path for get method are here
}

func (u *UserRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("",
		middleware.CheckJwtAuth, u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
}
