#!/usr/bin/env bash

test='{\"user\": \"foo\", \"time\": 1566429105, \"command\": \"echo bar\"}'

for i in {1..1000}
do
	curl -s -d "$test" -H "Content-Type: application/json" -H "uuid: secret" -X POST localhost:8080/logger >> /dev/null
done
