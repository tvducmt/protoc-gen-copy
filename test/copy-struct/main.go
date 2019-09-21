package main

import (
	"encoding/json"
	"fmt"

	"git.zapa.cloud/merchant-tools/helper/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"

	// helper "github.com/tvducmt/protoc-gen-copy/helper"
	// pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	// pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	// pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
)

func main() {
	cpSvc := pb_cp_svc.NewCopy()
	resp := &pb_svc.BODetailReconciliation{}
	err := cpSvc.ListCITransactionsRequest(&pb_midd.BODetailReconciliation{
		{&proto.Date{Year: 2019, Month: 9, Day: 23}, `{"reqDate":1569171600000}`},
		TransTime: &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
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
