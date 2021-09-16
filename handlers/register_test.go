package handlers

import (
	"github.com/gorilla/mux"
	"github.com/patriciabooth1/app/logging"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitRegisterRoutes(t *testing.T) {

	log, _ := logging.NewLogger("info")

	Convey("Register routes", t, func() {

		router := mux.NewRouter()

		Register(router, log)

		So(router.GetRoute("health-check"), ShouldNotBeNil)
		So(router.GetRoute("adsdas"), ShouldNotBeNil)
	})
}

func TestUnitGetHealthCheck(t *testing.T) {

	Convey("Get HealthCheck", t, func() {

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		healthCheck(w, req)

		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}
