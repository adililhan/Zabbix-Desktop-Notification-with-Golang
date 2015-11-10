package main

import "encoding/json"
import "strings"

func main() {

	url := Protocol + "://" + Address + "/zabbix/api_jsonrpc.php"

	jsonStr := getJson()

	body := GetContent(url, jsonStr)

	var Alert Alert
	var result []string
	var description []string

	json.Unmarshal(body, &Alert)

	for _, val := range Alert.Result {
		val.Description = strings.Replace(val.Description, "{HOST.NAME}",val.Hosts[0].Name, -1)
		description = append(description, "HOST: " + val.Hosts[0].Name + "\n" + val.Description + "\n" + val.Hosts[0].Error)
		result = append(result, "HOST: " + val.Hosts[0].Name + "\n" + val.Description)
	}

	if Notification == true {
		SendNotification(result)
	}

	if Console == true {
		SendNotificationConsole(description)
	}

}
