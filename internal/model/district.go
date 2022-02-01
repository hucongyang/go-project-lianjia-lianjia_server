package model

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/jinzhu/gorm"
)

type District struct {
	*Model
	Name              string `json:"name"`
	Quanpin           string `json:"quanpin"`
	LianjiaDistrictId int    `json:"lianjia_district_id"`
	Status            uint8  `json:"status"`
}

func (d District) TableName() string {
	return "lj_district"
}

type DistrictSwagger struct {
	List  []*District
	Pager *app.Pager
}

func (d District) Count(db *gorm.DB) (int, error) {
	var count int
	if d.Name != "" {
		db = db.Where("name = ?", d.Name)
	}
	db = db.Where("status = ?", d.Status)
	db = db.Where("is_del = ?", 0)
	if err := db.Model(&d).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d District) List(db *gorm.DB, pageOffset, pageSize int) ([]*District, error) {
	var districts []*District
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if d.Name != "" {
		db = db.Where("name = ?", d.Name)
	}
	db = db.Where("status = ?", d.Status)
	db = db.Where("is_del = ?", 0)
	if err = db.Find(&districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (d District) Get(db *gorm.DB) (District, error) {
	var district District
	err := db.Where("id = ? AND is_del = ? AND state = ?", d.ID, 0, d.Status).First(&district).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return d, err
	}

	return d, nil
}

func (d District) ListByIDs(db *gorm.DB, ids []uint32) ([]*District, error) {
	var districts []*District
	db = db.Where("state = ? AND is_del = ?", d.Status, 0)
	err := db.Where("id IN (?)", ids).Find(districts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return districts, nil
}

func (d District) Create(db *gorm.DB) error {
	return db.Create(&d).Error
}

func (d District) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&d).Where("id = ? AND is_del = ?", d.ID, 0).Updates(values).Error
}

func (d District) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", d.Model.ID, 0).Delete(&d).Error
}
