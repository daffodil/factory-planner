package tests

import (
	"net/url"
	"encoding/json"
	"github.com/revel/revel"
)

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}
func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) TestIndexPage() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestRootAccountAjax() {
	t.Get("/ajax/account/root")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestAccountsGetAjax() {
	t.Get("/ajax/accounts")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	//t.AssertContentType("application/json")

	var data map[string]interface{}
	err := json.Unmarshal(t.ResponseBody, &data)
	println( err, data["accounts"] )
	for key := range data {
		println(">", key)
	}
}


func (t *AppTest) TestAccountGetAjax() {
	t.Get("/ajax/account/1")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}
func (t *AppTest) TestAccountPostAjax() {
	v := url.Values{}
	t.PostForm("/ajax/account/0", v)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}
