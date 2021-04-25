
#server build
FROM golang:alpine
# create app directtory

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main 


#client build

FROM node:10

WORKDIR /app/client
ENV NODE_ENV=production
RUN npm install --also=dev
RUN npm run build

#run server executable
CMD ["/app/main"]