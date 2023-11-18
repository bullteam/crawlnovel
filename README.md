# crawlnovel 爬书小程序
    本项目是小说下载工具和转换工具。
## 命令
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
[gospider](https://github.com/nange/gospider)
[FictionDown](https://github.com/ma6254/FictionDown)