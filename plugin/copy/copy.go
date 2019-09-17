package copy

import (
	"fmt"
	"strings"

	"github.com/golang/glog"

	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	copier "github.com/tvducmt/protoc-gen-copy/protobuf"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

const (
	contextPkgPath = "context"
	grpcPkgPath    = "google.golang.org/grpc"
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
	grpcPkg    string
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
	grpcPkg = string(c.Generator.AddImport(grpcPkgPath))

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
func (c *copy) generateClientSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	reqArg := ", in *" + c.typeName(method.GetInputType())
	respName := "*" + c.typeName(method.GetOutputType())
	return fmt.Sprintf("%s(ctx %s.Context%s, opts ...%s.CallOption) (%s, error)", methName, contextPkg, reqArg, grpcPkg, respName)
}
func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

func (c *copy) getFieldsOfMsg(msg *descriptor.DescriptorProto, parent string) []interface{} {
	args := []interface{}{}
	for _, v := range msg.Field {
		if v.IsMessage() {
			mpMsg := make(map[string][]interface{})
			msg1 := c.ObjectNamed(v.GetTypeName()).(*generator.Descriptor).DescriptorProto
			// parent = parent + "." + strings.Title(v.GetJsonName())
			x := c.getFieldsOfMsg(msg1, parent)
			mpMsg[strings.Title(v.GetJsonName())] = x
			args = append(args, mpMsg)
		} else {
			args = append(args, strings.Title(v.GetJsonName()))
		}
	}
	return args
}

// func (c *copy) compareMap(mp map[string][]interface{}, toArg []interface{}, parent string) {
// 	for key, mp1Val := range mp {
// 		for _, v := range toArg {
// 			if mp2, oki := v.(map[string][]interface{}); oki {
// 				glog.Infoln("keykey", key)
// 				if mp2Val, ok := mp2[key]; ok {
// 					glog.Infoln("keykey1", key)
// 					parent = parent + key + "."
// 					for _, v := range mp1Val {
// 						if k, ok := v.(string); ok {
// 							if contains(mp2Val, k) {
// 								c.P(`out.`, parent, v, ` = in.`, parent, v)
// 							}
// 						} else if mp, oki := v.(map[string][]interface{}); oki {
// 							glog.Infoln("sdfs_mp", mp)
// 							c.compareMap(mp, toArg, parent)
// 						}
// 					}

// 				}
// 			}
// 		}
// 	}
// }
func (c *copy) compareTwoArrString(fromArg, toArg []string) {
	for _, v := range fromArg {
		if contains(toArg, v) {
			c.P(`out.`, v, ` = in.`, v)

		}
	}
}
func (c *copy) getArrString(input []interface{}, path string) []string {
	args := []string{}
	for _, v := range input {
		if k, ok := v.(string); ok {
			x := path + k
			args = append(args, x)
		} else if mp, oki := v.(map[string][]interface{}); oki {
			for key, mp1Val := range mp {
				path = path + key + "."
				h := c.getArrString(mp1Val, path)
				args = append(args, h...)
			}
		}
	}
	return args
}

func (c *copy) generateClientMethod(servName, fullServName, serviceDescVar string, method *pb.MethodDescriptorProto, descExpr string) {
	outType := c.typeName(method.GetOutputType())
	c.P("func (c *", unexport(servName), ") ", c.generateClientSignature(servName, method), "{")
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		c.P("out := new(", outType, ")")
		msgInput := c.ObjectNamed(method.GetInputType()).(*generator.Descriptor).DescriptorProto
		msgOutput := c.ObjectNamed(method.GetOutputType()).(*generator.Descriptor).DescriptorProto
		// glog.Infoln("msgOutput", msgOutput)
		fromArg := c.getFieldsOfMsg(msgInput, "")
		toArg := c.getFieldsOfMsg(msgOutput, "")
		glog.Infoln("fromArg", fromArg)
		// glog.Infoln("toArg", toArg)
		fromArrString := c.getArrString(fromArg, "")
		toArrString := c.getArrString(toArg, "")
		glog.Infoln("fromArrString", fromArrString)
		glog.Infoln("toArrString", toArrString)
		c.compareTwoArrString(fromArrString, toArrString)
		// c.compareTwoStruct(fromArg, toArg)

		c.P("return out, nil")
		c.P("}")
		c.P()
		return
	}
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
	// deprecated := service.GetOptions().GetDeprecated()
	// glog.Infoln("servName", servName)
	// glog.Infoln("deprecated", deprecated)
	// 	c.P()
	// 	c.P(fmt.Sprintf(`// %sClient is the client API for %s service.
	// //
	// // For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.`, servName, servName))

	// 	// Client interface.
	// 	if deprecated {
	// 		c.P("//")
	// 		c.P(deprecationComment)
	// 	}
	// c.P("type ", servName, "Client interface {")
	// for i, method := range service.Method {
	// 	c.Generator.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
	// 	c.P(c.generateClientSignature(servName, method))
	// }
	// c.P("}")
	// c.P()

	// Client structure.
	c.P("type ", unexport(servName), " struct {")
	// c.P("cc *", grpcPkg, ".ClientConn")
	c.P("}")
	c.P()

	// 	// NewClient factory.
	// 	if deprecated {
	// 		c.P(deprecationComment)
	// 	}
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
			descExpr = fmt.Sprintf("&%s.Streams[%d]", serviceDescVar, streamIndex)
			streamIndex++
		}
		c.generateClientMethod(servName, fullServName, serviceDescVar, method, descExpr)
	}

}

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}

