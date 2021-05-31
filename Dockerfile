FROM golang:1.16 AS gobuilder
ADD ./backend /backend
WORKDIR /backend
RUN go mod download
RUN cd cmd/backend && CGO_ENABLED=0 GOOS=linux go build

FROM node:12.1-alpine
ADD ./frontend /app
WORKDIR /app
COPY --from=gobuilder /backend/cmd/backend ./backend
ARG REACT_APP_PROXY
ENV REACT_APP_PROXY $REACT_APP_PROXY
RUN echo 'REACT_APP_PROXY=$REACT_APP_PROXY' >> .env
RUN npm install
RUN npm run build
RUN npm install -g serve

RUN echo '0 0 * * 0 /app/backend/backend --update-menus >> /app/backend/backend_cron.log 2>&1' >> /etc/crontabs/root

CMD sh -c './backend/backend >> backend.log 2>&1 &' \
 && serve -s build -l 3000

EXPOSE 3000
EXPOSE 8080
