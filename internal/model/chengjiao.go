package model

import (
	"github.com/jinzhu/gorm"
)

type Chengjiao struct {
	*Model
	XiaoquId        string `json:"xiaoquId"`
	FangziId        string `json:"fangziId"`
	Url             string `json:"url"`
	UnitPrice       string `json:"unitPrice"`
	TotalPrice      string `json:"totalPrice"`
	Title           string `json:"title"`
	PositionInfo    string `json:"positionInfo"`
	Img             string `json:"img"`
	HouseInfo       string `json:"houseInfo"`
	DealHouseInfo   string `json:"DealHouseInfo"`
	DealDate        string `json:"DealDate"`
	DealCycleePrice string `json:"DealCycleePrice"`
	DealCycleeDay   string `json:"DealCycleeDay"`
}

func (chengjiao Chengjiao) TableName() string {
	return "lj_chengjiao"
}

func (chengjiao Chengjiao) Count(db *gorm.DB) (int, error) {
	var count int
	db = db.Where("xiaoqu_id = ?", chengjiao.XiaoquId)
	if err := db.Model(&chengjiao).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (chengjiao Chengjiao) List(db *gorm.DB, pageOffset, pageSize int) ([]*Chengjiao, error) {
	var chengjiaos []*Chengjiao
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("xiaoqu_id = ?", chengjiao.XiaoquId)
	if err = db.Find(&chengjiaos).Error; err != nil {
		return nil, err
	}
	return chengjiaos, nil
}

func (chengjiao Chengjiao) Detail(db *gorm.DB) (Chengjiao, error) {
	var detail Chengjiao
	err := db.Where("fangzi_id = ?", chengjiao.FangziId).First(&detail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return detail, err
	}

	return detail, nil
}
