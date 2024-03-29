// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BYOCCluster struct {
	Metadata     *Metadata                                         `json:"metadata"`
	Spec         *GithubComKloudliteOperatorApisClustersV1BYOCSpec `json:"spec"`
	ClusterToken string                                            `json:"clusterToken"`
}

func (BYOCCluster) IsEntity() {}

type Cluster struct {
	Metadata     *Metadata                                            `json:"metadata"`
	Spec         *GithubComKloudliteOperatorApisClustersV1ClusterSpec `json:"spec,omitempty"`
	ClusterToken string                                               `json:"clusterToken"`
}

func (Cluster) IsEntity() {}

type GithubComKloudliteOperatorApisClustersV1BYOCSpec struct {
	AccountName string `json:"accountName"`
}

type GithubComKloudliteOperatorApisClustersV1ClusterSpec struct {
	AccountName string `json:"accountName"`
}

type Metadata struct {
	Name string `json:"name"`
}
