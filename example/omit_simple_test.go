package main

import (
	"fmt"
	"testing"

	"github.com/miloxing/json-filter/filter"
)

type OmitUser struct {
	Name    string
	Avatar  string `json:"avatar,omit(lang)"`
	LangAge LangAger
}

type LangAger struct {
	Name     string
	IsStatic bool `json:"is_static,omit(lang)"`
}

func NewOmitUser() OmitUser {
	return OmitUser{
		Name:   "boyan",
		Avatar: "avatar111",
		LangAge: LangAger{
			Name:     "go",
			IsStatic: true,
		},
	}
}

func TestOmitUser(t *testing.T) {
	fmt.Println(filter.Omit("lang", NewOmitUser()))
	//{"LangAge":{"Name":"go"},"Name":"boyan"}
}
