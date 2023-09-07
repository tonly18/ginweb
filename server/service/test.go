/* *
 * error code: 3000000000 ` 3000000199
 */

package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/samber/lo"
	"server/core/xerror"
	"server/service/model"
)

//TestService Struct
type TestService struct {
	ctx context.Context
}

func NewTestService(ctx context.Context) *TestService {
	return &TestService{
		ctx: ctx,
	}
}

func (s *TestService) Query(serverId, uid int) ([]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Query(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid desc")
	if err != nil {
		if errors.Is(err.GetErr(), sql.ErrNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:    300000000,
				Err:     err.GetErr(),
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    300000009,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) QueryMap(serverId, uid int) (map[int]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.QueryMap(serverId, uid, []string{"uid", "item", "expire", "itime"})
	if err != nil {
		if errors.Is(err.GetErr(), sql.ErrNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:    300000010,
				Err:     err.GetErr(),
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    300000011,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) Insert(serverId, uid int, params map[string]any) (int, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Insert(serverId, uid, params)
	if err != nil {
		if errors.Is(err.GetErr(), sql.ErrNoRows) {
			return 0, xerror.Wrap(err, &xerror.NewError{
				Code:    300000030,
				Err:     err.GetErr(),
				Message: "bag.Insert(service)",
			})
		}
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:    300000039,
			Err:     err.GetErr(),
			Message: "bag.Insert(service)",
		})
	}

	return data, nil
}

func (s *TestService) Modify(serverId, uid int, params map[string]any) (int, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Modify(serverId, uid, params)
	if err != nil {
		if errors.Is(err.GetErr(), sql.ErrNoRows) {
			return 0, xerror.Wrap(err, &xerror.NewError{
				Code:    300000040,
				Err:     err.GetErr(),
				Message: "bag.Modify(service)",
			})
		}
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:    300000049,
			Err:     err.GetErr(),
			Message: "bag.Modify(service)",
		})
	}

	return data, nil
}

func (s *TestService) FilterSlice(serverId, uid int) ([]int, xerror.Error) {
	stu := []int{1, 2, 3, 4, 5, 6}
	stu = lo.Filter(stu, func(item int, _ int) bool {
		return item%2 == 0
	})

	return stu, nil
}
