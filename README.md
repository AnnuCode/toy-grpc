# GRPC Unary API 

## Steps to run the project:
- to run the grpc server, run `go run main.go` in the `cmd` directory to run the `main.go` file.
- run `go run main.go` in the `client/user` directory to run the grpc client for `User` detail(`GetUser`).
- run `go run main.go` in the `client/users` directory to run the grpc client for a list of `User` details(`GetUsers`).
- to run the tests, run `go test` in the `cmd` directory to run the `main_test.go` file.

## Design Pattern
- I've used Repository pattern in this project to create a layer between database specific logic and business logic in the app.
- I've used an in-memory repository implementation.
- Using this pattern, we can treat repositories as adapters which can be changed in the future without having to touch the business logic(handlers). For example, to use MySQL in future, we can create a MySQL
    repository and use it instead of the in-memory repository. 
  

### GetUser response screenshot

![GetUser](https://github.com/AnnuCode/toy-grpc/assets/97174641/1e576188-0031-421b-8119-37401588befc)


## GetUsers response screenshot

![GetUsers](https://github.com/AnnuCode/toy-grpc/assets/97174641/d0ff6bdd-ea28-433b-bc89-aeea4c217237)
