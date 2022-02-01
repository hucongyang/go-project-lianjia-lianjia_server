package service

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 业务逻辑的封装

// 每个接口的参数各自定义一个校验参数
// form：表单的映射字段名
// binding：入参校验的规则内容

type CountDistrictRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type DistrictListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type CreateDistrictRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateDistrictRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	Status     uint8  `form:"status,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteDistrictRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountDistrict(param *CountDistrictRequest) (int, error) {
	return svc.dao.CountDistrict(param.Name, param.Status)
}

func (svc *Service) GetDistrictList(param *DistrictListRequest, pager *app.Pager) ([]*model.District, error) {
	return svc.dao.GetDistrictList(param.Name, param.Status, pager.Page, pager.PageSize)
}

func (svc *Service) CreateDistrict(param *CreateDistrictRequest) error {
	return svc.dao.CreateDistrict(param.Name, param.Status, param.CreatedBy)
}

func (svc *Service) UpdateDistrict(param *UpdateDistrictRequest) error {
	return svc.dao.UpdateDistrict(param.ID, param.Name, param.Status, param.ModifiedBy)
}

func (svc *Service) DeleteDistrict(param *DeleteDistrictRequest) error {
	return svc.dao.DeleteDistrict(param.ID)
}
