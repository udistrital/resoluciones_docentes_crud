package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAllResolucionEstado(t *testing.T) {
	Convey("Test: / - Resolucion Estado GetAll", t, func() {
		r, err := http.NewRequest("GET", "/v1/resolucion_estado", nil)
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

func TestGetOneResolucionEstadoById(t *testing.T) {
	Convey("Test: / - ResolucionEstado GetOne", t, func() {
		r, err := http.NewRequest("GET", "/v1/resolucion_estado/100", nil)
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
			r, err = http.NewRequest("GET", "/v1/resolucion_estado/5", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 404)
		})
	})

}
