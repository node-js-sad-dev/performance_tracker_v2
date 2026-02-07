.PHONY: run_fe run_be migrate-new be_swag_fmt

migrate-new:
	cd be && migrate create -ext sql -dir ./db/main-db/migrations -seq $(name)

run_fe:
	cd fe && npm run dev

run_be:
	cd be && swag init -g cmd/api/main.go --parseDependency --parseInternal && go run ./cmd/api

be_swag_fmt:
	cd be && swag fmt