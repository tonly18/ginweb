/* *
 * error code: 2000000000 ` 2000000199
 */

package model

import (
	"context"
	"fmt"
	"server/core/xerror"
	"server/service/model/dao"
)

//BagMode Struct
type BagMode struct {
	ctx context.Context
	dao *dao.BagDao
}

func NewBagMode(ctx context.Context) *BagMode {
	return &BagMode{
		ctx: ctx,
		dao: dao.NewBagDao(ctx),
	}
}

func (m *BagMode) Query(serverId, uid int, fields []string, order ...string) ([]map[string]any, xerror.Error) {
	data, err := m.dao.Query(serverId, uid, fields, "uid = 444", order...)
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:    200000000,
				Err:     ErrorNoRows,
				Message: "bag.query",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    200000009,
			Err:     err.GetErr(),
			Message: "bag.query",
		})
	}

	return data, nil
}

func (m *BagMode) QueryMap(serverId, uid int, fields []string) (map[int]map[string]any, xerror.Error) {
	data, err := m.dao.QueryMap(serverId, uid, fields, "uid < 40")
	if err != nil {
		if err.Is(dao.ErrorNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:    200000030,
				Err:     ErrorNoRows,
				Message: "bag.query map",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    200000040,
			Err:     err.GetErr(),
			Message: "bag.query map",
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
