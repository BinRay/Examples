package main

import (
	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestClientHttpGet(t *testing.T) {
	Convey("TestHttpGet", t, func() {
		// Define the mock response
		mockResponse := &http.Response{StatusCode: 201}

		client := &http.Client{
			Timeout: time.Second * 10, // Timeout after 10 seconds
		}

		// Apply the function patch
		patch := gomonkey.ApplyMethod(reflect.TypeOf(client), "Get", func(_ *http.Client, _ string) (*http.Response, error) {
			return mockResponse, nil
		})
		defer patch.Reset()

		// Call the function and check the response
		resp, err := http.Get("http://anyexample.com")
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}
