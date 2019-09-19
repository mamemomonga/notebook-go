#!/bin/bash
set -eu
exec docker run --rm -it --network redis_app_default \
	redis:5.0.5-alpine redis-cli -h redis