func contains(s []string, e string) bool {
	for _, k := range s {
		if k == e {
			return true
		}
	}
	return false
}

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {

	// ccTypeName := generator.CamelCaseSlice(message.TypeName())

	mapMsg := make(map[string][]string)
	args := []string{}
	for _, field := range message.Field {
		if !field.IsMessage() {
			return
		}
		fieldTypeName := field.GetTypeName()
		if strings.HasPrefix(fieldTypeName, ".") {
			// arr := strings.Split(fieldTypeName, ".")
			// glog.Infoln("arr[len(arr)-1]", arr[len(arr)-1])
			fieldCopy := c.getFieldCpIfAny(field)
			if fieldCopy != nil {
				if fieldCopy.GetEnableCpProto() != "" {
					// toArg = fieldCopy.GetEnableCpProto()
					msg := c.ObjectNamed(field.GetTypeName()).(*generator.Descriptor).DescriptorProto
					for _, v := range msg.Field {
						if _, ok := mapMsg[fieldTypeName]; ok {
							mapMsg[fieldTypeName] = append(mapMsg[fieldTypeName], strings.Title(v.GetJsonName()))
						} else {
							arr := []string{strings.Title(v.GetJsonName())}
							mapMsg[fieldTypeName] = arr
							args = append(args, fieldTypeName)
						}
					}
				}
			}
		} else {
		}
	}

	// toArg := []string{}
	// fromArg := []string{}

	// intoCheck := true
	// if len(mapMsg) == 2 {
	// 	for _, v := range mapMsg {
	// 		if intoCheck {
	// 			fromArg = v
	// 			intoCheck = false

	// 		} else {
	// 			toArg = v
	// 			// keys = append(keys, k)
	// 		}
	// 	}

	// }

	// c.P(`func (this *`, ccTypeName, `) Copy(to *`, trimFirstRune(args[1]), `, from *`, trimFirstRune(args[0]), ` ) {`) //, copier.CopyProto{}.A,from *`, trimFirstRune(field.GetTypeName()), `
	// c.P(c.flagPkg.Use(), `.Parse()`)
	// for _, v := range fromArg {
	// 	if contains(toArg, v) {
	// 		c.P(`to.`, v, ` = from.`, v)
	// 	}
	// }
	// c.P(`}`)
}

func (c *copy) getFieldCpIfAny(field *descriptor.FieldDescriptorProto) *copier.CopyProto {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, copier.E_Field)
		if err == nil && v.(*copier.CopyProto) != nil {
			return (v.(*copier.CopyProto))
		}
	}
	return nil
}
