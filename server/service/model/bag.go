/* *
 * error code: 2000000000 ` 2000000199
 */

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tonly18/xerror"
	"server/service/model/dao"
)

// BagMode Struct
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

func (m *BagMode) Query(uid int, fields []string, order ...string) ([]map[string]any, xerror.Error) {
	where := fmt.Sprintf(`uid=%v`, uid)
	data, xerr := m.dao.Query(uid, fields, where, order...)
	if xerr != nil {
		if xerr.Contain(sql.ErrNoRows) {
			return nil, xerror.Wrap(xerr, &xerror.NewError{
				Code:     200000000,
				RawError: xerr.GetRawError(),
				Message:  "bag.query",
			})
		}
		return nil, xerror.Wrap(xerr, &xerror.NewError{
			Code:     200000009,
			RawError: xerr.GetRawError(),
			Message:  "bag.query",
		})
	}

	return data, nil
}

func (m *BagMode) QueryMap(uid int, fields []string) (map[int]map[string]any, xerror.Error) {
	data, xerr := m.dao.QueryMap(uid, fields, "uid < 40")
	if xerr != nil {
		if xerr.Contain(sql.ErrNoRows) {
			return nil, xerror.Wrap(xerr, &xerror.NewError{
				Code:     200000030,
				RawError: xerr.GetRawError(),
				Message:  "bag.query map",
			})
		}
		return nil, xerror.Wrap(xerr, &xerror.NewError{
			Code:     200000040,
			RawError: xerr.GetRawError(),
			Message:  "bag.query map",
		})
	}

	return data, nil
}

func (m *BagMode) Insert(uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Insert(uid, params)
}

func (m *BagMode) Modify(uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Modify(uid, fmt.Sprintf(`uid=%v`, uid), params)
}

func (m *BagMode) Delete(uid int, where string) (int, xerror.Error) {
	return m.dao.Delete(uid, where)
}
