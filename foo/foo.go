package foo

import "github.com/kyokomi/redis-test-example/bar"

type Foo struct {
	Name string
	Bar  bar.Bar
}

func NewFoo() *Foo {
	f := Foo{}
	f.Name = "test"
	f.Bar = bar.Bar{Name: "barTest"}
	return &f
}
