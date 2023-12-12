// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/kloudlite/api/pkg/repos"
)

type App struct {
	ID                     repos.ID               `json:"id"`
	Pipelines              []*GitPipeline         `json:"pipelines"`
	CiCreateDockerPipeLine map[string]interface{} `json:"ci_createDockerPipeLine"`
	CiCreatePipeLine       map[string]interface{} `json:"ci_createPipeLine"`
}

func (App) IsEntity() {}

type DockerBuild struct {
	DockerFile string  `json:"dockerFile"`
	ContextDir string  `json:"contextDir"`
	BuildArgs  *string `json:"buildArgs"`
}

type GitDockerPipelineIn struct {
	Name          string                 `json:"name"`
	AccountID     string                 `json:"accountId"`
	ProjectID     string                 `json:"projectId"`
	AppID         string                 `json:"appId"`
	ProjectName   string                 `json:"projectName"`
	ContainerName string                 `json:"containerName"`
	GitProvider   string                 `json:"gitProvider"`
	GitRepoURL    string                 `json:"gitRepoUrl"`
	RepoName      string                 `json:"repoName"`
	GitBranch     string                 `json:"gitBranch"`
	DockerFile    string                 `json:"dockerFile"`
	ContextDir    string                 `json:"contextDir"`
	BuildArgs     string                 `json:"buildArgs"`
	ArtifactRef   *GitPipelineArtifactIn `json:"artifactRef"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type GitPipeline struct {
	ID          repos.ID               `json:"id"`
	Name        string                 `json:"name"`
	GitProvider string                 `json:"gitProvider"`
	GitRepoURL  string                 `json:"gitRepoUrl"`
	GitBranch   string                 `json:"gitBranch"`
	Build       *GitPipelineBuild      `json:"build"`
	Run         *GitPipelineRun        `json:"run"`
	DockerBuild *DockerBuild           `json:"dockerBuild"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type GitPipelineArtifact struct {
	DockerImageName *string `json:"dockerImageName"`
	DockerImageTag  *string `json:"dockerImageTag"`
}

type GitPipelineArtifactIn struct {
	DockerImageName *string `json:"dockerImageName"`
	DockerImageTag  *string `json:"dockerImageTag"`
}

type GitPipelineBuild struct {
	BaseImage *string `json:"baseImage"`
	Cmd       string  `json:"cmd"`
	OutputDir *string `json:"outputDir"`
}

type GitPipelineBuildIn struct {
	BaseImage string  `json:"baseImage"`
	Cmd       string  `json:"cmd"`
	OutputDir *string `json:"outputDir"`
}

type GitPipelineIn struct {
	Name          string                 `json:"name"`
	AccountID     string                 `json:"accountId"`
	ProjectID     string                 `json:"projectId"`
	AppID         string                 `json:"appId"`
	ProjectName   string                 `json:"projectName"`
	ContainerName string                 `json:"containerName"`
	GitProvider   string                 `json:"gitProvider"`
	GitRepoURL    string                 `json:"gitRepoUrl"`
	RepoName      string                 `json:"repoName"`
	GitBranch     string                 `json:"gitBranch"`
	Build         *GitPipelineBuildIn    `json:"build"`
	Run           *GitPipelineRunIn      `json:"run"`
	ArtifactRef   *GitPipelineArtifactIn `json:"artifactRef"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type GitPipelineRun struct {
	BaseImage *string `json:"baseImage"`
	Cmd       string  `json:"cmd"`
}

type GitPipelineRunIn struct {
	BaseImage *string `json:"baseImage"`
	Cmd       string  `json:"cmd"`
}

type HarborImageTagsResult struct {
	Name      string `json:"name"`
	Signed    bool   `json:"signed"`
	Immutable bool   `json:"immutable"`
}

type HarborSearchResult struct {
	ImageName string                   `json:"imageName"`
	UpdatedAt *time.Time               `json:"updatedAt"`
	CreatedAt *time.Time               `json:"createdAt"`
	Tags      []*HarborImageTagsResult `json:"tags"`
}

type PipelineRun struct {
	ID               repos.ID             `json:"id"`
	PipelineID       repos.ID             `json:"pipelineId"`
	CreationTime     time.Time            `json:"creationTime"`
	StartTime        *time.Time           `json:"startTime"`
	EndTime          *time.Time           `json:"endTime"`
	Success          bool                 `json:"success"`
	Message          *string              `json:"message"`
	State            *string              `json:"state"`
	GitProvider      string               `json:"gitProvider"`
	GitRepo          string               `json:"gitRepo"`
	GitBranch        string               `json:"gitBranch"`
	Build            *GitPipelineBuild    `json:"build"`
	Run              *GitPipelineRun      `json:"run"`
	DockerBuildInput *DockerBuild         `json:"dockerBuildInput"`
	ArtifactRef      *GitPipelineArtifact `json:"artifactRef"`
}
