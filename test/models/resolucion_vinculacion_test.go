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
	//var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	/*db, err = sql.Open("postgres", "dbname=myapp_test")
	if err != nil {
		fmt.Println("SE MURIO")
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),                   // You database connection
		testfixtures.Dialect("postgres"),            // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata/fixtures"), // the directory containing the YAML files
	)
	if err != nil {
		fmt.Println("SE MURIO")
	}

	os.Exit(t.Run())*/

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgUrls := os.Getenv("POSTGRES_HOST")
	pgDb := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_DB_PORT")
	pgSchema := os.Getenv("POSTGRES_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+":"+pgPort+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")
	//f(t)

}

//POSTGRES_SCHEMA=resoluciones POSTGRES_HOST=localhost POSTGRES_DB_PORT=5432 POSTGRES_DB=resoluciones_db POSTGRES_USER=resoluciones POSTGRES_PASSWORD=resoluciones go test -v ./...
func TestGetAllExpedidasVigenciaPeriodo(t *testing.T) {

	Convey("Subject: Test Contenido Resolucion Endpoint\n", t, func() {
		res_vin, err := models.GetAllExpedidasVigenciaPeriodo(2019)
		t.Log(res_vin)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		/*Convey("The Result Should Not Be Empty", func() {
			So(len(res_vin), ShouldBeGreaterThan, 0)
		})*/
	})
}
