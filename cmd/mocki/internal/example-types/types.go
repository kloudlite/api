package example_types

import (
	io2 "io"
)

type Example interface {
	ex1() string
	Ex2() string
}

type User struct {
	Name string
}

type Sample[T any, K any] interface {
	Name() string
	Age() int
	hello() string
	SetName(name string)
	SetUser(name string, age int, ex Example, writer io2.Writer)
	//SetAndGetUser(name string, age int, ex kloudlite.io/cmd/interface-impl/types.Example) (*User, error)
	SetAndGetUser(name string, age int, ex Example) *User
	// SetAndGetUser2(name string, age int, ex Example) (user types2.Sample)
	// SetAndGetUser3(name string, age int, ex Example, s1 T, s2 K) (user types2.Sample)
	// Example
}
