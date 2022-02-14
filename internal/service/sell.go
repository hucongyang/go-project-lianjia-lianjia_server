package service

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

type CountSellRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

type SellListRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

type SellGetDetailRequest struct {
	FangziId string `form:"fangzi_id" `
}

func (svc *Service) GetSellDetail(param *SellGetDetailRequest) (model.Sell, error) {
	return svc.dao.GetSellDetail(param.FangziId)
}

func (svc *Service) CountSell(param *CountSellRequest) (int, error) {
	return svc.dao.CountSell(param.XiaoquId)
}

func (svc *Service) GetSellList(param *SellListRequest, pager *app.Pager) ([]*model.Sell, error) {
	return svc.dao.GetSellList(param.XiaoquId, pager.Page, pager.PageSize)
}
