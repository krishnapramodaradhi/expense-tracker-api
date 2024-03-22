FROM golang:1.22.1-alpine3.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /expense-tracker ./cmd/expense-tracker/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /expense-tracker /expense-tracker

USER nonroot:nonroot

ENTRYPOINT ["/expense-tracker"]
