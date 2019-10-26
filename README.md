# logdel

过期日志删除工具

## 支持文件名

支持文件名包含以下格式的日志文件

* `YYYY-MM-DD`，如 `info-2019-10-11.log`
* `YYYY.MM.DD`，如 `info.2019.10.11.log`
* `YYYY_MM_DD`，如 `info_2019_10_11.log`
* `YYYYMMDD`，如 `info_20191011.log`

## 使用方法

1. 获取 `logdel`

    `go get -u go.guoyk.net/logdel`

    该操作会获取并编译 `logdel` 到 `$GOPATH/bin/logdel`

    **也可以从 [此处](https://github.com/go-guoyk/logdel/releases) 获取预编译好的二进制文件**

2. 复制 `logdel` 到 `/usr/bin/logdel`

3. 编辑配置文件

    创建文件夹 `/etc/logdel.d`，并写入配置文件，**文件名可以任意指定**

    格式如下

    ```
    # "#" 为注释起始标志
    # 格式为 "规则:保存天数"
    /home/logs/info*.log: 3
    /home/logs/error*.log: 5
    ```

4. 配置 `cron`

   建议使用 `cron` 作为启动方式，执行 `crontab -e` 并填写如下内容

   ```crontab
   # 每日 02:00 执行一次 logdel
   0 2 * * * /usr/bin/logdel
   ```

## Credits

Guo Y.K., MIT License
