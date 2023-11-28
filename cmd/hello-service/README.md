# Hello-service

Write a service with 3 endpoints using [go-chi/chi](https://github.com/go-chi/chi)

Endpoints:

1. GET `/hello/{Name}` - returns a json `{"Hello": {Name}}`
2. GET `/hello/count` - returns the number of 1st endpoint calls
3. PUT `/hello/{delete_count}` - decreases the number of calls by the value of `delete_count`
