.PHONY: regendb schema.sql dump.sql start

CONN_STR="postgres://postgres@ls-7f61d5801d5118c272bda3330c45e92366caec82.cnyyiyeuvyw9.us-east-1.rds.amazonaws.com:5432/postgres"

regendb:
	psql $(CONN_STR) -f ./schema.sql

schema.sql:
	pg_dump --if-exists  --clean --schema-only $(CONN_STR) -f ./$@

dump.sql:
	pg_dump --if-exists  --clean $(CONN_STR) -f ./$@

start:
	psql $(CONN_STR)
