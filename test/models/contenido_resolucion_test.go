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
func TestGetOneResolucionCompleta(t *testing.T) {

	resolucionCompletaTest := map[string]interface{}{
		"Data": []map[string]interface{}{
			map[string]interface{}{
				"Vinculacion": map[string]interface{}{					
					"Id": 0,
            		"FacultadId": 0,
            		"Dedicacion": "",
            		"NivelAcademico": "",
            		"Activo": false,
            		"FechaCreacion": "0001-01-01T00:00:00Z",
            		"FechaModificacion": "0001-01-01T00:00:00Z",
					
				},
				"Consideracion": "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.(Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 término Fijo por periodos académicos.\n\nQue la presente rige y se aplica únicamente a los docentes de Vinculación Especial Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional en Pregrado en lo pertinente únicamente al valor del punto.”\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue el artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue para efectos de pago de Salarios y prestaciones sociales, el periodo académico corresponde de conformidad con la Resolución No 182 de Diciembre 19 de 2017 emitida por el Consejo Académico, por la cual se expide para los proyectos curriculares de pregrado y posgrado de la Universidad Distrital Francisco José de Caldas el calendario académico para el año 2018 comprendido entre el 1° de Agosto de 2018 al 11 de Diciembre del 2018 equivalentes a 4,5 meses (mes de cuatro semanas) y el mes comprenderá (30) días.\n\nQue mediante Resolución No. 001 del 15 de Febrero de 2012 emitida por la vicerrectoría Académica se establecen los procesos de selección y vinculación de docentes de vinculación especial.\n\nQue de conformidad con el artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 072 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue para efectos presupuestales de la presente resolución se hará a cargo de la Disponibilidad presupuestal No XXXX del XX de XXXXX de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n",
				"Preambulo": "El decano de la FACULTAD TECNOLOGICA de la Universidad Distrital Francisco José de Caldas en uso de sus facultades legales y estatuarias y",
        		"Vigencia": 2018,
        		"Numero": "111",
        		"Id": 333,
				"Articulos": []map[string]interface{}{
					map[string]interface{}{
						"Id": 7869,
                		"Numero": 1,
                		"Texto": "Vincular para el Tercer Periodo Académico del año 2018 como docentes en la modalidad de Hora Cátedra de Vinculación Especial en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:\t",
                		"Paragrafos": nil,
					},
					map[string]interface{}{
						"Id": 7870,
                		"Numero": 2,
                		"Texto": "El pago de los servicios prestados por los profesores de Vinculación Especial según su escalafón, se cancelarán previa certificación de las horas efectivamente dictadas, expedida por el Decano(a) y/o Director(a) o quien haga las veces.",
						"Paragrafos": []map[string]interface{}{
							map[string]interface{}{
								"Id": 7871,
                        		"Numero": 1,
                        		"Texto": "El valor del punto en pesos para el reconocimiento y pago de los docentes de Hora Cátedra, será el que fije el Gobierno Nacional, mediante Decreto cada año y que la Universidad acoja mediante acto administrativo para los docentes de Vinculación Especial.",
							},
						},
					},
					map[string]interface{}{
						"Id": 7872,
                		"Numero": 3,
                		"Texto": "El docente deberá cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la Ley, en los Reglamentos de la Universidad y en los Planes de Trabajo entregados por el Profesor y aprobados por el Decano y/o Director.",
						"Paragrafos": []map[string]interface{}{
							map[string]interface{}{
								"Id": 7873,
                        		"Numero": 1,
                        		"Texto": "En caso de incumplimiento o retiro del docente, la Universidad mediante acto administrativo hará la liquidación con corte a la fecha del cumplido expedido por el Decano y se cancelarán las prestaciones sociales en la última liquidación del periodo académico que efectúe la División de Recursos Humanos.",
							},
						},
					},
					map[string]interface{}{
						"Id": 7874,
                		"Numero": 4,
                		"Texto": "El gasto que ocasione la presente resolución se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.",
						"Paragrafos": []map[string]interface{}{
							map[string]interface{}{
								"Id": 7875,
                        		"Numero": 1,
                        		"Texto": "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la Secretaría de Hacienda Distrital.",
							},
						},
					},
					map[string]interface{}{
						"Id": 7876,
                		"Numero": 5,
                		"Texto": "Comunicar de la presente resolución al Docente, quien manifestará bajo la gravedad de juramento que no se encuentra incurso en inhabilidades e incompatibilidades de ley para aceptar dicha vinculación especial, que no tienen con cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra en entidades de educación oficiales, siempre y cuando los honorarios sumados no correspondan a más de ocho (8) horas diarias de trabajo, que no ostente vinculación de Tiempo Completo en dos entidades ni Medio Tiempo Ocasional únicamente a excepción de hora cátedra.",
                		"Paragrafos": nil,
					},
					map[string]interface{}{
						"Id": 7877,
                		"Numero": 6,
                		"Texto": "En caso de declaratoria de suspensión de actividades académicas, por parte de los máximos organismos de gobierno de la Universidad y autorización previa del Ministerio de Trabajo, cesará para el docente la vinculación especial la obligación de prestar el servicio y para la Universidad la de pagar los salarios, pero persistirá para esta última, la de continuar efectuando los respectivos aportes a salud y pensión en el porcentaje que le corresponda.",
                		"Paragrafos": nil,
					},
					map[string]interface{}{
						"Id": 7878,
                		"Numero": 7,
                		"Texto": "La presente Resolución se expide a los XXXXXXX (XX) días del mes de XXXXXX de 2018 y surte efectos para el Tercer Periodo Académico del año 2018.",
                		"Paragrafos": nil,
					},
				},
				"Titulo": "“Por la cual se vinculan docentes para el Tercer Periodo Académico de 2018 en la modalidad de Docentes de HORA CÁTEDRA (Vinculación Especial) para la FACULTAD TECNOLOGICA de la Universidad Distrital Francisco José de Caldas en PREGRADO.”",
			},
		},
		"Message": "Request successful",
    	"Status": "200",
    	"Success": true,
	}
	
	Convey("Subject: Test Get One Resolucion Completa\n", t, func() {
		res_resolucion, err := models.GetOneResolucionCompleta("333")
		//var p map[string]interface{}
		//t.Log(res_resolucion)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {

			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			So(res_resolucion, ShouldEqual, resolucionCompletaTest)
		})
	})
}

