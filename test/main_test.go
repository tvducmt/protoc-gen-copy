package test

import (
	fmt "fmt"
	"testing"

	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	// helper "github.com/tvducmt/protoc-gen-copy/helper"
)

func TestSimple(t *testing.T) {
	req := pb_midd.ListCITransactionsRequest{
		MTransId:   "124235245",
		MId:        "23432",
		MerchantId: "342343",
	}
	resp := pb_svc.ListCITransactionsRequest{}

	x := pb_cp_svc.CopyProtoAB{}
	x.Copy(&resp, &req)
	fmt.Println("dsfsdf", resp.MTransId)
	fmt.Println("dsfsdf", resp.MId)

}
