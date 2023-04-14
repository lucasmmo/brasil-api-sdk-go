package banksv1_test

import (
	"errors"
	"fmt"
	"testing"

	banksv1 "github.com/lucasmmo/brasil-api-sdk-go/services/banks_v1"
)

func TestBank(t *testing.T) {
    var test_cases = []struct {
        Name   string
        Sample string
        Expected string
        ErrExpected error
    }{
        {
            Name:   "",
            Sample: "2",
            Expected: "",
            ErrExpected: errors.New("error to call /banks/v1/. status code: 404"),
        },
        {
            Name:   "should return BRB - BANCO DE BRASILIA S.A.",
            Sample: "70",
            Expected: "BRB - BANCO DE BRASILIA S.A.",
            ErrExpected: nil,
        },
        {
            Name:   "should return Banco do Brasil S.A.",
            Sample: "1",
            Expected: "Banco do Brasil S.A.",
            ErrExpected: nil,
        },
    }
    client := banksv1.NewClient()
	for _, tc := range test_cases {
		t.Run(tc.Name, func(t *testing.T) {
			output, err := client.FindBank(tc.Sample)
			if err != nil {
                if err.Error() != tc.ErrExpected.Error() {
                    t.Error(err)
                    return
                }
                return
			}
            if output.FullName != tc.Expected {
				t.Error(fmt.Errorf("Expected: %s, Received: %s", tc.Expected, output.FullName))
                return
            }
		})
	}
}


func TestBanks(t *testing.T) {
    var test_cases = []struct{
        Name string
        ErrExpected error
    }{
        {
            Name: "should return a list with all banks infos",
            ErrExpected: nil,
        },
    }
    client := banksv1.NewClient()
    for _, tc := range test_cases {
            t.Run(tc.Name, func(t *testing.T) {
                output, err := client.FindBanks()
                if err != nil {
                    if err.Error() != tc.ErrExpected.Error(){
                        t.Error(err)
                        return
                    }
                }
                if len(*output) == 0 {
                    t.Error(errors.New("the size of output need to be greater than 0"))
                }
            })
    }
    
}
