# Go Dating Service

RESTful API for a dating service.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Dependency](#dependency)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)

## About

Go Dating Service is a RESTful API prototype of a dating application. It provides authentication, user discovery, preference management, and premium package purchases.

## Features

List of key features or functionalities of the project:

- Authentication (Sign In / Sign Up)
- List potential users for date
- Set user preference (Pass / Like)
- Show my profile
- List of packages
- Purchase (Order) package for more swap freedom

## Dependency

Since this application uses Docker, you can run it simply with Docker. Here is the list of dependencies:

- Docker* (Required in the machine host)
- GoLang 1.21.3 (Optional in the machine host)
- MySQL 8 (Optional in the machine host)
- Protobuf / Buf (Optional in the machine host)
- CMAKE* or MAKE* (Required in the machine host)

## Installation

First of all setup the environment, copy file `.env.example` to `.env`. By default the values in the example are ready to use.

Next, start docker container for the service:

```bash
$ make start
```

If Make is not installed, run:

```bash
$ docker-compose up -d
```

After the service is started, migrate the default database structure and run the seeder:

```bash
$ make migrate
```

If Make is not installed, run:

```bash
$ docker-compose exec -it httpd go run database/migration.go
```

## Usage

After installation and setup is successful, this service will open 2 ports on the machine:

- 8080 (gRPC)
- 8081 (gRPC Gateway)

Since the API is RESTful, the common usage must be in `http://localhost:8081` as the entrypoint.

## Documentation

There is currently no automatic RESTful API documentation served in the browser, but generated Open API Swagger documentation is available inside `docs/schema.swagger.json`, which can be copied and pasted to `https://editor.swagger.io` or imported through Postman. There is also `docs/schema.har`, which can be imported as a documentation collection to Postman or Insomnia.
