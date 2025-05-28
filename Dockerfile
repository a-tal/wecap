FROM golang:alpine AS build-env
COPY . /src
WORKDIR /src
RUN go build -o wecap

FROM alpine
WORKDIR /app
COPY --from=build-env /src/wecap /app/
EXPOSE 8000
ENTRYPOINT ./wecap
