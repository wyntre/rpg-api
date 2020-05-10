# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.

## Environment Variables

* `JWT_PUBLIC_KEY` is used for the json web token received to authenticate to the API.  It is a file location of an OpenSSL public signing key.
* `JWT_PRIVATE_KEY` is used for the json web token received to authenticate to the API.  It is the file location on an OpenSSL private signing key.
* `JWT_TOKEN_EXPIRATION` is the time limit a JWT is valid.  It is in the form of time-metric.  Ex: 2h for 2 hours, 55m for 55 minutes, etc.

## JWT Signing Key Setup

```
$ cd keys/
$ ./gen_keys.sh
```

## Local Testing

Requirements:

* Postgres database
* gcc
* openssl

Starting the application:

* buffalo dev

## Docker

Requirements:

* docker

## Docker Compose

Requirements:

* docker
* docker-compose

Starting the application:

* docker-compose up
