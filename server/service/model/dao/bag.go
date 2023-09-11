/* *
 * error code: 100000000 ` 100000199
 */

package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/cast"
	"server/core/xerror"
	"server/global"
	"server/library/command"
)

//BagDao struct
type BagDao struct {
	ctx     context.Context
	db      *DBBase
	redis   *RedisPoolConn
	tbl     string   //表名
	fields  []string //表字段
	primary string   //表主键
}

func NewBagDao(ctx context.Context) *BagDao {
	return &BagDao{
		ctx:     ctx,
		db:      NewDBBase(ctx),
		redis:   NewRedis(ctx),
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

	rows, err := d.db.Query()
	if err != nil {
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    100000000,
			Err:     err.GetErr(),
			Message: "bag.Query(dao)",
		})
	}
	defer rows.Close()

	//字段 - 数据
	data := make([]map[string]any, 0, global.DefaultCount)
	entity := genEntity(len(fields))
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, xerror.Wrap(nil, &xerror.NewError{
				Code:    100000006,
				Err:     err,
				Message: fmt.Sprintf(`query uid:%v`, uid),
			})
		}
		data = append(data, genRecord(entity, fields))
	}
	if len(data) == 0 {
		return data, xerror.Wrap(nil, &xerror.NewError{
			Code: 100000009,
			Err:  sql.ErrNoRows,
		})
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
	rows, err := d.db.Query()
	if err != nil {
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    100000020,
			Message: fmt.Sprintf(`bag.QueryMap uid:%v`, uid),
		})
	}
	defer rows.Close()

	//字段 - 数据
	data := make(map[int]map[string]any, global.DefaultCount)
	entity := genEntity(len(fields))
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, xerror.Wrap(nil, &xerror.NewError{
				Code:    100000025,
				Err:     err,
				Message: fmt.Sprintf(`query map uid:%v`, uid),
			})
		}
		record := genRecord(entity, fields)
		data[cast.ToInt(record[d.primary])] = record
	}
	if len(data) == 0 {
		return data, xerror.Wrap(nil, &xerror.NewError{
			Code: 100000029,
			Err:  sql.ErrNoRows,
		})
	}

	return data, nil
}

func (d *BagDao) Insert(serverId, uid int, params map[string]any) (int, xerror.Error) {
	id, err := d.db.Table(getTableName(uid, d.tbl)).Insert(params).Exec(true)
	if err != nil {
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:    100000030,
			Err:     err.GetErr(),
			Message: fmt.Sprintf(`serverId:%v, uid:%v, params:%v`, serverId, uid, params),
		})
	}

	return id, nil
}

func (d *BagDao) Modify(serverId, uid int, where string, params map[string]any) (int, xerror.Error) {
	count, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Modify(params).Exec(false)
	if err != nil {
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:    100000040,
			Err:     err.GetErr(),
			Message: "bag.Modify",
		})
	}

	return count, nil
}

func (d *BagDao) Delete(serverId, uid int, where string) (int, xerror.Error) {
	count, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Delete().Exec(false)
	if err != nil {
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:    100000045,
			Err:     err.GetErr(),
			Message: "bag.Delete",
		})
	}

	return count, nil
}
