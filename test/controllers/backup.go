package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/udistrital/resoluciones_docentes_crud/models"
	_ "github.com/udistrital/resoluciones_docentes_crud/routers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

/*var o orm.Ormer

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
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	//regitro de bd postgres
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true

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

func testWithDb(t *testing.T, f func(t *testing.T, mock sqlmock.Sqlmock)) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to open stub database connection, error: %v", err)
	}

	orm.AddAliasWthDB("default", "postgres", db)
	defer db.Close()

	f(t, mock)
}

func TestDBWithMockedSqlDriver(t *testing.T) {
	testWithDb(t, func(t *testing.T, mock sqlmock.Sqlmock) {
		// setup mock
		columnasTabla := []string{}

		v := reflect.TypeOf(models.ResolucionVinculacion{})
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			tag := field.Tag.Get("orm")
			column := ""
			_, err := fmt.Sscanf(strings.Split(tag, ";")[0], "column(%s )", &column)
			if err != nil {
				beego.Error(err)
			}
			columnasTabla = append(columnasTabla, column[:len(column)-1])
		}

		/*mock.ExpectQuery("SELECT DISTINCT r.id id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.facultad_id facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM resoluciones.resolucion r, resoluciones.resolucion_vinculacion_docente rv, resoluciones.resolucion_estado re, resoluciones.estado_resolucion e WHERE r.id=rv.id AND re.resolucion_id=r.id AND re.estado_resolucion_id=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id) AND r.vigencia = $1 AND e.nombre_estado IN('Expedida') ORDER BY id desc;").
		WillReturnRows(
			sqlmock.NewRows(columnasTabla).
				FromCSVString("1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1").
				FromCSVString("2"))*/

		//mock.ExpectBegin()
		//mock.ExpectQuery("SELECT DISTINCT r.id id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.facultad_id facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM resoluciones.resolucion r, resoluciones.resolucion_vinculacion_docente rv, resoluciones.resolucion_estado re, resoluciones.estado_resolucion e WHERE r.id=rv.id AND re.resolucion_id=r.id AND re.estado_resolucion_id=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id) AND r.vigencia = $1 AND e.nombre_estado IN('Expedida') ORDER BY id desc;").
		//mock.ExpectQuery("SELECT  r\\.id id FROM resoluciones\\.resolucion r;").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM resoluciones.resolucion")).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(""))
		mock.ExpectCommit()

		/*mock.ExpectPrepare("INSERT INTO").
		WillBeClosed().
		ExpectExec().WithArgs("2").WillReturnResult(sqlmock.NewResult(1, 1))
		*/
		//***************************************
		// call function to test
		r, err := http.NewRequest("GET", "/v1/resolucion_vinculacion/expedidas_vigencia_periodo?vigencia=2019", nil)
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
		//****************************************
		//db.WhateverToTest()

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
