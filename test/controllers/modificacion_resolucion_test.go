package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"

	. "github.com/smartystreets/goconvey/convey"
)

/* func init() {
	pgUser := os.Getenv("RESOLUCIONES_CRUD_PGUSER")
	pgPass := os.Getenv("RESOLUCIONES_CRUD_PGPASS")
	pgUrls := os.Getenv("RESOLUCIONES_CRUD_PGHOST")
	pgDb := os.Getenv("RESOLUCIONES_CRUD_PGDB")
	pgPort := os.Getenv("RESOLUCIONES_CRUD_PGPORT")
	pgSchema := os.Getenv("RESOLUCIONES_CRUD_PGSCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")
} */

func TestGetAllModificacionResolucion(t *testing.T) {
	Convey("Test: / - Modificacion Resolucion GetAll Endpoint", t, func() {
		r, err := http.NewRequest("GET", "/v1/modificacion_resolucion", nil)
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

func TestGetOneModificacionResolucionById(t *testing.T) {
	Convey("Test: / - Modificacion Resolucion GetOne Endpoint", t, func() {
		r, err := http.NewRequest("GET", "/v1/modificacion_resolucion/100", nil)
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
			r, err = http.NewRequest("GET", "/v1/modificacion_resolucion/1000000", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 404)
		})
	})

}
