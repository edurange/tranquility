#!/usr/bin/env bash

for i in {1..1000}
do
	redis-cli ZADD "foo" 1566429105 "echo bar "$(uuidgen) >> /dev/null
done

