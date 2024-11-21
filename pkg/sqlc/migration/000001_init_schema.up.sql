CREATE TABLE "classes" (
  "id" uuid PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "subject" varchar(150),
  "grade" varchar(100),
  "code" char(6) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "units" (
  "id" uuid PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "class_id" uuid NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("class_id") REFERENCES "classes" ("id")
);

CREATE TABLE "roles" (
  "id" uuid PRIMARY KEY,
  "key" varchar(50) NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "members" (
  "id" uuid PRIMARY KEY,
  "role_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "class_id" uuid NOT NULL,
  "deleted_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id"),
  FOREIGN KEY ("class_id") REFERENCES "classes" ("id")
);