FROM golang:1.12-alpine AS build
WORKDIR /usr/src
RUN apk add --no-cache git gcc build-base
ARG GOVERSION=1.12
ARG COMMITHASH
ARG BUILDTIME
ARG VERSION
COPY . ./
RUN go build -o app -ldflags "-X main.VERSION=${VERSION} -X main.GOVERSION=${GOVERSION} -X main.BUILDTIME=${BUILDTIME} -X main.COMMITHASH=${COMMITHASH}"

FROM alpine
EXPOSE 3000
WORKDIR /usr/src
COPY --from=build /usr/src/app /usr/src/
CMD ["/usr/src/app"]
