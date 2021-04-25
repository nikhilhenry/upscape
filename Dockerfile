
#client build

FROM node:alpine AS client

WORKDIR /usr/src/app/client
COPY ./client/package*.json ./
ENV NODE_ENV=production
RUN npm install --also=dev
COPY ./client .
ENV NODE_ENV=production
RUN npm run build
RUN ls

#server build
FROM golang:alpine AS server
# create app directtory

WORKDIR /usr/src/app/server
COPY . .
RUN go build -o main



FROM alpine:latest
WORKDIR /root/
RUN mkdir /public
COPY --from=server /usr/src/app/server .
COPY --from=client /usr/src/app/client/dist ./public
RUN cd public && ls

#run server executable
CMD ["./main"]