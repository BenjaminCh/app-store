# App-store application

## Introduction
This application is a fake app-store search demo using [Algolia search](https://www.algolia.com/).

It's composed in two parts:
- A backend, written in Go
  * [Sources](https://github.com/BenjaminCh/app-store/tree/master/backend)
  * [Live endpoint](api.search.tests.sh)
- A frontend using [InfernoJS](https://github.com/infernojs/inferno)
  * [Sources](https://github.com/BenjaminCh/app-store/tree/master/frontend)
  * [Live endpoint](search.tests.sh)

## Table of content

1. **[Install](#install)**

    * [Clone the code](#code)
    * [Frontend (Yarn)](#frontend)
    * [Backend (Go)](#backend)

1. **[Quick Start](#quick-start)**

    * [Run Frontend](#run-frontend)
    * [Run Backend](#run-backend)

1. **[Tests](#tests)**

    * [Run the Frontend](#test-frontend)
    * [Run the Backend](#test-backend)


# Getting Started

## Install

### Code
Get the code on your machine cloning the repository.
```
git clone https://github.com/BenjaminCh/app-store.git
```

### Frontend
We will use [Yarn](https://yarnpkg.com) to run the frontend code (make sure you have it installed on your machine first).
Go in the repository (root level) and type the following commands:
```
cd frontend
yarn install
```

### Backend
We will use [Go](https://golang.org/) to run the backend code (make sure you have it installed on your machine first).
Go in the repository (root level) and type the following commands:
```
cd backend
go install
```
Now, we need to set Algolia's key and index in our configuration.
We need to specify it as environnement variable.
Example on Mac :
```
export APPSTORE_ALGOLIA_APPLICATIONID=[YOUR_APPLICATION_ID]
export APPSTORE_ALGOLIA_APIKEY=[YOUR_API_KEY]
```
There is also a way to specify those keys in the config file but for some reason (not found yet), the configuration doesn't seem to be taken into account when specified in the config file.
```
{
   "version":"0.0.1.0",
   "server":{
      "port":"5000"
   },
   "application":{
      "debug":false
   },
   "algolia":{
      "applicationID":"[YOUR_APPLICATION_ID]",
      "apiKey":"[YOUR_API_KEY]",
      "indexes":{
         "apps":"apps"
      }
   }
}
```

## Quick Start

### Run Frontend
We use [Yarn](https://yarnpkg.com) to launch locally our app.
Type the following command from the frontend folder:
```
yarn start
```

The server will start :
```
$ yarn start
yarn start v0.20.3
$ webpack-dev-server --quiet --config webpack.conf.js --host 0.0.0.0
Project is running at http://0.0.0.0:8080/
webpack output is served from http://localhost:8080/
Content not from webpack is served from ./
```

You can now go to http://localhost:8080/ and play with the app.

### Run Backend
We use [Go](https://golang.org/) to launch locally our app.
Type the following command from the backend folder:
```
go run *.go
```

The server will start.

You can now go to http://localhost:8080/ and play with the app.

#### Queries examples
##### Get an app from its identifier
```
curl -X GET "http://localhost:8080/api/1/apps/346374311"
```

Will produce :
```
{
   "name":"Alaska Airlines",
   "image":"http://a1.mzstatic.com/us/r1000/113/Purple/v4/78/29/10/78291078-1abf-e98b-668a-c4b8389c6746/mzl.vpxscteh.175x175-75.jpg",
   "link":"http://itunes.apple.com/us/app/alaska-airlines/id356143077?mt=8",
   "category":"Travel",
   "rank":80
}
```
##### Create a new app
```
curl --data '
{
    "name" : "WhatTheTVShow",
    "image" : "static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
    "link" : "http://whatthetvshow.com",
    "Category" : "Quiz",
    "rank" : 223
}
' http://localhost:8080/api/1/apps

Will produce an HTTP 200 if object was created and will returns its ID.

##### Delete an app from its identifier
```
curl -X DELETE "http://localhost:8080/api/1/apps/362950001"
```

Will produce an HTTP 200 if ok, an HTTP 404 otherwise.

## Test

### Frontend
Tests to be added on the frontend.

### Backend
We will use [Go](https://golang.org/) to test the backend code (make sure you have it installed on your machine first).
Go in the repository (root level) and type the following commands:
```
cd backend
go test ./...
```



