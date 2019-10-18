package main

import (
	"encoding/json"
	"fmt"

	// proto "git.zapa.cloud/merchant-tools/helper/proto"
	// proto "git.zapa.cloud/merchant-tools/helper/proto"

	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"

	// helper "github.com/tvducmt/protoc-gen-copy/helper"
	// pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	// pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	// pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
)

func test1(resp *pb_svc.BODetailReconciliation) {
	cpSvc := pb_cp_svc.NewCopy()
	err := cpSvc.ListCITransactionsRequest(&pb_midd.BODetailReconciliation{
		// FromDate:  &proto.Date{Year: 2019, Month: 9, Day: 23},
		// TransTime: &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
	}, resp)
	if err != nil {
		fmt.Println("Error", err)
	}
	src, _ := json.Marshal(resp)
	fmt.Println("resp", string(src))
}

func test2(resp *pb_svc.BODetailReconciliation) {
	cpSvc := pb_cp_svc.NewCopy()
	err := cpSvc.ListCITransactionsRequest(&pb_midd.BODetailReconciliation{
		// FromDate:     &proto.Date{Year: 2019, Month: 9, Day: 23},
		// TransTime:    &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
		// ZpSystemName: "ZpSystemName",
		// VoucherCode:  3,

		// TimeAttribute: &pb_midd.BODetailReconciliation_TimeAttribute{
		// 	RetryTime: &timestamp.Timestamp{Seconds: 50, Nanos: 1000},
		// },
		CountableAttribute: &pb_midd.BODetailReconciliation_CountableAttribute{
			RefundApi: 3,
			Data: &pb_midd.BODetailReconciliation_CountableAttribute_Transaction{
				K1: "K1",
				K2: "K2",
				Stm: &pb_midd.Statement{
					H1: "Data.H1",
					H2: "Data.H2",
				},
			},
			TpeBankCode: "TpeBankCode",
			Stm: &pb_midd.Statement{
				H1: "H1",
				H2: "H2",
			},
			ItemCount: 5,
		},
		// P1: &pb_midd.BODetailReconciliation_Profile{
		// 	Name: "P1",
		// },
	}, resp)
	if err != nil {
		fmt.Println("Error", err)
	}
	src, _ := json.Marshal(resp)
	fmt.Println("resp", string(src))
}

func main() {
	resp := &pb_svc.BODetailReconciliation{}
	test2(resp)

}
