package parser

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"sigs.k8s.io/yaml"

	// "github.com/maxatome/go-testdeep/td"
	// "github.com/andreyvit/diff"
	// "github.com/sergi/go-diff/diffmatchpatch"
	"github.com/h2non/gock"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"kloudlite.io/pkg/k8s"
	"kloudlite.io/pkg/types"
)

type ExampleJson struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              struct {
		ClusterName  string   `json:"clusterName"`
		NodePoolName string   `json:"nodePoolName"`
		NodeType     string   `json:"nodeType"`
		Taints       []string `json:"taints"`
	}
}

func exampleJsonSchema() ([]byte, error) {
	var x = `description: Node is the Schema for the nodes API
properties:
  apiVersion:
    description: 'sample description'
    type: string
  kind:
    description: 'sample description'
    type: string
  metadata:
    type: object
  spec:
    properties:
      clusterName:
        type: string
      nodePoolName:
        type: string
      nodeType:
        enum:
          - worker
          - master
          - cluster
        type: string
      taints:
        items:
          type: string
        type: array
    required:
      - nodeType
      - clusterName
      - nodePoolName
    type: object
required:
  - spec
type: object
`

	return yaml.YAMLToJSON([]byte(x))
}

func Test_GeneratedGraphqlSchema(t *testing.T) {
	kCli, err := func() (k8s.ExtendedK8sClient, error) {
		return k8s.NewExtendedK8sClient(&rest.Config{Host: "localhost:8080"})
	}()
	if err != nil {
		t.Error(err)
	}

	b, err := exampleJsonSchema()
	if err != nil {
		t.Error(err)
	}

	// defer gock.Off()
	gock.New("http://example.com").
		Get("/example-json-schema").
		Times(2).
		Reply(200).Body(bytes.NewBuffer(b))

	req, err := http.NewRequest(http.MethodGet, "http://example.com/example-json-schema", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("invalid error code")
	}
	t.Logf("healthcheck passed on route %s with status code: %d", req.URL.String(), resp.StatusCode)

	type fields struct {
		structs map[string]*Struct
		kCli    k8s.ExtendedK8sClient
	}
	type args struct {
		name           string
		data           any
		withPagination bool
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
				kCli:    kCli,
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
		{
			name: "test case 16 (with struct containing map field)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    kCli,
			},
			args: args{
				name: "User",
				data: struct {
					ID    int
					Name  string
					Email string
					Tags  map[string]string
					KVs   map[string]any `json:"kvs"`
				}{},
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"ID: Int!",
							"Name: String!",
							"Email: String!",
							"Tags: Map!",
							"kvs: Map!",
						},
					},
					Inputs: map[string][]string{
						"UserIn": {
							"ID: Int!",
							"Name: String!",
							"Email: String!",
							"Tags: Map!",
							"kvs: Map!",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 17 (with struct containing nested kloudlite CRD)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    kCli,
			},
			args: args{
				name: "Project",
				data: struct {
					AccountName string
					Project     crdsv1.Project `json:",inline" graphql:"uri=k8s://projects.crds.kloudlite.io"`
				}{},
			},
			want: map[string]*Struct{
				"Project": {
					Types: map[string][]string{
						"Project": {
							"AccountName: String!",
							"apiVersion: String!",
							"kind: String!",
							// "metadata: Metadata!",
							"metadata: Metadata! @goField(name: \"objectMeta\")",
							"spec: Github_com__kloudlite__operator__apis__crds__v1_ProjectSpec!",
							"status: Github_com__kloudlite__operator__pkg__operator_Status",
						},
					},
					Inputs: map[string][]string{
						"ProjectIn": {
							"AccountName: String!",
							"apiVersion: String",
							"kind: String",
							"metadata: MetadataIn!",
							"spec: Github_com__kloudlite__operator__apis__crds__v1_ProjectSpecIn!",
						},
					},
					Enums: map[string][]string{},
				},
				"common-types": {
					Types: map[string][]string{
						"Github_com__kloudlite__operator__apis__crds__v1_ProjectSpec": {
							"accountName: String!",
							"clusterName: String!",
							"displayName: String",
							"logo: String",
							"targetNamespace: String!",
						},
						"Github_com__kloudlite__operator__pkg__operator_Check": {
							"status: Boolean!",
							"message: String",
							"generation: Int",
						},
						"Github_com__kloudlite__operator__pkg__operator_ResourceRef": {
							"apiVersion: String",
							"kind: String",
							"namespace: String!",
							"name: String!",
						},
						"Github_com__kloudlite__operator__pkg__operator_Status": {
							"isReady: Boolean!",
							"resources: [Github_com__kloudlite__operator__pkg__operator_ResourceRef!]",
							"message: Github_com__kloudlite__operator__pkg__raw___json_RawJson",
							"checks: Map",
							"lastReconcileTime: Date",
						},
						"Github_com__kloudlite__operator__pkg__raw___json_RawJson": {
							"RawMessage: Any",
						},
						"Metadata": {
							"name: String!",
							"namespace: String",
							"labels: Map",
							"annotations: Map",
							"generation: Int!",
							"creationTimestamp: Date!",
							"deletionTimestamp: Date",
						},
					},
					Inputs: map[string][]string{
						"Github_com__kloudlite__operator__apis__crds__v1_ProjectSpecIn": {
							"accountName: String!",
							"clusterName: String!",
							"displayName: String",
							"logo: String",
							"targetNamespace: String!",
						},
						"MetadataIn": {
							"name: String!",
							"namespace: String",
							"labels: Map",
							"annotations: Map",
						},
					},
					Enums: map[string][]string{},
				},
			},
		},
		{
			name: "test case 18 (with pagination enabled)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    kCli,
			},
			args: args{
				name: "User",
				data: struct {
					ID       int
					Username string
					Gender   string
				}{},
				withPagination: true,
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"ID: Int!",
							"Username: String!",
							"Gender: String!",
						},
						"UserPaginatedRecords": {
							"totalCount: Int!",
							"edges: [UserEdge!]!",
							"pageInfo: PageInfo!",
						},
						"UserEdge": {
							"node: User!",
							"cursor: String!",
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
				"common-types": {
					Types: map[string][]string{
						"PageInfo": {
							"hasNextPage: Boolean!",
							"hasPreviousPage: Boolean!",
							"startCursor: String",
							"endCursor: String",
						},
					},
				},
			},
		},
		{
			name: "test case 19 (with graphql (noinput))",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    kCli,
			},
			args: args{
				name: "User",
				data: struct {
					SyncStatus types.SyncStatus `json:"syncStatus" graphql:"noinput"`
				}{},
			},
			want: map[string]*Struct{
				"User": {
					Types: map[string][]string{
						"User": {
							"syncStatus: Kloudlite_io__pkg__types_SyncStatus!",
						},
					},
					Inputs: map[string][]string{},
					Enums:  map[string][]string{},
				},
				"common-types": {
					Types: map[string][]string{
						"Kloudlite_io__pkg__types_SyncStatus": {
							"action: Kloudlite_io__pkg__types_SyncStatusAction!",
							"error: String",
							"recordVersion: Int!",
							"lastSyncedAt: Date",
							"state: Kloudlite_io__pkg__types_SyncStatusState!",
							"syncScheduledAt: Date",
						},
					},
					Enums: map[string][]string{
						"Kloudlite_io__pkg__types_SyncStatusAction": {
							"APPLY",
							"DELETE",
						},
						"Kloudlite_io__pkg__types_SyncStatusState": {
							"IDLE",
							"APPLIED_AT_AGENT",
							"ERRORED_AT_AGENT",
							"IN_QUEUE",
							"RECEIVED_UPDATE_FROM_AGENT",
						},
					},
				},
			},
		},
		{
			name: "test case 20 (with json schema http uri)",
			fields: fields{
				structs: map[string]*Struct{},
				kCli:    kCli,
			},
			args: args{
				name: "Example",
				data: struct {
					// Example ExampleJson `json:"example" graphql:"uri=http://localhost:30017/example-json-schema"`
					Example ExampleJson `json:"example" graphql:"uri=http://example.com/example-json-schema"`
				}{},
			},
			want: map[string]*Struct{
				"Example": {
					Types: map[string][]string{
						"Example": {
							"example: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJson!",
						},
					},
					Inputs: map[string][]string{
						"ExampleIn": {
							"example: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonIn!",
						},
					},
					Enums: map[string][]string{},
				},
				"common-types": {
					Types: map[string][]string{
						"Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJson": {
							"apiVersion: String!",
							"kind: String!",
							"metadata: Metadata! @goField(name: \"objectMeta\")",
							"spec: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpec!",
						},
						"Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpec": {
							"clusterName: String!",
							"nodePoolName: String!",
							"nodeType: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpecNodeType!",
							"taints: [String]",
						},
						"Metadata": {
							"annotations: Map",
							"labels: Map",
							"name: String!",
							"namespace: String",
							"creationTimestamp: Date!",
							"deletionTimestamp: Date",
							"generation: Int!",
						},
					},
					Inputs: map[string][]string{
						"Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonIn": {
							"apiVersion: String",
							"kind: String",
							"metadata: MetadataIn!",
							"spec: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpecIn!",
						},
						"Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpecIn": {
							"clusterName: String!",
							"nodePoolName: String!",
							"nodeType: Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpecNodeType!",
							"taints: [String]",
						},
						"MetadataIn": {
							"annotations: Map",
							"labels: Map",
							"name: String!",
							"namespace: String",
						},
					},
					Enums: map[string][]string{
						"Kloudlite_io__cmd__struct___to___graphql__pkg__parser_ExampleJsonSpecNodeType": {
							"worker",
							"master",
							"cluster",
						},
					},
				},
			},
		},
	}

	for _, _tt := range tests {
		tt := _tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &parser{
				structs: tt.fields.structs,
				kCli:    tt.fields.kCli,
			}

			p.LoadStruct(tt.args.name, tt.args.data)
			buf := new(bytes.Buffer)
			if tt.args.withPagination {
				p.WithPagination()
			}
			p.PrintSchema(buf)
			got := buf.String()

			buf2 := new(bytes.Buffer)
			p2 := &parser{
				structs: tt.want,
			}
			p2.PrintSchema(buf2)
			want := buf2.String()

			if got != want {
				dir := "/tmp/x"
				g, err2 := os.Create(filepath.Join(dir, "./got.txt"))
				if err2 != nil {
					t.Error(err2)
				}
				g.WriteString(got)

				w, err2 := os.Create(filepath.Join(dir, "./want.txt"))
				if err2 != nil {
					t.Error(err2)
				}
				w.WriteString(want)

				cmd := exec.Command("diff", filepath.Join(dir, "./got.txt"), filepath.Join(dir, "./want.txt"))
				b, err := cmd.CombinedOutput()
				if err != nil {
					t.Error(err)
				}

				t.Errorf("diff output:\n%s\n", string(b))
			}
		})
	}
}
