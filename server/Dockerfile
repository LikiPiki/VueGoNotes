FROM golang:latest
WORKDIR /build/
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=0 /build/app .
EXPOSE 3000
EXPOSE 5432
CMD ["./app"] 