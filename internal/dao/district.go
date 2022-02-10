package dao

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 数据访问对象的封装

func (d *Dao) CountDistrict(name string, status uint8) (int, error) {
	district := model.District{
		Name: name,
	}
	return district.Count(d.engine)
}

func (d *Dao) GetDistrict(id uint32, status uint8) (model.District, error) {
	district := model.District{Model: &model.Model{ID: id}}
	return district.Get(d.engine)
}

func (d *Dao) GetDistrictList(name string, status uint8, page, pageSize int) ([]*model.District, error) {
	district := model.District{Name: name}
	pageOffset := app.GetPageOffset(page, pageSize)
	return district.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetDistrictListByIDs(ids []uint32, status uint8) ([]*model.District, error) {
	district := model.District{}
	return district.ListByIDs(d.engine, ids)
}

func (d *Dao) CreateDistrict(name string, status uint8, createdBy string) error {
	tag := model.District{
		Name: name,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateDistrict(id uint32, name string, status uint8, modifiedBy string) error {
	tag := model.District{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"status":      status,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteDistrict(id uint32) error {
	tag := model.District{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
