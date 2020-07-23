package plugin

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/golang/glog"
)

const (
	contextPkgPath   = "context"
	reflectPkgPath   = "reflect"
	timeStampPkgPath = "github.com/golang/protobuf/ptypes/timestamp"
)

type copy struct {
	*generator.Generator
	generator.PluginImports
	protoPkg generator.Single
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
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
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

func (c *copy) generateClientSignature(servName string, method *pb.MethodDescriptorProto, isCpSlice bool) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	var reqArg string
	var respName string
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
	Name       string
	Type       string
	Val        string
	IsEnum     bool
	IsRepeated bool
	IsObject   bool
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
func (c *copy) covertFieldsToMap(msg *descriptor.DescriptorProto, path string) []interface{} {
	args := []interface{}{}
	for _, v := range msg.Field {
		isRepeated := false
		if v.IsMessage() {
			if v.IsRepeated() {
				isRepeated = true
			}
			mpMsg := make(map[ObjQueue][]interface{})
			msg1 := c.ObjectNamed(v.GetTypeName()).(*generator.Descriptor).DescriptorProto
			path = path + "." + strings.Title(v.GetJsonName())
			x := c.covertFieldsToMap(msg1, path)
			typeNested, _ := c.getNestedType(TrimFirstRune(v.GetTypeName()))
			mpMsg[ObjQueue{Name: strings.Title(v.GetJsonName()), Type: typeNested, Val: path, IsRepeated: isRepeated, IsObject: true}] = x
			args = append(args, mpMsg)
			last := path[strings.LastIndex(path, ".")+1:]
			path = path[:(len(path) - len(last) - 1)]
		} else if v.IsEnum() {
			typeNested, _ := c.getNestedType(TrimFirstRune(v.GetTypeName()))
			args = append(args, ObjQueue{Name: strings.Title(v.GetJsonName()), Type: typeNested, Val: path + "." + strings.Title(v.GetJsonName()), IsEnum: true})
		} else {
			if v.IsRepeated() {
				isRepeated = true
			}
			primitiveType := getTypeOfFied(v)
			args = append(args, ObjQueue{Name: strings.Title(v.GetJsonName()), Type: primitiveType, Val: path + "." + strings.Title(v.GetJsonName()), IsRepeated: isRepeated})
		}
	}
	return args
}

func (c *copy) generateCpFunc(v interface{}, path string, checkInner bool, checkOuter bool, fromArgString []interface{}, parentRepeated bool) {

	if k, ok := v.(ObjQueue); ok {
		if k.IsEnum {
			result := ExistInFromArrString(k.Val, fromArgString)
			if result.Name != "" {
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
		} else if k.IsRepeated {

			result := ExistInFromArrString(k.Val, fromArgString)
			if result.Name != "" {

				if checkInner {
					if k.IsObject {
						c.P(result.Name+`tmp := make([]*`, k.Type, `, len(from.`, TrimFirstRune(result.Val), `))`)
						c.P(`for i, v := range `, result.Type, ` {`)
						c.P(result.Name, `tmp[i] = &`, k.Type, `{`)
						c.P(k.Name, `: func(h *`, result.Type, `) `, k.Type, ` {`) //  `: from`, k.Val, `,`
						c.P(`	if h == nil {`)
						c.P(`		return reflect.Zero(reflect.TypeOf(reflect.`, strings.Title(k.Type), `)).Interface().(`, k.Type, `)`)
						c.P(`	}`)
						c.P(`	return h.`, k.Name)
						c.P(`}(v)`)
					} else {
						c.P(k.Name, `: func(h *`, result.Type, `) []`, k.Type, ` {`) //  `: from`, k.Val, `,`
						c.P(`	if h == nil {`)
						c.P(`		return nil`)
						c.P(`	}`)
						c.P(`	return h.`, k.Name)
						c.P(`}(v),`)
					}

				} else {
					c.P(`if !isNil(from`, k.Val, `){`)
					c.P(path+k.Name, `= from`, k.Val)
					c.P(`}`)
				}
			}
		} else {
			result := ExistInFromArrString(k.Val, fromArgString)
			glog.Infoln("Into here 1", k)
			if result.Name != "" {
				if checkInner {
					if parentRepeated {
						c.P(k.Name, `: func(h *`, result.Type, `) `, k.Type, ` {`) //  `: from`, k.Val, `,`
						c.P(`	if h == nil {`)
						c.P(`		return reflect.Zero(reflect.TypeOf(reflect.`, strings.Title(k.Type), `)).Interface().(`, k.Type, `)`)
						c.P(`	}`)
						c.P(`	return h.`, k.Name)
						c.P(`}(from.` + TrimFirstRune(result.Val) + `[i]),`)
					} else {
						c.P(k.Name, `: func(h *`, result.Type, `) `, k.Type, ` {`) //  `: from`, k.Val, `,`
						c.P(`	if h == nil {`)
						c.P(`		return reflect.Zero(reflect.TypeOf(reflect.`, strings.Title(k.Type), `)).Interface().(`, k.Type, `)`)
						c.P(`	}`)
						c.P(`	return h.`, k.Name)
						c.P(`}(from.`, TrimFirstRune(result.Val), `),`)
					}
				} else {
					c.P(`if !isNil(from`, k.Val, `){`)
					c.P(path+k.Name, `= from`, k.Val)
					c.P(`}`)
				}
			}
		}

	} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
		for key, mp1Val := range mp {
			if checkInner {
				if key.IsRepeated {
					glog.Infoln("Into here 0", key)
					c.P(`if !isNil(from`, key.Val, `){`)
					c.P(`tmp := make([]*`, key.Type, `, len(from`, key.Val, `))`)
					c.P(`for i, v := range from`, key.Val, ` {`)
					c.P(`glog.Infoln("v",v)`)
					c.P(`tmp[i] = &`, key.Type, `{`)
				} else {
					path = key.Name
					c.P(path, `: &`, key.Type, `{`)
				}
			} else {
				path = path + key.Name
				if key.IsRepeated {
					c.P(`if !isNil(from`, key.Val, `){`)
					c.P(`tmp := make([]*`, key.Type, `, len(from`, key.Val, `))`)
					c.P(`for i, v := range from`, key.Val, ` {`)
					c.P(`glog.Infoln("v",v)`)
					c.P(`tmp[i] = &`, key.Type, `{`)
				} else {
					c.P(`if !isNil(from`, key.Val, `){`)
					c.P(path, `= &`, (key.Type), `{`)
				}

			}

			checkInner = true
			var parentRepeated bool
			if key.IsRepeated {
				parentRepeated = true
			}
			for _, h := range mp1Val {

				c.generateCpFunc(h, path+".", checkInner, false, fromArgString, parentRepeated)
			}
			if checkOuter {
				c.P(`}`)
				c.P(`}`)
				if key.IsRepeated {
					c.P(`to`, key.Val, `= tmp`)
					c.P(`}`)
				}
			} else {
				c.P(`},`)
			}
			checkOuter = true
			path = "to."
		}

	}

}

func (c *copy) covertMapToArrString(fromArg []interface{}) []interface{} {
	args := []interface{}{}
	for _, v := range fromArg {
		if k, ok := v.(ObjQueue); ok {
			args = append(args, k.Val)
		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
			for Key, mp1Val := range mp {
				valMaping := c.covertMapToArrString(mp1Val)
				valMaping = append([]interface{}{Key}, valMaping...)
				args = append(args, valMaping)
			}
		}
	}
	return args
}
func (c *copy) generateClientMethod(servName, fullServName, serviceDescVar string, method *pb.MethodDescriptorProto) {
	outType := c.typeName(method.GetOutputType())
	c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, true), "{")

	c.P(`for _, v := range from {`)
	c.P(`resp := &`, outType, `{}`)
	c.P(`c.`, method.GetName(), `(v, resp)`)
	c.P(`*to = append(*to, resp)`)
	c.P(`}`)

	c.P("return  nil")
	c.P("}")

	c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, false), "{")

	msgInput := c.ObjectNamed(method.GetInputType()).(*generator.Descriptor).DescriptorProto
	msgOutput := c.ObjectNamed(method.GetOutputType()).(*generator.Descriptor).DescriptorProto
	fromMap := c.covertFieldsToMap(msgInput, "")
	toMap := c.covertFieldsToMap(msgOutput, "")
	glog.Infoln("fromMap", fromMap)
	glog.Infoln("toMap", toMap)
	fromArrString := c.covertMapToArrString(fromMap)
	glog.Infoln("fromArrString", fromArrString)
	for _, v := range toMap {
		c.generateCpFunc(v, "to.", false, true, fromArrString, false)
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
