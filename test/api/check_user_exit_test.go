package api

import (
	"fmt"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
	"testing"
)


func Test(t *testing.T) { TestingT(t) }


type CheckUserExist struct {
}

var _ = Suite(&CheckUserExist{})

var (
	urlCheckUserExist string = "http://localhost/api/v1/checkUserExist"
)

type CheckUserExistReq struct {
	Addr string `json:"addr"`
}

func (s *CheckUserExist) TestCheckUserExistCase00(goCheck *C) {
	//
	privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	sign := common.GetSign(privateHex)
	fmt.Println("sign: ", sign)
	req := CheckUserExistReq {
		Addr: "test_not_in_db",
	}
	resp, err := common.DoPost(urlCheckUserExist, common.ConvToJSON(req))
	goCheck.Assert(err, IsNil)
	fmt.Println(string(resp))
}

