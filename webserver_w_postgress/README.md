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

create sqlc.yaml, query.sql, schema.sql


go get github.com/lib/pq

todo
```