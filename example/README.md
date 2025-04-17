# Sample API (Go)

Provides a sample API environment using `docker-compose`.

## Quick Start

### Prerequisites

#### Install Docker

- Install Docker for Desktop. [Download Site](https://www.docker.com/products/docker-desktop)

#### Certificate Files
Obtain the `viron.crt`, `viron.csr`, and `viron.key` files and place them in the `viron-go/example/cert` directory.

#### .env File
```
cd viron-go/example
cp .env.template .env
```
Then, fill in the project's secret information.

### Start the Sample Database

```
cd viron-go/example

# When using MySQL
docker compose -f docker-compose-store.yaml up --build mysql

# When using Mongo
docker compose -f docker-compose-store.yaml up --build mongo
```

### Start the Sample Backend Application
```
cd viron-go

# When using MySQL
task example-app-mysql

# When using Mongo
task example-app-mongo
```
