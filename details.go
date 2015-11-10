package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
)


func getJson() string {

	if _, err := os.Stat(AuthFile); os.IsNotExist(err) {
		FindAuthKey()
	}

	data, err := ioutil.ReadFile(AuthFile)
	check(err)

	jsonStr := `
		{
		    "jsonrpc": "2.0",
		    "method": "trigger.get",
		    "params": {
		        "output": "extend",
		        "selectHosts": "extend",
		        "selectLastEvent": "extend",
		        "sortfield": "lastchange",
		        "monitored": "true",
		        "only_true": "true",
		        "maintenance": "false",
		        "filter": {
		            "value": 1
		        },
		        "limit": "20"
		    },
		    "auth": "` + string(data) + `",
		    "id": 1
		}

		`

	return jsonStr
}

func FindAuthKey() {
	var user User
	jsonStr :=
	`
		{
		    "jsonrpc": "2.0",
		    "method": "user.login",
		    "params": {
		        "user": "` + Username + `",
		        "password": "` + Password + `"
		    },
		    "id": 1
		}
	`

	url := Protocol + "://" + Address + "/zabbix/api_jsonrpc.php"

	body := GetContent(url, jsonStr)
	json.Unmarshal(body, &user)
	ioutil.WriteFile(AuthFile, []byte(user.AuthKey), 0644)
}