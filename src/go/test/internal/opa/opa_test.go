package opa

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

var opa *sdk.OPA
var ctx context.Context

func init() {
	opa, ctx = prepareOpa()
}
func TestAllow(t *testing.T) {
	// test
	input := map[string]interface{}{"open": "sesame"}
	var actual = Allow(input, opa, "authz", ctx)

	if !actual {
		t.Fatal("Expected true, got false")
	}
}

func TestAllowFail(t *testing.T) {
	// test
	input := map[string]interface{}{"open": "not sesame"}
	var actual = Allow(input, opa, "authz", ctx)

	if actual {
		t.Fatal("Expected false, got true")
	}
}

func prepareOpa() (*sdk.OPA, context.Context) {
	ctx := context.Background()

	// create a mock HTTP bundle server
	server, err := sdktest.NewServer(sdktest.MockBundle("/bundles/bundle.tar.gz", map[string]string{
		"example.rego": `
				package authz

				import future.keywords.if

				default allow := false

				allow if input.open == "sesame"
			`,
	}))
	if err != nil {
		log.Fatal("Initialize mock server failed")
	}

	config := []byte(fmt.Sprintf(`{
		"services": {
			"test": {
				"url": %q
			}
		},
		"bundles": {
			"test": {
				"resource": "/bundles/bundle.tar.gz"
			}
		},
		"decision_logs": {
			"console": true
		}
	}`, server.URL()))

	// create an instance of the OPA object
	opa, err := sdk.New(ctx, sdk.Options{
		ID:     "opa-test-1",
		Config: bytes.NewReader(config),
	})
	if err != nil {
		log.Fatal("Initialize OPA failed")
	}

	return opa, ctx
}
