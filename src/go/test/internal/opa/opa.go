package opa

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/sdk"
)

func Allow(input interface{}, opa *sdk.OPA, opaPath string, ctx context.Context) bool {
	result, err := opa.Decision(ctx, sdk.DecisionOptions{Path: fmt.Sprintf("%s/allow", opaPath), Input: input})

	if err != nil {
		log.Println(err)
		return false
	}

	return result.Result.(bool)
}
