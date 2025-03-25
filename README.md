# Easy Gin learning from scratch
- 学习Go中Web框架的设计思路
- 学习net/http标准库
- 实现了Gin框架的核心功能

## 技术点
1. 前缀树实现动态路由的绑定和解析
2. 分组控制路由, Group Control
3. 中间件插入点, c.Next()

## 架构


## 项目结构
```text
\---frame
        context.go
        engine.go
        go.mod
        logger.go
        recover.go
        router.go
        routergroup.go
        trie.go
```

## 学习笔记
### [学习记录](https://soongao.github.io/posts/webframe/)
- 参考gee