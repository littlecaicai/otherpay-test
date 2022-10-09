package api

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"otherpay-test/common"
)




type Register struct {
}

var _ = Suite(&Register{})

var (
	urlRegister string = "http://localhost/api/v1/register"
)

type RegisterReq struct {
	Addr string `json:"addr"`
	TimeStamp int64 `json:"timestamp"`
	Sign string `json:"sign"`

}

type RegisterResp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data RegisterRespData `json:"data"`
}

type RegisterRespData struct {
	Token string `json:"token"`
}

func (s *Register) TestRegisterCase00(goCheck *C) {
	//参数合法，可以注册成功
	privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	addr, sign, st:= common.GetSign(privateHex)
	req := RegisterReq {
		Addr: addr,
		TimeStamp: st,
		Sign: sign,
	}
	respStr, err := common.DoPost(urlRegister, common.ConvToJSON(req))
	var resp RegisterResp
	goCheck.Assert(err, IsNil)
	err = json.Unmarshal(respStr, &resp)
	goCheck.Assert(resp.Code, Equals, 0)
}

