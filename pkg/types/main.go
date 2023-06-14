package types

import (
	"encoding/base64"
	"encoding/json"
)

type M map[string]interface{}

type Pagination struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

type CursorSortBy struct {
	Field string `json:"field"`
	Order int    `json:"order"`
}

type Cursor struct {
	SortBy CursorSortBy `json:"sortBy,omitempty"`
	Value  string       `json:"value,omitempty"`
}

func CursorToBase64(c Cursor) string {
	b, _ := json.Marshal(c)
	return base64.StdEncoding.EncodeToString(b)
}

func CursorFromBase64(b string) (Cursor, error) {
	var c Cursor
	bb, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return Cursor{}, err
	}
	if err := json.Unmarshal(bb, &c); err != nil {
		return Cursor{}, err
	}

	return c, nil
}

type CursorPagination struct {
	First int64  `json:"first,omitempty"`
	After Cursor `json:"after,omitempty"`
}

func BuildCursorPagination(first *int, after *string) CursorPagination {
	c := CursorPagination{}

	c.First = func() int64 {
		if first == nil {
			return 10
		}
		return int64(*first)
	}()

	c.After = func() Cursor {
		if after == nil {
			return Cursor{
				SortBy: CursorSortBy{Field: "_id", Order: 1},
				Value:  "",
			}
		}
		aft, _ := CursorFromBase64(*after)
		return aft
	}()

	return c
}
