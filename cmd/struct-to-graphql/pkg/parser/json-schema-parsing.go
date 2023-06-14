package parser

import (
	"fmt"
	"sort"
	"strings"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func gqlTypeMap(jsonType string) string {
	switch jsonType {
	case "boolean":
		return "Boolean"
	case "integer":
		return "Int"
	case "object":
		return "Object"
	case "string":
		return "String"
	case "array":
		return "Array"
	default:
		return "Any"
	}
}

func genTypeName(n string) string {
	return strings.ToUpper(n[0:1]) + n[1:]
}

func genFieldEntry(k string, t string, required bool) string {
	if required {
		return fmt.Sprintf("%s: %s!", k, t)
	}
	return fmt.Sprintf("%s: %s", k, t)
}

func navigateTree(tree *v1.JSONSchemaProps, name string, schemaMap map[string][]string, depth ...int) {
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
		// fmt.Printf("[properties] %q type: %s\n", k, v.Type)

		if k == "apiVersion" || k == "kind" {
			fields = append(fields, genFieldEntry(k, "String!", m[k]))
			inputFields = append(inputFields, genFieldEntry(k, "String!", m[k]))
			continue
		}

		if v.Type == "array" {
			if v.Items.Schema != nil && v.Items.Schema.Type == "object" {
				fields = append(fields, genFieldEntry(k, fmt.Sprintf("[%s]", typeName+genTypeName(k)), m[k]))
				inputFields = append(inputFields, genFieldEntry(k, fmt.Sprintf("[%sIn]", typeName+genTypeName(k)), m[k]))
				// iVar += genFieldEntry(k, fmt.Sprintf("[%s]", typeName+genTypeName(k)+"In"), m[k])

				navigateTree(v.Items.Schema, typeName+genTypeName(k), schemaMap, currDepth+1)
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

				if k == "status" {
					fields = append(fields, genFieldEntry(k, "Status", m[k]))
					continue
				}
			}

			if len(v.Properties) == 0 {
				fields = append(fields, genFieldEntry(k, "Map", m[k]))
				inputFields = append(inputFields, genFieldEntry(k, "Map", m[k]))
				continue
			}

			fields = append(fields, genFieldEntry(k, typeName+genTypeName(k), m[k]))
			inputFields = append(inputFields, genFieldEntry(k, typeName+genTypeName(k)+"In", m[k]))
			navigateTree(&v, typeName+genTypeName(k), schemaMap, currDepth+1)
			continue
		}

		fields = append(fields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
		inputFields = append(inputFields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
	}

	sort.Strings(fields)
	schemaMap[fmt.Sprintf("type %s", name)] = fields
	sort.Strings(inputFields)
	schemaMap[fmt.Sprintf("input %sIn", name)] = inputFields
}

func Convert(schema *v1.JSONSchemaProps, name string, schemaMap map[string][]string) error {
	navigateTree(schema, name, schemaMap)
	return nil
}
