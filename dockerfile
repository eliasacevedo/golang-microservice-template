FROM golang:1.19.4-alpine as build

WORKDIR /app
RUN apk add make

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .
RUN make build

FROM alpine:3.14 as runner
WORKDIR /app
ARG PORT
ARG APPNAME
ENV APPNAME=$APPNAME

COPY --from=build /app/bin .
RUN chmod +x $APPNAME
EXPOSE $PORT

CMD ./$APPNAME