INSERT INTO resoluciones.estado_resolucion(fecha_registro, nombre_estado, activo) VALUES
    ('2017-08-13 00:00:00', 'Solicitada', true),
    ('2017-08-13 00:00:00', 'Expedida', true),
    ('2017-08-13 00:00:00', 'Cancelada', true),
    ('2017-11-03 00:00:00', 'RP Solicitado', true),
    ('2017-12-15 00:00:00', 'Aprobada', true),
    ('2018-02-07 00:00:00', 'Anulada', true);

INSERT INTO resoluciones.tipo_resolucion(nombre_tipo_resolucion, descripcion, activo) VALUES
    ('Vinculación', 'Resolución que permite vincular docentes', true),
    ('Cancelación', 'Resolución que se crea para anular una ya existente y que corta el vinculo de los docentes elegidos con la Universidad', true),
    ('Adición', 'Resolución que se crea para adicionar horas a los docentes elegidos de la resolución que reemplaza', true),
    ('Vinculación', 'Resolución que se crea para quitar horas a los docentes elegidos de la resolución a la que reemplaza ', true);

INSERT INTO resoluciones.dedicacion(nombre_dedicacion, descripcion, activo) VALUES
    ('HCH', 'HORA CATEDRA HONORARIOS', true),
    ('HCP', 'HORA CATEDRA PRESTACIONES', true),
    ('MTO', 'MEDIO TIEMPO OCASIONAL', true),
    ('TCO', 'TIEMPO COMPLETO OCASIONAL', true);