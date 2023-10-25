/* *
 * error code: 3000000000 ` 3000000199
 */

package service

import (
	"context"
	"database/sql"
	"github.com/tonly18/xerror"
	"server/service/model"
)

// TestService Struct
type TestService struct {
	ctx context.Context
}

func NewTestService(ctx context.Context) *TestService {
	return &TestService{
		ctx: ctx,
	}
}

func (s *TestService) Query(uid int) ([]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Query(uid, []string{"uid", "item", "expire", "itime"}, "uid desc")
	if err != nil {
		if err.Is(sql.ErrNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:     300000000,
				RawError: err.GetRawError(),
				Message:  "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:     300000009,
			RawError: err.GetRawError(),
			Message:  "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) QueryMap(uid int) (map[int]map[string]any, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.QueryMap(uid, []string{"uid", "item", "expire", "itime"})
	if err != nil {
		if err.Is(sql.ErrNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:     300000010,
				RawError: err.GetRawError(),
				Message:  "bag.Query(service)",
			})
		}
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:     300000011,
			RawError: err.GetRawError(),
			Message:  "bag.Query(service)",
		})
	}

	return data, nil
}

func (s *TestService) Insert(uid int, params map[string]any) (int, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Insert(uid, params)
	if err != nil {
		if err.Is(sql.ErrNoRows) {
			return 0, xerror.Wrap(err, &xerror.NewError{
				Code:     300000030,
				RawError: err.GetRawError(),
				Message:  "bag.Insert(service)",
			})
		}
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:     300000039,
			RawError: err.GetRawError(),
			Message:  "bag.Insert(service)",
		})
	}

	return data, nil
}

func (s *TestService) Modify(uid int, params map[string]any) (int, xerror.Error) {
	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Modify(uid, params)
	if err != nil {
		if err.Is(sql.ErrNoRows) {
			return 0, xerror.Wrap(err, &xerror.NewError{
				Code:     300000040,
				RawError: err.GetRawError(),
				Message:  "bag.Modify(service)",
			})
		}
		return 0, xerror.Wrap(err, &xerror.NewError{
			Code:     300000049,
			RawError: err.GetRawError(),
			Message:  "bag.Modify(service)",
		})
	}

	return data, nil
}
