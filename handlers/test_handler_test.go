package handlers

import (
	"bytes"
	c "context"
	"github.com/patriciabooth1/app/logging"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnitServeHTTP(t *testing.T) {

	log, _ := logging.NewLogger("trace")

	handler := &testHandler{
		log: log,
	}

	//Happy Path
	Convey("Given a valid request is made", t, func() {

		reqBody := []byte(`{}`)

		req := httptest.NewRequest(http.MethodPost, "/test", ioutil.NopCloser(bytes.NewReader(reqBody)))
		res := httptest.NewRecorder()

		Convey("Then an OK status should be returned", func() {

			handler.ServeHTTP(res, req.WithContext(c.Background()))

			So(res.Code, ShouldEqual, http.StatusOK)
		})
	})

	//Unhappy Path
	Convey("Given an invalid request is made", t, func() {

		reqBody := []byte(`{`)

		req := httptest.NewRequest(http.MethodPost, "/test", ioutil.NopCloser(bytes.NewReader(reqBody)))
		res := httptest.NewRecorder()

		Convey("Then a bad request status should be returned", func() {

			handler.ServeHTTP(res, req.WithContext(c.Background()))

			So(res.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
