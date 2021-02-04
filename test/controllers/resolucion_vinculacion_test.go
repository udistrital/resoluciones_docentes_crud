package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/udistrital/resoluciones_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

var o orm.Ormer
var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

type ResolucionVinculacion struct {
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
}

func init() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgUrls := os.Getenv("POSTGRES_HOST")
	pgDb := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_DB_PORT")
	pgSchema := os.Getenv("POSTGRES_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

}

// TestGetAll is a sample to run an endpoint test
/*func TestGetAll(t *testing.T) {
	r, err := http.NewRequest("GET", "/v1/resolucion-vinculacion/", nil)
	if err != nil {
		t.Fatal("error", err)
	}
	w := httptest.NewRecorder()
	//x := conn.ResolucionVinculacionController{}
	//x.GetAll()
	//handler := http.HandlerFunc(x.GetAll(x, y))
	fmt.Println(w)
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	fmt.Println(r)
	fmt.Println(w)

	var response = map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &response)
	t.Log(w.Body.String())
	t.Log(response)
	Convey("Subject: Test Contenido Resolucion Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

}*/

func TestGetAll(t *testing.T) {
	r, err := http.NewRequest("GET", "/v1/resolucion-vinculacion/expedidas_vigencia_periodo?vigencia=2019", nil)
	if err != nil {
		t.Fatal("error", err)
	}
	w := httptest.NewRecorder()

	fmt.Println(w)
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	fmt.Println(r)
	fmt.Println(w)

	var response = map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &response)
	t.Log(w.Body.String())
	t.Log(response)
	Convey("Subject: Test Contenido Resolucion Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
