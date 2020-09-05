package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New() //不使用gin的默认咯
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		//暂时不需要获取一个标签的需求
		apiv1.GET("/tags", tag.List) //标签列表
		apiv1.POST("/tags", tag.Update)
		apiv1.PUT("/tags/:id", tag.Create) //更新
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PATCH("/tags/:id/state", tag.Update)

		apiv1.POST("/articles", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update) //更新
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List) //拿所有文章
	}
	return r
}

func A(c *gin.Context) {}
