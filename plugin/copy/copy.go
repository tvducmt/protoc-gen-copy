package copy

import (
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

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {

	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	c.P(`func (this *`, ccTypeName, `) Copy(`, copier.CopyProto{}.A, `) {`)
	c.In()
	c.P(c.flagPkg.Use(), `.Parse()`)
	// c.P(`esMap := map[string]interface{}{}`)

	for _, field := range message.Field {
		fieldCopy := c.getFieldQueryIfAny(field)
		if fieldCopy == nil {
			continue
		}
		fieldName := c.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName
		c.P(` x:= `, variableName)

		// if field.IsMessage() {
		// 	c.generatePtrAndStructEs(variableName, ccTypeName, fieldName, fieldEs)
		// } else {
		// 	c.generateFieldEs(variableName, ccTypeName, fieldEs)
		// }

	}
	// c.P(`}`)
}

func (c *copy) getFieldQueryIfAny(field *descriptor.FieldDescriptorProto) *copier.CopyProto {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, copier.E_Field)
		if err == nil && v.(*copier.CopyProto) != nil {
			return (v.(*copier.CopyProto))
		}
	}
	return nil
}
