/* *
 * error code: 2000000000 ` 2000000199
 */

package model

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"server/core/xerror"
	"server/service/model/dao"
)

//BagTable Struct
type BagTable struct {
	Uid    int    `json:"uid"`
	Item   string `json:"item"`
	Expire int    `json:"expire"`
	Itime  int    `json:"itime"`
}

//BagMode Struct
type BagMode struct {
	ctx    context.Context
	dao    *dao.BagDao
	fields []string
}

func NewBagMode(ctx context.Context) *BagMode {
	return &BagMode{
		ctx:    ctx,
		dao:    dao.NewBagDao(ctx),
		fields: []string{"uid", "item", "expire", "itime"},
	}
}

func (m *BagMode) genTable(data map[string]any) *BagTable {
	entity := BagTable{}
	for k, v := range data {
		if k == "uid" {
			entity.Uid = cast.ToInt(v)
		} else if k == "item" {
			entity.Item = cast.ToString(v)
		} else if k == "expire" {
			entity.Expire = cast.ToInt(v)
		} else if k == "itime" {
			entity.Itime = cast.ToInt(v)
		}
	}

	//return
	return &entity
}

func (m *BagMode) Query(serverId, uid int, fields []string, order ...string) ([]*BagTable, xerror.Error) {
	if len(fields) == 0 {
		fields = m.fields
	}
	data, err := m.dao.Query(serverId, uid, fields, "uid < 444", order...)
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
				Code:    200000000,
				Err:     ErrorNoRows,
				Message: "bag.query",
			})
		}
		return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
			Code:    200000009,
			Err:     err.GetErr(),
			Message: "bag.query",
		})
	}

	dataList := make([]*BagTable, 0, len(data))
	for _, v := range data {
		dataList = append(dataList, m.genTable(v))
	}

	return dataList, nil
}

func (m *BagMode) QueryMap(serverId, uid int, fields []string) (map[int]*BagTable, xerror.Error) {
	if len(fields) == 0 {
		fields = m.fields
	}
	data, err := m.dao.QueryMap(serverId, uid, fields, "uid<40")
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
				Code:    200000030,
				Err:     ErrorNoRows,
				Message: "bag.query map",
			})
		}
		return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
			Code:    200000040,
			Err:     err.GetErr(),
			Message: "bag.query map",
		})
	}

	dataList := make(map[int]*BagTable, len(data))
	for _, v := range data {
		dataList[cast.ToInt(v["uid"])] = m.genTable(v)
	}

	return dataList, nil
}

func (m *BagMode) Get(serverId, uid int, fields []string, order ...string) ([]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = m.fields
	}
	data, err := m.dao.Query(serverId, uid, fields, "uid < 444", order...)
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
				Code:    200000030,
				Err:     ErrorNoRows,
				Message: "bag.get",
			})
		}
		return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
			Code:    200000039,
			Err:     err.GetErr(),
			Message: "bag.get",
		})
	}

	return data, nil
}

func (m *BagMode) GetMap(serverId, uid int, fields []string) (map[int]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = m.fields
	}
	data, err := m.dao.QueryMap(serverId, uid, fields, "uid < 444")
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
				Code:    200000040,
				Err:     ErrorNoRows,
				Message: "bag.get map",
			})
		}
		return nil, xerror.Wrap(m.ctx, err, &xerror.NewError{
			Code:    200000049,
			Err:     err.GetErr(),
			Message: "bag.get map",
		})
	}

	return data, nil
}

func (m *BagMode) Add(serverId, uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Insert(serverId, uid, params)
}

func (m *BagMode) ModifyByUserId(serverId, uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Modify(serverId, uid, fmt.Sprintf(`uid=%v`, uid), params)
}

func (m *BagMode) Delete(serverId, uid int, where string) (int, xerror.Error) {
	return m.dao.Delete(serverId, uid, where)
}
