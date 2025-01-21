## groot
[简体中文](./README.md) | English

A distributed task system based on the open source project [asynq](https://github.com/hibiken/asynq)

### New Features
[TODO](./TODO.md)
* Command-line tool: gctl
* MySQL task storage implementation
* Task notifications: Basic implementation of Feishu robot notifications; custom notifications can be implemented
* API authentication: JWT authentication

### [Architecture Diagram](groot.drawio)

### Deployment Dependencies
* Redis (>5.0)
* MySQL (>5.7.8)

### Task Implementation
* Specific task implementations are stored in the internal/apps/tasks directory.

### Dependencies
* OpenAPI: [oapi-codegen](https://github.com/deepmap/oapi-codegen)
  ```
  // Generate OpenAPI
  cd docs/v1
  sh codegen.sh

  // Directory where the generated code is located
  cd gen/
  ```
* Parameter validation: [libopenapi-validator](https://github.com/pb33f/libopenapi-validator)
