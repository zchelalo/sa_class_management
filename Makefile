URI_DB=postgresql://postgres:example@localhost:5435/sa_class_management?sslmode=disable
MIGRATE=migrate -path pkg/sqlc/migration -database "$(URI_DB)" -verbose

migrateup:
	$(MIGRATE) up

migrateup1:
	$(MIGRATE) up 1

migratedown:
	$(MIGRATE) down

migratedown1:
	$(MIGRATE) down 1

compose:
	docker compose -f ./.dockers/docker-compose.yml up

composebuild:
	docker compose -f ./.dockers/docker-compose.yml up --build

composebuilddetached:
	docker compose -f ./.dockers/docker-compose.yml up --build -d

sqlc:
	sqlc generate

protouser:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/user/service.proto

protoclass:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/class/service.proto

protomember:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/member/service.proto

proto:
	protouser
	protoclass
	protomember

.PHONY: migrateup migrateup1 migratedown migratedown1 compose composebuild sqlc protouser protoclass protomember proto