tidy:
	go mod tidy
.PHONY: run

run:
	go run greet.go -f etc/greet-api.yaml
greet:
	goctl api go -api greet.api -dir .

.PHONY: greet
user_model:
	goctl model mysql ddl -src user.sql -dir .

.PHONY: user_model
#/internal/adapter/database/user.sql  goctl model mysql ddl -src user.sql -dir .

MODEL_SRC=$(wildcard ./module/**/*.sql)
 
.PHONY: gen-model
gen-model:
	goctl model mysql ddl -src $(MODEL_SRC) -dir .
