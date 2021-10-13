package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAllResolucion(t *testing.T) {
	Convey("Test: / - Resolucion GetAll", t, func() {
		r, err := http.NewRequest("GET", "/v1/resolucion", nil)
		if err != nil {
			t.Fatal("error", err)
		}
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

}

func TestGetOneResolucionById(t *testing.T) {
	Convey("Test: / - Resolucion GetOne", t, func() {
		r, err := http.NewRequest("GET", "/v1/resolucion/333", nil)
		if err != nil {
			t.Fatal("error", err)
		}
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

		// petición con datos incorrectos
		Convey("Status Code Should Be 404", func() {
			r, err = http.NewRequest("GET", "/v1/resolucion/36", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 404)
		})
	})

}
