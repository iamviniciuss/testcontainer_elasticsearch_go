# Testcontainer Elasticsearch Go

This project is an example of how to implement clean architecture and unit and integration tests in Go. We use Elasticsearch and Testcontainer to create a real instance for tests without violating the principles of Clean Architecture.

## Project Structure

The project structure follows the principles of Clean Architecture, separating responsibilities into different layers:

- **Domain**: Contains the domain entities and interfaces.
- **Use Cases**: Contains the business logic and use cases.
- **Interface Adapters**: Contains the adapters and interfaces for communication with the outside world (e.g., controllers, gateways).
- **Frameworks & Drivers**: Contains the specific implementations of frameworks and drivers (e.g., Elasticsearch, Testcontainer).

## Configuration

### Prerequisites

- Go 1.16+
- Docker

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/iamviniciuss/testcontainer_elasticsearch_go.git
    cd testcontainer_elasticsearch_go
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Running the Tests

To run the unit and integration tests, use the command:

```sh
go test ./...