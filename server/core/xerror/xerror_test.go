package xerror

import (
	"database/sql"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"unsafe"
)

func TestXError(t *testing.T) {

	xErrorEntity := NewError{
		Code:       20005000,
		Err:        net.ErrClosed,
		Message:    "message",
		ErrorStack: nil,
	}
	size := unsafe.Sizeof(xErrorEntity.Code)
	sizeErr := unsafe.Sizeof(xErrorEntity.Err)
	sizeMsg := unsafe.Sizeof(xErrorEntity.Message)
	fmt.Println("size::::::", size, sizeErr, sizeMsg)
	fmt.Println("sizeTotal::::::", unsafe.Sizeof(xErrorEntity.ErrorStack))
	fmt.Println("sizeTotal::::::", unsafe.Sizeof(xErrorEntity))

	_, err := A(100)
	//fmt.Println("data::::::", data)
	//fmt.Println("err:::::::", err.GetCode(), err.GetErr(), err.GetMsg())
	contain := err.Contain(net.ErrClosed)
	fmt.Println("contain::::::", contain)

	if err.GetErr() == net.ErrClosed {
		errList := err.GetErrorStack()
		//fmt.Println("err::::::", len(err.GetErrorStack()))
		//fmt.Println("errorList::::::", errList)
		for _, e := range errList {
			fmt.Println("error::::::", e.GetCode(), e.GetMsg(), e.GetErr())
		}
	}

	fmt.Println("main")
}

func A(uid int) (int, Error) {
	data, err := B(uid)
	//fmt.Println("a-err::::::", err.GetCode(), err.GetErr(), err.GetMsg())
	if err.GetErr() == os.ErrClosed {
		//fmt.Println("a-err::::::", err.GetErr())
		xerr := Wrap(err, &NewError{
			Code:    20005000,
			Err:     net.ErrClosed,
			Message: "a-message",
		})
		//fmt.Println("a-err:::::::", xerr.GetCode(), xerr.GetErr(), xerr.GetMsg(), len(xerr.GetErrorStack()))
		return 0, xerr
	}

	return data, nil
}

func B(uid int) (int, Error) {
	_, err := C(uid)
	if err.GetErr() == sql.ErrNoRows {
		xerr := Wrap(err, &NewError{
			Code:    20005010,
			Err:     os.ErrClosed,
			Message: "b-message",
		})
		//fmt.Println("b-err:::::::", xerr.GetCode(), xerr.GetErr(), xerr.GetMsg(), len(xerr.GetErrorStack()))
		return 0, xerr
	}

	return 1, nil
}

func C(uid int) (int, Error) {
	_, err := D(uid)
	if err.GetErr() == io.ErrClosedPipe {
		xerr := Wrap(err, &NewError{
			Code:    20005020,
			Err:     sql.ErrNoRows,
			Message: "c-message",
		})
		//fmt.Println("c-err:::::::", xerr.GetCode(), xerr.GetErr(), xerr.GetMsg(), len(xerr.GetErrorStack()))
		return 0, xerr
	}

	return 1, nil
}

func D(uid int) (int, Error) {
	_, err := E(uid)
	//if err.GetErr() == os.ErrPermission {
	if err.Is(os.ErrPermission) {
		xerr := Wrap(err, &NewError{
			Code:    20005030,
			Err:     io.ErrClosedPipe,
			Message: "d-message",
		})
		//fmt.Println("d-err:::::::", xerr.GetCode(), xerr.GetErr(), xerr.GetMsg(), len(xerr.GetErrorStack()))
		return 0, xerr
	}

	return 1, nil
}

func E(uid int) (int, Error) {
	err := os.ErrPermission
	if err == os.ErrPermission {
		xerr := Wrap(nil, &NewError{
			Code:    20005040,
			Err:     err,
			Message: "e-message",
		})
		//fmt.Println("e-err:::::::", xerr.GetCode(), xerr.GetErr(), xerr.GetMsg(), len(xerr.GetErrorStack()))
		return 0, xerr
	}

	return 1, nil
}
