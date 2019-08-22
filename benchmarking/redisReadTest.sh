#!/usr/bin/env bash

redis-cli ZRANGE "foo" 0 -1 WITHSCORES > redisTest.out

