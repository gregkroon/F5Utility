package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	gjson "github.com/tidwall/gjson"
)



func main() {

	client := resty.New()

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`{"username": admin,"password": 7642098@,"loginProviderName": "tmos"}`).Post("https://10.1.1.170:8443/mgmt/shared/authn/login")

	//var resp = rest.Post("https://10.1.1.170:8443/mgmt/shared/authn/login", Body)
	//fmt.Println(resp , "this is not an error")
	fmt.Println(err)

	bs := resp.String()

	//byteValue, _ := ioutil.ReadAll(bs)
	var result map[string]interface{}
	json.Unmarshal([]byte(bs), &result)

	var token interface{}
	token = result["token"].(map[string]interface{})["token"]
	//test := result["token"].(map[string]interface{})["type"]

	fmt.Println(token)

	//misc(token.(string))

	//func misc(tokenarg string) {

	//client := resty.New()

	//client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	//resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{}`).Get("https://10.1.1.170:8443/mgmt/tm/ltm/pool/~Common~PoolA/members")

	//fmt.Println(err)
	//fmt.Println(resp)

	//}
	disablepool(token.(string))

}


func disablepool(tokenarg string) {



	type Items struct {
		Items string `json:"Items"`
	}

	//var items Items

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{}`).Get("https://10.1.1.170:8443/mgmt/tm/ltm/pool/~Common~PoolA/members")

	fmt.Println(resp)

	//json.Unmarshal([]byte(resp.Body()), &items)

	result := gjson.Get(resp.String(), "items.#.name")
	//println(value.String())

	fmt.Println(err)
	//fmt.Println(items)
	//fmt.Println(members.Name)

	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{"session":"user-disabled"}`).Put("https://10.1.1.170:8443/mgmt/tm/ltm/pool/~Common~PoolA/members/~Common~" + value.String())
		fmt.Println(err)
		fmt.Println(resp)
		return true // keep iterating
	})



}