// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"kloudlite.io/pkg/repos"
)

type Cluster struct {
	ID       repos.ID  `json:"id"`
	Name     string    `json:"name"`
	Endpoint *string   `json:"endpoint"`
	Devices  []*Device `json:"devices"`
}

type Device struct {
	ID            repos.ID `json:"id"`
	User          *User    `json:"user"`
	Name          string   `json:"name"`
	Cluster       *Cluster `json:"cluster"`
	Configuration string   `json:"configuration"`
}

type User struct {
	ID      repos.ID  `json:"id"`
	Devices []*Device `json:"devices"`
}
