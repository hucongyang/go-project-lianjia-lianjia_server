package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type District struct{}

func NewDistrict() District {
	return District{}
}

func (d District) Get(c *gin.Context) {

}

// List
// @Summary 获取多个商业圈
// @Produce  json
// @Param name query string false "商业圈名称" maxlength(200)
// @Param quanpin query string false "商业圈拼音" maxlength(200)
// @Param status query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.DistrictSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/districts [get]
func (d District) List(c *gin.Context) {
	param := struct {
		Name   string `form:"name" binding:"max=100"`
		Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	response.ToResponse(gin.H{})
	return
}

// Create
// @Summary 新增商业圈
// @Produce  json
// @Param name body string true "商业圈名称" minlength(3) maxlength(200)
// @Param quanpin body string true "商业圈拼音" minlength(3) maxlength(200)
// @Param lianjia_district_id body int true "链家网站原始区域ID"
// @Param status body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.District "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/districts [post]
func (d District) Create(c *gin.Context) {

}

func (d District) Update(c *gin.Context) {

}

func (d District) Delete(c *gin.Context) {

}
