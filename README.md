# fly.io

Setting up a free postgres server is a breeze - only 3 commands:

```console
➜ web git:(main) flyctl auth signup # opens browser
➜ web git:(main) flyctl auth login # opens browser
➜ web git:(main) flyctl postgres create --machines
? Choose an app name (leave blank to generate one): awair
automatically selected personal organization
? Select regions: Secaucus, NJ (US) (ewr)
? Select configuration: Development - Single node, 1x shared CPU, 256MB RAM, 1GB disk
Creating postgres cluster awair in personal organization
Creating app...
Setting secrets...
Provisioning 1 of 1 machines with image flyio/postgres:14.4
Waiting for machine to start...Machine is created
Username: postgres
Password:
Hostname: awair.internal
Proxy port: 5432
Postgres port: 5433
Save your credentials in a secure place -- you won't be able to see them again!

Connect to postgres
Any app within the personal organization can connect to this Postgres using the following credentials:
For example: postgres://postgres:***@awair.internal:5432
```

## Makefile

To run a local proxy on port 15432:

```console
make start
```

Some other useful targets:

```console
make PW=*** regendb
make PW=*** schema.sql
make PW=*** dump.sql
```
