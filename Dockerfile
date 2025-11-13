FROM golang:1.25-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM builder as build
COPY . ./
RUN CGO_ENABLED=0 go build -o main -ldflags "-s -w" .

FROM golang:1.25-alpine
COPY --from=build /app/main /main
CMD ["/main"]
