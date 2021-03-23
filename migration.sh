migrate -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DBNAME?sslmode=disable" \
 -path db/migrations up