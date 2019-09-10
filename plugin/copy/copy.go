package copy

import (

	// querier "git.zapa.cloud/merchant-tools/helper/protoc-gen-buildquery/protobuf"

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
	// proto3 := gogoproto.IsProto3(file.FileDescriptorProto)
	c.PluginImports = generator.NewPluginImports(c.Generator)

	c.glogPkg = c.NewImport("githuc.com/golang/glog")
	c.stringsPkg = c.NewImport("strings")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	c.elasticPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/search/elastic")
	c.reflectPkg = c.NewImport("reflect")
	c.timePkg = c.NewImport("time")
	c.flagPkg = c.NewImport("flag")

	for _, msg := range file.Messages() {
		// if msg.DescriptorProto.GetOptions().GetMapEntry() {
		// 	continue
		// }
		// c.generateRegexVars(file, msg)
		// if gogoproto.IsProto3(file.FileDescriptorProto) {
		c.generateProto3Message(file, msg)
		// }
	}
}

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	c.P(`func (this *`, ccTypeName, `) Copy(resp interface{}) {`)
	c.In()
	c.P(c.flagPkg.Use(), `.Parse()`)

	c.P(`reqVal := reflect.Indirect(reflect.ValueOf(l))`)
	c.P(`reqType := indirectType(reqVal.Type())`)
	c.P(`reqCoreServiceVal := reflect.Indirect(reflect.ValueOf(resp))`)
	c.P(`for i := 0; i < reqType.NumField(); i++ {`)
	c.P(`v := reqType.Field(i)`)
	c.P(`fmt.Println(":reqCoreServiceVal.Type()", reqCoreServiceVal.Type().Name())`)
	c.P(`if !strings.HasPrefix(v.Name, "XXX_") {`)
	c.P(`if f, ok := reqCoreServiceVal.Type().FieldByName(v.Name); ok {`)
	// reqCoreService.ZpTransID = l.GetZpTransID()
	c.P(`ListCITransactionsSvcRequest.ZpTransID =`)
	c.P(`fmt.Println("into here  v.Index", v.Name)`)
	c.P(`fmt.Println("into here  f.Index", f.Name)`)
	// fmt.Println("into here  v.Index", v.Index)
	// fields = append(fields, &searchField{indexFrom: v.Index, indexTo: f.Index})
	c.P(`}`)

	c.P(`}`)
	c.P(`}`)

	c.Out()
	c.P(`}`)
}
