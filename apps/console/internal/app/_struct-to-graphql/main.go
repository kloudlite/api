package main

import (
	kldj882pqtnzprrf2qjlfpc2cj7r2srr "github.com/kloudlite/api/apps/console/internal/domain"
	kl4kbmrs2vgrnzst4nd5gmgljc82vvvd "github.com/kloudlite/api/apps/console/internal/entities"
	kldwjrlxf7brdg65st8mz69xqmfcrclj "github.com/kloudlite/api/pkg/repos"
	klh4jn7rhvrhm5hksjbf997gkdbgdjdb "github.com/kloudlite/operator/apis/wireguard/v1"
	"github.com/kloudlite/api/cmd/struct-to-graphql/pkg/parser"
	"k8s.io/client-go/rest"
	"os"
	"golang.org/x/sync/errgroup"
	"context"
	"path"
	"flag"
	"fmt"
	"strings"

	"time"
	"net/http"
	"io"
	"encoding/json"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apiExtensionsV1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type schemaClient struct {
	kcli *clientset.Clientset
}

func (s schemaClient) GetK8sJsonSchema(name string) (*apiExtensionsV1.JSONSchemaProps, error) {
	ctx, cf := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cf()
	crd, err := s.kcli.ApiextensionsV1().CustomResourceDefinitions().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return crd.Spec.Versions[0].Schema.OpenAPIV3Schema, nil
}

func (s schemaClient) GetHttpJsonSchema(url string) (*apiExtensionsV1.JSONSchemaProps, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var m apiExtensionsV1.JSONSchemaProps
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &m, nil
}


func main() {
	var isDev bool
	var outDir string
	var withPagination string

	flag.BoolVar(&isDev, "dev", false, "--dev")
	flag.StringVar(&outDir, "out-dir", "struct-to-graphql", "--out-dir <dir-name>")
	flag.StringVar(&withPagination, "with-pagination", "", "--with-pagination <type1,type2,type3...>")
	flag.Parse()

	stat, err := os.Stat(outDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(outDir, 0755); err != nil {
				panic(err)
			}
		}
	}

	if stat != nil && !stat.IsDir() {
		panic(fmt.Errorf("out-dir (%s) is not a directory", outDir))
	}

	types := map[string]any{
		"App": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.App{},
		"Config": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.Config{},
		"ConfigKeyRef": &kldj882pqtnzprrf2qjlfpc2cj7r2srr.ConfigKeyRef{},
		"ConfigKeyValueRef": &kldj882pqtnzprrf2qjlfpc2cj7r2srr.ConfigKeyValueRef{},
		"CursorPagination": &kldwjrlxf7brdg65st8mz69xqmfcrclj.CursorPagination{},
		"Environment": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.Environment{},
		"ImagePullSecret": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.ImagePullSecret{},
		"ManagedResource": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.ManagedResource{},
		"MatchFilter": &kldwjrlxf7brdg65st8mz69xqmfcrclj.MatchFilter{},
		"Port": &klh4jn7rhvrhm5hksjbf997gkdbgdjdb.Port{},
		"Project": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.Project{},
		"ProjectManagedService": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.ProjectManagedService{},
		"Router": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.Router{},
		"Secret": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.Secret{},
		"SecretKeyRef": &kldj882pqtnzprrf2qjlfpc2cj7r2srr.SecretKeyRef{},
		"SecretKeyValueRef": &kldj882pqtnzprrf2qjlfpc2cj7r2srr.SecretKeyValueRef{},
		"VPNDevice": &kl4kbmrs2vgrnzst4nd5gmgljc82vvvd.VPNDevice{},
	}

	kcli, err := func() (*clientset.Clientset, error) {
		if isDev {
			return clientset.NewForConfig(&rest.Config{Host: "localhost:8080"})
		}

		cfg, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return clientset.NewForConfig(cfg)
	}()
	if err != nil {
		panic(err)
	}


	if err != nil {
		panic(err)
	}

	g, _ := errgroup.WithContext(context.TODO())

	g.Go(func() error {
		directives, err := parser.Directives()
		if err != nil {
			return err
		}
		return os.WriteFile(path.Join(outDir, "directives.graphqls"), directives, 0644)
	})

	g.Go(func() error {
		scalarTypes, err := parser.ScalarTypes()
		if err != nil {
			panic(err)
		}
		return os.WriteFile(path.Join(outDir, "scalars.graphqls"), scalarTypes, 0644)
	})

	p := parser.NewParser(&schemaClient{kcli: kcli})

	for k, v := range types {
		p.LoadStruct(k, v)
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}

	p.WithPagination(strings.Split(withPagination, ","))

	if err := p.DumpSchema(outDir); err != nil {
		panic(err)
	}
}
