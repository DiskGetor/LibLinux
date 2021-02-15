package net

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HttpRequestCtx struct {
	Client      *http.Client
	StopCode    int
	Method      string
	Url         string
	Head        map[string]string
	RequestBody interface{}
}

func (p *ObjNet) HttpRequest(data *HttpRequestCtx) (Response *http.Response, err error) {
	//data.Client = p.Client //临时操作，删错了代码，设代理的话工程结构体还是要定义客户端，或者在工程每个网站请求之前设置p对象里面的客户端的代码即可
	switch data.RequestBody.(type) {
	case string:
		p.RequestReader = strings.NewReader(data.RequestBody.(string))
	case []byte:
		p.RequestReader = bytes.NewReader(data.RequestBody.([]byte))
	}
	p.Request, err = http.NewRequest(data.Method, data.Url, p.RequestReader)
	if !W.CheckErr(err) {
		return
	}
	p.Request.Close = true
	//Request.Header.Add("Connection", "close")
	for k, v := range data.Head {
		p.Request.Header.Set(k, v)
	}
	return p.Client.Do(p.Request)
}

func (p *ObjNet) HttpReadRequestBody(data *HttpRequestCtx) (buf []byte, err error) {
	defer func() {
		if p.Response != nil {
			W.CheckErr(p.Response.Body.Close())
		}
	}()
	p.Response, err = p.HttpRequest(data)
	if !W.CheckErr(err) {
		return
	}
	if p.Response.StatusCode == data.StopCode {
		return ioutil.ReadAll(p.Response.Body)
	}
	return nil, errors.New(p.Response.Status + " != StopCode " + strconv.Itoa(data.StopCode))
}
