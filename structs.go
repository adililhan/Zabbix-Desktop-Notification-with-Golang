package main

type Alert struct {
	Result SubResult
}

type SubResult []struct {
	Description string
	Hosts SubHosts
}

type SubHosts []struct {
	Name string
	Error string
}

type User struct {
	AuthKey string `json:"result"`
}