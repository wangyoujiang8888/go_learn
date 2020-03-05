package main

import (
	"encoding/json"
	"example/util/log"
	"example/util/network"
	"fmt"
)

func main()  {
	url := "http://kandian.youth.cn/v3/user/userinfo.json?access=WIFI&androidid=c3de226d22656499&app_version=1.8.2&carrier=%E4%B8%AD%E5%9B%BD%E8%81%94%E9%80%9A&channel=c1005&device_id=39374144&device_model=MI%208%20Lite&device_platform=android&gt_uid=3d8e195c8b38f76ae36992c1e4e65452&iid=0&imei=868453048520615&inner_version=202001061736&mi=1&mobile_type=1&net_type=1&openudid=c3de226d22656499&os_api=28&os_version=PKQ1.181007.001&phone_sim=1&request_time=1579159656&sm_device_id=202001061741029f6e6d87bf73439af5c5a2bd9f7d390b01b99338a29982fd&uid=42185765&version_code=39&zqkey=MDAwMDAwMDAwMJCMpN-w09Wtg5-Bb36eh6CPqHualIejlq6bqWaxt5drhYyp4LDPyGl9onqkj3ZqYJa8Y898najWsJupY7Hdn7GFjKCXr6m6apqGcXY&zqkey_id=58c779d7e01abfdfa69bd62c0e19024d&sign=739032bb830dfe0b0a8515ab3412c1e3"
	resp,err := network.Get(url,nil,nil)
	if err !=nil {
		log.Logger.LogError("http get error:"+err.Error())
	}
	//out,_:= json.Marshal(body)
	//fmt.Println(out)
	type Item struct {
		UID string `json:"uid"`
		ISBlocked string `json:"is_blocked"`
	}
	type Ret struct {
		Success bool `json:"success"`
		ErrorCode string `json:"error_code"`
		Message string `json:"message"`
		Items Item `json:"items"`
	}
	var ret Ret
	json.Unmarshal(resp,&ret)
	out,_ := json.Marshal(ret)
	fmt.Println(string(out))
}

