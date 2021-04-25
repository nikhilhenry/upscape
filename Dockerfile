#ALL ARGS
#Server
ARG MONGODB_URI
ARG DATABASE
ARG PASSWORD_SECRET




#client build

FROM node:alpine AS client

WORKDIR /usr/src/app/client
COPY ./client/package*.json ./
ENV NODE_ENV=production
RUN npm install --also=dev
COPY ./client .
ENV NODE_ENV=production
RUN npm run build

#server build
FROM golang:alpine AS server
#Set arg to env
ENV MONGODB_URI=${MONGODB_URI}
ENV DATABASE=${DATABASE}
ENV PASSWORD_SECRET=${PASSWORD_SECRET}

# create app directtory

WORKDIR /usr/src/app/server
COPY . .
RUN go build -o main



FROM alpine:latest
WORKDIR /root/
RUN mkdir /public
COPY --from=server /usr/src/app/server .
COPY --from=client /usr/src/app/client/dist ./public

#run server executable
CMD ["./main"]