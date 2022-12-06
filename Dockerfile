FROM golang:1.19-buster as builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/flexible ./cmd/flexible

FROM gcr.io/distroless/static-debian10:nonroot as runtime

COPY --from=builder api/bin/flexible /

CMD ["/flexible"]