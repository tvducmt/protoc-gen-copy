package copy

import (
	"fmt"
	"strings"

	"github.com/golang/glog"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	cpier "github.com/tvducmt/protoc-gen-copy/protobuf"
)

const (
	contextPkgPath = "context"
)

type copy struct {
	*generator.Generator
	generator.PluginImports
	glogPkg    generator.Single
	protoPkg   generator.Single
	reflectPkg generator.Single
	timePkg    generator.Single
	flagPkg    generator.Single
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
	contextPkg string
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

	c.glogPkg = c.NewImport("githuc.com/golang/glog")
	c.stringsPkg = c.NewImport("strings")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	c.reflectPkg = c.NewImport("reflect")
	c.timePkg = c.NewImport("time")
	c.flagPkg = c.NewImport("flag")

	contextPkg = string(c.Generator.AddImport(contextPkgPath))

	c.P("// Reference imports to suppress errors if they are not otherwise used.")
	c.P("var _ ", contextPkg, ".Context")
	// c.P("var _ ", grpcPkg, ".ClientConn")
	c.P()

	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	for i, service := range file.FileDescriptorProto.Service { //Messages()
		c.generateService(file, service, i)
	}
}

// // generateClientSignature returns the client-side signature for a method.
func (c *copy) generateClientSignature(servName string, method *pb.MethodDescriptorProto, typeMethod bool) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	var reqArg string   //:= "from *" + c.typeName(method.GetInputType())
	var respName string //:= "to *" + c.typeName(method.GetOutputType())
	if typeMethod {
		methName = methName + "Slice"
		reqArg = "from []*" + c.typeName(method.GetInputType())
		respName = "to []*" + c.typeName(method.GetOutputType())
	} else {
		reqArg = "from *" + c.typeName(method.GetInputType())
		respName = "to *" + c.typeName(method.GetOutputType())
	}
	return fmt.Sprintf("%s(%s,%s) error", methName, reqArg, respName)
}
func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// ObjQueue ...
type ObjQueue struct {
	Name string
	Type string
	Val  string
}

func (c *copy) getFieldsOfMsg(msg *descriptor.DescriptorProto, path string) []interface{} {
	args := []interface{}{}
	for _, v := range msg.Field {
		if v.IsMessage() {
			// glog.Infoln("v.GetTypeName()", v.GetTypeName())
			mpMsg := make(map[ObjQueue][]interface{})
			msg1 := c.ObjectNamed(v.GetTypeName()).(*generator.Descriptor).DescriptorProto
			// parent = parent + "." + strings.Title(v.GetJsonName())
			path = path + "." + strings.Title(v.GetJsonName())
			x := c.getFieldsOfMsg(msg1, path)
			mpMsg[ObjQueue{Name: strings.Title(v.GetJsonName()), Type: v.GetTypeName(), Val: path}] = x
			args = append(args, mpMsg)
		} else {
			args = append(args, ObjQueue{Name: strings.Title(v.GetJsonName()), Type: "PrimitiveType", Val: path + "." + strings.Title(v.GetJsonName())})
		}
	}
	return args
}

func existInFromArr(input string, fromArg []interface{}) bool {
	for _, v := range fromArg {
		if _, ok := v.(string); ok {
			if input == v {
				return true
			}
		}
	}
	return false
}
func (c *copy) getArrObjQueue(toArg []interface{}, path string, checkInner bool, checkOuter bool, fromArgString []interface{}) []ObjQueue {
	args := []ObjQueue{}
	for _, v := range toArg {
		if k, ok := v.(ObjQueue); ok {
			if existInFromArr(k.Val, fromArgString) {
				if checkInner {
					c.P(k.Name, `: from`, k.Val, `,`)
				} else {
					c.P(path+k.Name, `= from`, k.Val)
				}
			}
			// args = append(args, ObjQueue{Name: k.Name, Type: k.Type})
		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
			for key, mp1Val := range mp {
				if checkInner {
					path = key.Name
					c.P(path, `: &`, trimFirstRune(key.Type), `{`)
				} else {
					path = path + key.Name
					c.P(path, `= &`, trimFirstRune(key.Type), `{`)
				}

				checkInner = true
				c.getArrObjQueue(mp1Val, path+".", checkInner, false, fromArgString)

				if checkOuter {
					c.P(`}`)
					checkOuter = false
				} else {
					c.P(`},`)
				}
			}
		}
	}
	return args
}

// func (c *copy) combineTwoArg(fromArg, toArg []interface{}) []interface{} {
// 	args := []interface{}{}
// 	for _, v := range toArg {
// 		if k, ok := v.(ObjQueue); ok {
// 			if existInFromArg(k, fromArg) {
// 				args = append(args, ObjQueue{Name: k.Name, Type: k.Type})
// 			}
// 		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
// 			for key, mp1Val := range mp {
// 				mpMsg := make(map[ObjQueue][]interface{})
// 				x := c.combineTwoArg(fromArg, mp1Val)
// 				mpMsg[ObjQueue{Name: key.Name, Type: key.Type}] = x
// 				args = append(args, mpMsg)
// 			}
// 		}
// 	}
// 	return args
// }

