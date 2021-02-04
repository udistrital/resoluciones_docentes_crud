# resoluciones_docentes_crud

El API provee la gestion de las diferentes procesos que requiere el sistema de resoluciones


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
# parametros de api
RESOLUCIONES_CRUD_HTTP_PORT=[Puerto de exposición del API]
# paramametros de bd
RESOLUCIONES_CRUD_PGUSER=[Usuario de BD]
RESOLUCIONES_CRUD_PGPASS=[Contraseña del usaurio de BD]
RESOLUCIONES_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
RESOLUCIONES_CRUD_PGPORT=[Puerto de la BD]
RESOLUCIONES_CRUD_PGDB=[Nombre de Base de Datos]
RESOLUCIONES_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con RESOLUCIONES_CRUD...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/resoluciones_docentes_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/resoluciones_docentes_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
RESOLUCIONES_CRUD_HTTP_PORT=8080 RESOLUCIONES_CRUD_SOME_VARIABLE bee run
```

### Ejecución Dockerfile


### Ejecución docker-compose


### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI


| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/resoluciones_docentes_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/resoluciones_docentes_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/resoluciones_docentes_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/resoluciones_docentes_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/resoluciones_docentes_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/resoluciones_docentes_crud) |


## Licencia

This file is part of resoluciones_docentes_crud.

resoluciones_docentes_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

resoluciones_docentes_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.

