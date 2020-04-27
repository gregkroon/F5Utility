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

	var poolnameflag = flag.String("poolname", "default", "F5 pool name")
	var actionflag = flag.String("action", "enablepool", "Enable or Disable F5 pool members")
	var usernameflag = flag.String("username", "admin", "Username for F5")
	var passwordflag = flag.String("password", "*******", "Password for F5")
	var hostflag = flag.String("host", "10.1.1.1", "F5 hostname and port combination ")

	flag.Parse()

	fmt.Println(*poolnameflag)
	fmt.Println(*actionflag)
	fmt.Println(*usernameflag)
	fmt.Println(*passwordflag)
	fmt.Println(*hostflag)

	token := auth(*usernameflag, *passwordflag, *hostflag)

	if *actionflag == "disablepool" {
		fmt.Println("Flag set to disable")
		disablepool(token, *poolnameflag, *hostflag)
	} else if *actionflag == "enablepool" {
		fmt.Println("Flag set to enable")
		enablepool(token, *poolnameflag, *hostflag)
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

	str := fmt.Sprintf("%v", token)
	return str
}

func disablepool(tokenarg string, poolname string, host string) {

	type Items struct {
		Items string `json:"Items"`
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{}`).Get("https://" + host + "/mgmt/tm/ltm/pool/~Common~" + poolname + "/members")

	fmt.Println(resp)

	result := gjson.Get(resp.String(), "items.#.name")

	fmt.Println(err)

	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{"session":"user-disabled"}`).Put("https://" + host + "/mgmt/tm/ltm/pool/~Common~" + poolname + "/members/~Common~" + value.String())
		fmt.Println(err)
		fmt.Println(resp)
		return true // keep iterating
	})

}

func enablepool(tokenarg string, poolname string, host string) {

	type Items struct {
		Items string `json:"Items"`
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{}`).Get("https://" + host + "/mgmt/tm/ltm/pool/~Common~" + poolname + "/members")

	fmt.Println(resp)

	result := gjson.Get(resp.String(), "items.#.name")

	fmt.Println(err)

	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("X-F5-Auth-Token", tokenarg).SetBody(`{"session":"user-enabled"}`).Put("https://" + host + "/mgmt/tm/ltm/pool/~Common~" + poolname + "/members/~Common~" + value.String())
		fmt.Println(err)
		fmt.Println(resp)
		return true // keep iterating
	})

}
