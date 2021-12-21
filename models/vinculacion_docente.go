package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type VinculacionDocente struct {
	Id                             int                           `orm:"column(id);pk;auto"`
	NumeroContrato                 string                        `orm:"column(numero_contrato);null"`
	Vigencia                       int                           `orm:"column(vigencia);null"`
	PersonaId                      float64                       `orm:"column(persona_id)"`
	NumeroHorasSemanales           int                           `orm:"column(numero_horas_semanales)"`
	NumeroSemanas                  int                           `orm:"column(numero_semanas)"`
	PuntoSalarialId                int                           `orm:"column(punto_salarial_id);null"`
	SalarioMinimoId                int                           `orm:"column(salario_minimo_id);null"`
	ResolucionVinculacionDocenteId *ResolucionVinculacionDocente `orm:"column(resolucion_vinculacion_docente_id);rel(fk)"`
	DedicacionId                   *Dedicacion                   `orm:"column(dedicacion_id);rel(fk)"`
	ProyectoCurricularId           int16                         `orm:"column(proyecto_curricular_id)"`
	ValorContrato                  float64                       `orm:"column(valor_contrato);null"`
	Categoria                      string                        `orm:"column(categoria);null"`
	Disponibilidad                 int                           `orm:"column(disponibilidad);null"`
	DependenciaAcademica           int                           `orm:"column(dependencia_academica);null"`
	NumeroRp                       float64                       `orm:"column(numero_rp);null"`
	VigenciaRp                     float64                       `orm:"column(vigencia_rp);null"`
	FechaInicio                    time.Time                     `orm:"column(fecha_inicio);type(timestamp without time zone);null"`
	Activo                         bool                          `orm:"column(activo);null"`
	FechaCreacion                  string                        `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion              string                        `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
}

func (t *VinculacionDocente) TableName() string {
	return "vinculacion_docente"
}

func init() {
	orm.RegisterModel(new(VinculacionDocente))
}

// AddVinculacionDocente insert a new VinculacionDocente into database and returns
// last inserted Id on success.
func AddVinculacionDocente(m *VinculacionDocente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVinculacionDocenteById retrieves VinculacionDocente by Id. Returns error if
// Id doesn't exist
func GetVinculacionDocenteById(id int) (v *VinculacionDocente, err error) {
	o := orm.NewOrm()
	v = &VinculacionDocente{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVinculacionDocente retrieves all VinculacionDocente matches certain condition. Returns empty list if
// no records exist
func GetAllVinculacionDocente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VinculacionDocente)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []VinculacionDocente
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateVinculacionDocente updates VinculacionDocente by Id and returns error if
// the record to be updated doesn't exist
func UpdateVinculacionDocenteById(m *VinculacionDocente) (err error) {
	o := orm.NewOrm()
	v := VinculacionDocente{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVinculacionDocente deletes VinculacionDocente by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVinculacionDocente(id int) (err error) {
	o := orm.NewOrm()
	v := VinculacionDocente{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&VinculacionDocente{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetVinculacionesAgrupadas(id string) (v []VinculacionDocente, er error) {
	o := orm.NewOrm()

	var temp []VinculacionDocente
	_, err := o.Raw("SELECT vd.* FROM resoluciones.vinculacion_docente vd JOIN (SELECT vd.persona_id, MAX(vd.id) AS id FROM resoluciones.vinculacion_docente vd JOIN (SELECT resolucion_vinculacion_docente_id, persona_id, MAX(numero_horas_semanales) FROM resoluciones.vinculacion_docente GROUP BY resolucion_vinculacion_docente_id, persona_id)TAB1 ON vd.resolucion_vinculacion_docente_id=TAB1.resolucion_vinculacion_docente_id AND vd.persona_id=TAB1.persona_id AND vd.numero_horas_semanales=TAB1.max WHERE vd.resolucion_vinculacion_docente_id=? GROUP BY vd.persona_id)TAB1 ON vd.id = TAB1.id", id).QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp, err
}

func GetValoresTotalesPorDisponibilidad(anio, periodo, id_disponibilidad string) (totales int, er error) {
	o := orm.NewOrm()
	var temp float64

	err := o.Raw("SELECT SUM(valor_contrato) FROM resoluciones.vinculacion_docente vd, resoluciones.resolucion res WHERE vd.resolucion_vinculacion_docente_id = res.id AND res.vigencia = ? AND res.periodo = ? AND vd.disponibilidad = ?;", anio, periodo, id_disponibilidad).QueryRow(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")

	}
	return int(temp), err
}

func GetTotalContratosXResolucion(id_resolucion string, dedicacion string) (totales int, err error) {
	o := orm.NewOrm()
	var temp float64
	query := "SELECT SUM(valor_contrato)  FROM resoluciones.vinculacion_docente where resolucion_vinculacion_docente_id=?"
	if dedicacion == "TCO|MTO" {
		query = "SELECT SUM(valor) FROM (SELECT SUM(DISTINCT(valor_contrato)) AS valor FROM resoluciones.vinculacion_docente WHERE resolucion_vinculacion_docente_id=? GROUP BY persona_id) AS vinculaciones"
	}
	err = o.Raw(query, id_resolucion).QueryRow(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
		fmt.Println(int(temp))
	}
	return int(temp), err
}

func AddConjuntoVinculaciones(m []VinculacionDocente) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		beego.Error(err)
	}
	for _, vinculacion := range m {
		vinculacion.Activo = true
		vinculacion.FechaCreacion = time_bogota.TiempoBogotaFormato()
		vinculacion.FechaModificacion = time_bogota.TiempoBogotaFormato()
		id, err = o.Insert(&vinculacion)
		fmt.Println("id de vinculacion insertada", id)
		if err != nil {
			beego.Error(err)
			err = o.Rollback()
			if err != nil {
				beego.Error(err)
			}
		}
	}
	err = o.Commit()
	if err != nil {
		beego.Error(err)
	}
	return
}
