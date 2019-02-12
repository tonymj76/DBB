FROM golang:1.11-alpine

RUN mkdir /app

# RUN apk add --no-cache libstdc++ gcc git musl-dev

WORKDIR /app

COPY . .

RUN go get

# GO SERVER PORT
EXPOSE 3500

# WEBPACK PORT
EXPOSE 8082

# START COMMAND
CMD [ "sh" ]
