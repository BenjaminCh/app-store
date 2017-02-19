# Backend

## Queries examples
### Get an app from its identifier
```
curl -X GET "http://localhost:8888/api/1/apps/346374311"
```
### Create a new app
```
curl --data '
{
    "name" : "WhatTheTVShow",
    "image" : "static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
    "link" : "http://whatthetvshow.com",
    "Category" : "Quiz",
    "rank" : 223
}
' http://localhost:8888/api/1/apps
```
### Delete an app from its identifier
```
curl -X DELETE "http://localhost:8888/api/1/apps/352641411"
```
