
#client build

FROM node:alpine AS client

WORKDIR /usr/src/app/client
COPY ./client .
ENV NODE_ENV=production
RUN npm install --also=dev
RUN npm run build

#server build
FROM golang:alpine AS server
# create app directtory

WORKDIR /usr/src/app/server
COPY . .
RUN go build -o main



FROM alpine:latest
WORKDIR /root/
COPY --from=server /usr/src/app/server .
COPY --from=client /usr/src/app/public /public

#run server executable
CMD ["./main"]