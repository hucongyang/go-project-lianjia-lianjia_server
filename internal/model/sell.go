package model

import (
	"github.com/jinzhu/gorm"
)

type Sell struct {
	*Model
	XiaoquId     string `json:"xiaoquId"`
	FangziId     string `json:"fangziId"`
	Img          string `json:"img"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	HouseInfo    string `json:"houseInfo"`
	PositionInfo string `json:"positionInfo"`
	Follow       string `json:"tagList"`
	TagList      string `json:"follow"`
	TotalPrice   string `json:"totalPrice"`
	UnitPrice    string `json:"unitPrice"`
}

func (sell Sell) TableName() string {
	return "lj_sell"
}

func (sell Sell) Count(db *gorm.DB) (int, error) {
	var count int
	db = db.Where("xiaoqu_id = ?", sell.XiaoquId)
	if err := db.Model(&sell).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (sell Sell) List(db *gorm.DB, pageOffset, pageSize int) ([]*Sell, error) {
	var sells []*Sell
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("xiaoqu_id = ?", sell.XiaoquId)
	if err = db.Find(&sells).Error; err != nil {
		return nil, err
	}
	return sells, nil
}

func (sell Sell) Detail(db *gorm.DB) (Sell, error) {
	var detail Sell
	err := db.Where("fangzi_id = ?", sell.FangziId).First(&detail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return detail, err
	}

	return detail, nil
}
