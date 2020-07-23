// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: copy.proto

package copy

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	core_service "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	middleware "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	_ "git.zapa.cloud/merchant-tools/protobuf/middleware/report"
	_ "git.zapa.cloud/merchant-tools/protobuf/report-service"
	_ "github.com/golang/protobuf/ptypes/empty"
	context "context"
	reflect "reflect"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context

func isNil(field interface{}) bool {
	zero := reflect.Zero(reflect.TypeOf(field)).Interface()
	if reflect.DeepEqual(field, zero) {
		return true
	}
	return false
}

type copy struct {
}

func NewCopy() *copy {
	return &copy{}
}

func (c *copy) CopyNestedFishSlice(from []*middleware.NestedFish, to *[]*core_service.NestedFish) error {
	for _, v := range from {
		resp := &core_service.NestedFish{}
		c.CopyNestedFish(v, resp)
		*to = append(*to, resp)
	}
	return nil
}
func (c *copy) CopyNestedFish(from *middleware.NestedFish, to *core_service.NestedFish) error {
	if !isNil(from.Fish) {
		tmp := make([]*core_service.Fish, len(from.Fish))
		for i, v := range from.Fish {
			glog.Infoln("v", v)
			tmp[i] = &core_service.Fish{
				Statement1: &core_service.Statement1{
					H1: func(h *middleware.Statement1) string {
						if h == nil {
							return reflect.Zero(reflect.TypeOf(reflect.String)).Interface().(string)
						}
						return h.H1
					}(from.Fish.Statement1[i]),
					H2: func(h *middleware.Statement1) string {
						if h == nil {
							return reflect.Zero(reflect.TypeOf(reflect.String)).Interface().(string)
						}
						return h.H2
					}(from.Fish.Statement1[i]),
				},
			}
		}
		to.Fish = tmp
	}
	return nil
}

func (c *copy) CopyNestedFish(from *middleware.NestedFish,to *core_service.NestedFish) error{
	if !isNil(from.Fish){
		tmp := make([]*core_service.Fish, len(from.Fish))
		for i, v := range from.Fish {
		glog.Infoln("v",v)
		tmp[i] = &core_service.Fish{
			Statement1: v.Statement1,
		}
			// if !isNil(from.Fish.Statement1){
				
			// 	tmp := make([]*core_service.Statement1, len(from.Fish.Statement1))
			// 	for i, v := range from.Fish.Statement1 {
			// 		glog.Infoln("v",v)
			// 		tmp[i] = &core_service.Statement1{
			// 			H1: func(h *middleware.Statement1) string {
			// 					if h == nil {
			// 							return reflect.Zero(reflect.TypeOf(reflect.String)).Interface().(string)
			// 					}
			// 					return h.H1
			// 			}(from.Fish.Statement1[i]),
			// 			H2: func(h *middleware.Statement1) string {
			// 					if h == nil {
			// 							return reflect.Zero(reflect.TypeOf(reflect.String)).Interface().(string)
			// 					}
			// 					return h.H2
			// 			}(from.Fish.Statement1[i]),
			// 		},
			// 	}
			// }
		}
	}
	to.Fish= tmp
	}
	return  nil
}
