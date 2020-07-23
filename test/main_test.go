package main

import (
	"encoding/json"
	reflect "reflect"
	"testing"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
)

func TestCpStructToStruct(t *testing.T) {
	resp := &pb_svc.BODetailReconciliation{}
	cpSvc := pb_cp_svc.NewCopy()
	err := cpSvc.ListCITransactionsRequest(&pb_midd.BODetailReconciliation{
		OutnerFiled: &pb_midd.OutnerField{
			Name: "duc",
		},
		// FromDate:     &proto.Date{Year: 2019, Month: 9, Day: 23},
		TransTime:    &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
		ZpSystemName: "ZpSystemName",
		VoucherCode:  3,

		TimeAttribute: &pb_midd.BODetailReconciliation_TimeAttribute{
			RetryTime: &timestamp.Timestamp{Seconds: 50, Nanos: 1000},
		},
		CountableAttribute: &pb_midd.BODetailReconciliation_CountableAttribute{
			RefundApi: &pb_midd.RefundApi{
				Stm1: &pb_midd.Statement1{
					H1: "GTH1",
					H2: "GTH2",
				},
			},
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
		P1: &pb_midd.BODetailReconciliation_Profile{
			Name: "P1",
		},
	}, resp)
	if err != nil {
		glog.Errorln("Error", err)
	}
	src, _ := json.Marshal(resp)
	got := string(src)
	want := `{"outner_filed":{"name":"duc"},"trans_time":{"seconds":60,"nanos":1000},"zp_system_name":"ZpSystemName","voucher_code":3,"time_attribute":{"retry_time":{"seconds":50,"nanos":1000}},"countable_attribute":{"refund_api":{"stm1":{"h1":"GTH1","h2":"GTH2"}},"data":{"k1":"K1","k2":"K2","stm":{"h1":"Data.H1","h2":"Data.H2"}},"tpe_bank_code":"TpeBankCode","stm":{"h1":"H1","h2":"H2"},"item_count":5}}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestCpSliceToSlice(t *testing.T) {
	resp := &[]*pb_svc.BODetailReconciliation{}
	cpSvc := pb_cp_svc.NewCopy()
	reqs := []*pb_midd.BODetailReconciliation{
		{
			TransTime:    &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
			ZpSystemName: "ZpSystemName",
			VoucherCode:  3,

			TimeAttribute: &pb_midd.BODetailReconciliation_TimeAttribute{
				RetryTime: &timestamp.Timestamp{Seconds: 50, Nanos: 1000},
			},
			CountableAttribute: &pb_midd.BODetailReconciliation_CountableAttribute{
				RefundApi: &pb_midd.RefundApi{
					Stm1: &pb_midd.Statement1{
						H1: "GTH1",
						H2: "GTH2",
					},
				},
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
			P1: &pb_midd.BODetailReconciliation_Profile{
				Name: "P1",
			},
		},
		{
			TransTime:    &timestamp.Timestamp{Seconds: 60, Nanos: 1000},
			ZpSystemName: "ZpSystemName",
		},
		{
			TimeAttribute: &pb_midd.BODetailReconciliation_TimeAttribute{
				RetryTime: &timestamp.Timestamp{Seconds: 50, Nanos: 1000},
			},
		},
		{
			CountableAttribute: &pb_midd.BODetailReconciliation_CountableAttribute{
				RefundApi: &pb_midd.RefundApi{
					Stm1: &pb_midd.Statement1{
						H1: "GTH1",
						H2: "GTH2",
					},
				},
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
		},
		{
			P1: &pb_midd.BODetailReconciliation_Profile{
				Name: "P1",
			},
		},
	}
	err := cpSvc.ListCITransactionsRequestSlice(reqs, resp)
	if err != nil {
		glog.Errorln("Error", err)
	}
	src, _ := json.Marshal(resp)
	got := string(src)
	want := `[{"trans_time":{"seconds":60,"nanos":1000},"zp_system_name":"ZpSystemName","voucher_code":3,"time_attribute":{"retry_time":{"seconds":50,"nanos":1000}},"countable_attribute":{"refund_api":{"stm1":{"h1":"GTH1","h2":"GTH2"}},"data":{"k1":"K1","k2":"K2","stm":{"h1":"Data.H1","h2":"Data.H2"}},"tpe_bank_code":"TpeBankCode","stm":{"h1":"H1","h2":"H2"},"item_count":5}},{"trans_time":{"seconds":60,"nanos":1000},"zp_system_name":"ZpSystemName"},{"time_attribute":{"retry_time":{"seconds":50,"nanos":1000}}},{"countable_attribute":{"refund_api":{"stm1":{"h1":"GTH1","h2":"GTH2"}},"data":{"k1":"K1","k2":"K2","stm":{"h1":"Data.H1","h2":"Data.H2"}},"tpe_bank_code":"TpeBankCode","stm":{"h1":"H1","h2":"H2"},"item_count":5}},{}]`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestRepeatedStringCp(t *testing.T) {
	resp := &pb_svc.Fish{}
	cpSvc := pb_cp_svc.NewCopy()
	err := cpSvc.CopyFish(&pb_midd.Fish{
		Name: []string{"hello1", "hello2"},
	}, resp)
	if err != nil {
		glog.Errorln("Error", err)
	}
	src, _ := json.Marshal(resp)
	got := string(src)
	glog.Errorln("got", got)
	want := `{"name":["hello1","hello2"]}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// func TestRepeatedObjectCp(t *testing.T) {
// 	resp := &pb_svc.Fish{}
// 	cpSvc := pb_cp_svc.NewCopy()
// 	err := cpSvc.CopyFish(&pb_midd.Fish{
// 		Name: []string{"hello1", "hello2"},
// 	}, resp)
// 	if err != nil {
// 		glog.Errorln("Error", err)
// 	}
// 	src, _ := json.Marshal(resp)
// 	got := string(src)
// 	glog.Errorln("got", got)
// 	want := `{"name":["hello1","hello2"]}`
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }
