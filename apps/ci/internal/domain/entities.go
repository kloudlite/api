package domain

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"kloudlite.io/pkg/repos"
)

type DockerBuildInput struct {
	DockerFile string `json:"docker_file,omitempty" bson:"docker_file,omitempty"`
	ContextDir string `json:"working_dir,omitempty" bson:"working_dir,omitempty"`
	BuildArgs  string `json:"build_args,omitempty" bson:"build_args,omitempty"`
}

type ContainerImageBuild struct {
	BaseImage string `json:"base_image,omitempty" bson:"base_image,omitempty"`
	Cmd       string `json:"cmd,omitempty" bson:"cmd,omitempty"`
	OutputDir string `json:"output_dir,omitempty" bson:"output_dir,omitempty"`
}

type ContainerImageRun struct {
	BaseImage string `json:"base_image,omitempty" bson:"base_image,omitempty"`
	Cmd       string `json:"cmd,omitempty" bson:"cmd,omitempty"`
}

type ArtifactRef struct {
	DockerImageName string `json:"docker_image_name,omitempty" bson:"docker_image_name,omitempty"`
	DockerImageTag  string `json:"docker_image_tag,omitempty" bson:"docker_image_tag,omitempty"`
}

type GitlabWebhookId int
type GithubWebhookId int64

type Pipeline struct {
	repos.BaseEntity `bson:",inline"`
	Name             string `json:"name,omitempty" bson:"name"`
	ProjectId        string `json:"project_id" bson:"project_id"`
	AccountId        string `json:"account_id" bson:"account_id"`
	AppId            string `json:"app_id" bson:"app_id"`
	ProjectName      string `json:"project_name" bson:"project_name"`
	ContainerName    string `json:"container_name" bson:"container_name"`

	GitProvider string `json:"git_provider,omitempty" bson:"git_provider"`
	GitRepoUrl  string `json:"git_repo_url,omitempty" bson:"git_repo_url"`
	GitBranch   string `json:"git_branch" bson:"git_branch"`

	GitlabTokenId *repos.ID `json:"gitlab_token,omitempty" bson:"gitlab_token_id"`

	Build            *ContainerImageBuild `json:"build,omitempty" bson:"build,omitempty"`
	Run              *ContainerImageRun   `json:"run,omitempty" bson:"run,omitempty"`
	DockerBuildInput *DockerBuildInput    `json:"docker_build_input,omitempty" bson:"docker_build_input,omitempty"`

	ArtifactRef ArtifactRef `json:"artifact_ref,omitempty" bson:"artifact_ref,omitempty"`

	GithubWebhookId *GithubWebhookId `json:"github_webhook_id,omitempty" bson:"github_webhook_id,omitempty"`
	GitlabWebhookId *GitlabWebhookId `json:"gitlab_webhook_id,omitempty" bson:"gitlab_webhook_id,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty" bson:"metadata"`
}

type TektonVars struct {
	GitRepo     string `json:"git-repo"`
	GitUser     string `json:"git-user"`
	GitPassword string `json:"git-password"`

	GitRef        string `json:"git-ref"`
	GitCommitHash string `json:"git-commit-hash"`

	IsDockerBuild    bool    `json:"is-docker-build"`
	DockerFile       *string `json:"docker-file"`
	DockerContextDir *string `json:"docker-context-dir"`
	DockerBuildArgs  *string `json:"docker-build-args"`

	BuildBaseImage string `json:"build-base-image"`
	BuildCmd       string `json:"build-cmd"`
	BuildOutputDir string `json:"build-output-dir"`

	RunBaseImage string `json:"run-base-image"`
	RunCmd       string `json:"run-cmd"`

	TaskNamespace string `json:"task-namespace"`

	ArtifactDockerImageName string `json:"artifact_ref-docker_image_name"`
	ArtifactDockerImageTag  string `json:"artifact_ref-docker_image_tag"`
}

func (t *TektonVars) ToJson() (map[string]any, error) {
	marshal, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(marshal, &m); err != nil {
		return nil, err
	}
	return m, nil
}

var PipelineIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

type HarborAccount struct {
	repos.BaseEntity `bson:",inline"`
	HarborId         int    `json:"harbor_id"`
	ProjectName      string `json:"project_name"`
	Username         string `json:"username"`
	Password         string `json:"password"`
}

type AccessToken struct {
	Id       repos.ID       `json:"_id"`
	UserId   repos.ID       `json:"user_id" bson:"user_id"`
	Email    string         `json:"email" bson:"email"`
	Provider string         `json:"provider" bson:"provider"`
	Token    *oauth2.Token  `json:"token" bson:"token"`
	Data     map[string]any `json:"data" bson:"data"`
}
