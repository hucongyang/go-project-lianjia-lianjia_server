package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hucongyang/go-project-lianjia-lianjia_server/docs"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/middleware"
	v1 "github.com/hucongyang/go-project-lianjia-lianjia_server/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册中间件：国际化处理
	r.Use(middleware.Translations())

	// 区域
	district := v1.NewDistrict()
	// 小区
	xiaoqu := v1.NewXiaoqu()
	// 小区房子
	fangzi := v1.NewSell()
	// 成交房子
	chengjiao := v1.NewChengjiao()

	// 添加 swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/districts", district.Create)
		apiV1.DELETE("/districts/:id", district.Delete)
		apiV1.PUT("/districts/:id", district.Update)
		apiV1.PATCH("/districts/:id/status", district.Update)
		apiV1.GET("/districts", district.List)

		apiV1.POST("/search/xiaoqus", xiaoqu.Search)
		apiV1.POST("/xiaoqus", xiaoqu.Create)
		apiV1.DELETE("/xiaoqus/:id", xiaoqu.Delete)
		apiV1.PUT("/xiaoqus/:id", xiaoqu.Update)
		apiV1.PATCH("/xiaoqus/:id/status", xiaoqu.Update)
		apiV1.GET("/xiaoqus", xiaoqu.List)
		apiV1.GET("/xiaoqus/:xiaoquId", xiaoqu.Get)
		apiV1.GET("/xiaoqus/:xiaoquId/history", xiaoqu.GetHistory)

		apiV1.GET("/fangzis", fangzi.List)
		apiV1.GET("/fangzis/detail/:fangziId", fangzi.Get)

		apiV1.GET("/chengjiaos", chengjiao.List)
		apiV1.GET("/chengjiaos/detail/:fangziId", chengjiao.Get)

	}

	return r
}
