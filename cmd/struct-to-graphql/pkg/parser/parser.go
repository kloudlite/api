package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/sanity-io/litter"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/k8s"
)

type Parser interface {
	GenerateGraphQLSchema(name string, t reflect.Type)
	Debug()
}

type GraphqlType string

const (
	Type  GraphqlType = "type"
	Input GraphqlType = "input"
	Enum  GraphqlType = "enum"
)

var scalarMappings = map[reflect.Type]string{
	reflect.TypeOf(metav1.Time{}):      "Date",
	reflect.TypeOf(&metav1.Time{}):     "Date",
	reflect.TypeOf(time.Time{}):        "Date",
	reflect.TypeOf(&time.Time{}):       "Date",
	reflect.TypeOf(json.RawMessage{}):  "Any",
	reflect.TypeOf(&json.RawMessage{}): "Any",
}

var kindMap = map[reflect.Kind]string{
	reflect.Int:   "Int",
	reflect.Int8:  "Int",
	reflect.Int16: "Int",
	reflect.Int32: "Int",
	reflect.Int64: "Int",

	reflect.Uint:   "Int",
	reflect.Uint8:  "Int",
	reflect.Uint16: "Int",
	reflect.Uint32: "Int",
	reflect.Uint64: "Int",

	reflect.Float32: "Float",
	reflect.Float64: "Float",

	reflect.Bool:      "Boolean",
	reflect.Interface: "Any",
	// reflect.String: "String",
}

type parser struct {
	Types  map[string][]string
	Inputs map[string][]string
	Enums  map[string][]string

	kCli k8s.ExtendedK8sClient

	CommonTypes   map[string][]string
	CommonInpuuts map[string][]string
	CommonEnums   map[string][]string
}

type JsonTag struct {
	Value     string
	OmitEmpty bool
	Inline    bool
}

func parseJsonTag(field reflect.StructField) JsonTag {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return JsonTag{Value: field.Name, OmitEmpty: false, Inline: false}
	}

	var jt JsonTag
	sp := strings.Split(jsonTag, ",")
	jt.Value = sp[0]

	if jt.Value == "" {
		jt.Value = field.Name
	}

	for i := 1; i < len(sp); i++ {
		if sp[i] == "omitempty" {
			jt.OmitEmpty = true
		}
		if sp[i] == "inline" {
			jt.Inline = true
		}
	}

	return jt
}

type GraphqlTag struct {
	Uri        *string
	Enum       []string
	CommonType *bool
}

func parseGraphqlTag(field reflect.StructField) GraphqlTag {
	tag := field.Tag.Get("graphql")
	if tag == "" {
		return GraphqlTag{}
	}

	var gt GraphqlTag
	sp := strings.Split(tag, ",")
	for i := range sp {
		kv := strings.Split(sp[i], "=")
		if len(kv) != 2 {
			return GraphqlTag{}
		}

		switch kv[0] {
		case "uri":
			{
				gt.Uri = &kv[1]
			}
		case "enum":
			{
				enumVals := strings.Split(kv[1], ";")
				gt.Enum = enumVals
			}
		case "common":
			{
				gt.CommonType = fn.New(kv[1] == "true")
			}
		}
	}

	return gt
}

func toFieldType(fieldType string, isRequired bool) string {
	if isRequired {
		return fieldType + "!"
	}
	return fieldType
}

func (p *parser) mergeParser(other *parser, overKey string) (fields []string, inputFields []string) {
	for k, v := range other.Types {
		if k == overKey {
			fields = append(fields, v...)
			continue
		}
		p.Types[k] = v
	}

	for k, v := range other.Inputs {
		if k == overKey+"In" {
			inputFields = append(inputFields, v...)
			continue
		}
		p.Inputs[k] = v
	}

	for k, v := range other.Enums {
		p.Enums[k] = v
	}

	return fields, inputFields
}

