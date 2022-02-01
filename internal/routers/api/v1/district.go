package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/service"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/convert"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type District struct{}

func NewDistrict() District {
	return District{}
}

// 数据操作流程：router路由 -> server服务 -> dao数据封装接口 -> model模型数据

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
	param := service.DistrictListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountDistrict(&service.CountDistrictRequest{Name: param.Name, Status: param.Status})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountDistrict err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetDistrictList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
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
	param := service.CreateDistrictRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateDistrict(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateDistrict err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (d District) Update(c *gin.Context) {
	param := service.UpdateDistrictRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateDistrict(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateDistrict err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (d District) Delete(c *gin.Context) {
	param := service.DeleteDistrictRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteDistrict(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteDistrict err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
