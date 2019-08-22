#!/usr/bin/env bash

curl -s -H "Content-Type: application/json" -H "uuid: secret" -X GET localhost:8080/results/foo > clientTest.log


