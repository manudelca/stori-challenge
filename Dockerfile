FROM golang:1.23 as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -v -o /usr/local/bin/app ./cmd/api/main.go

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /usr/local/bin/app ./app
ENTRYPOINT [ "./app" ]
