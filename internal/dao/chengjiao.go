package dao

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 数据访问对象的封装

func (d *Dao) CountChengjiao(xiaoquId string) (int, error) {
	Chengjiao := model.Chengjiao{
		XiaoquId: xiaoquId,
	}
	return Chengjiao.Count(d.engine)
}

func (d *Dao) GetChengjiaoList(xiaoquId string, page, pageSize int) ([]*model.Chengjiao, error) {
	Chengjiao := model.Chengjiao{XiaoquId: xiaoquId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return Chengjiao.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetChengjiaoDetail(fangziId string) (model.Chengjiao, error) {
	Chengjiao := model.Chengjiao{FangziId: fangziId}
	return Chengjiao.Detail(d.engine)
}
