package model

import (
	"github.com/jinzhu/gorm"
)

type XiaoquLog struct {
	*Model
	FetchTime      string `json:"fetch_time"`
	XiaoquId       string `json:"xiaoqu_id"`
	ChengjiaoNum   string `json:"chengjiao_num"`
	ChuzuNum       string `json:"chuzu_num"`
	TotalPrice     string `json:"totalPrice"`
	TotalSellCount string `json:"totalSellCount"`
}

func (xiaoqulog XiaoquLog) TableName() string {
	return "lj_xiaoqu_log"
}

func (xiaoqulog XiaoquLog) List(db *gorm.DB, pageOffset, pageSize int) ([]*XiaoquLog, error) {
	var xiaoqulogs []*XiaoquLog
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if xiaoqulog.XiaoquId != "" {
		db = db.Where("xiaoqu_id = ?", xiaoqulog.XiaoquId)
	}
	if err = db.Order("fetch_time desc").Find(&xiaoqulogs).Error; err != nil {
		return nil, err
	}
	return xiaoqulogs, nil
}
