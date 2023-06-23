# XM microservice

This repository contains the `XM microservice`, a Golang application that handles companies. The `microservice` provides operations such as `Create, Patch, Delete, and Get (one)` for managing company data.

## Technical Requirements
To run the `XM microservice`, make sure you have the following installed:

1. Golang (v1.15+)
2. Docker (optional, for database setup and running as a container)


## Getting Started

Follow the steps below to set up and run the `XM microservice`:

1. Clone this repository to your local machine:

```
git clone https://github.com/SXerox007/XM.git
```

2. Change into the project directory:
```
cd xm
```

3. Build and run the application:
```
go run main.go
```

### OR
Alternatively, you can build and run the application using `Docker`:
```
docker build -t xm-microservice .
docker run -p 8080:8080 xm-microservice
```

## Usage.
Start the servers:
`gRPC`: localhost:50051
`REST`: localhost:5051
Use your preferred gRPC client to interact with the available endpoints.

## MakeFile
`To run the gPRC server`

```
make app 
```

`To run the REST server`

```
make rest
```


## Configuration

The XM microservice can be configured using the `config.yaml` file. You can adjust settings such as the database connection details and authentication settings in this file.

## Endpoints

The XM microservice exposes the following endpoints:

1. POST /companies: Create a new company.
2. PATCH /companies/{id}: Update an existing company.
3. DELETE /companies/{id}: Delete a company.
4. GET /companies/{id}: Retrieve a specific company.


Only authenticated users have access to the create, update, and delete operations. Make sure to include the necessary authentication headers when making requests.

## Additional Features

The `XM microservice `also includes the following additional features:

1. Produces an event on each mutating operation.
2. Supports authentication using JWT.
3. Supports event streaming using Kafka.
3. Uses a PostgreSQL database for storing company data.

## Testing

To run the tests for the XM microservice, use the following command:

```
go test ./...
```

## Contributing
If you'd like to contribute to the XM microservice, please follow these steps:

1. Fork this repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes.
4. Push your branch to your forked repository.
5. Submit a pull request to the main repository.


## License
OrderServiceApp is licensed under the Apache License.

## Contact
For any inquiries or feedback, please contact us at sumitthakur769@gmail.com.

