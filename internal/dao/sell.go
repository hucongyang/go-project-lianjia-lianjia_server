package dao

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 数据访问对象的封装

func (d *Dao) CountSell(xiaoquId string) (int, error) {
	Sell := model.Sell{
		XiaoquId: xiaoquId,
	}
	return Sell.Count(d.engine)
}

func (d *Dao) GetSellList(xiaoquId string, page, pageSize int) ([]*model.Sell, error) {
	Sell := model.Sell{XiaoquId: xiaoquId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return Sell.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetSellDetail(fangziId string) (model.Sell, error) {
	Sell := model.Sell{FangziId: fangziId}
	return Sell.Detail(d.engine)
}