func (p *parser) GenerateGraphQLSchema(name string, t reflect.Type) {
	var fields []string
	var inputFields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		jt := parseJsonTag(field)
		if jt.Value == "-" {
			continue
		}

		fieldType := ""
		inputFieldType := ""

		if scalar, ok := scalarMappings[field.Type]; ok {
			fieldType = toFieldType(scalar, !jt.OmitEmpty)
			inputFieldType = toFieldType(scalar, !jt.OmitEmpty)
		}

		if v, ok := kindMap[field.Type.Kind()]; ok {
			fieldType = toFieldType(v, !jt.OmitEmpty)
			inputFieldType = toFieldType(v, !jt.OmitEmpty)
		}

		gt := parseGraphqlTag(field)

		if fieldType == "" {
			switch field.Type.Kind() {
			case reflect.String:
				{
					if gt.Enum != nil {
						sort.Strings(gt.Enum)
						p.Enums[field.Name] = gt.Enum

						fieldType = toFieldType(field.Name, !jt.OmitEmpty)
						inputFieldType = toFieldType(field.Name, !jt.OmitEmpty)
						break
					}
					fieldType = toFieldType("String", !jt.OmitEmpty)
					inputFieldType = toFieldType("String", !jt.OmitEmpty)
				}
			case reflect.Struct:
				{
					fmt.Println("structName: ", field.Type.String(), field.Type.PkgPath())

					if gt.Uri != nil {
						if strings.HasPrefix(*gt.Uri, "k8s://") {
							k8sCrdName := strings.Split(*gt.Uri, "k8s://")[1]
							jsonSchema, err := p.kCli.GetCRDJsonSchema(context.TODO(), k8sCrdName)
							if err != nil {
								panic(err)
							}

							func() {
								childType := "K8s" + jt.Value
								if jt.Inline {
									// TODO: call json parsing and create type, input, and enum for all those types, using jsonSchema, with fields being inline
									p2 := newParser(p.kCli)
									p2.GenerateFromJsonSchema(childType, jsonSchema)

									fields2, inputFields2 := p.mergeParser(p2, childType)
									fields = append(fields, fields2...)
									inputFields = append(inputFields, inputFields2...)

									return
								}

								fieldType = toFieldType(childType, !jt.OmitEmpty)
								inputFieldType = toFieldType(childType+"In", !jt.OmitEmpty)

								// TODO: call json parsing and create type, input, and enum for all those types, using jsonSchema
								p.GenerateFromJsonSchema(childType, jsonSchema)
							}()
						}
						break
					}

					func() {
						if jt.Inline {
							p2 := newParser(p.kCli)
							childType := name + field.Name
							p2.GenerateGraphQLSchema(childType, field.Type)

							fields2, inputFields2 := p.mergeParser(p2, childType)
							fields = append(fields, fields2...)
							inputFields = append(inputFields, inputFields2...)

							return
						}

						fieldType = toFieldType(name+field.Name, jt.OmitEmpty)
						inputFieldType = toFieldType(name+field.Name+"In", jt.OmitEmpty)
						// p.GenerateGraphQLSchema(name+field.Name, field.Type)
						p.GenerateGraphQLSchema(fieldType, field.Type)
					}()
				}
			case reflect.Slice:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						// TODO: if common, take out the name+ field,
						fieldType = toFieldType(fmt.Sprintf("[%s]", name+field.Type.Elem().Name()), !jt.OmitEmpty)
						inputFieldType = toFieldType(fmt.Sprintf("[%sIn]", name+field.Type.Elem().Name()), !jt.OmitEmpty)
						p.GenerateGraphQLSchema(name+field.Type.Elem().Name(), field.Type.Elem())
						break
					}
					fieldType = toFieldType(fmt.Sprintf("[%s]", kindMap[field.Type.Elem().Kind()]), !jt.OmitEmpty)
					inputFieldType = toFieldType(fmt.Sprintf("[%s]", kindMap[field.Type.Elem().Kind()]), !jt.OmitEmpty)
				}
			case reflect.Ptr:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						fieldType = name + field.Type.Elem().Name()
						inputFieldType = name + field.Type.Elem().Name() + "In"
						p.GenerateGraphQLSchema(name+field.Type.Elem().Name(), field.Type.Elem())
						break
					}
					fieldType = kindMap[field.Type.Elem().Kind()]
					inputFieldType = kindMap[field.Type.Elem().Kind()]
				}
			case reflect.Map:
				{
					fieldType = "Map"
					inputFieldType = "Map"
					if field.Type.Elem().Kind() == reflect.Struct {
						// p.GenerateGraphQLSchema(name+field.Name, field.Type.Elem())
						p.GenerateGraphQLSchema(name+field.Type.Elem().Name(), field.Type.Elem())
					}
				}
			default:
				{
					fmt.Printf("default: name: %v (field-name: %v), type: %v, kind: %v\n", jt.Value, field.Name, field.Type, field.Type.Kind())
				}
			}
		}

		if fieldType != "" {
			fields = append(fields, fmt.Sprintf("%s: %s", jt.Value, fieldType))
			inputFields = append(inputFields, fmt.Sprintf("%s: %s", jt.Value, inputFieldType))
			continue
		}
	}

	p.Types[name] = fields
	p.Inputs[name+"In"] = inputFields
}