func TestGetTemplateResolucion(t *testing.T) {
	paragrafos := map[string]interface{}{
		"Id": 0,
		"Numero": 0,
		"Texto": "string",
	}
	articulos := map[string]interface{}{
		"Id": 0,
		"Numero": 0,
		"Paragrafos": paragrafos,
		"Texto": "string",
	  }
	vinculación := map[string]interface{}{
		"Activo": true,
		"Dedicacion": "string",
		"FacultadId": 0,
		"FechaCreacion": "string",
		"FechaModificacion": "string",
		"Id": 0,
		"NivelAcademico": "string",
	}
	resolucionCompleta := map[string]interface{}{
		"Articulos": articulos,
		"Consideracion": "string",
		"Id": 0,
		"Numero": "string",
		"Preambulo": "string",
		"Titulo": "string",
		"Vigencia": 0,
		"Vinculacion": vinculación,
	}
	Convey("Subject: Test Get Template Resolucion\n", t, func() {
		res_template_resolucion, err := models.GetTemplateResolucion("asd", "2")
		//t.Log(res_template_resolucion)
		Convey("Error is null", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("The Result Should Not Be Empty", func() {

			//t.Log("Arreglo: ", res_vin)
			//t.Log("Tamano arreglo: ", len(res_vin))
			So(res_template_resolucion, ShouldBeGreaterThan, resolucionCompleta)
		})
	})
}
