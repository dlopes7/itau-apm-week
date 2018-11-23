#!/usr/bin/env bash

docker ps -a --format "{{.Names}}" | grep shopizer | xargs docker rm -f 