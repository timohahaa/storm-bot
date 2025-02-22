########## BUILD STAGE ##########
FROM golang:1.24 AS builder

WORKDIR /src

# deps
COPY go.mod .
RUN go mod download

# source code
COPY . .

# compile
RUN make build

########## RUN STAGE ##########
FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/.build ./
COPY --from=builder /src/migrations/ ./migrations/

ENTRYPOINT [ "./bot" ]
