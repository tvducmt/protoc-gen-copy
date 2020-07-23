package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"time"

	// helperMidReport "git.zapa.cloud/merchant-tools/protobuf/middleware/report"
	// helperProtoReport "git.zapa.cloud/merchant-tools/protobuf/report-service"
	// "github.com/golang/glog"
	"github.com/golang/glog"
	ggProto "github.com/golang/protobuf/proto"
	pb_cp_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/copy-service"
	pb_svc "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	pb_midd "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	"gitlab.com/core-grpc-project/helper/copier"
)

func main() {
	TestCopy()
}

func TestCopy() {
	cpuprofile := flag.String("cpuprofile", "cpu.prof", "write cpu profile to `file`")

	data, err := ioutil.ReadFile("msg.dat")
	if err != nil {
		glog.Errorln(err)
		return
	}
	coreResp := &pb_svc.ListTransactionResponse{}
	now := time.Now()
	if err = ggProto.Unmarshal(data, coreResp); err != nil {
		glog.Errorln(err)
		return
	}
	glog.Infoln("resp.first", coreResp.GetData()[0])
	glog.Infoln("len(coreResp)", len(coreResp.GetData()))
	glog.Infoln("Decode ", time.Since(now))
	midResp := &pb_midd.ListTransLogResponse{}
	// resp1 = nil
	glog.Infoln(*midResp)
	// return

	//ggProto.Merge(resp1, resp)
	// copier.Copy(midResp, coreResp)
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if false {
		if err = copier.Copy(midResp, coreResp); err != nil {
			glog.Errorln(err)
			return
		}
		writeFile(midResp.Data, "copier.txt")
	} else {
		// test func protoc-gen-copy
		glog.Infoln("midResp Ingto", len(midResp.Data))
		cpSvc := pb_cp_svc.NewCopy()
		err = cpSvc.ListCITransactionResponse(coreResp, midResp)
		if err != nil {
			glog.Errorln("Error", err)
		}
		writeFile(midResp.Data, "protoc-gen-copy.txt")
	}
	if n := len(midResp.Data); n > 0 {
		glog.Infoln(midResp.Data[n-1])
	}

	glog.Infoln("Copy ", time.Since(now))
	glog.Infoln("len(resp1.Data)", len(midResp.Data))
}

func writeFile(data interface{}, filename string) {
	src, _ := json.Marshal(data)
	dataString := string(src)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(dataString)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
