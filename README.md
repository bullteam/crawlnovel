# crawlnovel 爬书小程序
    本项目由golang + gin + vue-element-admin 开发,是一个爬小说程序，也可以改造成用在爬其他数据。
    同时也是zeus 的实例应用之一，也是为了大家更好的理解zeus 应用所做的一个开源示例产品。
    欢迎大家fork 并递交PR、issue
## 命令

## 演示
[http://crawlnovel.bullteam.cn](http://crawlnovel.bullteam.cn)

> 安装mysql 并创建 `crawlnovel` 库，再创建如下表
```SQL
CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `task_name` varchar(64) NOT NULL COMMENT '任务名称',
  `task_rule_name` varchar(64) NOT NULL COMMENT '任务规则名',
  `task_desc` varchar(512) NOT NULL DEFAULT '' COMMENT '任务描述',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0 未开始 1 开始',
  `counts` int(11) NOT NULL DEFAULT '0' COMMENT '次数',
  `cron_spec` varchar(64) NOT NULL DEFAULT '' COMMENT '定时任务',
  `output_type` varchar(64) NOT NULL COMMENT '导出类型',
  `opt_user_agent` varchar(128) NOT NULL DEFAULT '' COMMENT '用户代理',
  `opt_max_depth` int(11) NOT NULL DEFAULT '0' COMMENT '爬虫最大深度',
  `opt_allowed_domains` varchar(512) NOT NULL DEFAULT '' COMMENT '允许访问的域名',
  `opt_url_filters` varchar(512) NOT NULL DEFAULT '' COMMENT 'URL过滤',
  `opt_max_body_size` int(11) NOT NULL DEFAULT '0' COMMENT '最大body值',
  `opt_request_timeout` int(11) NOT NULL DEFAULT '10' COMMENT '请求超时时间',
  `auto_migrate` tinyint(4) NOT NULL DEFAULT '0',
  `limit_enable` tinyint(4) NOT NULL DEFAULT '0' COMMENT '频率限制',
  `limit_domain_regexp` varchar(128) NOT NULL DEFAULT '' COMMENT '域名glob匹配regexp',
  `limit_domain_glob` varchar(128) NOT NULL DEFAULT '' COMMENT '域名glob匹配',
  `limit_delay` int(11) NOT NULL DEFAULT '0' COMMENT '延迟',
  `limit_random_delay` int(11) NOT NULL DEFAULT '0' COMMENT '随机延迟',
  `limit_parallelism` int(11) NOT NULL DEFAULT '0' COMMENT '请求并发度',
  `proxy_urls` varchar(2048) NOT NULL DEFAULT '' COMMENT '代理列表',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_task_name` (`task_name`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务表'

```
### 设置变量
```bash
export GOPROXY=https://goproxy.cn
export GO111MODULE=on
export CRAWLNOVEL_MYSQL_USERNAME=root
export CRAWLNOVEL_MYSQL_PASSWORD=123456
export CRAWLNOVEL_MYSQL_HOST=127.0.0.1
export CRAWLNOVEL_MYSQL_DB=crawlnovel
export CRAWLNOVEL_MYSQL_PORT=3306
```
```bash
# 编译 (需要golang 1.11+ 环境)
go build -o crawlnovel

# 执行task

./crawlnovel task -c ./config/in-local.yaml --type 1

# 运行web

./crawlnovel server -c ./config/in-local.yaml -p 8081
# 单独执行爬虫


go build -o crawlnovel && ./crawlnovel task -c ./config/in-local.yaml --type 3
```

## 支持小说下载
1.下载小说 -d 支持none,chromedp,phantomjs 三种方式
```
./crawlnovel download -u https://www.biquge5200.cc/0_195/ -d 
```
2.转换小说 -f 可以转换txt,md,epub 三种格式

```
./crawlnovel convert -n 超品相师.crawnovel -f txt

```
3.搜索小说

```
./crawlnovel search -k 超品相师
```

## 特别鸣谢

[gin](https://gin-gonic.com)
[gospider](https://github.com/nange/gospider)
[vue-element-ui](https://panjiachen.gitee.io/vue-element-admin)
[FictionDown](https://github.com/ma6254/FictionDown)