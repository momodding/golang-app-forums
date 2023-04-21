Forum APP

## Migration
Tutorial: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

ENV

#

export POSTGRESQL_FORUM_APP_URL='postgres://username:password@winhost:5432/forum-app?sslmode=disable'

#

Create migration

#

migrate create -ext sql -dir db/migrations -seq create_category_table

#

Migrate migration

#

migrate -database ${POSTGRESQL_FORUM_APP_URL} -path db/migrations up

#