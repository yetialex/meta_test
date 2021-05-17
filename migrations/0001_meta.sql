CREATE TABLE "signals" (
                           "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                           "value_type_id" bigint,
                           "class_id" bigint,
                           "sources_id" bigint,
                           "status_id" int DEFAULT 0,
                           "created_at" timestamp DEFAULT (now()),
                           "updated_at" timestamp DEFAULT (now()),
                           "acl" json
);

CREATE TABLE "status" (
                          "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                          "name" text
);

CREATE TABLE "sources" (
                           "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                           "parent_id" bigint,
                           "name" text UNIQUE NOT NULL,
                           "source_classes_id" bigint NOT NULL,
                           "comment" text
);

CREATE TABLE "source_classes" (
                                  "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                  "name" text UNIQUE,
                                  "comment" text
);

CREATE TABLE "signal_classes" (
                                  "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                  "name" text
);

CREATE TABLE "value_types" (
                               "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               "name" text UNIQUE
);

CREATE TABLE "nodes" (
                         "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                         "name" text,
                         "parent_id" bigint,
                         "description" text,
                         "comment" text,
                         "meta" text,
                         "full_name" text,
                         "directory_id" bigint,
                         "signal_id" bigint,
                         "created_at" timestamp DEFAULT (now()),
                         "updated_at" timestamp DEFAULT (now()),
                         "acl" json
);

CREATE TABLE "directories" (
                               "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               "name" text,
                               "description" text,
                               "created_at" timestamp DEFAULT (now()),
                               "updated_at" timestamp DEFAULT (now()),
                               "acl" json
);

CREATE TABLE "sap_functional_locations" (
                                            "tplnr" text PRIMARY KEY,
                                            "zex" text NOT NULL,
                                            "zex_ktext" text NOT NULL,
                                            "pltxt" text NOT NULL,
                                            "fltyp" text NOT NULL,
                                            "eqart" text NOT NULL,
                                            "eartx" text NOT NULL,
                                            "serge" text,
                                            "invnr" text,
                                            "kostl" text,
                                            "stort" text,
                                            "anlnr" text,
                                            "anlun_sub" text,
                                            "arbpl" text,
                                            "rfr_ladle" int,
                                            "erdat" text NOT NULL,
                                            "aedat" text,
                                            "zz_tm_rfid" text,
                                            "zz_tm_tag" text,
                                            "objnr" text NOT NULL,
                                            "status_object" text NOT NULL,
                                            "status" text NOT NULL,
                                            "timestamp" text NOT NULL,
                                            "beber" text,
                                            "swerk" text,
                                            "iwerk" text,
                                            "tplma" text
);

CREATE TABLE "iba_rt_topics" (
                                 "iba_server_id" bigint,
                                 "topic" text UNIQUE,
                                 "comment" text
);

CREATE TABLE "iba_gates" (
                             "id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                             "name" text UNIQUE NOT NULL,
                             "comment" text
);

CREATE TABLE "iba_gates_sources" (
                                     "source_id" bigint NOT NULL,
                                     "gate_id" bigint NOT NULL,
                                     "mnt" text NOT NULL,
                                     "comment" text
);

ALTER TABLE "signals" ADD FOREIGN KEY ("value_type_id") REFERENCES "value_types" ("id");

ALTER TABLE "signals" ADD FOREIGN KEY ("class_id") REFERENCES "signal_classes" ("id");

ALTER TABLE "signals" ADD FOREIGN KEY ("sources_id") REFERENCES "sources" ("id");

ALTER TABLE "signals" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "sources" ADD FOREIGN KEY ("parent_id") REFERENCES "sources" ("id");

ALTER TABLE "sources" ADD FOREIGN KEY ("source_classes_id") REFERENCES "source_classes" ("id");

ALTER TABLE "nodes" ADD FOREIGN KEY ("parent_id") REFERENCES "nodes" ("id");

ALTER TABLE "nodes" ADD FOREIGN KEY ("directory_id") REFERENCES "directories" ("id");

ALTER TABLE "nodes" ADD FOREIGN KEY ("signal_id") REFERENCES "signals" ("id");

ALTER TABLE "iba_gates_sources" ADD FOREIGN KEY ("source_id") REFERENCES "sources" ("id");

ALTER TABLE "iba_gates_sources" ADD FOREIGN KEY ("gate_id") REFERENCES "iba_gates" ("id");

CREATE UNIQUE INDEX ON "iba_gates_sources" ("gate_id", "mnt");

