# InMemory Data Store API

This project mimics an **InMemory Data Store** (similar to Redis) using Go.  
It provides a RESTful API to store, update, retrieve, and delete key-value data — supporting both string and list data types.

---

## Tech Stack

- **Language:** Go  
- **Framework:** Gin  
- **Storage:** Go Map (in-memory store)   // NOTE: I have not implemenetd the actual DB to store the data
- **HTTP API Spec:** OpenAPI (see `openapi.yaml`) 
- **API Client:** refer the client_api_doc
- **Testing:**  
  - `testing` — Go’s standard testing package  
  - `mockery` — for mocking interfaces  
  - `gock` — for HTTP mocking in client unit tests  

---

## Testing

- **Unit Tests:** for internal API handler and client methods using mocks and stubs.  
- **Integration Tests:** for REST endpoints via HTTP requests (partial coverage to demonstrate examples).  
- Run tests with:  
  ```bash
  go test ./...


## How to run this project
- run go run cmd/main.go : this will run it on localhost at port 8081
- run the go run client_main.go : to run the client function. I have added the example of those client function in it. 
