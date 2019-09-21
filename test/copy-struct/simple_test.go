package main

import (
	"encoding/json"
	"reflect"
	"testing"

	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
)

func covertToJSON(resp interface{}) string {
	src, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	// fmt.Println(string(queryJSON))
	return string(src)
}

func TestStringType(t *testing.T) {
	cpSvc := pb_cp_svc.NewCopy()
	resp := &pb_svc.ListCITransactionsRequest{}
	cpSvc.ListCITransactionsRequest(&pb_midd.ListCITransactionsRequest{
		MTransId: "124235245",
	}, resp)
	got := covertToJSON(resp)
	want := `{"m_trans_id":"124235245"}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestLoop(t *testing.T) {
	cpSvc := pb_cp_svc.NewCopy()
	resp := &pb_svc.ListCITransactionsRequest{}
	cpSvc.ListCITransactionsRequest(&pb_midd.ListCITransactionsRequest{
		Data: &pb_midd.Transaction{},
	}, resp)
	got := covertToJSON(resp)
	want := `{"m_trans_id":"124235245"}`
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
