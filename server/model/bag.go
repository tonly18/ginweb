/* *
 * bag
 * error code: 20005000 ` 20005999
 */

package model

import (
	"fmt"
	"server/core/request"
	"server/core/xerror"
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

func (m *Bag) Query(serverId, uid int) ([]*dao.BagTable, xerror.Error) {
	data, err := m.dao.Query(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid = 44", "uid DESC")
	if err.Is(dao.ErrorNoRows) {
		return nil, xerror.Wrap(m.req, err, &xerror.TempError{
			Code:    20005000,
			Err:     ErrorNoRows,
			Message: "bag.Query(model)",
		})
	}
	if err != nil {
		return nil, xerror.Wrap(m.req, nil, &xerror.TempError{
			Code:    20005009,
			Err:     err.GetErr(),
			Message: "bag.Query(model)",
		})
	}

	return data, nil
}

func (m *Bag) QueryMap(serverId, uid int) (map[int]*dao.BagTable, xerror.Error) {
	data, err := m.dao.QueryMap(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid<40", "uid DESC")
	if err.Is(dao.ErrorNoRows) {
		return nil, xerror.Wrap(m.req, err, &xerror.TempError{
			Code:    20005010,
			Err:     err.GetErr(),
			Message: "bag.QueryMap",
			Type:    1,
		})
	}
	if err != nil {
		return nil, xerror.Wrap(m.req, err, &xerror.TempError{
			Code:    20005019,
			Err:     err.GetErr(),
			Message: "bag.QueryMap",
			Type:    1,
		})
	}

	return data, nil
}

func (m *Bag) Add(serverId, uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Insert(serverId, uid, params)
}

func (m *Bag) ModifyByUserId(serverId, uid int, params map[string]any) (int, xerror.Error) {
	return m.dao.Modify(serverId, uid, fmt.Sprintf(`uid=%v`, uid), params)
}

func (m *Bag) Delete(serverId, uid int, where string) (int, xerror.Error) {
	return m.dao.Delete(serverId, uid, where)
}
