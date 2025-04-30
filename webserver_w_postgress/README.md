```
docker exec -it db bash

psql -U myuser

\l          #list databases
\c myuser   #connect to the database
\dt         #list tables
```

webservice setup

install sqlc
```
brew install sqlc

create sqlc.yaml, query.sql, schema.sql (moved under sqlc folder)

sqlc generate


go get github.com/lib/pq

docker exec -it db bash
psql -U myuser
\l
create database sqlctest;           #create test database manually
\c sqlctest
create table from scema.sql

#insert dummy element
sqlctest=# INSERT INTO authors (
  name, bio
) VALUES (
  'dummyname', 'dummybio'
);

SELECT * from authors;
```



```
brew install golang-migrate

# create migration files
migrate create -ext sql -dir webserver/database/migration/ -seq init_mg
```