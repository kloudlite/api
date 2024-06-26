package domain

import (
	"bytes"
	"fmt"
	"github.com/kloudlite/api/pkg/errors"
	"net/url"
	"text/template"

	dbv1 "github.com/kloudlite/operator/apis/distribution/v1"

	"github.com/kloudlite/api/apps/container-registry/templates"
	text_templates "github.com/kloudlite/api/pkg/text-templates"
	common_types "github.com/kloudlite/operator/apis/common-types"
)

func BuildUrl(repo, pullToken string) (string, error) {
	parsedURL, err := url.Parse(repo)
	if err != nil {
		fmt.Println("Error parsing Repo URL:", err)
		return "", errors.NewE(err)
	}

	parsedURL.User = url.User(pullToken)

	return parsedURL.String(), nil
}

type BuildJobTemplateData struct {
	AccountName string
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string

	Registry     dbv1.Registry
	Caches       []dbv1.Cache
	Resource     dbv1.Resource
	GitRepo      dbv1.GitRepo
	BuildOptions *dbv1.BuildOptions

	CredentialsRef common_types.SecretRef
}

func getTemplate(obj BuildJobTemplateData) ([]byte, error) {
	b, err := templates.ReadBuildJobTemplate()
	if err != nil {
		return nil, errors.NewE(err)
	}

	tmpl := text_templates.WithFunctions(template.New("build-job-template"))

	tmpl, err = tmpl.Parse(string(b))
	if err != nil {
		return nil, errors.NewE(err)
	}

	out := new(bytes.Buffer)
	err = tmpl.Execute(out, obj)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return out.Bytes(), nil
}

func (*Impl) GetBuildTemplate(obj BuildJobTemplateData) ([]byte, error) {
	return getTemplate(obj)
}
