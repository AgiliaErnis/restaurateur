FROM golang:1.16 
ADD . /app
WORKDIR /app
RUN go mod download
RUN cd cmd/backend && go build

CMD ./backend

EXPOSE 8080
