/* *
 * error code: 3000000000 ` 3000000199
 */

package service

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"server/core/xerror"
	"server/service/model"
	"server/service/model/dao"
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

func (s *TestService) Query(serverId, uid int) ([]*model.BagTable, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Query(serverId, uid, []string{"uid", "item", "expire", "itime"}, "uid desc")
	if err != nil {
		if errors.Is(err.GetErr(), dao.ErrorNoRows) {
			return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
				Code:    300000000,
				Err:     model.ErrorNoRows,
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
			Code:    300000009,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) QueryMap(serverId, uid int) (map[int]*model.BagTable, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.QueryMap(serverId, uid, []string{"uid", "item"})
	if err != nil {
		if errors.Is(err.GetErr(), dao.ErrorNoRows) {
			return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
				Code:    300000010,
				Err:     model.ErrorNoRows,
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
			Code:    300000011,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) Get(serverId, uid int) ([]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Get(serverId, uid, []string{"uid", "item", "expire", "abc"}, "uid desc")
	if err != nil {
		if errors.Is(err.GetErr(), dao.ErrorNoRows) {
			return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
				Code:    300000020,
				Err:     model.ErrorNoRows,
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
			Code:    300000021,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) GetMap(serverId, uid int) (map[int]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.GetMap(serverId, uid, []string{"uid", "item", "expire"})
	if err != nil {
		if errors.Is(err.GetErr(), dao.ErrorNoRows) {
			return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
				Code:    300000030,
				Err:     model.ErrorNoRows,
				Message: "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(s.ctx, err, &xerror.NewError{
			Code:    300000031,
			Err:     err.GetErr(),
			Message: "bag.Query(service)",
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
