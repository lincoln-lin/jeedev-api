package myhttp

import "github.com/kirinlabs/HttpRequest"

func HttpPost(url string , params map[string]interface{})  (resp *HttpRequest.Response, err error ){
	req := HttpRequest.NewRequest()
	req.SetTimeout(5)
	return req.Get(url, params)
}
