.PHONY: RUN_FE RUN_BE SWAGGER_BE

RUN_FE:
	cd fe && npm run dev

RUN_BE:
	cd be && swag init -g ./cmd/api/main.go -o ./docs/swagger && go run ./cmd/api

SWAGGER_BE:
	cd be && swag init -g ./cmd/api/main.go -o ./docs/swagger