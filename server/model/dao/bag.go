/* *
 * bag
 * error code: 30005000 ` 30005999
 */

package dao

import (
	"fmt"
	"server/config"
	"server/core/merror"
	"server/core/request"
	"server/global"
	"strings"
	"time"
)

//bag table struct
type BagTable struct {
	Uid    int    `json:"uid"`
	Item   string `json:"item"`
	Expire string `json:"expire"`
	Itime  int    `json:"itime"`
}

//Bag struct
type Bag struct {
	req      *request.Request
	tbl      string   //表名
	fields   []string //字段
	redisKey string   //redis key

	db    *DBBase
	redis *RedisPoolConn
}

func NewBag(req *request.Request) *Bag {
	return &Bag{
		req:    req,
		tbl:    "bag",
		fields: []string{"uid", "itedm", "expire", "itime"},
		db:     NewDBBase(req),
		redis:  NewRedis(req),
	}
}

func (d *Bag) Get(uid int) (map[int]*BagTable, merror.Error) {
	bagdata := make(map[int]*BagTable, global.MapInitCount)

	//redis
	//redisBag, err := d.redis.HGetRd(config.REDIS_KEY_BAG, strconv.Itoa(uid))
	//if err == nil {
	//	if err := json.Unmarshal([]byte(redisBag), bagData); err != nil {
	//		//logger.Error(d.req.GinCtx, "[bag get] redis json.unmarshal, table:%v, uid: %v, error: %v", d.tbl, uid, err.Error())
	//		return nil, err
	//	}
	//	return bagData, nil
	//}

	//db
	data, err := d.Query(uid, fmt.Sprintf(`uid = %v`, uid))
	if err != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code:  30000000,
			Error: fmt.Errorf(`[bag get uid: %v error:%v]`, uid, err.GetError()),
			Type:  1,
		})
	}
	if len(data) == 0 {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code:  30000002,
			Error: fmt.Errorf(`[bag get uid: %v error:%v]`, uid, ErrorNoRows),
			Type:  1,
		})
	}
	bagdata[uid] = data[uid]

	//if err := json.Unmarshal([]byte(cast.ToString(dbBag)), bagData); err != nil {
	//	logger.Error(d.req.GinCtx, fmt.Sprintf(`[bag get] db.get json.unmarshal, table:%v, uid: %v, error: %v`, d.tbl, uid, err))
	//	return nil, err
	//}
	////redis
	//if err := d.redis.HSetRd(config.REDIS_KEY_BAG, uid, dbBag); err != nil {
	//	logger.Error(d.req.GinCtx, fmt.Sprintf(`[bag get] redis.HSetRd, table:%v, uid: %v, error: %v`, d.tbl, uid, err))
	//}

	return bagdata, nil
}

func (d *Bag) Query(id int, where string) (map[int]*BagTable, merror.Error) {
	bagList := make(map[int]*BagTable, global.MapInitCount)

	rawsql := fmt.Sprintf(`SELECT %v FROM %v WHERE %v`, strings.Join(d.fields, ","), getTableName(id, d.tbl), where)
	rows, err := d.db.Query(rawsql)
	if err != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code:  30000020,
			Error: fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
			Type:  1,
		})
	}
	defer rows.Close()

	for rows.Next() {
		bag := &BagTable{}
		if err := rows.Scan(&bag.Uid, &bag.Item, &bag.Expire, &bag.Itime); err != nil {
			return nil, merror.NewError(d.req, &merror.ErrorTemp{
				Code:  30000021,
				Error: fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
				Type:  1,
			})
		}
		bagList[bag.Uid] = bag
	}
	if rows.Err() != nil {
		return nil, merror.NewError(d.req, &merror.ErrorTemp{
			Code:  30000029,
			Error: fmt.Errorf(`[query sql: %v, error:%v]`, rawsql, err),
			Type:  1,
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
				Code:  30000010,
				Error: fmt.Errorf(`[bag modify redis.HSetRd, key:%v, uid: %v, error:%v]`, config.REDIS_KEY_BAG, uid, err),
				Type:  1,
			})
		}
		return nil
	}

	return merror.NewError(d.req, &merror.ErrorTemp{
		Code:  30000012,
		Error: fmt.Errorf(`[bag modify table:%v, uid: %v, error:%v]`, d.tbl, uid, err),
		Type:  1,
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
			Code:  30000015,
			Error: fmt.Errorf(`[bag init sql:%v, uid:%v, error:%v]`, rawsql, id, err),
			Type:  1,
		})
	}

	return bag, nil
}
