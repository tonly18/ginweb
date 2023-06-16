/* *
 * bag
 * error code: 20005000 ` 20005999
 */

package model

import (
	"fmt"
	"server/core/merror"
	"server/core/request"
	"server/model/dao"
)

//Bag Model
type Bag struct {
	req *request.Request
	dao *dao.Bag
}

func NewBag(req *request.Request) *Bag {
	return &Bag{
		req: req,
		dao: dao.NewBag(req),
	}
}

func (m *Bag) Query(serverId, uid int) ([]*dao.BagTable, merror.Error) {
	data, err := m.dao.Query(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid < 40", "uid DESC")
	if err == dao.ErrorNoRows {
		return nil, merror.NewError(m.req, &merror.ErrorTemp{
			Code: 20005000,
			Err:  fmt.Errorf(`[bag query uid: %v error: %w]`, uid, err.GetError()),
			Type: 1,
		})
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Bag) QueryMap(serverId, uid int) (map[int]*dao.BagTable, merror.Error) {
	data, err := m.dao.QueryMap(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid<40", "uid DESC")
	if err == dao.ErrorNoRows {
		return nil, merror.NewError(m.req, &merror.ErrorTemp{
			Code: 20005010,
			Err:  fmt.Errorf(`[bag query map uid: %v error: %w]`, uid, err.GetError()),
			Type: 1,
		})
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Bag) Add(serverId, uid int, params map[string]any) (int, merror.Error) {
	return m.dao.Insert(serverId, uid, params)
}

func (m *Bag) ModifyByUserId(serverId, uid int, params map[string]any) (int, merror.Error) {
	return m.dao.Modify(serverId, uid, fmt.Sprintf(`uid=%v`, uid), params)
}

func (m *Bag) Delete(serverId, uid int, where string) (int, merror.Error) {
	return m.dao.Delete(serverId, uid, where)
}
