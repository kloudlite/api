package types

import (
	"encoding/base64"
)

type M map[string]interface{}

type Pagination struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

type CursorSortBy struct {
	Field     string        `json:"field"`
	Direction SortDirection `json:"sortDirection"`
}

//type Cursor struct {
//	SortBy CursorSortBy `json:"sortBy,omitempty"`
//	Value  string       `json:"value,omitempty"`
//}

type Cursor string

func CursorToBase64(c Cursor) string {
	return base64.StdEncoding.EncodeToString([]byte(c))
}

func CursorFromBase64(b string) (Cursor, error) {
	b2, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return Cursor(""), err
	}
	return Cursor(b2), nil
}

type CursorPagination struct {
	First int64   `json:"first"`
	After *string `json:"after,omitempty"`

	Last   int64   `json:"last,omitempty"`
	Before *string `json:"before,omitempty"`

	OrderBy       string        `json:"orderBy,omitempty"`
	SortDirection SortDirection `json:"sortDirection,omitempty" graphql:"enum=ASC;DESC"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

// func BuildCursorPagination(first *int, after *string) CursorPagination {
// 	c := CursorPagination{}
//
// 	c.First = func() int64 {
// 		if first == nil {
// 			return 10
// 		}
// 		return int64(*first)
// 	}()
//
// 	c.After = func() Cursor {
// 		if after == nil {
// 			return Cursor{
// 				SortBy: CursorSortBy{Field: "_id", Order: 1},
// 				Value:  "",
// 			}
// 		}
// 		aft, _ := CursorFromBase64(*after)
// 		return aft
// 	}()
//
// 	return c
// }
