FROM golang:latest as build
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o desafio-01 .

FROM scratch
COPY --from=build /app/desafio-01 .
CMD ["./desafio-01"]