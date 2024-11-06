CREATE TABLE "classes" (
  "id" char(36) PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "subject" varchar(150),
  "grade" varchar(100),
  "code" char(8) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "units" (
  "id" char(36) PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "class_id" char(36) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("class_id") REFERENCES "classes" ("id")
);

CREATE TABLE "roles" (
  "id" char(36) PRIMARY KEY,
  "key" varchar(50) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "members" (
  "id" char(36) PRIMARY KEY,
  "role_id" char(36) NOT NULL,
  "user_id" char(36) NOT NULL,
  "class_id" char(36) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id"),
  FOREIGN KEY ("class_id") REFERENCES "classes" ("id")
);