package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/udistrital/resoluciones_docentes_crud/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

/* type ResolucionVinculacion struct {
	Id                 int       `orm:"column(id );pk;auto"`
	Estado             string    `orm:"column(estado )"`
	Numero             string    `orm:"column(numero )"`
	Vigencia           int       `orm:"column(vigencia )"`
	Facultad           int       `orm:"column(facultad )"`
	NivelAcademico     string    `orm:"column(nivel_academico )"`
	Dedicacion         string    `orm:"column(dedicacion )"`
	FechaExpedicion    time.Time `orm:"column(fecha_expedicion );type(timestamp without time zone)"`
	NumeroSemanas      int       `orm:"column(numero_semanas )"`
	Periodo            int       `orm:"column(periodo )"`
	TipoResolucion     string    `orm:"column(tipo_resolucion )"`
	IdDependenciaFirma int       `orm:"column(dependencia_firma )"`
	PeriodoCarga       int       `orm:"column(periodo_carga )"`
	VigenciaCarga      int       `orm:"column(vigencia_carga )"`
} */

/* func init() {
	pgUser := os.Getenv("RESOLUCIONES_CRUD_PGUSER")
	pgPass := os.Getenv("RESOLUCIONES_CRUD_PGPASS")
	pgUrls := os.Getenv("RESOLUCIONES_CRUD_PGHOST")
	pgDb := os.Getenv("RESOLUCIONES_CRUD_PGDB")
	pgPort := os.Getenv("RESOLUCIONES_CRUD_PGPORT")
	pgSchema := os.Getenv("RESOLUCIONES_CRUD_PGSCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

}
*/
func TestGetAllGetAll(t *testing.T) {

	Convey("Test: / - Resolucion vinculación  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/resolucion_vinculacion/", nil)
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

		/*Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})*/
	})

}

func TestGetAllGetAllAprobada(t *testing.T) {

	Convey("Test: /Aprobada - Resolucion vinculación  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/resolucion_vinculacion/Aprobada", nil)
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

		/*Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})*/
	})

}

// Test al endpoint Expedidas vigencia periodo del controlador Resolucion vinculacion
func TestGetAllExpedidasVigenciaPeriodo(t *testing.T) {
	//t.Log(w.Body.String())
	//t.Log(response)
	Convey("Test: /expedidas_vigencia_periodo - Resolucion vinculacion  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/resolucion_vinculacion/expedidas_vigencia_periodo?vigencia=2019", nil)
		if err != nil {
			t.Fatal("error", err)
		}
		w := httptest.NewRecorder()

		beego.BeeApp.Handlers.ServeHTTP(w, r)
		//var response = map[string]interface{}{}
		//json.Unmarshal(w.Body.Bytes(), &response)

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
			r, err = http.NewRequest("GET", "/v1/resolucion_vinculacion/expedidas_vigencia_periodo?vigencia=201a", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})
		/*Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})*/
	})

}

func TestGetAllExpedidasVigenciaPeriodoVinculacion(t *testing.T) {

	Convey("Test: /expedidas_vigencia_periodo_vinculacion - Resolucion vinculación  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/resolucion_vinculacion/expedidas_vigencia_periodo_vinculacion?vigencia=2019", nil)
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
			r, err = http.NewRequest("GET", "/v1/resolucion_vinculacion/expedidas_vigencia_periodo_vinculacion?vigencia=201a", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})
		/*Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})*/
	})

}
