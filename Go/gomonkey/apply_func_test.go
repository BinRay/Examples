package main

import (
	"github.com/agiledragon/gomonkey"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHttpGet(t *testing.T) {
	Convey("TestHttpGet", t, func() {
		// Define the mock response
		mockResponse := &http.Response{StatusCode: 201}

		// Apply the function patch
		patch := gomonkey.ApplyFunc(http.Get, func(_ string) (*http.Response, error) {
			return mockResponse, nil
		})
		defer patch.Reset()

		// Call the function and check the response
		resp, err := http.Get("http://anyexample.com")
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}
