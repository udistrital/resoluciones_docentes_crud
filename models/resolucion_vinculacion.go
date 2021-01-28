package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
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

const DEFAULTMAXTEMS = 10000

var (
	columnNames = make(map[string]string)
)

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
	t := reflect.TypeOf(ResolucionVinculacion{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("orm")
		column := ""
		_, err := fmt.Sscanf(strings.Split(tag, ";")[0], "column(%s )", &column)
		if err != nil {
			beego.Error(err)
		}
		columnNames[field.Name] = column[:len(column)-1]
	}
}

// se quita fields, stortBy y order porque no se usan
//func GetAllResolucionVinculacion(query map[string]string, fields []string, sortby []string, order []string,
func GetAllResolucionVinculacion(query map[string]string, offset int64, limit int64) (ml []ResolucionVinculacion, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Hola")
			outputError = map[string]interface{}{"funcion": "/GetAllResolucionVinculacion", "err": err, "status": "500"}
			return
		}
	}()

	o := orm.NewOrm()

	if limit == 0 {
		limit = DEFAULTMAXTEMS
	}

	qs := make([]Operation, 0)
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = append(qs, filter(k, (v == "true" || v == "1")))
		} else {
			qs = append(qs, filter(k, v))
		}
	}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionVinculacion1", "err": err.Error(), "status": "500"}
		return
	}

	// _, err = o.Raw(`
	qb.Select(
		"DISTINCT r.id id",
		"e.nombre_estado estado",
		"r.numero_resolucion numero",
		"r.vigencia vigencia",
		"r.periodo periodo",
		"rv.facultad_id facultad",
		"rv.nivel_academico nivel_academico",
		"rv.dedicacion dedicacion",
		"r.numero_semanas numero_semanas",
		"r.fecha_expedicion fecha_expedicion",
		"tipo.nombre_tipo_resolucion tipo_resolucion",
		"r.dependencia_firma_id dependencia_firma",
		"r.vigencia_carga vigencia_carga",
		"r.periodo_carga periodo_carga").
		From(
			"resoluciones.resolucion r",
			"resoluciones.resolucion_vinculacion_docente rv",
			"resoluciones.resolucion_estado re",
			"resoluciones.estado_resolucion e",
			"resoluciones.tipo_resolucion tipo").
		Where("r.id=rv.id").
		And("re.resolucion_id=r.id").
		And("r.tipo_resolucion_id=tipo.id").
		And("re.estado_resolucion_id=e.id").
		And("re.estado_resolucion_id!=6").
		And("re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id)")

	qb2, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionVinculacion2", "err": err.Error(), "status": "500"}
		return ml, outputError
	}
	qb2.Select("*").From(qb.Subquery(qb.String(), "T"))

	// query externo
	flag := true
	for _, v := range qs {
		columnName, ok := columnNames[v.Field]
		if !ok {
			err := errors.New("Inexistent field in query")
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/GetAllResolucionVinculacion3", "err": err, "status": "500"}
			return ml, outputError
		}
		tmp := fmt.Sprintf("T.%s::VARCHAR %s", columnName, v.Op)
		if flag {
			qb2.Where(tmp)
			flag = false
		} else {
			qb2.And(tmp)
		}
	}
	qb2.OrderBy("id").
		Desc().
		Limit(int(limit)).
		Offset(int(offset))

	_, err = o.Raw(qb2.String()).QueryRows(&ml)
	if err != nil {

		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionVinculacion4", "err": err.Error(), "status": "500"}
		return ml, outputError
	}

	for x, resoluciones := range ml {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		ml[x].FechaExpedicion = resoluciones.FechaExpedicion
	}
	return ml, nil
}

