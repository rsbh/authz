package controller

import (
	"context"
	"testing"
)

type Input struct {
	User     string `json:"user"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func BenchmarkIsAuthorized(b *testing.B) {
	policies := []string{"../policies/policy.rego"}
	ctx := context.Background()
	data := GetJsonFileData("../data.json")
	path := "data.authz.allow"
	ac := New(ctx, policies, path, data)
	input := Input{
		User:     "u1",
		Resource: "r1",
		Action:   "read",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ac.IsAuthorized(ctx, input)
	}
}
