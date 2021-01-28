
CREATE SCHEMA IF NOT EXISTS resoluciones;

SET search_path TO pg_catalog,public,resoluciones;

CREATE TABLE IF NOT EXISTS resoluciones.resolucion (
	id serial NOT NULL,
	numero_resolucion character varying NOT NULL DEFAULT 10,
	fecha_expedicion timestamp,
	vigencia integer NOT NULL,
	dependencia_id integer NOT NULL,
	tipo_resolucion_id integer NOT NULL,
	preambulo_resolucion character varying(3000) NOT NULL,
	consideracion_resolucion character varying(3000) NOT NULL,
	fecha_registro timestamp NOT NULL,
	numero_semanas integer NOT NULL DEFAULT 0,
	periodo integer NOT NULL DEFAULT 0,
	titulo character varying(2000),
	dependencia_firma_id integer,
	vigencia_carga integer,
	periodo_carga integer,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion PRIMARY KEY (id)

);



CREATE TABLE IF NOT EXISTS resoluciones.dedicacion (
	id serial NOT NULL,
	nombre_dedicacion character varying(100) NOT NULL,
	descripcion character varying(5000),
	activo bool DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_dedicacion PRIMARY KEY (id)

);




CREATE TABLE IF NOT EXISTS resoluciones.vinculacion_docente (
	id serial NOT NULL,
	numero_contrato character varying(10),
	vigencia integer,
	persona_id numeric(15,0) NOT NULL,
	numero_horas_semanales integer NOT NULL,
	numero_semanas integer NOT NULL,
	punto_salarial_id integer,
	salario_minimo_id integer,
	resolucion_vinculacion_docente_id integer NOT NULL,
	dedicacion_id integer NOT NULL,
	proyecto_curricular_id smallint NOT NULL,
	fecha_registro timestamp NOT NULL,
	valor_contrato numeric(16,3),
	categoria character varying(15),
	disponibilidad integer,
	dependencia_academica integer,
	numero_rp numeric(6,0),
	vigencia_rp numeric(4,0),
	fecha_inicio timestamp,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_contrato_docente PRIMARY KEY (id),
	CONSTRAINT uq_numero_contrato_vinculacion_docente UNIQUE (numero_contrato,vigencia)

);



CREATE TABLE IF NOT EXISTS resoluciones.resolucion_vinculacion_docente (
	id serial NOT NULL,
	facultad_id integer NOT NULL,
	dedicacion character varying(12) NOT NULL,
	nivel_academico character varying(15) NOT NULL,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion_vinculacion_docente PRIMARY KEY (id)

);




CREATE TABLE IF NOT EXISTS resoluciones.componente_resolucion (
	id serial NOT NULL,
	numero integer NOT NULL,
	resolucion_id smallint NOT NULL,
	texto character varying(3000) NOT NULL,
	tipo_componente character varying(15) NOT NULL,
	componente_resolucion_padre integer,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_componente_resolucion PRIMARY KEY (id)

);




CREATE TABLE IF NOT EXISTS resoluciones.estado_resolucion (
	id serial NOT NULL,
	fecha_registro timestamp NOT NULL,
	nombre_estado character varying NOT NULL DEFAULT 20,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_estado_resolucion PRIMARY KEY (id)

);



CREATE TABLE IF NOT EXISTS resoluciones.resolucion_estado (
	id serial NOT NULL,
	fecha_registro timestamp NOT NULL,
	usuario character varying(50),
	estado_resolucion_id integer NOT NULL,
	resolucion_id integer NOT NULL,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_resolucion_estado PRIMARY KEY (id)

);





CREATE TABLE IF NOT EXISTS resoluciones.modificacion_resolucion (
	id serial NOT NULL,
	resolucion_nueva_id integer NOT NULL,
	resolucion_anterior_id integer NOT NULL,
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_modificacion_resolucion PRIMARY KEY (id)

);



CREATE TABLE IF NOT EXISTS resoluciones.modificacion_vinculacion (
	id serial NOT NULL,
	modificacion_resolucion_id integer NOT NULL,
	vinculacion_docente_cancelada_id integer NOT NULL,
	vinculacion_docente_registrada_id integer,
	horas numeric(2,0),
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_modificacion_vinculacion PRIMARY KEY (id)

);



CREATE TABLE IF NOT EXISTS resoluciones.tipo_resolucion (
	id serial NOT NULL,
	nombre_tipo_resolucion character varying(150) NOT NULL,
	descripcion character varying(5000),
	activo bool DEFAULT True,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_tipo_resolucion PRIMARY KEY (id)

);



ALTER TABLE resoluciones.resolucion ADD CONSTRAINT fk_resolucion_tipo_resolucion FOREIGN KEY (tipo_resolucion_id)
REFERENCES resoluciones.tipo_resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.vinculacion_docente ADD CONSTRAINT fk_vinculacion_docente_dedicacion FOREIGN KEY (dedicacion_id)
REFERENCES resoluciones.dedicacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.vinculacion_docente ADD CONSTRAINT fk_vinculacion_docente_resolucion_vinculacion_docente FOREIGN KEY (resolucion_vinculacion_docente_id)
REFERENCES resoluciones.resolucion_vinculacion_docente (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.resolucion_vinculacion_docente ADD CONSTRAINT fk_resolucion_vinculacion_docente_resolucion FOREIGN KEY (id)
REFERENCES resoluciones.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.componente_resolucion ADD CONSTRAINT fk_componente_resolucion_resolucion FOREIGN KEY (resolucion_id)
REFERENCES resoluciones.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.componente_resolucion ADD CONSTRAINT fk_componente_resolucion_componente_resolucion_padre FOREIGN KEY (componente_resolucion_padre)
REFERENCES resoluciones.componente_resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.resolucion_estado ADD CONSTRAINT fk_resolucion_estado_estado_resolucion FOREIGN KEY (estado_resolucion_id)
REFERENCES resoluciones.estado_resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.resolucion_estado ADD CONSTRAINT fk_resolucion_estado_resolucion FOREIGN KEY (resolucion_id)
REFERENCES resoluciones.resolucion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE resoluciones.modificacion_resolucion ADD CONSTRAINT fk_modificacion_resolucion_resolucion_nueva FOREIGN KEY (resolucion_nueva_id)
REFERENCES resoluciones.resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE resoluciones.modificacion_resolucion ADD CONSTRAINT fk_modificacion_resolucion_resolucion_anterior FOREIGN KEY (resolucion_anterior_id)
REFERENCES resoluciones.resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE resoluciones.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_vinculacion_docente_cancelada FOREIGN KEY (vinculacion_docente_cancelada_id)
REFERENCES resoluciones.vinculacion_docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE resoluciones.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_vinculacion_docente_registrada FOREIGN KEY (vinculacion_docente_registrada_id)
REFERENCES resoluciones.vinculacion_docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE resoluciones.modificacion_vinculacion ADD CONSTRAINT fk_modificacion_vinculacion_modificacion_resolucion FOREIGN KEY (modificacion_resolucion_id)
REFERENCES resoluciones.modificacion_resolucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;