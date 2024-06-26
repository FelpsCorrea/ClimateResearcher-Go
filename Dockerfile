FROM golang:1.22 as build 
WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/weathersystem/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
COPY cmd/weathersystem/.env .
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["./cloudrun"]
