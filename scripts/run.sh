#!/bin/bash

export CRAWLNOVEL_MYSQL_USERNAME=root
export CRAWLNOVEL_MYSQL_PASSWORD=123456
export CRAWLNOVEL_MYSQL_HOST=127.0.0.1
export CRAWLNOVEL_MYSQL_DB=crawlnovel
export CRAWLNOVEL_MYSQL_PORT=3306
export CRAWLNOVEL_REDIS_HOST=127.0.0.1
export CRAWLNOVEL_REDIS_PORT=6379
export CRAWLNOVEL_REDIS_PASSWORD=""
export CRAWLNOVEL_MONGO_HOST=127.0.0.1
export CRAWLNOVEL_MONGO_PORT=27017
export CRAWLNOVEL_MONGO_DB=crawnovel
export CRAWLNOVEL_MONGO_USERNAME=
export CRAWLNOVEL_MONGO_PASSWORD=
export CRAWLNOVEL_MONGO_AUTHSOURCE="admin"

./crawlnovel server ./config/in-local.yaml -p 8081