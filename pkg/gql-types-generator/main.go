package gqltypesgenerator

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateGraphQLTypes(inputs []interface{}, notReqTypes []string) string {
	// var sb strings.Builder
	seenTypes := make(map[reflect.Type]bool)

	if notReqTypes == nil {
		notReqTypes = []string{"Time", "-", "RawJson", "Location", "zone", "zoneTrans", "creation_time", "update_time", "FieldsV1"}
	}

	res := ""

	for i, _ := range inputs {
		op1, op2 := generateGraphQLTypesRecursive(reflect.TypeOf(inputs[i]), seenTypes, false, notReqTypes)
		res += op1 + op2
	}

	return res

}

func isNeeded(fieldName string, notReqTypes []string) bool {

	for _, v := range notReqTypes {
		if fieldName == v {
			return false
		}
	}

	return true
}

func generateGraphQLTypesRecursive(t reflect.Type, seenTypes map[reflect.Type]bool, onlyMembers bool, notReqTypes []string) (string, string) {
	// fmt.Printf("%s", seenTypes)

	// main string
	res := ""

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if seenTypes[t] {
		return res, ""
	}

	// fmt.Println(t)

	seenTypes[t] = true

	if t.Kind() == reflect.Struct {

		typeName := t.Name()

		if typeName == "" {
			// If the type doesn't have a name, use the name of the parent struct and the field name
			parentType := t.Field(0).Type
			if parentType.Kind() == reflect.Ptr {
				parentType = parentType.Elem()
			}
			typeName = fmt.Sprintf("%s%s", parentType.Name(), strings.Title(t.Field(0).Name))
		}

		res2 := ""
		res3 := ""
		res4 := ""
		res5 := ""
		res6 := ""

		// fmt.Println(sb)
		for i := 0; i < t.NumField(); i++ {

			field := t.Field(i)
			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = field.Name
			} else if fieldName == ",inline" {
				op1, op2 := generateGraphQLTypesRecursive(field.Type, seenTypes, true, notReqTypes)
				// fmt.Println(op2, field.Name)

				res4 += op1
				res6 += op2
				continue
			} else if s := strings.Split(fieldName, ","); len(s) == 2 && s[1] == "omitempty" {
				// fmt.Println("yes", fieldName)
				fieldName = s[0]
			}

			if !isNeeded(fieldName, notReqTypes) {
				continue
			}

			// fmt.Println("here1", typeName)
			fieldType := getGraphQLType(field.Type, seenTypes)

			if !isNeeded(fieldType, notReqTypes) {
				// either continue
				// continue
				// or use String
				fieldType = "String"
			}

			res2 += fmt.Sprintf("\t%s: %s%s\n", fieldName, fieldType, getGraphQLRequired(field))

			mT := field.Type
			for {
				if mT.Kind() != reflect.Ptr {
					break
				}
				mT = mT.Elem()
			}

			if mT.Kind() == reflect.Struct {
				op1, op2 := generateGraphQLTypesRecursive(field.Type, seenTypes, false, notReqTypes)
				res3 += op1
				res6 += op2
				// + op2
			} else if mT.Kind() == reflect.Slice || mT.Kind() == reflect.Array {

				op1, op2 := generateGraphQLTypesRecursive(mT.Elem(), seenTypes, false, notReqTypes)
				res3 += op1
				res6 += op2
			}

		}

		if isNeeded(typeName, notReqTypes) {
			if !onlyMembers {
				if res4 != "" || res2 != "" {
					res += fmt.Sprintf("type %s {\n", typeName)

					res += res2
					// res += res5
					res += res6

					res += "}\n\n"

					res += res4
				} else {
					res += fmt.Sprintf("type %s String\n\n", typeName)
				}
			} else {
				// fmt.Println(res2, typeName)
				res5 += res2
				res5 += res6
				res5 += res4
			}
		} else {
			res5 += res6
		}

		// fmt.Println(res5)

		res += res3

		return res, res5
	}

	return "", ""
}

func getGraphQLType(t reflect.Type, seenTypes map[reflect.Type]bool) string {
	tObj := t
	if t.Kind() == reflect.Ptr {
		tObj = t.Elem()
	}

	// switch tObj.Name() {
	// case "Time":
	// 	return "String"
	// }

	switch tObj.Kind() {
	case reflect.String:
		return "String"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "Int"
	case reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint32:
		return "Int"
	case reflect.Map:
		return "String"
	case reflect.Float32, reflect.Float64:
		return "Float"
	case reflect.Bool:
		return "Boolean"
	case reflect.Slice, reflect.Array:
		return "[" + getGraphQLType(tObj.Elem(), seenTypes) + "]"
	case reflect.Struct:
		typeName := tObj.Name()
		if typeName == "" {
			parentType := tObj.Field(0).Type
			if parentType.Kind() == reflect.Ptr {
				parentType = parentType.Elem()
			}
			typeName = fmt.Sprintf("%s%s", parentType.Name(), strings.Title(tObj.Field(0).Name))
		}
		if seenTypes[tObj] {
			return typeName
		}
		// fmt.Println("here", typeName)
		// generateGraphQLTypesRecursive(tObj, seenTypes)
		return typeName
	default:
		fmt.Println("unknown type:", tObj.Kind())
		return ""
	}
}

func getGraphQLRequired(t reflect.StructField) string {
	req := true
	fieldName := t.Tag.Get("json")
	s := strings.Split(fieldName, ",")
	if len(s) == 2 && s[1] == "omitempty" {
		// fmt.Println("yes", fieldName)
		req = false
	}
	if t.Type.Kind() != reflect.Ptr && req {
		return "!"
	}
	return ""
}