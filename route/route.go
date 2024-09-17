package route

//加载gin
import (
	"GoEasyApi/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	//健康接口
	r.GET("/health", controller.Health)
	r.GET("/api/*path", controller.GetApi)
	r.POST("/api/*path", controller.PostApi)
	r.POST("/manger_login", controller.MangerLogin)
}
