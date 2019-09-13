package test

import (
	fmt "fmt"
	"testing"

	helper "github.com/tvducmt/protoc-gen-copy/helper"
)

func TestSimple(t *testing.T) {
	req := helper.ListCITransactionsRequestA{
		MTransId:   "124235245",
		MId:        "23432",
		MerchantId: "342343",
	}
	resp := helper.ListCITransactionsRequestB{}
	x := CopyProtoAB{}
	x.Copy(&resp, &req)
	fmt.Println("dsfsdf", resp.MTransId)
	fmt.Println("dsfsdf", resp.MId)

}
