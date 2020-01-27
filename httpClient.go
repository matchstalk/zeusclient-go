package zeusclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	ApiPermList  = "/user/perm/list"
	ApiPermCheck = "/user/perm/check"
	ApiMenuList  = "/user/menu"
)

type ZeusHttpClient struct {
	client        *http.Client
	centerService string
	domain        string
	token         string
}

func NewZeusHttpClient(centerService, domain, token string) ZeusHttpClient {
	return ZeusHttpClient{
		client:        &http.Client{},
		centerService: centerService,
		domain:        domain,
		token:         token,
	}
}

//检查用户权限
func (e ZeusHttpClient) CheckPerm(perm string) (can bool, err error) {
	data := map[string]interface{}{
		"domain": e.domain,
		"perm":   "/taskmanage/list",
	}
	var rb []byte
	rb, err = e.send(http.MethodPost, e.centerService+ApiPermCheck, data)
	if err != nil {
		return
	}
	var body CheckPermRsp
	err = json.Unmarshal(rb, &body)
	if err != nil {
		return
	}
	can = body.Data.Result
	return
}

//获取用户权限
func (e ZeusHttpClient) GetUserPerms() (perms []string, err error) {
	var rb []byte
	rb, err = e.send(http.MethodGet, e.centerService+ApiPermList+"?domain="+e.domain, nil)
	if err != nil {
		return
	}
	var rsp UserPermsRsp
	err = json.Unmarshal(rb, &rsp)
	if err != nil {
		return
	}
	perms = rsp.Data.Result
	return
}

//获取用户菜单
func (e ZeusHttpClient) GetUserMenus() (menus []Menus, err error) {
	var rb []byte
	rb, err = e.send(http.MethodGet, e.centerService+ApiMenuList+"?domain="+e.domain, nil)
	if err != nil {
		return
	}
	var rsp UserMenusRsp
	err = json.Unmarshal(rb, &rsp)
	if err != nil {
		return
	}
	menus = rsp.Data.Result
	return
}

//发送请求
func (e ZeusHttpClient) send(method, u string, body interface{}) (data []byte, err error) {
	var bytesData []byte
	var r *bytes.Reader
	bytesData, err = json.Marshal(data)
	if err != nil {
		return
	}
	r = bytes.NewReader(bytesData)
	var req *http.Request
	req, err = http.NewRequest(
		method,
		u,
		r)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+e.token)
	req.Header.Set("Content-Type", "application/json")
	var rsp *http.Response
	rsp, err = e.client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	data, err = ioutil.ReadAll(rsp.Body)
	return
}
