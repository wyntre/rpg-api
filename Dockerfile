# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:v0.16.5 as builder

RUN mkdir -p $GOPATH/src/github.com/wyntre
WORKDIR $GOPATH/src/github.com/wyntre

ADD . .
ENV GO111MODULE=on
RUN go get ./...
RUN buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

RUN apk add openssl
RUN mkdir -p /keys
ADD keys/gen_keys.sh /keys
WORKDIR /keys
RUN /keys/gen_keys.sh
WORKDIR /

# Uncomment to run the migrations before running the binary:
CMD /bin/app
# CMD exec /bin/app
