package api

import (
	. "gopkg.in/check.v1"
	"testing"
)


func Test(t *testing.T) { TestingT(t) }


type CheckUserExist struct {
}

var _ = Suite(&CheckUserExist{})

var (
	urlCheckUserExist string = "http://localhost/api/v1/checkUserExist"
)

func (s *CheckUserExist) TestCheckUserExistCase00(goCheck *C) {
	//
	goCheck.Assert(nil, Equals, nil)
}

func (s *CheckUserExist) TestCheckUserExistCase01(goCheck *C) {
	goCheck.Assert(a, Equals, 1)
}
