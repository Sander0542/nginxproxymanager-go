#!/bin/sh

set -e

openapi-generator-cli generate \
-i "http://localhost:3081/api/schema" \
-g go \
-o . \
-c config.yaml
