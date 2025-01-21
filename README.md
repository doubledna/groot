## groot
简体中文 | [English](./README_ENG.md)

基于开源项目 [asynq](https://github.com/hibiken/asynq) 二次封装

### 新增功能
[TODO](./TODO.md)
* 命令行工具：[gctl](./internal/gctl/README.md)
* 实现 mysql 存储任务
* 任务通知：简单实现飞书 robot 通知，可自行实现通知
* API 认证：jwt 认证

### [架构图](groot.drawio)

### 部署依赖组件
* redis(>5.0)  
* mysql(>5.7.8)

### 任务实现
* 具体任务实现存放在 internal/apps/tasks 目录

### 依赖库
* openapi: [oapi-codegen](https://github.com/deepmap/oapi-codegen)
  ```
  // 生成 openapi
  cd docs/v1
  sh codegen.sh
  
  // 生成代码所在目录
  cd gen/
  ```
* 参数校验: [libopenapi-validator](https://github.com/pb33f/libopenapi-validator)
