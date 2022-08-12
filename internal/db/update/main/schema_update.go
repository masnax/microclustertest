package main

// DO NOT EDIT BY HAND
//
// This code was generated by the schema.DotGo function. If you need to
// modify the database schema, please add a new schema update to update.go
// and the run 'make update-schema'.
const freshSchema = `
CREATE TABLE internal_cluster_members (
  id                   INTEGER   PRIMARY  KEY    AUTOINCREMENT  NOT  NULL,
  name                 TEXT      NOT      NULL,
  address              TEXT      NOT      NULL,
  certificate          TEXT      NOT      NULL,
  schema               INTEGER   NOT      NULL,
  heartbeat            DATETIME  NOT      NULL,
  role                 TEXT      NOT      NULL,
  UNIQUE(name),
  UNIQUE(certificate)
);
CREATE TABLE internal_token_records (
  id           INTEGER         PRIMARY  KEY    AUTOINCREMENT  NOT  NULL,
  name         TEXT            NOT      NULL,
  token        TEXT            NOT      NULL,
  UNIQUE       (name),
  UNIQUE       (token)
);

INSERT INTO schemas (version, updated_at) VALUES (1, strftime("%s"))
`
