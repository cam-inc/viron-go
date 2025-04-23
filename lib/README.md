# Viron-Go Library

This library provides basic functionalities for creating backend applications using [Viron](https://github.com/cam-inc/viron).

## Features

- **Authentication**: Standard support for email-password authentication, Google authentication, and OpenID Connect authentication.
- **Admin Account Management**: Provides functionality for managing Viron accounts used to log in to Viron.
- **Admin Role Management**: Provides RBAC-based role management for Viron accounts.

## Directory Structure

- `api/`: Directory for storing OAS (OpenAPI Specification) definitions used by Viron.
- `configs/`: Directory for managing type definitions related to configurations.
- `constant/`: Directory for managing constants used in the application.
- `domains/`: Directory for implementing common business logic.
- `errors/`: Directory for managing type definitions related to errors.
- `helpers/`: Directory for providing utility functions and auxiliary features.
- `linter/`: Directory for storing settings and scripts for static code analysis and style checks.
- `logging/`: Directory for managing settings and logic related to application log output.
- `repositories/`: Directory for implementing the repository layer that abstracts interactions with databases and external APIs.
- `routes/`: Directory for defining application routing and managing the processing of each endpoint.

## Usage

To use the libraries in this directory, import them into your code as follows:

```go
import "github.com/cam-inc/viron-go/lib/<library-name>"
```

Replace `<library-name>` with the specific library name you want to use.

## Contribution

For details, please refer to the [CONTRIBUTING](../CONTRIBUTING.md) file.

## License

This project is licensed under the MIT License. For details, please refer to the [LICENSE](../LICENSE) file.
