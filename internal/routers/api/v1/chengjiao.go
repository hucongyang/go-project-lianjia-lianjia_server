package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/service"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type Chengjiao struct{}

func NewChengjiao() Chengjiao {
	return Chengjiao{}
}

func (chengjiao Chengjiao) Get(c *gin.Context) {
	param := service.ChengjiaoGetDetailRequest{}
	param.FangziId = c.Param("fangziId")
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	sell, err := svc.GetChengjiaoDetail(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetChengjiaoDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponse(sell)
	return
}

func (chengjiao Chengjiao) List(c *gin.Context) {
	param := service.ChengjiaoListRequest{}
	param.XiaoquId = c.Query("xiaoquId")
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountChengjiao(&service.CountChengjiaoRequest{XiaoquId: param.XiaoquId})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountSell err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	chengjiaos, err := svc.GetChengjiaoList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetChengjiaoList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(chengjiaos, totalRows)
	return
}
