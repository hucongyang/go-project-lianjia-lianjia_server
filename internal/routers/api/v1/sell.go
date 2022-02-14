package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/service"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type Sell struct{}

func NewSell() Sell {
	return Sell{}
}

func (s Sell) Get(c *gin.Context) {
	param := service.SellGetDetailRequest{}
	param.FangziId = c.Param("fangziId")
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	sell, err := svc.GetSellDetail(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetSellDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponse(sell)
	return
}

func (s Sell) List(c *gin.Context) {
	param := service.SellListRequest{}
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
	totalRows, err := svc.CountSell(&service.CountSellRequest{XiaoquId: param.XiaoquId})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountSell err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	sells, err := svc.GetSellList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetSellList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(sells, totalRows)
	return
}
