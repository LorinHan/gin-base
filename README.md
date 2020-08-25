# gin的基本框架
- 结合以下：
    - gorm
    - zap 高性能日志
    - lumberjack 日志归档
    - yml配置方式
    - jwt-go （用户、角色等权限相关的校验可在中间件中自己实现）
    - gin-swagger（需要先安装`swaggo`）
    
- 拉取下来后，重命名，比如将本项目重命名为test
    - 先把项目根目录 gin-base改为 test
    - 然后把go.mod中的module 改为 test
    - goLand中 `ctrl + shift + r`或`command + shift + r`，全局替换 `"gin-base/` 为 `"test/`
    
- 2020-08-12
    - `logrus` 日志换为 `zap`
    - 加入`lumberjack` 日志归档
    
- 2020-08-25
    - 加入`swagger`文档
    - 使用前需要先安装`swaggo`
    ```
    go get -u github.com/swaggo/swag/cmd/swag
    swag init -h 可以查看相关命令
    ```
