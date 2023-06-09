package parser

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	// "github.com/maxatome/go-testdeep/td"
	// "github.com/andreyvit/diff"
	// "github.com/sergi/go-diff/diffmatchpatch"
	"kloudlite.io/pkg/k8s"
)

func Test_GeneratedGraphqlSchema(t *testing.T) {
	type fields struct {
		structs map[string]*Struct
		kCli    k8s.ExtendedK8sClient
	}
	type args struct {
		name string
		data any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]*Struct
	}{
		{
			name: "test case 1 (without any json tag)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "User",
				data: struct {
					ID       int
					Username string
					Gender   string
				}{},
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"ID: Int!",
							"Username: String!",
							"Gender: String!",
						},
					},
					Inputs: map[string][]string{
						"UserIn": {
							"ID: Int!",
							"Username: String!",
							"Gender: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 2 (with json tags, for naming)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "User",
				data: struct {
					ID       int    `json:"id,omitempty"`
					Username string `json:"username"`
					Gender   string `json:"gender"`
				}{},
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"id: Int",
							"username: String!",
							"gender: String!",
						},
					},
					Inputs: map[string][]string{
						"UserIn": {
							"id: Int",
							"username: String!",
							"gender: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 3 (with json tags for naming, and graphql enum tags)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "User",
				data: struct {
					ID       int    `json:"id,omitempty"`
					Username string `json:"username"`
					Gender   string `json:"gender" graphql:"enum=MALE;FEMALE"`
				}{},
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"id: Int",
							"username: String!",
							"gender: UserGender!",
						},
					},
					Inputs: map[string][]string{
						"UserIn": {
							"id: Int",
							"username: String!",
							"gender: UserGender!",
						},
					},
					Enums: map[string][]string{
						"UserGender": {
							"FEMALE",
							"MALE",
						},
					},
				},
			},
		},
		{
			name: "test case 5 (with struct containing slice field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Post",
				data: struct {
					ID      int
					Title   string
					Content string
					Tags    []string
				}{},
			},
			want: map[string]*Struct{
				"Post": {
					Types: map[string][]string{
						"Post": {
							"ID: Int!",
							"Title: String!",
							"Content: String!",
							"Tags: [String!]!",
						},
					},
					Inputs: map[string][]string{
						"PostIn": {
							"ID: Int!",
							"Title: String!",
							"Content: String!",
							"Tags: [String!]!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 6 (with struct containing pointer field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Address",
				data: struct {
					Street  string
					City    string
					Country *string
				}{},
			},
			want: map[string]*Struct{
				"Address": {
					Types: map[string][]string{
						"Address": {
							"Street: String!",
							"City: String!",
							"Country: String",
						},
					},
					Inputs: map[string][]string{
						"AddressIn": {
							"Street: String!",
							"City: String!",
							"Country: String",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 7 (with struct containing nested anonymous struct field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Employee",
				data: struct {
					ID      int
					Name    string
					Address struct {
						Street string
						City   string
					}
				}{},
			},
			want: map[string]*Struct{
				"Employee": {
					Types: map[string][]string{
						"Employee": {
							"ID: Int!",
							"Name: String!",
							"Address: EmployeeAddress!",
						},
						"EmployeeAddress": {
							"Street: String!",
							"City: String!",
						},
					},
					Inputs: map[string][]string{
						"EmployeeIn": {
							"ID: Int!",
							"Name: String!",
							"Address: EmployeeAddressIn!",
						},
						"EmployeeAddressIn": {
							"Street: String!",
							"City: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 8 (with struct containing nested struct field with json tags)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Employee",
				data: struct {
					ID      int
					Name    string
					Address struct {
						Street string `json:"street"`
						City   string `json:"city"`
					} `json:"address"`
				}{},
			},
			want: map[string]*Struct{
				"Employee": {
					Types: map[string][]string{
						"Employee": {
							"ID: Int!",
							"Name: String!",
							"address: EmployeeAddress!",
						},
						"EmployeeAddress": {
							"street: String!",
							"city: String!",
						},
					},
					Inputs: map[string][]string{
						"EmployeeIn": {
							"ID: Int!",
							"Name: String!",
							"address: EmployeeAddressIn!",
						},
						"EmployeeAddressIn": {
							"street: String!",
							"city: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 9 (with struct containing struct pointer field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Company",
				data: struct {
					ID      int
					Name    string
					Address *struct {
						Street string
						City   string
					}
				}{},
			},
			want: map[string]*Struct{
				"Company": {
					Types: map[string][]string{
						"Company": {
							"ID: Int!",
							"Name: String!",
							"Address: CompanyAddress",
						},
						"CompanyAddress": {
							"Street: String!",
							"City: String!",
						},
					},
					Inputs: map[string][]string{
						"CompanyIn": {
							"ID: Int!",
							"Name: String!",
							"Address: CompanyAddressIn",
						},
						"CompanyAddressIn": {
							"Street: String!",
							"City: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 11 (with struct containing struct slice field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Organization",
				data: struct {
					ID        int
					Name      string
					Employees []struct {
						ID   int
						Name string
					}
				}{},
			},
			want: map[string]*Struct{
				"Organization": {
					Types: map[string][]string{
						"Organization": {
							"ID: Int!",
							"Name: String!",
							"Employees: [OrganizationEmployees!]!",
						},
						"OrganizationEmployees": {
							"ID: Int!",
							"Name: String!",
						},
					},
					Inputs: map[string][]string{
						"OrganizationIn": {
							"ID: Int!",
							"Name: String!",
							"Employees: [OrganizationEmployeesIn!]!",
						},
						"OrganizationEmployeesIn": {
							"ID: Int!",
							"Name: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 12 (with struct containing struct slice field with json tags)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Organization",
				data: struct {
					ID        int
					Name      string
					Employees []struct {
						ID   int    `json:"employee_id"`
						Name string `json:"employee_name"`
					} `json:"employees"`
				}{},
			},
			want: map[string]*Struct{
				"Organization": {
					Types: map[string][]string{
						"Organization": {
							"ID: Int!",
							"Name: String!",
							"employees: [OrganizationEmployees!]!",
						},
						"OrganizationEmployees": {
							"employee_id: Int!",
							"employee_name: String!",
						},
					},
					Inputs: map[string][]string{
						"OrganizationIn": {
							"ID: Int!",
							"Name: String!",
							"employees: [OrganizationEmployeesIn!]!",
						},
						"OrganizationEmployeesIn": {
							"employee_id: Int!",
							"employee_name: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 13 (with struct containing enum field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Product",
				data: struct {
					ID       int
					Name     string
					Category string `graphql:"enum=ELECTRONICS;FASHION;SPORTS"`
				}{},
			},
			want: map[string]*Struct{
				"Product": {
					Types: map[string][]string{
						"Product": {
							"ID: Int!",
							"Name: String!",
							"Category: ProductCategory!",
						},
					},
					Inputs: map[string][]string{
						"ProductIn": {
							"ID: Int!",
							"Name: String!",
							"Category: ProductCategory!",
						},
					},
					Enums: map[string][]string{
						"ProductCategory": {
							"ELECTRONICS",
							"FASHION",
							"SPORTS",
						},
					},
				},
			},
		},
		{
			name: "test case 14 (with struct containing struct slice to pointer of a inline struct)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    nil,
			},
			args: args{
				name: "Organization",
				data: struct {
					ID        int
					Name      string
					Employees []*struct {
						ID   int    `json:"employee_id"`
						Name string `json:"employee_name"`
					} `json:"employees"`
				}{},
			},
			want: map[string]*Struct{
				"Organization": {
					Types: map[string][]string{
						"Organization": {
							"ID: Int!",
							"Name: String!",
							"employees: [OrganizationEmployees]!",
						},
						"OrganizationEmployees": {
							"employee_id: Int!",
							"employee_name: String!",
						},
					},
					Inputs: map[string][]string{
						"OrganizationIn": {
							"ID: Int!",
							"Name: String!",
							"employees: [OrganizationEmployeesIn]!",
						},
						"OrganizationEmployeesIn": {
							"employee_id: Int!",
							"employee_name: String!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				structs: tt.fields.structs,
				kCli:    tt.fields.kCli,
			}

			p.LoadStruct(tt.args.name, tt.args.data)
			buf := new(bytes.Buffer)
			p.PrintSchema(buf)
			got := buf.String()

			buf2 := new(bytes.Buffer)
			p2 := &parser{
				structs: tt.want,
			}
			p2.PrintSchema(buf2)
			want := buf2.String()

			// if !td.CmpString(t, got, want) {
			// 	t.Errorf("Failed")
			// 	// t.Errorf("GeneratedGraphqlSchema() = \n***\n%v\n***\n but want \n***\n%v\n***\n", got, want)
			// }

			if got != want {
				// t.Errorf("Result not as expected:\n%v", diff.LineDiff(got, want))
				g, err2 := os.Create("./got.txt")
				if err2 != nil {
					t.Error(err2)
				}
				g.WriteString(got)

				w, err2 := os.Create("./want.txt")
				if err2 != nil {
					t.Error(err2)
				}
				w.WriteString(want)

				cmd := exec.Command("delta", "./got.txt", "./want.txt", "-s")
				b, err := cmd.CombinedOutput()
				if err != nil {
					t.Error(err)
				}

				t.Errorf(string(b))

				// dmp := diffmatchpatch.New()
				// diffs := dmp.DiffMain(got, want, false)
				// t.Errorf(dmp.Diff)
				// t.Errorf(dmp.DiffPrettyText(diffs))
				// t.Errorf("GeneratedGraphqlSchema() = \n***\n%v\n***\n but want \n***\n%v\n***\n", got, want)
			}
		})
	}
}
