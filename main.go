package main

import (
	"context"
	"fmt"

	"github.com/rsbh/authz/controller"
)

type Input struct {
	User     string `json:"user"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func main() {
	policies := []string{"./policies/policy.rego"}
	ctx := context.Background()
	data := controller.GetJsonFileData("./data.json")
	path := "data.authz.allow"
	ac := controller.New(ctx, policies, path, data)

	input := Input{
		User:     "u1",
		Resource: "r1",
		Action:   "read",
	}

	ok, err := ac.IsAuthorized(ctx, input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)

}
