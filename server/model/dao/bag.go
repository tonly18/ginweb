/* *
 * error code: 30001000 ` 30001999
 */

package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/cast"
	"server/core/merror"
)

//BagTable Struct
type BagTable struct {
	Uid    int    `json:"uid"`
	Item   string `json:"item"`
	Expire int    `json:"expire"`
	Itime  int    `json:"itime"`
}

//Bag struct
type Bag struct {
	ctx      context.Context
	tbl      string //表名
	db       *DBBase
	redisKey string //redis key
	redis    *RedisPoolConn
}

func NewBag(ctx context.Context) *Bag {
	return &Bag{
		ctx:   ctx,
		tbl:   "bag",
		db:    NewDBBase(ctx),
		redis: NewRedis(ctx),
	}
}

func (d *Bag) GenBagTable(fields []string, values []any) *BagTable {
	entity := BagTable{}
	for k, v := range fields {
		if v == "uid" {
			entity.Uid = cast.ToInt(b2String(*(values[k]).(*sql.RawBytes)))
		}
		if v == "item" {
			entity.Item = b2String(*(values[k]).(*sql.RawBytes))
		}
		if v == "expire" {
			entity.Expire = cast.ToInt(b2String(*(values[k]).(*sql.RawBytes)))
		}
		if v == "itime" {
			entity.Itime = cast.ToInt(b2String(*(values[k]).(*sql.RawBytes)))
		}
	}

	//return
	return &entity
}

func (d *Bag) Query(serverId, uid int, fields []string, where string, order ...string) ([]*BagTable, merror.Error) {
	d.db.Table(getTableName(uid, d.tbl)).Fields(fields...).Where(where)
	if len(order) > 0 {
		d.db.OrderBy(order[0])
	}
	rows, err := d.db.Query()
	if err != nil {
		return nil, merror.NewError(d.ctx, &merror.ErrorTemp{
			Code: 30001000,
			Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, err),
			Type: 1,
		})
	}
	defer rows.Close()

	//字段 - 数据
	data := make([]*BagTable, 0, 10)
	entity := make([]any, 0, len(fields))
	for i := 0; i < len(fields); i++ {
		entity = append(entity, new(sql.RawBytes))
	}
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, merror.NewError(d.ctx, &merror.ErrorTemp{
				Code: 30001010,
				Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, err),
				Type: 1,
			})
		}
		data = append(data, d.GenBagTable(fields, entity))
	}

	return data, nil
}

func (d *Bag) QueryMap(serverId, uid int, fields []string, where string, order ...string) (map[int]*BagTable, merror.Error) {
	d.db.Table(getTableName(uid, d.tbl)).Fields(fields...).Where(where)
	if len(order) > 0 {
		d.db.OrderBy(order[0])
	}
	rows, err := d.db.Query()
	if err != nil {
		return nil, merror.NewError(d.ctx, &merror.ErrorTemp{
			Code: 30001020,
			Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, err),
			Type: 1,
		})
	}
	defer rows.Close()

	//字段 - 数据
	data := make(map[int]*BagTable, 50)
	entity := make([]any, 0, len(fields))
	for i := 0; i < len(fields); i++ {
		entity = append(entity, new(sql.RawBytes))
	}
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, merror.NewError(d.ctx, &merror.ErrorTemp{
				Code: 30001025,
				Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, err),
				Type: 1,
			})
		}
		record := d.GenBagTable(fields, entity)
		data[record.Uid] = record
	}

	return data, nil
}

func (d *Bag) Insert(serverId, uid int, params map[string]any) (int, merror.Error) {
	id, err := d.db.Table(getTableName(uid, d.tbl)).Insert(params).Exec()
	if err != nil {
		return 0, merror.NewError(d.ctx, &merror.ErrorTemp{
			Code: 30001030,
			Err:  fmt.Errorf(`[bag insert uid: %v error: %w]`, uid, err),
			Type: 1,
		})
	}

	return id, nil
}

func (d *Bag) Modify(serverId, uid int, where string, params map[string]any) (int, merror.Error) {
	count, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Modify(params).Exec()
	if err != nil {
		return 0, merror.NewError(d.ctx, &merror.ErrorTemp{
			Code: 30001040,
			Err:  fmt.Errorf(`[bag modify uid: %v error: %w]`, uid, err),
			Type: 1,
		})
	}

	return count, nil
}

func (d *Bag) Delete(serverId, uid int, where string) (int, merror.Error) {
	count, err := d.db.Table(getTableName(uid, d.tbl)).Where(where).Delete().Exec()
	if err != nil {
		return 0, merror.NewError(d.ctx, &merror.ErrorTemp{
			Code: 30001040,
			Err:  fmt.Errorf(`[bag delete uid: %v error: %w]`, uid, err),
			Type: 1,
		})
	}

	return count, nil
}
