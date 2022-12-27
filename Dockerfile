FROM golang:alpine AS build
WORKDIR /app
COPY fibCalc.go .
RUN go build -o FibCalc fibCalc.go

FROM scratch
COPY --from=build /app/FibCalc /app/

ENTRYPOINT ["/app/FibCalc"]