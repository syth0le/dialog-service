run:
	docker-compose up -d

rebuild:
	docker-compose up -d --build

# make make-migration ARGS="name"
make-migration:
	migrate create -ext sql -dir migrations -seq $(ARGS)

# migrate tool
migrate-master:
	migrate -path ./migrations -database "postgresql://social-network-local-admin:eephayl3eaph8Xo@localhost:6432/social-network-local?sslmode=disable" -verbose up

migrate-slave:
	migrate -path ./migrations -database "postgresql://social-network-local-admin:eephayl3eaph8Xo@localhost:7432/social-network-local?sslmode=disable" -verbose up

migrate:
	make migrate-master & make migrate-slave