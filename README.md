
# KingNet

一个基于事件驱动使用go语言实现的TCP网络库，用来处理长连接推送、游戏服务器等场景。其中项目的主体框架参考和模仿了zinx的实现方式，结合自己对于go的理解和框架的掌握做了一些修改。项目的目录结构和设计思想遵循Golang的规范。

## Features

- 基于go协程和事件驱动来处理TCP请求，通过连接池来处理消息，通过读协程和写协程分离来提高处理能力。
- 使用viper和全局对象来管理配置信息；
- 集成静态代码检查功能；
- 基于zap日志包封装项目日志库，提供多种级别的日志输出；
- 提供docker容器化部署方式;
- 遵循golang项目结构和编写规范。

## Installation

使用make编译运行该项目:

```bash
  git clone https://github.com/lkzhang98/KingNet.git
  cd KingNet
  make
```

编译生成容器镜像

```bash
make image IMAGES="kingnet"
```

## Documentation

+ [用户手册](./docs/guide/zh-CN/README.md)
+ [开发手册](./docs/devel/zh-CN/README.md)


## Feedback

如果您在使用中有任何问题，请反馈我：`lkzhang98@163.com`。


## Authors

- [@lkzhang98](https://www.github.com/lkzhang98)


## Reference and Thanks

+ 项目主要参考了[@刘丹冰](https://github.com/aceld/)老师的开源项目[zinx](https://github.com/aceld/zinx), 感谢大佬提供的优质的开源项目、文档和课程；
+ 同时还要感谢[@孔令飞](https://github.com/marmotedu/)老师的开源项目[miniblog](https://github.com/marmotedu/miniblog), 从中学到了很多企业级开发的工具和规范，对于开发一个项目的帮助非常大。

## License

本项目遵循 [MIT许可证](https://opensource.org/licenses/MIT)。