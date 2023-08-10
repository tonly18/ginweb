/* *
 * error code: 3000000000 ` 3000000199
 */

package service

import (
	"context"
	"fmt"
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

func (s *TestService) GetDataList(serverId, uid int) ([]int, xerror.Error) {
	//data, err := s.model.Query(serverId, uid, nil, "uid < 44", "uid DESC")
	//if err != nil {
	//	if err.Is(dao.ErrorNoRows) {
	//		return nil, xerror.Wrap(m.req, err, &xerror.TempError{
	//			Code:    20005000,
	//			Err:     ErrorNoRows,
	//			Message: "bag.Query(model)",
	//		})
	//	}
	//	return nil, xerror.Wrap(m.req, err, &xerror.TempError{
	//		Code:    20005009,
	//		Err:     err.GetErr(),
	//		Message: "bag.Query(model)",
	//	})
	//}

	bagMode := model.NewBagMode(s.ctx)
	data, err := bagMode.Query(serverId, uid)
	if err != nil {
		return nil, xerror.Wrap(s.ctx, err, &xerror.TempError{
			Code:    300000000,
			Err:     model.ErrorNoRows,
			Message: "bag.Query(service)",
		})
	}
	fmt.Println("data:::::", data)

	stu := []int{1, 2, 3, 4, 5, 6}
	stu = lo.Filter(stu, func(item int, _ int) bool {
		return item%2 == 0
	})

	return stu, nil
}
