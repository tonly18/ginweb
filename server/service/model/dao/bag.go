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
	"server/library/command"
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

func (d *BagDao) Query(serverId, uid int, fields []string, where string, order ...string) ([]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = d.fields
	}
	if !command.SliceContains(fields, d.primary) {
		fields = append(fields, d.primary)
	}

	d.db.Table(getTableName(uid, d.tbl)).Fields(fields...).Where(where)
	if len(order) > 0 {
		d.db.OrderBy(order[0])
	}
	data, err := d.db.Query()
	if err != nil {
		return nil, xerror.Wrap(&xerror.NewError{
			Code:    100000000,
			Err:     err,
			Message: "bag.Query(dao)",
		}, nil)
	}

	return data, nil
}

func (d *BagDao) QueryMap(serverId, uid int, fields []string, where string) (map[int]map[string]any, xerror.Error) {
	if len(fields) == 0 {
		fields = d.fields
	}
	if !command.SliceContains(fields, d.primary) {
		fields = append(fields, d.primary)
	}

	d.db.Table(getTableName(uid, d.tbl)).Fields(fields...).Where(where)
	data, err := d.db.QueryMap()
	if err != nil {
		return nil, xerror.Wrap(&xerror.NewError{
			Code:    100000020,
			Err:     err,
			Message: fmt.Sprintf(`bag.QueryMap uid:%v`, uid),
		}, nil)
	}

	return data, nil
}

func (d *BagDao) Insert(serverId, uid int, params map[string]any) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Insert(params).Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code: 100000030,
			Err:  fmt.Errorf(`serverId:%v, uid:%v, params:%v`, serverId, uid, params),
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:    100000033,
			Err:     err,
			Message: fmt.Sprintf(`serverId:%v, uid:%v, params:%v`, serverId, uid, params),
		}, nil)
	}
	if count > 0 {
		newId, err := result.LastInsertId()
		if err != nil {
			return 0, xerror.Wrap(&xerror.NewError{
				Code: 100000034,
				Err:  fmt.Errorf(`serverId:%v, uid:%v, params:%v`, serverId, uid, params),
			}, nil)
		}
		return int(newId), nil
	}

	return 0, xerror.Wrap(&xerror.NewError{
		Code: 100000035,
		Err:  errors.New("insert error"),
	}, nil)
}

func (d *BagDao) Modify(serverId, uid int, where string, params map[string]any) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Modify(params).Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:    100000040,
			Err:     err,
			Message: "bag.Modify",
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:    100000042,
			Err:     err,
			Message: fmt.Sprintf(`serverId:%v, uid:%v, params:%v`, serverId, uid, params),
		}, nil)
	}

	return int(count), nil
}

func (d *BagDao) Delete(serverId, uid int, where string) (int, xerror.Error) {
	result, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Delete().Exec()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:    100000045,
			Err:     err,
			Message: "bag.Delete",
		}, nil)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, xerror.Wrap(&xerror.NewError{
			Code:    100000042,
			Err:     err,
			Message: fmt.Sprintf(`serverId:%v, uid:%v`, serverId, uid),
		}, nil)
	}

	return int(count), nil
}
