package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hucongyang/go-project-lianjia-lianjia_server/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	district := v1.NewDistrict()
	xiaoqu := v1.NewXiaoqu()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/districts", district.Create)
		apiV1.DELETE("/districts/:id", district.Delete)
		apiV1.PUT("/districts/:id", district.Update)
		apiV1.PATCH("/districts/:id/status", district.Update)
		apiV1.GET("/districts", district.List)

		apiV1.POST("/xiaoqus", xiaoqu.Create)
		apiV1.DELETE("/xiaoqus/:id", xiaoqu.Delete)
		apiV1.PUT("/xiaoqus/:id", xiaoqu.Update)
		apiV1.PATCH("/xiaoqus/:id/status", xiaoqu.Update)
		apiV1.GET("/xiaoqus", xiaoqu.List)
	}

	return r
}
