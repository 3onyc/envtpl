package main

import (
	"testing"
)

func Test_EnvMap(t *testing.T) {
	e := []string{"FOO=bar", "QUX=baz"}
	r := EnvMap(e)

	if r["FOO"] != "bar" || r["QUX"] != "baz" {
		t.Errorf("Env not properly parsed: %+v\n", r)
	}
}