COMMENT ON TABLE "signals" IS 'реестр сигналов';

COMMENT ON COLUMN "status"."name" IS 'Имя статуса';

COMMENT ON TABLE "sources" IS 'источники данных';

COMMENT ON COLUMN "sources"."name" IS 'Имя источника';

COMMENT ON TABLE "source_classes" IS 'Классы источников';

COMMENT ON COLUMN "source_classes"."name" IS 'Example: iba, iba_rt, zifra_adp';

COMMENT ON COLUMN "signal_classes"."name" IS 'analog, discrete, virtual';

COMMENT ON COLUMN "value_types"."name" IS 'Тип данных сигнала: boolean,integer,real,double,undefined';

COMMENT ON TABLE "nodes" IS 'логическая топология';

COMMENT ON COLUMN "nodes"."full_name" IS 'root/sub/leaf/node, заполнять при редактировании узла автоматически';

COMMENT ON TABLE "directories" IS 'справочниках логической топологии';

COMMENT ON TABLE "sap_functional_locations" IS 'Справочник ТехМест';

COMMENT ON COLUMN "sap_functional_locations"."tplnr" IS 'Номер ТМ';

COMMENT ON COLUMN "sap_functional_locations"."zex" IS 'Цех';

COMMENT ON COLUMN "sap_functional_locations"."zex_ktext" IS 'Наименование цеха';

COMMENT ON COLUMN "sap_functional_locations"."pltxt" IS 'Наименование ТМ';

COMMENT ON COLUMN "sap_functional_locations"."fltyp" IS 'Тип';

COMMENT ON COLUMN "sap_functional_locations"."eqart" IS 'Тип объекта';

COMMENT ON COLUMN "sap_functional_locations"."eartx" IS 'Наименование типа объекта';

COMMENT ON COLUMN "sap_functional_locations"."serge" IS 'Серийный номер изготовителя';

COMMENT ON COLUMN "sap_functional_locations"."invnr" IS 'Инвентарный номер';

COMMENT ON COLUMN "sap_functional_locations"."kostl" IS 'МВЗ';

COMMENT ON COLUMN "sap_functional_locations"."stort" IS 'Местоположение';

COMMENT ON COLUMN "sap_functional_locations"."anlnr" IS 'Основное средство';

COMMENT ON COLUMN "sap_functional_locations"."anlun_sub" IS 'Субномеросн.средства';

COMMENT ON COLUMN "sap_functional_locations"."arbpl" IS 'Отв.раб.место';

COMMENT ON COLUMN "sap_functional_locations"."rfr_ladle" IS 'Номер ковша';

COMMENT ON COLUMN "sap_functional_locations"."erdat" IS 'Дата cоздания';

COMMENT ON COLUMN "sap_functional_locations"."aedat" IS 'Дата зменения';

COMMENT ON COLUMN "sap_functional_locations"."zz_tm_rfid" IS 'ID метки';

COMMENT ON COLUMN "sap_functional_locations"."zz_tm_tag" IS 'Обозначение метки';

COMMENT ON COLUMN "sap_functional_locations"."objnr" IS 'Номер объекта';

COMMENT ON COLUMN "sap_functional_locations"."status_object" IS 'Статус объекта';

COMMENT ON COLUMN "sap_functional_locations"."status" IS 'Статус';

COMMENT ON COLUMN "sap_functional_locations"."timestamp" IS 'Дата загрузки';

COMMENT ON COLUMN "sap_functional_locations"."beber" IS 'Произв Участок';

COMMENT ON COLUMN "sap_functional_locations"."swerk" IS 'Завод расположения';

COMMENT ON COLUMN "sap_functional_locations"."iwerk" IS 'Завод планирующий';

COMMENT ON COLUMN "sap_functional_locations"."tplma" IS 'Вышестоящее ТМ';

COMMENT ON TABLE "iba_rt_topics" IS 'Реестр топиков с RT данными';

COMMENT ON COLUMN "iba_rt_topics"."comment" IS 'заметки';

COMMENT ON TABLE "iba_gates" IS 'Реестр экспортеров Dat файлов в S3';

COMMENT ON COLUMN "iba_gates"."name" IS 'gate_id';

COMMENT ON COLUMN "iba_gates"."comment" IS 'заметки';

COMMENT ON COLUMN "iba_gates_sources"."mnt" IS 'имя папки шары';

COMMENT ON COLUMN "iba_gates_sources"."comment" IS 'заметки';
