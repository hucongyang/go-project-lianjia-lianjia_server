package model

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/jinzhu/gorm"
)

type Xiaoqu struct {
	*Model
	XiaoquId       string `json:"xiaoqu_id"`
	Url            string `json:"url"`
	BuildDate      string `json:"build_date"`
	ChengJiaoText  string `json:"cheng_jiao_text"`
	ChengJiaoUrl   string `json:"cheng_jiao_url"`
	DistrictId     uint32 `json:"district_id"`
	Img            string `json:"img"`
	HouseSellUrl   string `json:"sell_url"`
	TagList        string `json:"tag_list"`
	Title          string `json:"title"`
	TotalPrice     string `json:"total_price"`
	TotalSellCount string `json:"total_sell_count"`
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
