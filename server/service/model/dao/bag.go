/* *
 * error code: 100000000 ` 100000199
 */

package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonly18/xerror"
	"github.com/tonly18/xsql"
)

// BagDao struct
type BagDao struct {
	ctx     context.Context
	db      *xsql.XSQL
	redis   *RedisPoolConn
	tbl     string   //表名
	fields  []string //表字段
	primary string   //表主键
}

func NewBagDao(ctx context.Context) *BagDao {
	return &BagDao{
		ctx:     ctx,
		db:      xsql.NewXSQL(ctx, dbConfig),
		redis:   NewRedis(ctx, rdConfig),
		tbl:     "bag",
		fields:  []string{"uid", "item", "expire", "itime"},
		primary: "uid",
	}
}

func (d *BagDao) Query(uid int, fields []string, where string, order ...string) ([]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = d.fields
	}
	d.db.Table(getTableName(uid, d.tbl)).Primary(d.primary).Fields(fields...).Where(where)
	if len(order) > 0 {
		d.db.OrderBy(order[0])
	}
	data, err := d.db.Query()
	if err != nil {
		return nil, xerror.Wrap(&xerror.NewError{
			Code:     100000000,
			RawError: err,
			Message:  "bag.Query(dao)",
		}, nil)
	}

	return data, nil
}

func (d *BagDao) QueryMap(uid int, fields []string, where string) (map[int]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = d.fields
	}
	d.db.Table(getTableName(uid, d.tbl)).Primary(d.primary).Fields(fields...).Where(where)
	data, err := d.db.QueryMap()
	if err != nil {
		return nil, xerror.Wrap(&xerror.NewError{
			Code:     100000020,
			RawError: err,
			Message:  fmt.Sprintf(`bag.QueryMap uid:%v`, uid),
		}, nil)
	}

	return data, nil
}

func (d *BagDao) Insert(uid int, params map[string]any) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Insert(params).Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000030,
			RawError: fmt.Errorf(`uid:%v, params:%v`, uid, params),
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000033,
			RawError: err,
			Message:  fmt.Sprintf(`uid:%v, params:%v`, uid, params),
		}, nil)
	}
	if count > 0 {
		newId, err := result.LastInsertId()
		if err != nil {
			return 0, xerror.Wrap(&xerror.NewError{
				Code:     100000034,
				RawError: fmt.Errorf(` uid:%v, params:%v`, uid, params),
			}, nil)
		}
		return int(newId), nil
	}

	return 0, xerror.Wrap(&xerror.NewError{
		Code:     100000035,
		RawError: errors.New("insert error"),
	}, nil)
}

func (d *BagDao) Modify(uid int, where string, params map[string]any) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Modify(params).Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000040,
			RawError: err,
			Message:  "bag.Modify",
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000042,
			RawError: err,
			Message:  fmt.Sprintf(`uid:%v, params:%v`, uid, params),
		}, nil)
	}

	return int(count), nil
}

func (d *BagDao) Delete(uid int, where string) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Delete().Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000045,
			RawError: err,
			Message:  "bag.Delete",
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:     100000042,
			RawError: err,
			Message:  fmt.Sprintf(` uid:%v`, uid),
		}, nil)
	}

	return int(count), nil
}
