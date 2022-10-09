package common

import (
	"bytes"
	"net/http"
	"io/ioutil"
)







func DoPost(url, data string) (resp []byte, err error) {
	body := bytes.NewReader([]byte(data))
	contentType := "text/plain"
	r, err := http.Post(url, contentType, body)
	resp, err = ioutil.ReadAll(r.Body)
	return
}
