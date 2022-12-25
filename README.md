# zup-message-service

### Message service for zup application.

---
![GitHub top language](https://img.shields.io/github/languages/top/cccaaannn/zup-message-service?style=flat-square) ![](https://img.shields.io/github/repo-size/cccaaannn/zup-message-service?style=flat-square) [![GitHub license](https://img.shields.io/github/license/cccaaannn/zup-message-service?style=flat-square)](https://github.com/cccaaannn/zup-message-service/blob/master/LICENSE)

### zup is a messaging application, built by microservice architecture.
### Related repos
- [Frontend](https://github.com/cccaaannn/zup-frontend)
- [User service](https://github.com/cccaaannn/zup-user-service)
- [Message service](https://github.com/cccaaannn/zup-message-service) (This project)
- [K8s configurations](https://github.com/cccaaannn/zup-k8s)

<hr>

### Configurations
1. Add a `.env` file at the root of the project.
2. Fill it up with using example file. `doc/.env-template`

<hr>

## Running with Docker
1. Build
```shell
docker build -t cccaaannn/zup-message-service:latest .
```

2. Run
```shell
docker run -d --name zup-message-service -p 8081:8081 cccaaannn/zup-message-service:latest
```

## Running Native
1. Build
```shell
go mod download
go build -o zup-message-service
```

2. Run
```shell
go run zup-message-service
```
