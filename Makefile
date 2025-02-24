BUILD_ENVPARMS:=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
BUILD_TIME:=$(shell date +%FT%T%z)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
HEAD_COMMIT=$(shell git rev-parse HEAD)
APP_VERSION?=$(shell git name-rev --tags --name-only ${HEAD_COMMIT})

#LDFLAGS:=-X 'github.com/timohahaa/gw.ReleaseDate=$(BUILD_TIME)'\
#		 -X 'github.com/timohahaa/gw.GitCommit=$(GIT_COMMIT)'\
#		 -X 'github.com/timohahaa/gw.GitBranch=$(GIT_BRANCH)'\
#		 -X 'github.com/timohahaa/gw.AppVersion=$(APP_VERSION)'\

build:
	@[ -d .build ] || mkdir -p .build
	@$(BUILD_ENVPARMS) go build -ldflags "-s -w $(LDFLAGS)" -o .build/bot cmd/bot/main.go
	@#@file  .build/bot
	@#@du -h .build/bot

run:
	POSTGRES_DSN=postgres://postgres:password@localhost:5432/main?sslmode=disable \
	go run cmd/bot/main.go

migrate-down:
	go run cmd/db/main.go --dsn 'postgres://postgres:password@0.0.0.0:5433/main?sslmode=disable'

compose:
	docker compose --env-file deploy/.env up --build
