FROM golang:1.16 AS gobuilder
ADD ./backend /backend
WORKDIR /backend
RUN go mod download
RUN cd cmd/backend && CGO_ENABLED=0 GOOS=linux go build

FROM node:12.1-alpine
ADD ./frontend /app
WORKDIR /app
COPY --from=gobuilder /backend/cmd/backend ./backend
RUN npm install
RUN npm run build
RUN npm install -g serve

CMD sh -c './backend/backend &' \
 && serve -s build -l 3000

EXPOSE 3000
EXPOSE 8080
