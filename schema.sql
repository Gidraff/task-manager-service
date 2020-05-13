DROP TABLE IF EXISTS users;


CREATE TABLE users (
  id serial primary key,
  username varchar(255) not null,
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp not null
);

GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO taskmuser;
GRANT ALL PRIVILEGES ON TABLE users TO taskmuser;
