package xerror

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkBXError(b *testing.B) {
	err := &NewError{
		Code:    500000000,
		Err:     errors.New("test handler bag.query"),
		Message: "test bag.query",
	}

	//run b.N times
	for n := 0; n < b.N; n++ {
		Wrap(err, &NewError{
			Code:    uint32(n * 100000000),
			Err:     fmt.Errorf(`test handler bag.query:%d`, n),
			Message: "test stack",
		})
	}

	fmt.Println("over::::::", len(err.GetStack()), err.GetCode(), err.GetMsg())
}
