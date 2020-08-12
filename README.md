# gin的基本框架
- 结合以下：
    - gorm
    - logrus
    - yml配置方式
    - jwt （用户、角色等权限相关的校验可在中间件中自己实现）
    
- 拉取下来后，重命名，比如将本项目重命名为test
    - 先把项目根目录 gin-base改为 test
    - 然后把go.mod中的module 改为 test
    - goLand中 `ctrl + shift + r`或`command + shift + r`，全局替换 `"gin-base/` 为 `"test/`
    
- 2020-08-12
    - `logrus` 换为 `zap`
    - 加入`lumberjack` 日志归档
