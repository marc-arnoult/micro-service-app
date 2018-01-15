FROM golang:1.9-alpine

WORKDIR /go/src/app
#RUN apk --no-cache add curl git && \
#   curl https://glide.sh/get | sh


COPY . .
RUN go build -o todo-service

CMD /go/src/app/todo-service