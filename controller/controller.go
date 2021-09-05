package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

type AuthzController struct {
	rego              *rego.Rego
	preparedEvalQuery rego.PreparedEvalQuery
}

func (ac *AuthzController) IsAuthorized(ctx context.Context, input interface{}) (bool, error) {
	result, err := ac.preparedEvalQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		panic(err)
	}
	return result.Allowed(), nil
}

func GetJsonFileData(filePath string) map[string]interface{} {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var json map[string]interface{}

	err = util.UnmarshalJSON(byteValue, &json)
	if err != nil {
		log.Fatal(err)
	}
	return json
}

func New(ctx context.Context, policies []string, path string, data map[string]interface{}) *AuthzController {
	store := inmem.NewFromObject(data)
	txn, err := store.NewTransaction(ctx, storage.WriteParams)

	if err != nil {
		log.Fatal(err)
	}

	r := rego.New(
		rego.Query(path),
		rego.Load(policies, nil),
		rego.Store(store),
		rego.Transaction(txn),
	)

	eq, err := r.PrepareForEval(ctx)
	if err != nil {
		panic(err)
	}

	return &AuthzController{
		rego:              r,
		preparedEvalQuery: eq,
	}
}
