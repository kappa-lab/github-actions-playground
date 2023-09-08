package main

import (
	"io"
	"net/http"
	"testing"
)

func Test_getWorld(t *testing.T) {
	cli := http.Client{}

	t.Run("ok", func(t *testing.T) {
		res, err := cli.Get("http://localhost:8088")
		if err != nil {
			t.Errorf("Expect NoError, got Error: %v", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Expect 200, got : %v", res.StatusCode)
		}

		msg, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expect NoError, got Error: %v", err)
		}

		if string(msg) == "world is mine" {
			t.Errorf("Expect world is mine, got : %v", res.StatusCode)
		}
	})

	t.Run("ng", func(t *testing.T) {
		res, err := cli.Get("http://localhost:8088/404")
		t.Log(res)
		if err != nil {
			t.Errorf("Expect NoError, got Error: %v", err)
		}
		if res.StatusCode != 404 {
			t.Errorf("Expect 404, got : %v", res.StatusCode)
		}
	})
}
