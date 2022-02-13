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

type XiaoquListRequest struct {
	DistrictId uint32 `form:"district_id" `
}

func (svc *Service) CountXiaoqu(param *CountXiaoquRequest) (int, error) {
	return svc.dao.CountXiaoqu(param.DistrictId)
}

func (svc *Service) GetXiaoquList(param *XiaoquListRequest, pager *app.Pager) ([]*model.Xiaoqu, error) {
	return svc.dao.GetXiaoquList(param.DistrictId, pager.Page, pager.PageSize)
}
