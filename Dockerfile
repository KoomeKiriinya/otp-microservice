## We'll choose the incredibly lightweight
## Go alpine image to work with
FROM golang:1.15-buster AS builder

## We create an /app directory in which
## we'll put all of our project code


## build ARGS

 


WORKDIR  /go/src/otp
COPY go.* ./
RUN go mod download


COPY . ./


## We want to build our application's binary executable

RUN  CGO_ENABLED=0 GOOS=linux go build -o main .

## the lightweight scratch image we'll
## run our application within
FROM alpine:latest AS production
## We have to copy the output from our
## builder stage to our production stage

ARG REDIS_DSN
ARG REDIS_PASSWORD

ENV REDIS_DSN=${REDIS_DSN}
ENV REDIS_PASSWORD=${REDIS_PASSWORD}



COPY --from=builder /go/src/otp .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD /wait && ./main