func (c *copy) msgToString(fromArg []interface{}) []interface{} {
	args := []interface{}{}
	for _, v := range fromArg {
		if k, ok := v.(ObjQueue); ok {
			args = append(args, k.Val)

		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
			for _, mp1Val := range mp {
				valMaping := c.msgToString(mp1Val)
				for _, v := range valMaping {
					args = append(args, v)
				}
			}
		}
	}
	return args
}
func (c *copy) generateClientMethod(servName, fullServName, serviceDescVar string, method *pb.MethodDescriptorProto, descExpr string) {
	// outType := c.typeName(method.GetOutputType())

	// c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, true), "{")
	// c.P(`JSONCopy(to, from)`)
	// c.P("return  nil")
	// c.P("}")

	c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method, false), "{")
	// c.P(`JSONCopy(to, from)`)

	// c.P("out := new(", outType, ")")
	msgInput := c.ObjectNamed(method.GetInputType()).(*generator.Descriptor).DescriptorProto
	msgOutput := c.ObjectNamed(method.GetOutputType()).(*generator.Descriptor).DescriptorProto
	// glog.Infoln("msgOutput", msgOutput)
	fromArg := c.getFieldsOfMsg(msgInput, "")
	toArg := c.getFieldsOfMsg(msgOutput, "")

	fromArgString := c.msgToString(fromArg)
	toArgString := c.msgToString(toArg)
	glog.Infoln("fromArgString", fromArgString)
	glog.Infoln("toArgString", toArgString)
	glog.Infoln()
	glog.Infoln("fromArg", fromArg)
	glog.Infoln()
	glog.Infoln()
	glog.Infoln("toArg", toArg)
	glog.Infoln()
	glog.Infoln()
	// c.combineTwoArg(fromArgString, toArg)
	// glog.Infoln("toArg", toArg)

	c.getArrObjQueue(toArg, "to.", false, true, fromArgString)
	// fromArrString := c.getArrObjQueue(fromArg, "")
	// toArrString := c.getArrObjQueue(toArg, "")
	// c.funcTestCompareMapObj(fromArg, toArg, "")

	// glog.Infoln("fromArrS", fromArrString)
	// // glog.Infoln("getArrObjQueue", toArrString)
	// c.compareTwoArrString(fromArrString, toArrString)
	// // c.compareTwoStruct(fromArg, toArg)

	c.P("return  nil")
	c.P("}")
	c.P()
	return
}

// generateService generates all the code for the named service.
func (c *copy) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	// path := fmt.Sprintf("6,%d", index) // 6 means service.
	// glog.Infoln("path", path)
	origServName := service.GetName()
	// glog.Infoln("origServName", origServName)
	fullServName := origServName
	// glog.Infoln("fullServName", origServName)
	if pkg := file.GetPackage(); pkg != "" {
		fullServName = pkg + "." + fullServName
	}
	servName := generator.CamelCase(origServName)

	c.P("type ", unexport(servName), " struct {")
	// c.P("cc *", grpcPkg, ".ClientConn")
	c.P("}")
	c.P()

	c.P("func New", servName, "() *", unexport(servName), " {")
	c.P("return &", unexport(servName), "{}")
	c.P("}")
	c.P()

	var methodIndex, streamIndex int
	serviceDescVar := "_" + servName + "_serviceDesc"
	// Client method implementations.
	for _, method := range service.Method {
		var descExpr string
		if !method.GetServerStreaming() && !method.GetClientStreaming() {
			// Unary RPC method
			descExpr = fmt.Sprintf("&%s.Methods[%d]", serviceDescVar, methodIndex)
			methodIndex++
		} else {
			// Streaming RPC method
			descExpr = fmt.Sprintf("&%s.Streams[%d]", serviceDescVar, methodIndex)
			streamIndex++
		}
		c.generateClientMethod(servName, fullServName, serviceDescVar, method, descExpr)
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
func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}
func checkExistInFromArgWth(f ObjQueue, fromArg []interface{}) bool {
	for _, v := range fromArg {
		if k, ok := v.(ObjQueue); ok {
			glog.Infoln("k.Name", f.Name, k.Name, (k.Name == f.Name))
			if k.Name == f.Name {
				return true
			}
		} else if mp, oki := v.(map[ObjQueue][]interface{}); oki {
			for key, mp1Val := range mp {
				if key.Name == f.Name {
					checkExistInFromArgWth(f, mp1Val)
					// return true
				}
				// checkExistInFromArgWth(f, mp1Val)
			}
		}
	}
	return false
}
