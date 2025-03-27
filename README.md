# Easy Gin learning from scratch
- 学习Go Web框架的设计思路
- 学习net/http标准库
- 实现了Gin框架的核心功能
- 新添加了一个快速注册RESTful路由的方法
  - 只需一条配置即可完成一套RESTful API注册

## 技术点
1. `前缀树`实现动态路由的绑定和解析
2. 分组控制路由, `Group Control`
3. `中间件`配置
   - 自定义中间件
   - c.Next()
4. 新增快速注册`RESTful`路由的方法
   - 用户只需实现RESTful interface即可完成RESTful api的注册
  
## 项目结构
```text
\---frame
    context.go
    engine.go
    logger.go // logger中间件
    recover.go // recover中间件
    router.go 
    routergroup.go
    trie.go // 前缀树, 动态路由匹配
```

## 学习笔记
### [学习记录](https://soongao.github.io/posts/webframe/)
- 参考gee