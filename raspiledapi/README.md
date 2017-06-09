# REST API server in golang

## Links
- https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/


## Example 1
- [Example1](example1/)

### GET
~~~
curl http://localhost:12345/people
~~~

### POST
~~~
curl -d '{"id":"3","firstname":"Thomas","lastname":"Müller","address":{"city":"Bern","state":"BE"}}' -H "Content-Type: application/json" -X POST http://localhost:12345/people/3
~~~

### DELETE
~~~
curl -X DELETE http://localhost:12345/people/3
~~~
