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

/* var o orm.Ormer
var (
	db *sql.DB
) */

/*
type Paragrafo struct {
	Id     int
	Numero int
	Texto  string
}

type Articulo struct {
	Id         int
	Numero     int
	Texto      string
	Paragrafos []Paragrafo
}

type ResolucionCompleta struct {
	Vinculacion   ResolucionVinculacionDocente
	Consideracion string
	Preambulo     string
	Vigencia      int
	Numero        string
	Id            int
	Articulos     []Articulo
	Titulo        string
}
*/
/* func init() {
	pgUser := os.Getenv("RESOLUCIONES_CRUD_PGUSER")
	pgPass := os.Getenv("RESOLUCIONES_CRUD_PGPASS")
	pgUrls := os.Getenv("RESOLUCIONES_CRUD_PGHOST")
	pgDb := os.Getenv("RESOLUCIONES_CRUD_PGDB")
	pgPort := os.Getenv("RESOLUCIONES_CRUD_PGPORT")
	pgSchema := os.Getenv("RESOLUCIONES_CRUD_PGSCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

} */

func TestResolucionTemplate(t *testing.T) {
	//t.Log(w.Body.String())
	//t.Log(response)
	Convey("Test: /Resolucion_template - Contenido resolucion  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/contenido_resolucion/resolucion_template/HPC/PREGRADO", nil)
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
		//data
		// petición con datos incorrectos
	})

}

func TestGetOne(t *testing.T) {
	//t.Log(w.Body.String())
	//t.Log(response)
	Convey("Test: /Get_one - Contenido resolucion  Endpoint\n", t, func() {
		// peticion correcta
		r, err := http.NewRequest("GET", "/v1/contenido_resolucion/333", nil)
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
			r, err = http.NewRequest("GET", "/v1/contenido_resolucion/36", nil)
			w = httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)
			So(w.Code, ShouldEqual, 500)
		})
		/*Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})*/
	})

}
