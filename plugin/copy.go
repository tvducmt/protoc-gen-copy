package plugin

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/golang/glog"
	cpier "github.com/tvducmt/protoc-gen-copy/protobuf"
)

const (
	contextPkgPath   = "context"
	reflectPkgPath   = "reflect"
	timeStampPkgPath = "github.com/golang/protobuf/ptypes/timestamp"
)

type copy struct {
	*generator.Generator
	generator.PluginImports
	protoPkg   generator.Single
	stringsPkg generator.Single
}

// NewCopy ...
func NewCopy() generator.Plugin {
	return &copy{}
}

func (c *copy) Name() string {
	return "copy"
}

var (
	contextPkg   string
	reflectPkg   string
	timeStampPkg string
)

func (c *copy) Init(g *generator.Generator) {
	c.Generator = g
}
func (c *copy) objectNamed(name string) generator.Object {
	c.Generator.RecordTypeUse(name)
	return c.Generator.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (c *copy) typeName(str string) string {
	return c.Generator.TypeName(c.objectNamed(str))
}

func (c *copy) Generate(file *generator.FileDescriptor) {
	c.PluginImports = generator.NewPluginImports(c.Generator)
	c.stringsPkg = c.NewImport("strings")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	// c.reflectPkg = c.NewImport("reflect")
	contextPkg = string(c.Generator.AddImport(contextPkgPath))
	reflectPkg = string(c.Generator.AddImport(reflectPkgPath))
	c.P("// Reference imports to suppress errors if they are not otherwise used.")
	c.P("var _ ", contextPkg, ".Context")
	c.P()

	c.P(`func isNil(field interface{}) bool {`)
	c.P(`zero := reflect.Zero(reflect.TypeOf(field)).Interface()`)
	c.P(`if reflect.DeepEqual(field, zero) {`)
	c.P(`return true`)
	c.P(`}`)
	c.P(`return false`)
	c.P(`}`)

	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	for i, service := range file.FileDescriptorProto.Service { //Messages()
		c.generateService(file, service, i)
	}
}

// // generateClientSignature returns the client-side signature for a method.
func (c *copy) generateClientSignature(servName string, method *pb.MethodDescriptorProto, isCpSlice bool) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	var reqArg string   //:= "from *" + c.typeName(method.GetInputType())
	var respName string //:= "to *" + c.typeName(method.GetOutputType())
	if isCpSlice {
		methName = methName + "Slice"
		reqArg = "from []*" + c.typeName(method.GetInputType())
		respName = "to *[]*" + c.typeName(method.GetOutputType())
	} else {
		reqArg = "from *" + c.typeName(method.GetInputType())
		respName = "to *" + c.typeName(method.GetOutputType())
	}
	return fmt.Sprintf("%s(%s,%s) error", methName, reqArg, respName)
}
func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// ObjQueue ...
type ObjQueue struct {
	Name   string
	Type   string
	Val    string
	IsEnum bool
}

func getTypeOfFied(field *descriptor.FieldDescriptorProto) string {
	if field.IsBool() {
		return "bool"
	} else if field.IsString() {
		return "string"
	}
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "float"
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "double"
	}
	return "PrimitiveType"
}
func (c *copy) getNestedType(input string) (string, bool) {
	if input == "google.protobuf.Timestamp" {
		timeStampPkg = string(c.Generator.AddImport(timeStampPkgPath))
		return "timestamp.Timestamp", false
	}
	if input == "google.type.Date" {
		return c.protoPkg.Use() + `.Date`, false
	}
	arr := strings.Split(input, ".")
	if len(arr) > 2 {
		concatSrting := arr[0] + "."
		for i := 1; i < len(arr)-1; i++ {
			concatSrting += arr[i] + `_`
		}
		concatSrting += arr[len(arr)-1]
		return concatSrting, true
	}
	return input, false
}
func (c *copy) getFieldsOfMsg(msg *descriptor.DescriptorProto, path string) []interface{} {
	args := []interface{}{}
	for _, v := range msg.Field {

		if v.IsMessage() {
			mpMsg := make(map[ObjQueue][]interface{})
			msg1 := c.ObjectNamed(v.GetTypeName()).(*generator.Descriptor).DescriptorProto
			path = path + "." + strings.Title(v.GetJsonName())
			x := c.getFieldsOfMsg(msg1, path)
			typeNested, _ := c.getNestedType(TrimFirstRune(v.GetTypeName()))
			mpMsg[ObjQueue{Name: strings.Title(v.GetJsonName()), Type: typeNested, Val: path}] = x
			args = append(args, mpMsg)
			path = ""
		} else if v.IsEnum() {
			// glog.Infoln("genumfdgfd", v.GetTypeName())
			// path = path + "." + strings.Title(v.GetJsonName())
			typeNested, _ := c.getNestedType(TrimFirstRune(v.GetTypeName()))
			args = append(args, ObjQueue{Name: strings.Title(v.GetJsonName()), Type: typeNested, Val: path + "." + strings.Title(v.GetJsonName()), IsEnum: true})
		} else {
			primitiveType := getTypeOfFied(v)
			args = append(args, ObjQueue{Name: strings.Title(v.GetJsonName()), Type: primitiveType, Val: path + "." + strings.Title(v.GetJsonName())})
		}
	}
	return args
}

