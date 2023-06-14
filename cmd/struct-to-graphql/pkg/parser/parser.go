package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/pkg/k8s"
)

var typeMap = map[reflect.Type]string{
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

	reflect.Bool:   "Boolean",
	reflect.String: "String",
}

func genFieldType(fieldType string, omitempty bool) string {
	if omitempty {
		return fieldType
	}
	return fieldType + "!"
}

func GenerateGraphQLSchema(name string, data interface{}, kCli k8s.ExtendedK8sClient) (map[string][]string, error) {
	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	schemaMap := map[string][]string{}
	commonTypesMap := map[string][]string{}

	getGraphQLFields(ty, name, schemaMap, commonTypesMap, kCli)

	return schemaMap, nil
}

func getGraphQLFields(t reflect.Type, name string, dataMap map[string][]string, commonTypesMap map[string][]string, kCli k8s.ExtendedK8sClient) {
	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			// unexported field
			continue
		}

		graphqlTag := field.Tag.Get("graphql")

		jsonSchemaTag := field.Tag.Get("json-schema")
		fieldName := field.Tag.Get("json")
		sp := strings.Split(fieldName, ",")

		omitempty := false
		inline := false

		if len(sp) >= 1 && sp[0] == "-" {
			// this field does not want to be included in the schema
			continue
		}

		// iterating from 1 as the first element is the field name, it would always be there
		for i := 1; i < len(sp); i++ {
			if sp[i] == "omitempty" {
				omitempty = true
			}
			if sp[i] == "inline" {
				inline = true
			}
		}

		if len(sp) > 1 {
			fieldName = sp[0]
		}

		if fieldName == "" {
			fieldName = field.Name
		}

		fieldType := ""

		hasSpecialCase := false

		if tf, ok := typeMap[field.Type]; ok {
			hasSpecialCase = true
			fieldType = tf
		}

		if !hasSpecialCase {
			switch field.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fieldType = genFieldType("Int", omitempty)
			case reflect.Float32, reflect.Float64:
				fieldType = genFieldType("Float", omitempty)
			case reflect.Bool:
				fieldType = genFieldType("Boolean", omitempty)
			case reflect.String:
				if jsonSchemaTag != "" {
					jsonSchemaEnum := ""

					if strings.HasPrefix(jsonSchemaTag, "enum=") {
						jsonSchemaEnum = strings.Split(jsonSchemaTag, "enum=")[1]
					}

					if jsonSchemaEnum != "" {
						enums := strings.Split(jsonSchemaEnum, ",")
						fieldType = genFieldType(name+field.Name, omitempty)

						fields := make([]string, len(enums))
						copy(fields, enums)

						sort.Strings(fields)
						dataMap[fmt.Sprintf("enum %s", name+field.Name)] = fields
					}
				} else {
					fieldType = genFieldType("String", omitempty)
				}

			case reflect.Struct:
				fieldType = field.Name

				typeIsCommon := false

				if graphqlTag != "" {
					sp := strings.Split(graphqlTag, " ")
					for i := range sp {
						if strings.HasPrefix(sp[i], "common=") {
							if strings.SplitN(sp[i], "common=", 2)[1] == "true" {
								typeIsCommon = true
								getGraphQLFields(field.Type, fieldType, commonTypesMap, commonTypesMap, kCli)
							}
						}
					}
				}

				if jsonSchemaTag != "" {
					jsonSchemaUri := ""
					if strings.HasPrefix(jsonSchemaTag, "uri=") {
						jsonSchemaUri = strings.Split(jsonSchemaTag, "uri=")[1]
					}

					if strings.HasPrefix(jsonSchemaUri, "k8s://") {
						crdName := strings.Split(jsonSchemaUri, "k8s://")[1]
						jp, err := kCli.GetCRDJsonSchema(context.TODO(), crdName)
						if err != nil {
							panic(err)
						}

						if inline {
							dMap := map[string][]string{}
							Convert(jp, field.Type.Name(), dMap)
							for k, v := range dMap {
								if k == field.Type.Name() {
									fields = append(fields, v...)
									continue
								}
								sort.Strings(v)
								dataMap[k] = v
							}
							continue
						} else {
							fieldType = genFieldType("K8s"+field.Type.Name(), omitempty)
							Convert(jp, "K8s"+field.Type.Name(), dataMap)
						}
					}
				} else {
					if inline {
						dMap := map[string][]string{}
						getGraphQLFields(field.Type, field.Name, dMap, commonTypesMap, kCli)
						for k, v := range dMap {
							if k == field.Name {
								fields = append(fields, v...)
								continue
							}
							sort.Strings(v)
							dataMap[fmt.Sprintf("type %s", k)] = v
						}
						continue
					} else {
						if typeIsCommon {
							fieldType = field.Name
						} else {
							fieldType = name + field.Name
							getGraphQLFields(field.Type, name+field.Name, dataMap, commonTypesMap, kCli)
						}
					}
				}
			case reflect.Slice:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						if jsonSchemaTag != "" {
							jsonSchemaUri := ""
							if strings.HasPrefix(jsonSchemaTag, "uri=") {
								jsonSchemaUri = strings.Split(jsonSchemaTag, "uri=")[1]
							}
							if strings.HasPrefix(jsonSchemaUri, "k8s://") {
								fieldType = genFieldType(fmt.Sprintf("[%s]", "K8s"+field.Type.Name()), omitempty)
								crdName := strings.Split(jsonSchemaUri, "k8s://")[1]
								jp, err := kCli.GetCRDJsonSchema(context.TODO(), crdName)
								if err != nil {
									panic(err)
								}
								Convert(jp, "K8s"+field.Type.Name(), dataMap)
							}
						} else {
							getGraphQLFields(field.Type.Elem(), name+field.Name, dataMap, commonTypesMap, kCli)
						}
					} else {
						fieldType = fmt.Sprintf("[%s]", kindMap[field.Type.Elem().Kind()])
					}
				}
			case reflect.Ptr:
				{
					if field.Type.Elem().Kind() == reflect.Struct {
						fieldType = field.Name
						getGraphQLFields(field.Type.Elem(), name+field.Name, dataMap, commonTypesMap, kCli)
					} else {
						fieldType = kindMap[field.Type.Elem().Kind()]
					}
				}
			case reflect.Map:
				{
					fieldType = "Map"
					if field.Type.Elem().Kind() == reflect.Struct {
						getGraphQLFields(field.Type.Elem(), name+field.Name, dataMap, commonTypesMap, kCli)
					}
				}
			case reflect.Interface:
				{
					fieldType = "Any"
				}
			default:
				{
					fmt.Printf("default: name: %v (%v), type: %v, kind: %v\n", fieldName, field.Name, field.Type, field.Type.Kind())
				}
			}
		}

		if fieldType != "" {
			fields = append(fields, fmt.Sprintf("%s: %s", fieldName, fieldType))
			continue
		}
	}

	sort.Strings(fields)
	dataMap[fmt.Sprintf("type %s", name)] = fields
}

func WriteSchema(schema map[string][]string, writer io.Writer, shareable ...bool) error {
	for k, v := range schema {
		if strings.HasPrefix(k, "type") && len(shareable) > 0 {
			if _, err := fmt.Fprintf(writer, "%s shareable {\n", k); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(writer, "%s {\n", k); err != nil {
				return err
			}
		}
		for _, f := range v {
			if _, err := fmt.Fprintf(writer, "\t%s\n", f); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(writer, "}\n\n"); err != nil {
			return err
		}
	}
	return nil
}
