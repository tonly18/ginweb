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

/* --------------------------------------------------------------------------------- */

/*
func (d *Bag) Get(uid int) (map[int]*BagTable, merror.Error) {
	bagdata := make(map[int]*BagTable, global.MapInitCount)

	//redis
	//redisBag, err := d.redis.HGetRd(config.REDIS_KEY_BAG, strconv.Itoa(uid))
	//if err == nil {
	//	if err := json.Unmarshal([]byte(redisBag), bagData); err != nil {
	//		//logger.Err(d.req.GinCtx, "[bag get] redis json.unmarshal, table:%v, uid: %v, error: %v", d.tbl, uid, err.Err())
	//		return nil, err
	//	}
	//	return bagData, nil
	//}

	//db
	data, err := d.Query(uid, fmt.Sprintf(`uid = %v`, uid))
	if err != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code: 30000000,
			Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, err.GetError()),
			Type: 1,
		})
	}
	if len(data) == 0 {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code: 30000002,
			Err:  fmt.Errorf(`[bag get uid: %v error: %w]`, uid, ErrorNoRows),
			Type: 1,
		})
	}
	bagdata[uid] = data[uid]

	//if err := json.Unmarshal([]byte(cast.ToString(dbBag)), bagData); err != nil {
	//	logger.Err(d.req.GinCtx, fmt.Sprintf(`[bag get] db.get json.unmarshal, table:%v, uid: %v, error: %v`, d.tbl, uid, err))
	//	return nil, err
	//}
	////redis
	//if err := d.redis.HSetRd(config.REDIS_KEY_BAG, uid, dbBag); err != nil {
	//	logger.Err(d.req.GinCtx, fmt.Sprintf(`[bag get] redis.HSetRd, table:%v, uid: %v, error: %v`, d.tbl, uid, err))
	//}

	return bagdata, nil
}

func (d *Bag) Query(id int, where string) (map[int]*BagTable, merror.Error) {
	bagList := make(map[int]*BagTable, global.MapInitCount)

	rawsql := fmt.Sprintf(`SELECT %v FROM %v WHERE %v`, strings.Join(d.fields, ","), getTableName(id, d.tbl), where)
	rows, err := d.db.Query(rawsql)
	if err != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code: 30000020,
			Err:  fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
			Type: 1,
		})
	}
	defer rows.Close()

	for rows.Next() {
		bag := &BagTable{}
		if err := rows.Scan(&bag.Uid, &bag.Item, &bag.Expire, &bag.Itime); err != nil {
			return nil, merror.NewError(d.req, &merror.ErrorTemp{
				Code: 30000021,
				Err:  fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
				Type: 1,
			})
		}
		bagList[bag.Uid] = bag
	}
	if rows.Err() != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code: 30000029,
			Err:  fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
			Type: 1,
		})
	}

	return bagList, nil
}

func (d *Bag) Modify(uid int, data []byte) merror.Error {
	err := d.db.Modify(d.tbl, uid, data)
	if err == nil {
		if err := d.redis.HSetRd(config.REDIS_KEY_BAG, uid, data); err != nil {
			//d.redis.HDelRd(config.REDIS_KEY_BAG, strconv.Itoa(uid))
			return merror.NewError(d.req, &merror.ErrorTemp{
				Code: 30000010,
				Err:  fmt.Errorf(`[bag modify redis.HSetRd, key:%v, uid: %v, error:%v]`, config.REDIS_KEY_BAG, uid, err),
				Type: 1,
			})
		}
		return nil
	}

	return merror.NewError(d.req, &merror.ErrorTemp{
		Code: 30000012,
		Err:  fmt.Errorf(`[bag modify table:%v, uid: %v, error:%v]`, d.tbl, uid, err),
		Type: 1,
	})
}

//初始化数量
func (d *Bag) initData(id int) (*BagTable, merror.Error) {
	bag := &BagTable{
		Uid: id,
	}
	rawsql := fmt.Sprintf(`INSERT INTO %v(uid, gold, diamond, cash, life, stime) VALUES(?, ?, ?, ?, ?, ?)`, getTableName(id, d.tbl))
	if err := d.db.Exec(rawsql, id, 100, 100, 100, 100, time.Now().Unix()); err != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code: 30000015,
			Err:  fmt.Errorf(`[bag init sql:%v, uid:%v, error:%v]`, rawsql, id, err),
			Type: 1,
		})
	}

	return bag, nil
}
*/
