run:
	docker-compose up -d

rebuild:
	docker-compose up -d --build

# make make-migration ARGS="name"
make-migration:
	migrate create -ext sql -dir migrations -seq $(ARGS)

# migrate tool
migrate-master:
	#migrate -path ./migrations -database "postgresql://social-network-local-admin:eephayl3eaph8Xo@localhost:6432/social-network-local?sslmode=disable" -verbose up

migrate-slave:
	#migrate -path ./migrations -database "postgresql://social-network-local-admin:eephayl3eaph8Xo@localhost:7432/social-network-local?sslmode=disable" -verbose up

migrate:
	make migrate-master & make migrate-slave

proto:
	protoc -I proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/internalapi/dialog_service.proto --go-grpc_out=require_unimplemented_servers=false:.
