package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/service"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type Xiaoqu struct{}

func NewXiaoqu() Xiaoqu {
	return Xiaoqu{}
}

func (x Xiaoqu) Get(c *gin.Context) {
	param := service.XiaoquGetDetailRequest{}
	param.XiaoquId = c.Param("xiaoquId")
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	xiaoqu, err := svc.GetXiaoquDetail(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetXiaoquDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponse(xiaoqu)
	return
}

func (x Xiaoqu) List(c *gin.Context) {
	param := service.XiaoquListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountXiaoqu(&service.CountXiaoquRequest{DistrictId: param.DistrictId})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountDistrict err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetXiaoquList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetXiaoquList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

func (x Xiaoqu) Create(c *gin.Context) {}

func (x Xiaoqu) Update(c *gin.Context) {}

func (x Xiaoqu) Delete(c *gin.Context) {}

func (x Xiaoqu) Search(c *gin.Context) {
	param := service.SearchXiaoquRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	xiaoqus, err := svc.SearchXiaoquList(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.SearchXiaoqu err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponseList(xiaoqus, 10)
	return
}

func (x Xiaoqu) GetHistory(c *gin.Context) {
	param := service.XiaoquGetHistoryRequest{}
	param.XiaoquId = c.Param("xiaoquId")
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: 1, PageSize: 6}
	xiaoqulogs, err := svc.GetXiaoquHistoryList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetXiaoquHistoryList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	// 数据做处理
	response.ToResponse(xiaoqulogs)
	return
}
