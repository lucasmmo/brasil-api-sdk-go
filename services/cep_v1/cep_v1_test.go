package cepv1_test

import (
	"fmt"
	"testing"

	cepv1 "github.com/lucasmmo/brasil-api-sdk-go/service/cep_v1"
)

var test_cases = []struct {
	Name   string
	Sample string
}{
	{
		Name:   "should return a valid infos from a valid cep",
		Sample: "89010025",
	},
}

func TestCep(t *testing.T) {
    client := cepv1.NewClient()

	for _, tc := range test_cases {
		t.Run(tc.Name, func(t *testing.T) {
			cep, err := client.FindByCepWithMultiplesProviders(tc.Sample)
			if err != nil {
				t.Error(err)
				return
			}
			if cep.Cep != tc.Sample {
				t.Error(fmt.Errorf("Expected: %s, Received: %s", tc.Sample, cep.Cep))
				return
			}
		})
	}
}
