FROM node:12.1-alpine
ADD . .
ARG HTTP_PROXY
ENV HTTP_PROXY $HTTP_PROXY
RUN sed -i "s/backend.*:/$HTTP_PROXY:/" package.json
RUN npm install
RUN npm run build

CMD npm start

EXPOSE 3000