func (p *parser) NavigateTree(name string, tree *v1.JSONSchemaProps, depth ...int) {
	currDepth := func() int {
		if len(depth) == 0 {
			return 1
		}
		return depth[0]
	}()

	m := map[string]bool{}
	for i := range tree.Required {
		m[tree.Required[i]] = true
	}

	typeName := genTypeName(name)

	fields := make([]string, 0, len(tree.Properties))
	inputFields := make([]string, 0, len(tree.Properties))

	for k, v := range tree.Properties {
		if currDepth == 1 {
			if k == "apiVersion" || k == "kind" {
				fields = append(fields, genFieldEntry(k, "String!", m[k]))
				inputFields = append(inputFields, genFieldEntry(k, "String!", m[k]))
				continue
			}
		}

		if v.Type == "array" {
			if v.Items.Schema != nil && v.Items.Schema.Type == "object" {
				fields = append(fields, genFieldEntry(k, fmt.Sprintf("[%s]", typeName+genTypeName(k)), m[k]))
				inputFields = append(inputFields, genFieldEntry(k, fmt.Sprintf("[%sIn]", typeName+genTypeName(k)), m[k]))
				p.NavigateTree(typeName+genTypeName(k), v.Items.Schema, currDepth+1)
				continue
			}

			fields = append(fields, genFieldEntry(k, fmt.Sprintf("[%s]", genTypeName(v.Items.Schema.Type)), m[k]))
			inputFields = append(inputFields, genFieldEntry(k, fmt.Sprintf("[%s]", genTypeName(v.Items.Schema.Type)), m[k]))
			continue
		}

		if v.Type == "object" {
			if currDepth == 1 {
				// these types are common across all the types that will be generated
				if k == "metadata" {
					fields = append(fields, genFieldEntry(k, "Metadata! @goField(name: \"objectMeta\")", false))
					inputFields = append(fields, genFieldEntry(k, "MetadataIn!", false))
					continue
				}

				// if k == "status" {
				// 	fields = append(fields, genFieldEntry(k, "Status", m[k]))
				// 	continue
				// }
			}

			if len(v.Properties) == 0 {
				fields = append(fields, genFieldEntry(k, "Map", m[k]))
				inputFields = append(inputFields, genFieldEntry(k, "Map", m[k]))
				continue
			}

			fields = append(fields, genFieldEntry(k, typeName+genTypeName(k), m[k]))
			inputFields = append(inputFields, genFieldEntry(k, typeName+genTypeName(k)+"In", m[k]))
			p.NavigateTree(typeName+genTypeName(k), &v, currDepth+1)
			continue
		}

		fields = append(fields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
		inputFields = append(inputFields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
	}

	p.Types[typeName] = fields
	p.Inputs[typeName+"In"] = inputFields
}

func (p *parser) GenerateFromJsonSchema(name string, schema *v1.JSONSchemaProps) {
	p.NavigateTree(name, schema)
}

func (p *parser) LoadStruct(name string, data any) {
	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	p.GenerateGraphQLSchema(name, ty)
}

func (p *parser) Debug() {
	fmt.Println("Types:")
	litter.Dump(p.Types)
	fmt.Println("Inputs:")
	litter.Dump(p.Inputs)
	fmt.Println("Enums:")
	litter.Dump(p.Enums)
}

func newParser(kCli k8s.ExtendedK8sClient) *parser {
	return &parser{
		Types:         map[string][]string{},
		Inputs:        map[string][]string{},
		Enums:         map[string][]string{},
		kCli:          kCli,
		CommonTypes:   map[string][]string{},
		CommonInpuuts: map[string][]string{},
		CommonEnums:   map[string][]string{},
	}
}

func NewParser(kCli k8s.ExtendedK8sClient) Parser {
	return newParser(kCli)
}
