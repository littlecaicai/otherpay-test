package index3

import (
	//"encoding/json"
	"fmt"
	"otherpay-test/client"
	"otherpay-test/common"
	. "gopkg.in/check.v1"
	"time"
)




type LoginOrRegister struct {
}

var _ = Suite(&LoginOrRegister{})

var (
	urlLoginOrRegister string = "http://localhost:8765/login_or_register"
)

type CheckInfo struct {
	Address string `json:"address" example:"0x82198867e32e4405f0d4ff46a5a1d214c9b1d474"`
	Msg string  `json:"msg" example:"I am signing my one-time nonce: ABCDEF"`
	Sign string `json:"sign" example:"0xa9481518a8d279a5936735b50f1bda60f2998e745acfd406a6d6dd7e25786465119fb2c1bdc58fcc7987f996c8e787be094acbbeda2d05ea3507a129e9c26bb21c"`
}

type RequestLoginOrRegister struct {
	CheckInfo
	LoginTime int64 `json:"login_time"`
}

type ResponseToken struct {
	Token string `json:"token"`
}







func (s *LoginOrRegister) TestRegisterCase00(goCheck *C) {
	//参数合法，可以注册成功
	privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	msg := "I am registing for index3 "
	addr, sign := common.GetSignNew(privateHex, msg)
	sql := fmt.Sprintf("delete from mirror_user where addr = \"%s\"", addr)
	_, err := client.MysqlClientIndex3().Exec(sql)
	goCheck.Assert(err, IsNil)
	req := RequestLoginOrRegister{
		CheckInfo: CheckInfo{
			Address: addr,
			Msg:     msg,
			Sign:    sign,
		},
		LoginTime: time.Now().Unix(),
	}
	respStr, err := common.DoPost(urlLoginOrRegister, common.ConvToJSON(req))
	goCheck.Assert(err, IsNil)
	fmt.Println(string(respStr))
}

//func (s *LoginOrRegister) TestRegisterCase01(goCheck *C) {
//	//参数不合法，注册失败
//	privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
//	addr, _, st:= common.GetSign(privateHex)
//	sql := fmt.Sprintf("delete from otherpay_addr where addr = \"%s\"", addr)
//	_, err := client.MysqlClient().Exec(sql)
//	goCheck.Assert(err, IsNil)
//	req := RegisterReq {
//		Addr: addr,
//		TimeStamp: st,
//		Sign: "incrrect_sign",
//	}
//	respStr, err := common.DoPost(urlRegister, common.ConvToJSON(req))
//	var resp RegisterResp
//	goCheck.Assert(err, IsNil)
//	err = json.Unmarshal(respStr, &resp)
//	goCheck.Assert(resp.Code, Not(Equals), 0)
//}


