package parser

import (
	"fmt"
	"reflect"
)

func getGraphqlType(parentName string, fieldName string, t reflect.Type, jt JsonTag, gt GraphqlTag) (ctype string) {
	switch t.Kind() {
	// case reflect.String:
	// 	{
	// 		pkgPath := sanitizePackagePath(t)
	// 		if gt.Enum != nil {
	// 			if pkgPath == "" {
	// 				return parentName + fieldName
	// 			}
	// 			return pkgPath + "_" + t.Name()
	// 		}
	// 		return "String"
	// 	}
	case reflect.Struct:
		{
			pkgPath := sanitizePackagePath(t)

			childType := func() string {
				if pkgPath != "" {
					return genTypeName(pkgPath + "_" + t.Name())
				}
				return genTypeName(parentName + fieldName)
			}()

			return childType
		}
	default:
		panic(fmt.Sprintf("unsupported type %s", t.Kind()))
	}
}
