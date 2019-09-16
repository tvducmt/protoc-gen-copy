package copy

import (

	//copier "github.com/tvducmt/protoc-gen-copy/protobuf"

	"strings"

	copier "github.com/tvducmt/protoc-gen-copy/protobuf"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

const (
	protobufPkg = "github.com/tvducmt/protoc-gen-copy/protobuf"
)

type copy struct {
	*generator.Generator
	generator.PluginImports
	// querierPkg generator.Single
	glogPkg    generator.Single
	protoPkg   generator.Single
	elasticPkg generator.Single
	reflectPkg generator.Single
	timePkg    generator.Single
	flagPkg    generator.Single
	stringsPkg generator.Single
	// testPkg    generator.Single
}

// NewCopy ...
func NewCopy() generator.Plugin {
	return &copy{
		// query: query,
	}
}

func (c *copy) Name() string {
	return "copy"
}

func (c *copy) Init(g *generator.Generator) {
	c.Generator = g
}

func (c *copy) Generate(file *generator.FileDescriptor) {
	c.PluginImports = generator.NewPluginImports(c.Generator)

	c.glogPkg = c.NewImport("githuc.com/golang/glog")
	c.stringsPkg = c.NewImport("strings")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	c.elasticPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/search/elastic")
	c.reflectPkg = c.NewImport("reflect")
	c.timePkg = c.NewImport("time")
	c.flagPkg = c.NewImport("flag")

	protobufPkg = string(c.gen.AddImport(protobufPkgPath))

	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ ", protobufPkg, ".ClientConn")
	g.P()
	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			c.generateProto3Message(file, msg)
		}
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
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {

	ccTypeName := generator.CamelCaseSlice(message.TypeName())

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

	toArg := []string{}
	fromArg := []string{}

	intoCheck := true
	if len(mapMsg) == 2 {
		for _, v := range mapMsg {
			if intoCheck {
				fromArg = v
				intoCheck = false

			} else {
				toArg = v
				// keys = append(keys, k)
			}
		}

	}

	c.P(`func (this *`, ccTypeName, `) Copy(to *`, trimFirstRune(args[1]), `, from *`, trimFirstRune(args[0]), ` ) {`) //, copier.CopyProto{}.A,from *`, trimFirstRune(field.GetTypeName()), `
	c.P(c.flagPkg.Use(), `.Parse()`)
	for _, v := range fromArg {
		if contains(toArg, v) {
			c.P(`to.`, v, ` = from.`, v)
		}
	}
	c.P(`}`)
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
