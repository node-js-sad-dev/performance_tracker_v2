.PHONY: RUN_FE RUN_BE

RUN_FE:
	cd be && npm run dev

RUN_BE:
	cd be && swag init -g cmd/api/main.go --parseDependency --parseInternal && go run ./cmd/api