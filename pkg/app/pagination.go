package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/convert"
)

// 分页处理

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

// GetPageSize 获取每页数量
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

// GetPageOffset 获取页码下的偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
