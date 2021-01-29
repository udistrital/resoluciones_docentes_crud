package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_crud/models"
)

type ResolucionCompletaController struct {
	beego.Controller
}

func (c *ResolucionCompletaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("ResolucionTemplate", c.ResolucionTemplate)
}

// GetOne ...
// @Title Get Template
// @Description Arma una resolucion a partir de la dedicacion y el nivel.(ResolucionTemplate)
// @Param	dedicacion	path 	string	true		"Nombre de la dedicacion (HCP, HCH ...)"
// @Param	nivel	path 	string	true		"Nivel de la dedicacion (PEGRADO, POSGRADO ....)"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 404 Not found
// @Failure 500 Internal Server Error
// @Failure 502 Bad Gateway
// @router /resolucion-template/:dedicacion/:nivel [get]
func (c *ResolucionCompletaController) ResolucionTemplate() {
	defer ErrorControl(c.Controller, "ResolucionVinculacionController")

	dedicacion := c.Ctx.Input.Param(":dedicacion")
	nivel := c.Ctx.Input.Param(":nivel")

	fmt.Println("dedicacion", dedicacion, nivel)
	resolucion, err := models.GetTemplateResolucion(dedicacion, nivel)
	if err != nil {
		panic(err)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": resolucion}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ResolucionCompleta by id
// @Param	idResolucion		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 403 :idResolucion is empty
// @Failure 404 Not found
// @Failure 500 Internal Server Error
// @Failure 502 Bad Gateway
// @router /:idResolucion [get]
func (c *ResolucionCompletaController) GetOne() {
	defer ErrorControl(c.Controller, "ResolucionVinculacionController")

	idResolucion := c.Ctx.Input.Param(":idResolucion")
	resolucion, err := models.GetOneResolucionCompleta(idResolucion)

	if err != nil {
		panic(err)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": resolucion}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ResolucionCompleta
// @Param	idResolucion		path 	string	true		"The id you want to update"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 403 :idResolucion is not int
// @Failure 404 Not found
// @Failure 500 Internal Server Error
// @Failure 502 Bad Gateway
// @router /:idResolucion [put]
func (c *ResolucionCompletaController) Put() {
	defer ErrorControl(c.Controller, "ResolucionVinculacionController")

	idResolucionStr := c.Ctx.Input.Param(":idResolucion")
	idResolucion, _ := strconv.Atoi(idResolucionStr)

	v := models.ResolucionCompleta{Id: idResolucion}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateResolucionCompletaById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": ""}
		} else {
			panic(err)
		}
	} else {
		logs.Error(err)
		outputError := map[string]interface{}{"funcion": "/Put", "err": err.Error(), "status": "500"}
		panic(outputError)
	}

	c.ServeJSON()
}

/*func (c *ResolucionCompletaController) errorControl() {
	if err := recover(); err != nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ResolucionCompletaController" + "/" + (localError["funcion"]).(string))
		c.Data["data"] = (localError["err"])
		if status, ok := localError["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("404")
		}
	}
}*/
