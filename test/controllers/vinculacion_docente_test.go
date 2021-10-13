package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	pgUser := os.Getenv("RESOLUCIONES_CRUD_PGUSER")
	pgPass := os.Getenv("RESOLUCIONES_CRUD_PGPASS")
	pgUrls := os.Getenv("RESOLUCIONES_CRUD_PGHOST")
	pgDb := os.Getenv("RESOLUCIONES_CRUD_PGDB")
	pgPort := os.Getenv("RESOLUCIONES_CRUD_PGPORT")
	pgSchema := os.Getenv("RESOLUCIONES_CRUD_PGSCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")
}

func TestGetAllVinculacionDocente(t *testing.T) {
	Convey("Test: / - Vinculacion docente GetAll", t, func() {
		r, err := http.NewRequest("GET", "/v1/vinculacion_docente", nil)
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

func TestGetOneVinculacionDocenteById(t *testing.T) {
	Convey("Test: / - Vinculacion docente GetOne", t, func() {
		r, err := http.NewRequest("GET", "/v1/vinculacion_docente/1789", nil)
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
			r, err = http.NewRequest("GET", "/v1/vinculacion_docente/1t89", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 404)
		})
	})

}

func TestGetValoresTotalesXDisponibilidad(t *testing.T) {

	Convey("Test: /get_valores_totales_x_disponibilidad - Vinculacion docente Endpoint", t, func() {
		r, err := http.NewRequest("GET", "/v1/vinculacion_docente/get_valores_totales_x_disponibilidad/2020/1/100556", nil)
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
		Convey("Status Code Should Be 500", func() {
			r, err = http.NewRequest("GET", "/v1/vinculacion_docente/get_valores_totales_x_disponibilidad/202o/1/1oo556", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})

	})

}

func TestGetVinculacionesAgrupadas(t *testing.T) {

	Convey("Test: /get_vinculaciones_agrupadas - Vinculacion docente Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/vinculacion_docente/get_vinculaciones_agrupadas/1789", nil)
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
		//data
		// petición con datos incorrectos

		Convey("Status Code Should Be 500", func() {
			r, err = http.NewRequest("GET", "/v1/vinculacion_docente/get_vinculaciones_agrupadas/1t89", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})
	})
}

func TestGetTotalContratosResolucion(t *testing.T) {

	Convey("Test: /get_total_contratos_x_resolucion - Vinculacion docente Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/vinculacion_docente/get_total_contratos_x_resolucion/1789/HCH", nil)
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
		//data
		// petición con datos incorrectos

		Convey("Status Code Should Be 500", func() {
			r, err = http.NewRequest("GET", "/v1/vinculacion_docente/get_total_contratos_x_resolucion/1t89/MCO", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})
	})
}
