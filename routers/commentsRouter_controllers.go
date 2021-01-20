package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
