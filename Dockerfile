#ALL ARGS
#Server
ARG MONGODB_URI
ARG DATABASE
ARG PASSWORD_SECRET




#client build

FROM node:lts-alpine AS client

WORKDIR /usr/src/app/client
COPY ./client-next/package*.json ./
ENV NODE_ENV=production
RUN npm install --also=dev
RUN ls
COPY ./client-next .
ENV NODE_ENV=production
RUN npm run build

#server build
FROM golang:latest AS server

ARG MONGODB_URI
ARG DATABASE
ARG PASSWORD_SECRET
#Set arg to env
ENV MONGODB_URI=${MONGODB_URI}
ENV DATABASE=${DATABASE}
ENV PASSWORD_SECRET=${PASSWORD_SECRET}
ENV GIN_MODE=release


# create app directtory

WORKDIR /usr/src/app/server
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux go build -o  main -tags netgo -a -v



FROM alpine:latest
ARG MONGODB_URI
ARG DATABASE
ARG PASSWORD_SECRET

WORKDIR /root/
RUN mkdir /public
COPY --from=server /usr/src/app/server .
COPY --from=server /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip
COPY --from=client /usr/src/app/client/dist ./public

ENV MONGODB_URI=${MONGODB_URI}
ENV DATABASE=${DATABASE}
ENV PASSWORD_SECRET=${PASSWORD_SECRET}
ENV GIN_MODE=release

RUN echo ${MONGODB_URI}

#run server executable
CMD ["./main"]