package copy

import (
	"strings"

	"github.com/golang/glog"
	//copier "github.com/tvducmt/protoc-gen-copy/protobuf"

	copier "github.com/tvducmt/protoc-gen-copy/protobuf"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
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

type Object struct {
	Fields []string
	Name   string
}

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {

	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	fieldsA := []Object{}

	// listFieldTo :=
	for _, field := range message.Field {
		if !field.IsMessage() {
			return
		}
		fieldCopy := c.getFieldCpIfAny(field)
		if fieldCopy != nil {
			if fieldCopy.GetEnableCpProto() != "" {
				obj := Object{
					Fields: strings.Split(fieldCopy.GetEnableCpProto(), ";"),
					Name:   field.GetTypeName(),
				}
				fieldsA = append(fieldsA, obj)
			}
		}
	}
	if len(fieldsA) < 2 {
		glog.Errorln("len < 2")
	}
	c.P(`func (this *`, ccTypeName, `) Copy(to *`, trimFirstRune(fieldsA[1].Name), `, from *`, trimFirstRune(fieldsA[0].Name), ` ) {`) //, copier.CopyProto{}.A,from *`, trimFirstRune(field.GetTypeName()), `
	c.P(c.flagPkg.Use(), `.Parse()`)
	for _, v := range fieldsA[0].Fields {
		if contains(fieldsA[1].Fields, v) {
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
