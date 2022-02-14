package model

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/jinzhu/gorm"
)

type Xiaoqu struct {
	*Model
	XiaoquId       string `json:"xiaoquId"`
	Url            string `json:"url"`
	BuildDate      string `json:"buildDate"`
	ChengJiaoText  string `json:"cheng_jiao_text"`
	ChengJiaoUrl   string `json:"cheng_jiao_url"`
	DistrictId     uint32 `json:"district_id"`
	DistrictTitle  string `json:"district_title"`
	Img            string `json:"img"`
	HouseSellUrl   string `json:"sell_url"`
	TagList        string `json:"tagList"`
	Title          string `json:"title"`
	TotalPrice     string `json:"totalPrice"`
	TotalSellCount string `json:"totalSellCount"`
}

func (xiaoqu Xiaoqu) TableName() string {
	return "lj_xiaoqu"
}

type XiaoquSwagger struct {
	List  []*Xiaoqu
	Pager *app.Pager
}

func (xiaoqu Xiaoqu) Count(db *gorm.DB) (int, error) {
	var count int
	if xiaoqu.DistrictId > 0 {
		db = db.Where("district_id = ?", xiaoqu.DistrictId)
	}
	if err := db.Model(&xiaoqu).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (xiaoqu Xiaoqu) List(db *gorm.DB, pageOffset, pageSize int) ([]*Xiaoqu, error) {
	var xiaoqus []*Xiaoqu
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if xiaoqu.DistrictId > 0 {
		db = db.Where("district_id = ?", xiaoqu.DistrictId)
	}
	if err = db.Find(&xiaoqus).Error; err != nil {
		return nil, err
	}
	return xiaoqus, nil
}

func (xiaoqu Xiaoqu) Detail(db *gorm.DB) (Xiaoqu, error) {
	var detail Xiaoqu
	err := db.Where("xiaoqu_id = ?", xiaoqu.XiaoquId).First(&detail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return detail, err
	}

	return detail, nil
}

func (xiaoqu Xiaoqu) GetHistory(db *gorm.DB, pageOffset, pageSize int) ([]*XiaoquLog, error) {
	var xiaoqulogs []*XiaoquLog
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if xiaoqu.XiaoquId != "" {
		db = db.Where("xiaoqu_id = ?", xiaoqu.XiaoquId)
	}
	if err = db.Order("fetch_time desc").Find(&xiaoqulogs).Error; err != nil {
		return nil, err
	}
	return xiaoqulogs, nil
}
