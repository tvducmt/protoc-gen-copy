package test

import (
	"context"
	"fmt"
	"testing"

	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	// pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	// helper "github.com/tvducmt/protoc-gen-copy/helper"
)

func TestSimple(t *testing.T) {
	cpSvc := pb_cp_svc.NewCopy()
	resp, err := cpSvc.ListCITransactionsRequest(context.Background(), &pb_midd.ListCITransactionsRequest{
		MTransId:   "124235245",
		MId:        "23432",
		MerchantId: "342343",
		ZpTransId:  "3143254",
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("ZpTransId", resp.ZpTransId)
	// x.Copy(&resp, &req)
	fmt.Println("MTransId", resp.MTransId)
	fmt.Println("MId", resp.MId)
	// fmt.Println("MerchantId", resp.MerchantId)

}
