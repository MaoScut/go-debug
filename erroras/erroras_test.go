package erroras

import (
	"errors"
	"testing"
)

type Err1 struct {
	myCode string
	myMsg  string
}

type NotFoundErr1 struct {
	msg string
}

func (n *NotFoundErr1) Error() string {
	return n.msg
}

type NotFoundErr2 struct {
	msg string
}

func (n *NotFoundErr2) Error() string {
	return n.msg
}

func (e *Err1) Error() string {
	return e.myMsg
}

func TestErrorAsSameStruct(t *testing.T) {
	var targetErr *NotFoundErr1
	notFoundErr2 := &NotFoundErr2{
		msg: "I'm not found err 2",
	}
	ok := errors.As(notFoundErr2, &targetErr)
	if ok {
		t.Fatal("same data struct but different type should not be ok")
	}
	notFoundErr1 := &NotFoundErr1{
		msg: "I'm notFoundErr1",
	}
	ok = errors.As(notFoundErr1, &targetErr)
	if ok {
		t.Log(targetErr.msg)
	}
	if !ok {
		t.Fatal("same type should be ok")
	}
}
