package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// type Input struct {
// 	User     string `json:"user"`
// 	Resource string `json:"resource"`
// 	Action   string `json:"action"`
// }

type Data struct {
	Users     map[string][]string            `json:"users"`
	Resources map[string]map[string][]string `json:"resources"`
	Roles     map[string][]string            `json:"roles"`
}

func prepareData() Data {
	data := Data{
		Users:     map[string][]string{},
		Resources: map[string]map[string][]string{},
		Roles:     map[string][]string{},
	}
	groups := []string{}
	resources := []string{}
	users := []string{}
	roles := []string{}
	actions := []string{}

	for i := 0; i < 1000; i++ {
		groups = append(groups, fmt.Sprintf("group-%d", i))
	}

	for i := 0; i < 1000; i++ {
		resources = append(resources, fmt.Sprintf("resource-%d", i))
	}

	for i := 0; i < 1000; i++ {
		users = append(users, fmt.Sprintf("user-%d", i))
	}

	for i := 0; i < 100; i++ {
		roles = append(roles, fmt.Sprintf("user-%d", i))
	}

	for i := 0; i < 100; i++ {
		actions = append(actions, fmt.Sprintf("act-%d", i))
	}

	for _, user := range users {
		data.Users[user] = groups
	}

	for _, resource := range resources {
		data.Resources[resource] = map[string][]string{}
		for _, group := range groups {
			data.Resources[resource][group] = roles
		}
	}

	for _, role := range roles {
		data.Roles[role] = actions
	}

	return data
}

func main() {
	data := prepareData()
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
