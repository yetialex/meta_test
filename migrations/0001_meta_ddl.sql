CREATE TABLE Signals (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text,
  "description" text,
  "parent_id" bigint,
  "unit" text,
  "value_type_id" bigint,
  "class_id" bigint,
  "external_id" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "acl" json
);

-- CREATE TABLE Signals_history (
--   "id" bigint,
--   "name" text,
--   "description" text,
--   "parent_id" bigint,
--   "unit" text,
--   "value_type_id" bigint,
--   "class_id" bigint,
--   "external_id" text,
--   "created_at" timestamp DEFAULT (now()),
--   "updated_at" timestamp DEFAULT (now()),
--   "acl" json
-- );

CREATE TABLE signal_classes (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text
);

CREATE TABLE value_types (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text
);

CREATE TABLE Nodes (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text,
  "parent_id" bigint,
  "description" text,
  "comment" text,
  "meta" text,
  "full_name" text,
  "signal_id" bigint,
  "directory_id" bigint,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "acl" json
);

-- CREATE TABLE Nodes_history (
--   "id" BIGSERIAL PRIMARY KEY,
--   "name" text,
--   "parent_id" bigint,
--   "description" text,
--   "comment" text,
--   "meta" text,
--   "full_name" text,
--   "signal_id" bigint,
--   "directory_id" bigint,
--   "created_at" timestamp DEFAULT (now()),
--   "updated_at" timestamp DEFAULT (now()),
--   "acl" json
-- );

CREATE TABLE Directories (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text,
  "description" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "acl" json
);

-- CREATE TABLE sap_functional_locations (
--   "tplnr" text PRIMARY KEY,
--   "zex" text NOT NULL,
--   "zex_ktext" text NOT NULL,
--   "pltxt" text NOT NULL,
--   "fltyp" text NOT NULL,
--   "eqart" text NOT NULL,
--   "eartx" text NOT NULL,
--   "serge" text,
--   "invnr" text,
--   "kostl" text,
--   "stort" text,
--   "anlnr" text,
--   "anlun_sub" text,
--   "arbpl" text,
--   "rfr_ladle" int,
--   "erdat" text NOT NULL,
--   "aedat" text,
--   "zz_tm_rfid" text,
--   "zz_tm_tag" text,
--   "objnr" text NOT NULL,
--   "status_object" text NOT NULL,
--   "status" text NOT NULL,
--   "timestamp" text NOT NULL,
--   "beber" text,
--   "swerk" text,
--   "iwerk" text,
--   "tplma" text
-- );

CREATE TABLE iba_servers (
  "id" SERIAL PRIMARY KEY,
  "name" text NOT NULL,
  "zex" text,
  "stort" text,
  "comment" text
);

CREATE TABLE iba_rt_topics (
  "iba_server_id" bigint,
  "topic" text UNIQUE,
  "comment" text
);

CREATE TABLE iba_gates (
  "id" text,
  "comment" text
);

-- CREATE TABLE iba_batch_links (
--   "iba_server_id" bigint,
--   "gate_id" text,
--   "mnt" text,
--   "comment" text
-- );

ALTER TABLE Signals ADD FOREIGN KEY ("parent_id") REFERENCES Signals ("id");

ALTER TABLE Signals ADD FOREIGN KEY ("value_type_id") REFERENCES value_types ("id");

ALTER TABLE Signals ADD FOREIGN KEY ("class_id") REFERENCES signal_classes ("id");

-- ALTER TABLE Signals_history ADD FOREIGN KEY ("parent_id") REFERENCES Signals_history ("id");
--
-- ALTER TABLE Signals_history ADD FOREIGN KEY ("value_type_id") REFERENCES value_types ("id");
--
-- ALTER TABLE Signals_history ADD FOREIGN KEY ("class_id") REFERENCES signal_classes ("id");

ALTER TABLE Nodes ADD FOREIGN KEY ("parent_id") REFERENCES Nodes ("id");

ALTER TABLE Nodes ADD FOREIGN KEY ("signal_id") REFERENCES Signals ("id");

ALTER TABLE Nodes ADD FOREIGN KEY ("directory_id") REFERENCES Directories ("id");

-- ALTER TABLE "Nodes_history" ADD FOREIGN KEY ("parent_id") REFERENCES Nodes_history ("id");
--
-- ALTER TABLE "Nodes_history" ADD FOREIGN KEY ("signal_id") REFERENCES "Signals" ("id");
--
-- ALTER TABLE "Nodes_history" ADD FOREIGN KEY ("directory_id") REFERENCES "Directories" ("id");

-- ALTER TABLE "iba_servers" ADD FOREIGN KEY ("zex") REFERENCES "sap_functional_locations" ("zex");
--
-- ALTER TABLE "iba_servers" ADD FOREIGN KEY ("stort") REFERENCES "sap_functional_locations" ("stort");

-- ALTER TABLE "iba_servers" ADD FOREIGN KEY ("id") REFERENCES "iba_batch_links" ("iba_server_id");
--
-- ALTER TABLE "iba_gates" ADD FOREIGN KEY ("id") REFERENCES "iba_batch_links" ("gate_id");

CREATE UNIQUE INDEX ON Signals ("external_id");

