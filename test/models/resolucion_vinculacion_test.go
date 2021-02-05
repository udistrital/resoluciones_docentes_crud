package models

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/resoluciones_crud/models"
	_ "github.com/udistrital/resoluciones_crud/routers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgUrls := os.Getenv("POSTGRES_HOST")
	pgDb := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_DB_PORT")
	pgSchema := os.Getenv("POSTGRES_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

}

//POSTGRES_SCHEMA=resoluciones POSTGRES_HOST=localhost POSTGRES_DB_PORT=5432 POSTGRES_DB=resoluciones_db POSTGRES_USER=resoluciones POSTGRES_PASSWORD=resoluciones go test -v ./...
//POSTGRES_SCHEMA=resoluciones POSTGRES_HOST=pgtst.udistritaloas.edu.co POSTGRES_DB_PORT=5432 POSTGRES_DB=udistrital_administrativa POSTGRES_USER=desarrollooas POSTGRES_PASSWORD=W7Sz1lbWFwfE798b go test -v ./...

// Prueba unitaria para el metodo GetAllResolucionVinculacion
func TestGetAllResolucionVinculacion(t *testing.T) {

	query := make(map[string]string)
	//query["v1"] = "v1"
	Convey("Subject: Test GetAllResolucionVinculacion Models \n", t, func() {
		res_vin, err := models.GetAllResolucionVinculacion(query, 0, 0)
		t.Log(res_vin)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			//So(len(res_vin), ShouldBeGreaterThan, 0)
		})
	})
}

// Prueba unitaria para el metodo GetAllResolucionAprobada
func TestGetAllResolucionAprobada(t *testing.T) {

	query := make(map[string]string)
	//query["v1"] = "v1"
	Convey("Subject: Test GetAllResolucionAprobada Models \n", t, func() {
		res_vin, err := models.GetAllResolucionAprobada(query, 0, 0)
		t.Log(res_vin)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			//So(len(res_vin), ShouldBeGreaterThan, 0)
		})
	})
}

// Prueba unitaria para el metodo GetAllExpedidasVigenciaPeriodo
func TestGetAllExpedidasVigenciaPeriodo(t *testing.T) {

	Convey("Subject: Test GetAllExpedidasVigenciaPeriodo Models \n", t, func() {
		res_vin, err := models.GetAllExpedidasVigenciaPeriodo(2019)
		t.Log(res_vin)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			//So(len(res_vin), ShouldBeGreaterThan, 0)
		})
	})
}

// Prueba unitaria para el metodo GetAllExpedidasVigenciaPeriodoVinculacion
func TestGetAllExpedidasVigenciaPeriodoVinculacion(t *testing.T) {

	Convey("Subject: Test GetAllExpedidasVigenciaPeriodoVinculacion - Models\n", t, func() {
		res_vin, err := models.GetAllExpedidasVigenciaPeriodoVinculacion(2019)
		t.Log(res_vin)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {
			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			//So(len(res_vin), ShouldBeGreaterThan, 0)
		})
	})
}
