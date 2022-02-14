package service

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 业务逻辑的封装

// 每个接口的参数各自定义一个校验参数
// form：表单的映射字段名
// binding：入参校验的规则内容

type CountXiaoquRequest struct {
	DistrictId uint32 `form:"district_id" `
}

type XiaoquGetDetailRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

type XiaoquListRequest struct {
	DistrictId uint32 `form:"district_id" `
}

type XiaoquGetHistoryRequest struct {
	XiaoquId string `form:"xiaoqu_id" `
}

func (svc *Service) GetXiaoquDetail(param *XiaoquGetDetailRequest) (model.Xiaoqu, error) {
	return svc.dao.GetXiaoquDetail(param.XiaoquId)
}

func (svc *Service) CountXiaoqu(param *CountXiaoquRequest) (int, error) {
	return svc.dao.CountXiaoqu(param.DistrictId)
}

func (svc *Service) GetXiaoquList(param *XiaoquListRequest, pager *app.Pager) ([]*model.Xiaoqu, error) {
	return svc.dao.GetXiaoquList(param.DistrictId, pager.Page, pager.PageSize)
}

func (svc *Service) GetXiaoquHistoryList(param *XiaoquGetHistoryRequest, pager *app.Pager) ([]*model.XiaoquLog, error) {
	return svc.dao.GetXiaoquHistoryList(param.XiaoquId, pager.Page, pager.PageSize)
}
