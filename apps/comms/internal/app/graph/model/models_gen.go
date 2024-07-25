// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/kloudlite/api/apps/comms/types"
)

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesEmail struct {
	Enabled     bool   `json:"enabled"`
	MailAddress string `json:"mailAddress"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesEmailIn struct {
	Enabled     bool   `json:"enabled"`
	MailAddress string `json:"mailAddress"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesSlack struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesSlackIn struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesTelegram struct {
	ChatID  string `json:"chatId"`
	Enabled bool   `json:"enabled"`
	Token   string `json:"token"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesTelegramIn struct {
	ChatID  string `json:"chatId"`
	Enabled bool   `json:"enabled"`
	Token   string `json:"token"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesWebhook struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

type GithubComKloudliteAPIAppsCommsInternalDomainEntitiesWebhookIn struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

type GithubComKloudliteAPIAppsCommsTypesNotifyContent struct {
	Body    string `json:"body"`
	Image   string `json:"image"`
	Link    string `json:"link"`
	Subject string `json:"subject"`
	Title   string `json:"title"`
}

type Mutation struct {
}

type NotificationEdge struct {
	Cursor string              `json:"cursor"`
	Node   *types.Notification `json:"node"`
}

type NotificationPaginatedRecords struct {
	Edges      []*NotificationEdge `json:"edges"`
	PageInfo   *PageInfo           `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage *bool   `json:"hasNextPage,omitempty"`
	HasPrevPage *bool   `json:"hasPrevPage,omitempty"`
	StartCursor *string `json:"startCursor,omitempty"`
}

type Query struct {
}

type GithubComKloudliteAPIAppsCommsTypesNotificationType string

const (
	GithubComKloudliteAPIAppsCommsTypesNotificationTypeAlert        GithubComKloudliteAPIAppsCommsTypesNotificationType = "alert"
	GithubComKloudliteAPIAppsCommsTypesNotificationTypeNotification GithubComKloudliteAPIAppsCommsTypesNotificationType = "notification"
)

var AllGithubComKloudliteAPIAppsCommsTypesNotificationType = []GithubComKloudliteAPIAppsCommsTypesNotificationType{
	GithubComKloudliteAPIAppsCommsTypesNotificationTypeAlert,
	GithubComKloudliteAPIAppsCommsTypesNotificationTypeNotification,
}

func (e GithubComKloudliteAPIAppsCommsTypesNotificationType) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIAppsCommsTypesNotificationTypeAlert, GithubComKloudliteAPIAppsCommsTypesNotificationTypeNotification:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIAppsCommsTypesNotificationType) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIAppsCommsTypesNotificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIAppsCommsTypesNotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___apps___comms___types__NotificationType", str)
	}
	return nil
}

func (e GithubComKloudliteAPIAppsCommsTypesNotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteAPIPkgReposMatchType string

const (
	GithubComKloudliteAPIPkgReposMatchTypeArray      GithubComKloudliteAPIPkgReposMatchType = "array"
	GithubComKloudliteAPIPkgReposMatchTypeExact      GithubComKloudliteAPIPkgReposMatchType = "exact"
	GithubComKloudliteAPIPkgReposMatchTypeNotInArray GithubComKloudliteAPIPkgReposMatchType = "not_in_array"
	GithubComKloudliteAPIPkgReposMatchTypeRegex      GithubComKloudliteAPIPkgReposMatchType = "regex"
)

var AllGithubComKloudliteAPIPkgReposMatchType = []GithubComKloudliteAPIPkgReposMatchType{
	GithubComKloudliteAPIPkgReposMatchTypeArray,
	GithubComKloudliteAPIPkgReposMatchTypeExact,
	GithubComKloudliteAPIPkgReposMatchTypeNotInArray,
	GithubComKloudliteAPIPkgReposMatchTypeRegex,
}

func (e GithubComKloudliteAPIPkgReposMatchType) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIPkgReposMatchTypeArray, GithubComKloudliteAPIPkgReposMatchTypeExact, GithubComKloudliteAPIPkgReposMatchTypeNotInArray, GithubComKloudliteAPIPkgReposMatchTypeRegex:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIPkgReposMatchType) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIPkgReposMatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIPkgReposMatchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___pkg___repos__MatchType", str)
	}
	return nil
}

func (e GithubComKloudliteAPIPkgReposMatchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}