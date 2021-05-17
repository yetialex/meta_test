truncate table iba_gates cascade;
truncate table iba_gates_sources cascade;
truncate table nodes cascade;
truncate table directories cascade;
truncate table value_types cascade;
truncate table source_classes cascade;
truncate table signal_classes cascade;


-- default directories
insert into directories (id,name,description)
OVERRIDING SYSTEM VALUE
values
(1,'sap_tm', 'справочник техмест'),
(2,'iba_metadata', 'метаданные от IBA серверов')
;

-- default value_types
insert into value_types(id,name)
OVERRIDING SYSTEM VALUE
values
(1,'bool'),
(2,'byte'),
(3,'int'),
(4,'long'),
(5,'real'),
(6,'double'),
(7,'string')
;

-- default source_classes
insert into source_classes(id,name,comment)
OVERRIDING SYSTEM VALUE
values
(1,'core','Core signals'),
(2,'iba','IBA batch mode'),
(3,'iba-rt','IBA realtime mode')
;


-- default signal_classes
insert into signal_classes(id,name)
OVERRIDING SYSTEM VALUE
values
(1,'analog'),
(2,'discrete'),
(3,'analog calculated'),
(4,'discrete calculated'),
(5,'virtual')
;
