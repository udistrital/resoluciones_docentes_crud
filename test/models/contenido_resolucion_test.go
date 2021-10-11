package models

import (
	"encoding/json"
	"testing"

	"github.com/udistrital/resoluciones_docentes_crud/models"
	_ "github.com/udistrital/resoluciones_docentes_crud/routers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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

//POSTGRES_SCHEMA=resoluciones POSTGRES_HOST=localhost POSTGRES_DB_PORT=5432 POSTGRES_DB=resoluciones_db POSTGRES_USER=resoluciones POSTGRES_PASSWORD=resoluciones go test -v ./...
//POSTGRES_SCHEMA=resoluciones POSTGRES_HOST=pgtst.udistritaloas.edu.co POSTGRES_DB_PORT=5432 POSTGRES_DB=udistrital_administrativa POSTGRES_USER=desarrollooas POSTGRES_PASSWORD=W7Sz1lbWFwfE798b go test -v ./...
func TestGetOneResolucionCompleta(t *testing.T) {

	m := make(map[string]interface{})
	p := map[string]interface{}(nil)

	Convey("Subject: Test Get One Resolucion Completa\n", t, func() {
		res_resolucion, err := models.GetOneResolucionCompleta("333")
		temp, _ := json.Marshal(res_resolucion)
		json.Unmarshal(temp, &m)
		Convey("Error is null", func() {
			So(err, ShouldResemble, p)
		})
		Convey("The Result Should Not Be Empty", func() {

			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			So(len(m), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetTemplateResolucion(t *testing.T) {

	m := make(map[string]interface{})

	Convey("Subject: Test Get Template Resolucion\n", t, func() {
		res_template_resolucion, err := models.GetTemplateResolucion("HPC", "PREGRADO")
		temp, _ := json.Marshal(res_template_resolucion)
		json.Unmarshal(temp, &m)
		//t.Log(res_template_resolucion)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {

			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			So(len(m), ShouldBeGreaterThan, 0)
		})
	})
}
