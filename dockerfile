FROM golang:1.11-alpine

RUN mkdir /app

# RUN apk add --no-cache libstdc++ gcc git musl-dev

WORKDIR /app

COPY . .
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
  echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
  apt-get update -y && apt-get install google-cloud-sdk -y
  
RUN go get

# GO SERVER PORT
EXPOSE 3500

# WEBPACK PORT
EXPOSE 8082

# START COMMAND
CMD [ "sh" ]
