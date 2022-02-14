package dao

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
)

// 数据访问对象的封装

func (d *Dao) SearchXiaoquList(keyword string) ([]*model.Xiaoqu, error) {
	xiaoqu := model.Xiaoqu{Title: keyword}
	return xiaoqu.Search(d.engine)
}

func (d *Dao) CountXiaoqu(districtId uint32) (int, error) {
	xiaoqu := model.Xiaoqu{
		DistrictId: districtId,
	}
	return xiaoqu.Count(d.engine)
}

func (d *Dao) GetXiaoquList(districtId uint32, page, pageSize int) ([]*model.Xiaoqu, error) {
	xiaoqu := model.Xiaoqu{DistrictId: districtId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return xiaoqu.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetXiaoquDetail(xiaoquId string) (model.Xiaoqu, error) {
	xiaoqu := model.Xiaoqu{XiaoquId: xiaoquId}
	return xiaoqu.Detail(d.engine)
}

func (d *Dao) GetXiaoquHistoryList(xiaoquId string, page, pageSize int) ([]*model.XiaoquLog, error) {
	xiaoqulog := model.XiaoquLog{XiaoquId: xiaoquId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return xiaoqulog.List(d.engine, pageOffset, pageSize)
}
