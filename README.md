#ginblog
gin + vue 全栈项目

##包管理
- config --> 配置参数
- model --> 管理数据库的参数读写等操作
- api --> 前后端分离，接口入口，下面分包，管理不同版本
- middleware --> 中间件， 跨域，登录验证
- routes --> 路由接口
- utils --> 工具包，公共功能全局使用
- upload --> 上传下载
- web --> 托管前端页面


# errmsg
1. code 1000... 用户模块的错误
2. code 2000... 分类模块的错误
3. code 3000... 文章模块的错误


# 项目大概流程
1. 先写好项目结构
2. 写好项目需要的config信息
3. 连接，构建数据库
4. 错误处理，api分析结构
5. 数据库处理数据
6. 接口实现
7. 实现登录接口(使用jwt实现)
8. 实现日志功能(logrus)