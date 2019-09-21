package main

import (
	"encoding/json"
	"fmt"

	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	// helper "github.com/tvducmt/protoc-gen-copy/helper"
	// pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	// pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	// pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
)

func main() {
	cpSvc := pb_cp_svc.NewCopy()
	resp := &pb_svc.ListCITransactionsRequest{}
	err := cpSvc.ListCITransactionsRequest(&pb_midd.ListCITransactionsRequest{
		Data: &pb_midd.Transaction{},
	}, resp)
	if err != nil {
		fmt.Println("Error", err)
	}
	src, _ := json.Marshal(resp)
	fmt.Println("resp", string(src))

}

// MTransId:   "124235245",
// 		MId:        "23432",
// 		MerchantId: "342343",
// 		ZpTransId:  "3143254",
// 		Data: &pb_midd.Transaction{
// 			MA:  "ma",
// 			ZpB: "zpb",
// 			Hello: &pb_midd.Hello{
// 				KA: "ka",
// 				KB: "kb",
// 			},
// 		},