func (c *copy) generateStruct(v interface{}, path string, checkInner bool, checkOuter bool, fromArgString []interface{}) {

	if k, ok := v.(ObjQueue); ok {
		if k.IsEnum {
			result := ExistInFromArr(k.Val, fromArgString)
			if result.Name != "" {
				glog.Infoln("into herre", result.Name)
				if checkInner {
					c.P(k.Name, `: func(h *`, result.Type, `) `, k.Type, ` {`) //  `: from`, k.Val, `,`
					c.P(`	if h == nil {`)
					c.P(`		return 0`)
					c.P(`	}`)
					c.P(`	return `, k.Type, `(h.`, k.Name, `)`)
					c.P(`}(from.`, TrimFirstRune(result.Val), `),`)

				} else {
					c.P(`if !isNil(from`, k.Val, `){`)
					c.P(path+k.Name, `= `, k.Type, `(from`, k.Val, `)`)
					c.P(`}`)
				}
			}
		} else {
			result := ExistInFromArr(k.Val, fromArgString)
			if result.Name != "" {
				if checkInner {
					c.P(k.Name, `: func(h *`, result.Type, `) `, k.Type, ` {`) //  `: from`, k.Val, `,`
					c.P(`	if h == nil {`)
					c.P(`		return reflect.Zero(reflect.TypeOf(reflect.`, strings.Title(k.Type), `)).Interface().(`, k.Type, `)`)
					c.P(`	}`)
					c.P(`	return h.`, k.Name)
					c.P(`}(from.`, TrimFirstRune(result.Val), `),`)

				} else {
					c.P(`if !isNil(from`, k.Val, `){`)
					c.P(path+k.Name, `= from`, k.Val)
					c.P(`}`)
				}
			}
		}

	} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
		//
		for key, mp1Val := range mp {
			if checkInner {
				path = key.Name
				c.P(path, `: &`, key.Type, `{`)
			} else {
				path = path + key.Name
				c.P(`if !isNil(from`, key.Val, `){`)
				c.P(path, `= &`, (key.Type), `{`)

			}

			checkInner = true
			for _, h := range mp1Val {
				c.generateStruct(h, path+".", checkInner, false, fromArgString)
			}
			if checkOuter {
				c.P(`}`)
				c.P(`}`)
			} else {
				c.P(`},`)
			}
			checkOuter = true
			path = "to."
		}

	}

}

func (c *copy) msgToString(fromArg []interface{}) []interface{} {
	args := []interface{}{}
	// args = append(args, k.Val)
	for _, v := range fromArg {
		if k, ok := v.(ObjQueue); ok {
			// if !k.IsEnum {
			args = append(args, k.Val)
			// } else {
			// 	valMaping := []interface{}{
			// 		k.Val,
			// 	}
			// 	glog.Infoln("valMaping", valMaping)
			// 	valMaping = append([]interface{}{v}, valMaping...)
			// 	args = append(args, valMaping)
			// }

		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
			for Key, mp1Val := range mp {
				valMaping := c.msgToString(mp1Val)
				valMaping = append([]interface{}{Key}, valMaping...)
				// for _, v := range valMaping {
				args = append(args, valMaping)
				// }
			}
		}
	}
	return args
}
func (c *copy) generateClientMethod(servName, fullServName, serviceDescVar string, method *pb.MethodDescriptorProto) {
	// outType := c.typeName(method.GetOutputType())
	// glog.Infoln("outType", outType)
	// c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, true), "{")

	// c.P(`for _, v := range from {`)
	// c.P(`resp := &`, outType, `{}`)
	// c.P(`c.ListCITransactionsRequest(v, resp)`)
	// c.P(`*to = append(*to, resp)`)
	// c.P(`}`)

	// c.P("return  nil")
	// c.P("}")

	c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, false), "{")

	msgInput := c.ObjectNamed(method.GetInputType()).(*generator.Descriptor).DescriptorProto
	msgOutput := c.ObjectNamed(method.GetOutputType()).(*generator.Descriptor).DescriptorProto
	fromArg := c.getFieldsOfMsg(msgInput, "")
	toArg := c.getFieldsOfMsg(msgOutput, "")

	fromArgString := c.msgToString(fromArg)
	glog.Infoln()
	glog.Infoln()
	glog.Infoln("toArg", toArg)
	glog.Infoln()
	glog.Infoln()
	glog.Infoln("fromArgString", fromArgString)
	for _, v := range toArg {
		c.generateStruct(v, "to.", false, true, fromArgString)
	}
	c.P("return  nil")
	c.P("}")
	c.P()
}

// generateService generates all the code for the named service.
func (c *copy) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {

	origServName := service.GetName()
	fullServName := origServName
	if pkg := file.GetPackage(); pkg != "" {
		fullServName = pkg + "." + fullServName
	}
	servName := generator.CamelCase(origServName)

	c.P("type ", unexport(servName), " struct {")
	c.P("}")
	c.P()

	c.P("func New", servName, "() *", unexport(servName), " {")
	c.P("return &", unexport(servName), "{}")
	c.P("}")
	c.P()

	serviceDescVar := "_" + servName + "_serviceDesc"

	for _, method := range service.Method {

		c.generateClientMethod(servName, fullServName, serviceDescVar, method)
	}

}

func contains(s []ObjQueue, e string) bool {
	for _, k := range s {
		if k.Name == e {
			return true
		}
	}
	return false
}

func (c *copy) getFieldCpIfAny(field *descriptor.MethodDescriptorProto) *cpier.CpRule {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, cpier.E_Cp)
		if err == nil && v.(*cpier.CpRule) != nil {
			return (v.(*cpier.CpRule))
		}
	}
	return nil
}
