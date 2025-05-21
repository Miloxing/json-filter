package main

import (
	"encoding/json"
	"testing"

	"github.com/liu-cn/json-filter/filter"
)

func TestAnonymousAuto(t *testing.T) {

	assertAnonymous := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	type Page struct {
		PageInfo int `json:"pageInfo,select($any)"`
		PageNum  int `json:"pageNum,select($any)"`
	}

	type Article struct {
		Title string `json:"title,select(article)"`
		Page         // 期望库能自动深入Page结构内部，识别其字段的select标签，并根据当前的筛选组（如"article"）决定是否包含
		// Page   `json:"page,select(article)"` // 注意这里tag里标注了匿名结构体的字段名，所以解析时会解析成对象，不会展开
		Author string `json:"author,select(admin),omit(admin)"`
	}

	article := Article{
		Title: "c++从研发到脱发",
		Page: Page{
			PageInfo: 999,
			PageNum:  1,
		},
		Author: "abc",
	}

	t.Run("anonymous article", func(t *testing.T) {
		gotFilter := filter.Select("article", article)
		gotJson, _ := json.Marshal(gotFilter)
		got := string(gotJson)
		want := `{"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}`
		assertAnonymous(t, got, want)
	})

	t.Run("anonymous admin", func(t *testing.T) {
		gotFilter := filter.Select("admin", article)
		gotJson, _ := json.Marshal(gotFilter)
		got := string(gotJson)
		want := `{"author":"abc","pageInfo":999,"pageNum":1}`
		assertAnonymous(t, got, want)
	})

	t.Run("anonymous omit", func(t *testing.T) {
		gotFilter := filter.Omit("admin", article)
		gotJson, _ := json.Marshal(gotFilter)
		got := string(gotJson)
		want := `{"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}`
		assertAnonymous(t, got, want)
	})

}
