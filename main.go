package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"

	"fmt"
	resty "github.com/go-resty/resty/v2"
	gjson "github.com/tidwall/gjson"
)

func main() {

	var actionflag = flag.String("action", "enable", "Enable or Disable F5 pool members")
	var usernameflag = flag.String("username", "admin", "Username for F5")
	var passwordflag = flag.String("password", "*******", "Password for F5")
	var hostflag = flag.String("host", "10.1.1.1", "F5 hostname and port combination ")

	//fmt.Println(os.Args[0],os.Args[1],os.Args[2])
	flag.Parse()

	fmt.Println(*actionflag)
	fmt.Println(*usernameflag)
	fmt.Println(*passwordflag)
	fmt.Println(*hostflag)

	token := auth(*usernameflag, *passwordflag, *hostflag)

	if *actionflag == "disable" {
		fmt.Println("Flag set to disable")
		disablepool(token)
	} else if *actionflag == "enable" {
		fmt.Println("Flag set to enable")
		enablepool(token)
	}

}

func auth(username string, password string, host string) string {

	client := resty.New()

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`{"username":` + username + `,"password":` + password + `,"loginProviderName": "tmos"}`).Post("https://" + host + "/mgmt/shared/authn/login")

	fmt.Println(err)

	bs := resp.String()

	var result map[string]interface{}
	json.Unmarshal([]byte(bs), &result)

	var token interface{}
	token = result["token"].(map[string]interface{})["token"]

	//fmt.Println(token)
	//tokenstring := token
	str := fmt.Sprintf("%v", token)
	return str
}

func disablepool(tokenarg string) {

	type Items struct {
		Items string `json:"Items"`
	}

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

func enablepool(tokenarg string) {

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
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{"session":"user-enabled"}`).Put("https://10.1.1.170:8443/mgmt/tm/ltm/pool/~Common~PoolA/members/~Common~" + value.String())
		fmt.Println(err)
		fmt.Println(resp)
		return true // keep iterating
	})

}
