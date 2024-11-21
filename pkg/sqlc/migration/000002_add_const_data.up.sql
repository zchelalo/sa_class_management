CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO roles (
  id,
  key
) VALUES 
  (uuid_generate_v4(), 'student'),
  (uuid_generate_v4(), 'teacher')
;