// Hay parametros que no se usan y se eliminan fields, stortBy y order
//func GetAllResolucionAprobada(query map[string]string, fields []string, sortby []string, order []string,
func GetAllResolucionAprobada(query map[string]string, offset int64, limit int64) (arregloIDs []ResolucionVinculacion, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Hola")
			outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada", "err": err, "status": "500"}
			return
		}
	}()

	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	//_, err := o.Raw("SELECT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, oikos.dependencia d, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.id_tipo_resolucion=1 ORDER BY id desc;").QueryRows(&temp)
	//TODO: dar soporte a query (sin dejar que sea vulnerable a SQL injection)

	if limit == 0 {
		limit = DEFAULTMAXTEMS
	}

	qs := make([]Operation, 0)
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = append(qs, filter(k, (v == "true" || v == "1")))
		} else {
			qs = append(qs, filter(k, v))
		}
	}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada1", "err": err.Error(), "status": "500"}
		return arregloIDs, outputError
	}
	qb.Select(
		"DISTINCT r.id id",
		"e.nombre_estado estado",
		"r.numero_resolucion numero",
		"r.vigencia vigencia",
		"r.periodo periodo",
		"rv.facultad_id facultad",
		"rv.nivel_academico nivel_academico",
		"rv.dedicacion dedicacion",
		"r.numero_semanas numero_semanas",
		"r.fecha_expedicion fecha_expedicion",
		"tr.nombre_tipo_resolucion tipo_resolucion",
		"r.dependencia_firma_id dependencia_firma",
	).
		From(
			"resoluciones.resolucion r",
			"resoluciones.resolucion_vinculacion_docente rv",
			"resoluciones.resolucion_estado re",
			"resoluciones.estado_resolucion e",
			"resoluciones.tipo_resolucion tr",
		).
		Where("r.id=rv.id").
		And("re.resolucion_id=r.id").
		And("re.estado_resolucion_id=e.id").
		And("re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id) AND e.nombre_estado IN('Aprobada','Expedida')").
		And("tr.id=r.tipo_resolucion_id")

	qb2, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada2", "err": err.Error(), "status": "500"}
		return arregloIDs, outputError
	}
	qb2.Select("*").
		From(qb.Subquery(qb.String(), "T"))

	// query externo
	flag := true
	for _, v := range qs {
		columnName, ok := columnNames[v.Field]
		beego.Debug(columnName)
		if !ok {
			err := errors.New("Inexistent field in query")
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada2", "err": err.Error(), "status": "500"}
			return arregloIDs, outputError
		}
		tmp := fmt.Sprintf("T.%s::VARCHAR %s", columnName, v.Op)
		if flag {
			qb2.Where(tmp)
			flag = false
		} else {
			qb2.And(tmp)
		}
	}

	qb2.OrderBy("id").
		Desc().
		Limit(int(limit)).
		Offset(int(offset))

	_, err = o.Raw(qb2.String()).QueryRows(&temp)
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada2", "err": err.Error(), "status": "500"}
		return arregloIDs, outputError
	}
	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}
	return temp, nil
}

func GetAllExpedidasVigenciaPeriodo(vigencia int) (arregloIDs []ResolucionVinculacion, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada", "err": err, "status": "500"}
			return
		}
	}()

	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT DISTINCT r.id id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.facultad_id facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM resoluciones.resolucion r, resoluciones.resolucion_vinculacion_docente rv, resoluciones.resolucion_estado re, resoluciones.estado_resolucion e WHERE r.id=rv.id AND re.resolucion_id=r.id AND re.estado_resolucion_id=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id) AND r.vigencia = ? AND e.nombre_estado IN('Expedida') ORDER BY id desc;", vigencia).QueryRows(&temp)

	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllResolucionAprobada1", "err": err.Error(), "status": "500"}
		return arregloIDs, outputError
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}

	return temp, nil
}

func GetAllExpedidasVigenciaPeriodoVinculacion(vigencia int) (arregloIDs []ResolucionVinculacion, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetAllExpedidasVigenciaPeriodoVinculacion", "err": err, "status": "500"}
			return
		}
	}()

	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT DISTINCT r.id id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.facultad_id facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM resoluciones.resolucion r, resoluciones.resolucion_vinculacion_docente rv, resoluciones.resolucion_estado re, resoluciones.estado_resolucion e WHERE r.id=rv.id AND re.resolucion_id=r.id AND re.estado_resolucion_id=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM resoluciones.resolucion_estado re_aux WHERE re_aux.resolucion_id=r.id) AND r.vigencia = ? AND e.nombre_estado IN('Expedida') AND r.tipo_resolucion_id IN (1,3,4) ORDER BY id desc;", vigencia).QueryRows(&temp)

	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetAllExpedidasVigenciaPeriodoVinculacion", "err": err.Error(), "status": "500"}
		return arregloIDs, outputError
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}

	return temp, nil
}
