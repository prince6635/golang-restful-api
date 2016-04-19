http://thenewstack.io/make-a-restful-json-api-go/
Making a RESTful JSON API in Go

* build if it's not under $GOPATH
  http://stackoverflow.com/questions/14155122/how-to-call-function-from-another-file-in-go-language
  $ go run basic_web_server.go main.go

  curl to post data:
  ```
  curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
  Now, if you go to http://localhost/todos we should see the following response:
  [
      {
          "id": 1,
          "name": "Write presentation",
          "completed": false,
          "due": "0001-01-01T00:00:00Z"
      },
      {
          "id": 2,
          "name": "Host meetup",
          "completed": false,
          "due": "0001-01-01T00:00:00Z"
      },
      {
          "id": 3,
          "name": "New Todo",
          "completed": false,
          "due": "0001-01-01T00:00:00Z"
      }
  ]
  ```
