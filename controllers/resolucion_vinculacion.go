package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_crud/models"
	"strings"
)

type ResolucionVinculacionController struct {
	beego.Controller
}

func (c *ResolucionVinculacionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllAprobada", c.GetAllAprobada)
	c.Mapping("GetAllExpedidasVigenciaPeriodo", c.GetAllExpedidasVigenciaPeriodo)
	c.Mapping("GetAllExpedidasVigenciaPeriodoVinculacion", c.GetAllExpedidasVigenciaPeriodoVinculacion)
}

// GetAll ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ResolucionVinculacionDocente
// @Failure 403
// @router / [get]
func (c *ResolucionVinculacionController) GetAll() {
	defer c.errorControl()

	//var fields []string
	//var sortby []string
	//var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	/*if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}*/
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	/*if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}*/
	// order: desc,asc
	/*if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}*/
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	//l, err := models.GetAllResolucionVinculacion(query, fields, sortby, order, offset, limit)
	l, err := models.GetAllResolucionVinculacion(query, offset, limit)
	if err != nil {
		panic(err)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": l}
	}
	c.ServeJSON()
}

// GetAllAprobada ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	int		false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int		false	"Start position of result set. Must be an integer"
// @Success 201 {object} models.ResolucionVinculacionDocente
// @Failure 403
// @router /Aprobada [get]
func (c *ResolucionVinculacionController) GetAllAprobada() {

	defer c.errorControl()

	//var fields []string
	//var sortby []string
	//var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	/*if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}*/
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	/*if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}*/
	// order: desc,asc
	/*if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}*/
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err := errors.New("Invalid query key/value pair")
				logs.Error(err)
				outputError := map[string]interface{}{"funcion": "/GetAllAprobada", "err": err, "status": "500"}
				panic(outputError)
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	if listaResoluciones, err := models.GetAllResolucionAprobada(query, offset, limit); err != nil {
		panic(err)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": listaResoluciones}
	}
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodo ...
// @Title GetAllExpedidasVigenciaPeriodo
// @Description Agrupa los contratos de una preliquidacion segun mes, año y nomina para preliquidaicones en estado CERRADA
// @Param vigencia query string false "nomina a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodo() {

	vigencia, err := c.GetInt("vigencia")
	if err == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodo(vigencia)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodoVinculacion ...
// @Title GetAllExpedidasVigenciaPeriodoVinculacion
// @Description Muestra resoluciones de tipo vinculación para cancelar y modificar
// @Param vigencia query string false "nomina a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo_vinculacion [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodoVinculacion() {

	vigencia, err := c.GetInt("vigencia")
	if err == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodoVinculacion(vigencia)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err)
		c.Abort("403")
	}
	c.ServeJSON()
}

func (c *ResolucionVinculacionController) errorControl() {
	if err := recover(); err != nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ResolucionVinculacionController" + "/" + (localError["funcion"]).(string))
		c.Data["data"] = (localError["err"])
		if status, ok := localError["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("404")
		}
	}
}
