# Built With

Go 1.22  
PostgreSQL 14

# Database setup

create and setup DB

```
$ psql -U <admin_user> -d postgres
=> CREATE DATABASE greenlight;
=> \c greenlight
=> CREATE ROLE greenlight WITH LOGIN PASSWORD <some_pass>;
=> CREATE EXTENSION IF NOT EXISTS citext;
```

add variable for dsn

```
# add user password from prev step
$ echo "GREENLIGHT_DB_DSN='postgres://greenlight:<some_pass>@localhost/greenlight'" >> $HOME/.profile
```

connect

```
# using dsn
$ psql $GREENLIGHT_DB_DSN
```

### Migrations

Using [migrate](https://github.com/golang-migrate/migrate) tool to manage database.

```
brew install golang-migrate
```

run

```
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```
