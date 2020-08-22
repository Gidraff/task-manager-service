DROP TABLE IF EXISTS accounts;


CREATE TABLE accounts (
  user_id serial primary key,
  username varchar(50) not null,
  email varchar(50) unique not null,
  active boolean not null,
  password varchar(255) not null,
  created_on timestamp not null,
  last_login timestamp
);

GRANT USAGE, SELECT ON SEQUENCE accounts_user_id_seq TO taskmuser;
GRANT ALL PRIVILEGES ON TABLE accounts TO taskmuser;
