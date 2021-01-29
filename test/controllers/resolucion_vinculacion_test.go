package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/udistrital/resoluciones_crud/routers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

var o orm.Ormer

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
		columns := []string{"id"}
		mock.ExpectQuery("SELECT (.+) FROM `XXX`").
			WillReturnRows(
				sqlmock.NewRows(columns).
					FromCSVString("1").
					FromCSVString("2"))
			//***************************************
		// call function to test
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
		//****************************************
		//db.WhateverToTest()

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
