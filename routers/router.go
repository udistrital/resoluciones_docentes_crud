// @APIVersion 1.1.0
// @Title Resoluciones CRUD API
// @Description Crud del aplicativo resoluciones, se usa para generar las resoluciones que afectan a los docentes de vinculacion especial, las cuales se usan para vincular, adicionar, reducir un contrato.
// @Contact computo@udistrital.edu.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Author Oficina Asesora de Sistemas, Universidad Distrital
package routers

import (
	"github.com/udistrital/resoluciones_docentes_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/contenido_resolucion",
			beego.NSInclude(
				&controllers.ResolucionCompletaController{},
			),
		),

		beego.NSNamespace("/estado_resolucion",
			beego.NSInclude(
				&controllers.EstadoResolucionController{},
			),
		),

		beego.NSNamespace("/resolucion_estado",
			beego.NSInclude(
				&controllers.ResolucionEstadoController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion",
			beego.NSInclude(
				&controllers.ResolucionVinculacionController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/modificacion_resolucion",
			beego.NSInclude(
				&controllers.ModificacionResolucionController{},
			),
		),

		beego.NSNamespace("/modificacion_vinculacion",
			beego.NSInclude(
				&controllers.ModificacionVinculacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
