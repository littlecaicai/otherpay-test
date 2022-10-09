package common

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
	"strings"
	"time"
)


func GetSign(privateHex string) string {
	//privateHex := "ae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	privateKey, err := crypto.HexToECDSA(privateHex)

	if err != nil {
		return ""
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("err2", err)
	}

	addr := strings.ToLower(crypto.PubkeyToAddress(*publicKeyECDSA).String())
	fmt.Println("addr:", addr)

	extTime := time.Now().Unix()
	message := []byte(addr + ":" + strconv.FormatInt(extTime, 10))
	sign, _ := crypto.Sign(signHash(message), privateKey)
	return hexutil.Encode(sign)
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}