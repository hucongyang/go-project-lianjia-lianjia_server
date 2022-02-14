package service

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

type CountChengjiaoRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

type ChengjiaoListRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

type ChengjiaoGetDetailRequest struct {
	FangziId string `form:"fangzi_id" `
}

func (svc *Service) GetChengjiaoDetail(param *ChengjiaoGetDetailRequest) (model.Chengjiao, error) {
	return svc.dao.GetChengjiaoDetail(param.FangziId)
}

func (svc *Service) CountChengjiao(param *CountChengjiaoRequest) (int, error) {
	return svc.dao.CountChengjiao(param.XiaoquId)
}

func (svc *Service) GetChengjiaoList(param *ChengjiaoListRequest, pager *app.Pager) ([]*model.Chengjiao, error) {
	return svc.dao.GetChengjiaoList(param.XiaoquId, pager.Page, pager.PageSize)
}
