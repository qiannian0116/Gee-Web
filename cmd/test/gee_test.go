package test

import (
	"testing"

	"github.com/qiannian0116/Gee-Web/geek/gee"
)

func TestNestedGroup(t *testing.T) {
	r := gee.New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.GetGroupPrefix() != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.GetGroupPrefix() != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
