package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/k8s"
)

type Parser interface {
	GenerateGraphQLSchema(structName string, name string, t reflect.Type)
	LoadStruct(name string, data any)
	PrintSchema(w io.Writer)
	DumpSchema(dir string) error
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

	reflect.String: "String",
}

type Struct struct {
	Types  map[string][]string
	Inputs map[string][]string
	Enums  map[string][]string
}

func newStruct() *Struct {
	return &Struct{
		Types:  map[string][]string{},
		Inputs: map[string][]string{},
		Enums:  map[string][]string{},
	}
}

const (
	commonLabel = "common-types"
)

type parser struct {
	structs map[string]*Struct
	kCli    k8s.ExtendedK8sClient
}

type JsonTag struct {
	Value     string
	OmitEmpty bool
	Inline    bool
}

func sanitizePackagePath(t reflect.Type) string {
	pkgPath := t.PkgPath()
	pkgPath = strings.ReplaceAll(pkgPath, "/", "__")
	pkgPath = strings.ReplaceAll(pkgPath, ".", "_")

	return pkgPath
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

func (s *Struct) mergeParser(other *Struct, overKey string) (fields []string, inputFields []string) {
	for k, v := range other.Types {
		if k == overKey {
			fields = append(fields, v...)
			continue
		}
		s.Types[k] = v
	}

	for k, v := range other.Inputs {
		if k == overKey+"In" {
			inputFields = append(inputFields, v...)
			continue
		}
		s.Inputs[k] = v
	}

	for k, v := range other.Enums {
		s.Enums[k] = v
	}

	return fields, inputFields
}

func (p *parser) GenerateGraphQLSchema(structName string, name string, t reflect.Type) {
	var fields []string
	var inputFields []string

	if _, ok := p.structs[structName]; !ok {
		p.structs[structName] = newStruct()
	}

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

		if field.Type.Kind() != reflect.String {
			if v, ok := kindMap[field.Type.Kind()]; ok {
				fieldType = toFieldType(v, !jt.OmitEmpty)
				inputFieldType = toFieldType(v, !jt.OmitEmpty)
			}
		}

		gt := parseGraphqlTag(field)

		if fieldType == "" {
			switch field.Type.Kind() {
			case reflect.String:
				{
					childType := name + field.Name
					if gt.Enum != nil {
						p.structs[structName].Enums[childType] = gt.Enum
						fieldType = toFieldType(childType, !jt.OmitEmpty)
						inputFieldType = toFieldType(childType, !jt.OmitEmpty)
						break
					}

					fieldType = toFieldType("String", !jt.OmitEmpty)
					inputFieldType = toFieldType("String", !jt.OmitEmpty)
				}
			case reflect.Struct:
				{
					if gt.Uri != nil {
						if strings.HasPrefix(*gt.Uri, "k8s://") {
							k8sCrdName := strings.Split(*gt.Uri, "k8s://")[1]
							jsonSchema, err := p.kCli.GetCRDJsonSchema(context.TODO(), k8sCrdName)
							if err != nil {
								panic(err)
							}

							func() {
								childType := getGraphqlType(name, field.Name, field.Type, jt, gt)
								if jt.Inline {
									// TODO: call json parsing and create type, input, and enum for all those types, using jsonSchema, with fields being inline
									p2 := newParser(p.kCli)
									p2.structs[structName] = newStruct()
									p2.structs[structName].GenerateFromJsonSchema(childType, jsonSchema)

									// TODO
									fields2, inputFields2 := p.structs[structName].mergeParser(p2.structs[structName], childType)
									fields = append(fields, fields2...)
									inputFields = append(inputFields, inputFields2...)

									return
								}

								fieldType = toFieldType(childType, !jt.OmitEmpty)
								inputFieldType = toFieldType(childType+"In", !jt.OmitEmpty)

								// TODO: call json parsing and create type, input, and enum for all those types, using jsonSchema
								p.structs[structName].GenerateFromJsonSchema(childType, jsonSchema)
							}()
						}
						break
					}

					func() {
						pkgPath := sanitizePackagePath(field.Type)
						childType := getGraphqlType(name, field.Name, field.Type, jt, gt)

						if jt.Inline {
							p2 := newParser(p.kCli)
							p2.GenerateGraphQLSchema(structName, childType, field.Type)
							p2.structs[structName] = newStruct()

							fields2, inputFields2 := p.structs[structName].mergeParser(p2.structs[structName], childType)
							fields = append(fields, fields2...)
							inputFields = append(inputFields, inputFields2...)

							return
						}

						fieldType = toFieldType(childType, !jt.OmitEmpty)
						inputFieldType = toFieldType(childType+"In", !jt.OmitEmpty)

						if pkgPath == "" {
							p.GenerateGraphQLSchema(structName, childType, field.Type)
							return
						}
						p.GenerateGraphQLSchema(commonLabel, childType, field.Type)
					}()
				}
			case reflect.Slice:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						childType := getGraphqlType(name, field.Name, field.Type.Elem(), jt, gt)
						pkgPath := sanitizePackagePath(field.Type.Elem())

						fieldType = toFieldType(fmt.Sprintf("[%s]", toFieldType(childType, true)), !jt.OmitEmpty)
						inputFieldType = toFieldType(fmt.Sprintf("[%s]", toFieldType(childType+"In", true)), !jt.OmitEmpty)

						if pkgPath == "" {
							p.GenerateGraphQLSchema(structName, childType, field.Type.Elem())
							break
						}
						p.GenerateGraphQLSchema(commonLabel, childType, field.Type.Elem())
						break
					}

					fieldType = toFieldType(fmt.Sprintf("[%s]", toFieldType(kindMap[field.Type.Elem().Kind()], true)), !jt.OmitEmpty)
					inputFieldType = toFieldType(fmt.Sprintf("[%s]", toFieldType(kindMap[field.Type.Elem().Kind()], true)), !jt.OmitEmpty)
				}
			case reflect.Ptr:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						childType := getGraphqlType(name, field.Name, field.Type.Elem(), jt, gt)
						pkgPath := sanitizePackagePath(field.Type.Elem())

						fieldType = childType
						inputFieldType = childType + "In"

						if pkgPath == "" {
							p.GenerateGraphQLSchema(structName, childType, field.Type.Elem())
							break
						}
						p.GenerateGraphQLSchema(commonLabel, childType, field.Type.Elem())
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
						childType := getGraphqlType(name, field.Name, field.Type.Elem(), jt, gt)
						pkgPath := sanitizePackagePath(field.Type.Elem())
						if pkgPath == "" {
							p.GenerateGraphQLSchema(structName, childType, field.Type.Elem())
							break
						}
						p.GenerateGraphQLSchema(commonLabel, childType, field.Type.Elem())
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

	p.structs[structName].Types[name] = fields
	p.structs[structName].Inputs[name+"In"] = inputFields
}

func (s *Struct) NavigateTree(name string, tree *v1.JSONSchemaProps, depth ...int) {
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
				s.NavigateTree(typeName+genTypeName(k), v.Items.Schema, currDepth+1)
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
			s.NavigateTree(typeName+genTypeName(k), &v, currDepth+1)
			continue
		}

		fields = append(fields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
		inputFields = append(inputFields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
	}

	s.Types[typeName] = fields
	s.Inputs[typeName+"In"] = inputFields
}

func (s *Struct) GenerateFromJsonSchema(name string, schema *v1.JSONSchemaProps) {
	s.NavigateTree(name, schema)
}

func (p *parser) LoadStruct(name string, data any) {
	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	p.GenerateGraphQLSchema(name, name, ty)
}

func (s *Struct) WriteSchema(w io.Writer) {
	keys := make([]string, 0, len(s.Types))
	for k := range s.Types {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("type %s {\n", keys[i]))
		sort.Slice(s.Types[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Types[keys[i]][p]) < strings.ToLower(s.Types[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Types[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}

	keys = make([]string, 0, len(s.Inputs))
	for k := range s.Inputs {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("input %s {\n", keys[i]))
		sort.Slice(s.Inputs[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Inputs[keys[i]][p]) < strings.ToLower(s.Inputs[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Inputs[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}

	keys = make([]string, 0, len(s.Enums))
	for k := range s.Enums {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("enum %s {\n", keys[i]))
		sort.Slice(s.Enums[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Enums[keys[i]][p]) < strings.ToLower(s.Enums[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Enums[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}
}

func (p *parser) PrintSchema(w io.Writer) {
	for _, v := range p.structs {
		v.WriteSchema(w)
	}
}

func (p *parser) DumpSchema(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(dir, 0o766); err != nil {
			return err
		}
	}

	for k, v := range p.structs {
		f, err := os.Create(filepath.Join(dir, strings.ToLower(k)+".graphqls"))
		if err != nil {
			return err
		}

		v.WriteSchema(f)
		f.Close()
	}
	return nil
}

func newParser(kCli k8s.ExtendedK8sClient) *parser {
	return &parser{
		structs: map[string]*Struct{
			commonLabel: {
				Types:  map[string][]string{},
				Inputs: map[string][]string{},
				Enums:  map[string][]string{},
			},
		},
		kCli: kCli,
	}
}

func NewParser(kCli k8s.ExtendedK8sClient) Parser {
	return newParser(kCli)
}
