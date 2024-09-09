// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Mutation struct {
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage *bool   `json:"hasNextPage,omitempty"`
	HasPrevPage *bool   `json:"hasPrevPage,omitempty"`
	StartCursor *string `json:"startCursor,omitempty"`
}

type Query struct {
}

type GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus string

const (
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusFailed  GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus = "failed"
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusPending GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus = "pending"
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusSuccess GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus = "success"
)

var AllGithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus = []GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus{
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusFailed,
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusPending,
	GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusSuccess,
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusFailed, GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusPending, GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatusSuccess:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___apps___finance___internal___entities__ChargeStatus", str)
	}
	return nil
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus string

const (
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusFailed  GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus = "failed"
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusPending GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus = "pending"
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusSuccess GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus = "success"
)

var AllGithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus = []GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus{
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusFailed,
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusPending,
	GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusSuccess,
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusFailed, GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusPending, GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatusSuccess:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___apps___finance___internal___entities__PaymentStatus", str)
	}
	return nil
}

func (e GithubComKloudliteAPIAppsFinanceInternalEntitiesPaymentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
