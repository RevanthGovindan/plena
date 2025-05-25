# Access Key Management

This repository provides tools and utilities for managing access keys securely and efficiently.

## Features

### Admin
- **Key Generation**: Generate secure access keys.
- **Expiry and rate limit updatation**: Rotate keys periodically to enhance security.
- **delete key**: Generate secure access keys.
- **get all keys**: Generate secure access keys.

### User

- **disable access**: Generate secure access keys.
- **get specific Key data**: Generate secure access keys.

## Getting Started

1. build access-key-management repo using docker
    ```
    docker build -t key-management:latest .
    ```
2. build token info repo using docker
    ```
    docker build -t token-info:latest .
    ```
3. For event streaming **redis** pubsub has been used. Attached docker-compose.yaml for starting both microservice and redis as docker containers.

4. There is no external database used, only inmemory cache and its an interface so it can be extended database without much code changes and just by implementing methods implicitly.
Refer internal/database/mysql/MySql.go.

5. logs will be printing to a new file, log rotation and log type can be added.

6. Unit test added for few files, databases and services were only added. Mocks were used for integration tests.