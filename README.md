Lets GO！DDD
========================================
大大小小的项目做了不少，但是几乎每个项目都无法经历时间的冲击，甚至在开发过程中，因为需求的频繁变动，项目就开始变得臃肿、难以维护。
进行反思和学习后，从DDD中找到了解决这个问题的答案！所以想自己维护一套go语言框架——Go！DDD！

DCI 项目架构
--------------
下图介绍项目的架构
```
  .
  ├── api
  │   ├── http
  │   ├── mq
  │   ├── timers
  └── biz
  └── domain
  │   ├── entity
  │   ├── repository
  │   ├── service
  └── infrastructure
  │   ├── commons
  │   │   ├── adapter
  │   │   ├── driver
  │   │   │   ├── httpserver
  │   │   │   ├── mqclient
  │   │   ├── persistent
  │   │   ├── reference
  │   │   ├── serviceimpl
  │   ├── configs
  │   ├── tools
```

项目结构解释
--------------
- api 对外暴露的接口层
  - http（接口、路由的配置）
  - mq （消息队列的消费者监听）
- biz 主要流程层，放置整个项目要做的主要任务流程，调用domain层中的接口
- domain 领域层，在这个层级进行领域建模
  - entity
  - repository
  - service
- infrastructure 基础设施层，为其他层提供支持，不会调用其他层的代码