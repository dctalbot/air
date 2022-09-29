.PHONY: regendb schema.sql dump.sql start

CONN_STR="postgres://postgres:${PW}@localhost:15432"

regendb:
	psql $(CONN_STR) -f ./schema.sql

schema.sql:
	pg_dump --if-exists  --clean --schema-only $(CONN_STR) -f ./$@

dump.sql:
	pg_dump --if-exists  --clean $(CONN_STR) -f ./$@

start: 
	flyctl proxy 15432:5432 -a awair
