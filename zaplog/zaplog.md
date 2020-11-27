- 日志记录的原则
    debug: 没问题，就看看堆栈, 一般 dev 使用这个级别
    Info: 提示一切正常, 一般生产环境使用这个或更高级别
    Warn: 记录一下，某事又发生了
    Error: 跟遇到的用户说对不起，可能有bug
    Panic: 服务器内部严重 bug，但是会 recovery
    Fatal: 网站挂了，或者极度不正常

TODO:
format 输出 (是否有必要)
以 interface 方式 (是否有必要)
打印链路

链路：
https://github.com/xiaomeng79/go-log
https://github.com/zhufuyi/logger

基于改造：
https://github.com/hwholiday/learning_tools/blob/master/all_packaged_library/logtool/log.go
https://github.com/chalvern/sugar

zap 源码分析：
https://blog.csdn.net/xz_studying/article/details/104637513


----------------
目前建议的日志配置(基于 yml 配置文件)：
```
server:
  env: dev

log:
  level: debug
  dir: /tmp/iot-export/
  file: 
    max_size: 11
    max_age: 7
    max_backups: 10
```
```
var (
    l zaplog.Level
    e zaplog.Env
)

zaplog.NewLogger(
    zaplog.SetInitialFields("source", "Manage"), // 全局打印的字段
    zaplog.SetInitialFields("appID", "DataExport"),

    zaplog.SetDevelopment(e.Unmarshal(config.GetConfig().Server.Env)), // 环境: dev pro
    zaplog.SetLevel(l.Unmarshal(config.GetConfig().Log.Level)),        // log level
    zaplog.SetLogFileDir(config.GetConfig().Log.Dir),                  // log file directory
    zaplog.SetMaxSize(config.GetConfig().Log.File.MaxSize),            // 每个文件切分大小
    zaplog.SetMaxAge(config.GetConfig().Log.File.MaxAge),              // 文件最大生命周期
    zaplog.SetMaxBackups(config.GetConfig().Log.File.MaxBackups),      // 备份数
)
```


- 问题
  初始化字段也可以用：globalLogger.With()
  现在真正的代码行数：globalLogger.WithOptions(zap.AddCallerSkip(skip))