COMMENT ON TABLE Signals IS 'метаданные о сигналах';

COMMENT ON COLUMN Signals."external_id" IS 'ключ сигнала во внешней системе';

-- COMMENT ON TABLE "Signals_history" IS 'метаданные о сигналах';
--
-- COMMENT ON COLUMN "Signals_history"."external_id" IS 'ключ сигнала во внешней системе';

COMMENT ON COLUMN signal_classes."name" IS 'analog, discrete, virtual';

COMMENT ON COLUMN value_types."name" IS 'Тип данных сигнала: boolean,integer,real,double,undefined';

COMMENT ON TABLE Nodes IS 'логическая топология';

COMMENT ON COLUMN Nodes."full_name" IS 'root/sub/leaf/node, заполнять при редактировании узла автоматически';

-- COMMENT ON TABLE "Nodes_history" IS 'логическая топология';

-- COMMENT ON COLUMN "Nodes_history"."full_name" IS 'root/sub/leaf/node, заполнять при редактировании узла автоматически';

COMMENT ON TABLE Directories IS 'справочниках логической топологии';

-- COMMENT ON TABLE "sap_functional_locations" IS 'Справочник ТехМест';
--
-- COMMENT ON COLUMN "sap_functional_locations"."tplnr" IS 'Номер ТМ';
--
-- COMMENT ON COLUMN "sap_functional_locations"."zex" IS 'Цех';
--
-- COMMENT ON COLUMN "sap_functional_locations"."zex_ktext" IS 'Наименование цеха';
--
-- COMMENT ON COLUMN "sap_functional_locations"."pltxt" IS 'Наименование ТМ';
--
-- COMMENT ON COLUMN "sap_functional_locations"."fltyp" IS 'Тип';
--
-- COMMENT ON COLUMN "sap_functional_locations"."eqart" IS 'Тип объекта';
--
-- COMMENT ON COLUMN "sap_functional_locations"."eartx" IS 'Наименование типа объекта';
--
-- COMMENT ON COLUMN "sap_functional_locations"."serge" IS 'Серийный номер изготовителя';
--
-- COMMENT ON COLUMN "sap_functional_locations"."invnr" IS 'Инвентарный номер';
--
-- COMMENT ON COLUMN "sap_functional_locations"."kostl" IS 'МВЗ';
--
-- COMMENT ON COLUMN "sap_functional_locations"."stort" IS 'Местоположение';
--
-- COMMENT ON COLUMN "sap_functional_locations"."anlnr" IS 'Основное средство';
--
-- COMMENT ON COLUMN "sap_functional_locations"."anlun_sub" IS 'Субномеросн.средства';
--
-- COMMENT ON COLUMN "sap_functional_locations"."arbpl" IS 'Отв.раб.место';
--
-- COMMENT ON COLUMN "sap_functional_locations"."rfr_ladle" IS 'Номер ковша';
--
-- COMMENT ON COLUMN "sap_functional_locations"."erdat" IS 'Дата cоздания';
--
-- COMMENT ON COLUMN "sap_functional_locations"."aedat" IS 'Дата зменения';
--
-- COMMENT ON COLUMN "sap_functional_locations"."zz_tm_rfid" IS 'ID метки';
--
-- COMMENT ON COLUMN "sap_functional_locations"."zz_tm_tag" IS 'Обозначение метки';
--
-- COMMENT ON COLUMN "sap_functional_locations"."objnr" IS 'Номер объекта';
--
-- COMMENT ON COLUMN "sap_functional_locations"."status_object" IS 'Статус объекта';
--
-- COMMENT ON COLUMN "sap_functional_locations"."status" IS 'Статус';
--
-- COMMENT ON COLUMN "sap_functional_locations"."timestamp" IS 'Дата загрузки';
--
-- COMMENT ON COLUMN "sap_functional_locations"."beber" IS 'Произв Участок';
--
-- COMMENT ON COLUMN "sap_functional_locations"."swerk" IS 'Завод расположения';
--
-- COMMENT ON COLUMN "sap_functional_locations"."iwerk" IS 'Завод планирующий';
--
-- COMMENT ON COLUMN "sap_functional_locations"."tplma" IS 'Вышестоящее ТМ';

COMMENT ON TABLE iba_servers IS 'Реестр IBA серверов';

COMMENT ON COLUMN iba_servers."name" IS 'имя сервера';

COMMENT ON COLUMN iba_servers."zex" IS 'цех';

COMMENT ON COLUMN iba_servers."stort" IS 'место устанкови';

COMMENT ON COLUMN iba_servers."comment" IS 'заметки';

COMMENT ON TABLE iba_rt_topics IS 'Реестр топиков с RT данными';

COMMENT ON COLUMN iba_rt_topics."comment" IS 'заметки';

COMMENT ON TABLE iba_gates IS 'Реестр экспортеров Dat файлов в S3';

COMMENT ON COLUMN iba_gates."comment" IS 'заметки';

-- COMMENT ON COLUMN "iba_batch_links"."mnt" IS 'имя папки шары';
--
-- COMMENT ON COLUMN "iba_batch_links"."comment" IS 'заметки';

---- create above / drop below ----
DROP TABLE Signals;
