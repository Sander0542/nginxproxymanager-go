#!/bin/sh

set -e

rm -f api_*.go
rm -f model_*.go
rm -r docs/

url="https://raw.githubusercontent.com/Sander0542/nginx-proxy-manager/fix/openapi-schema/backend/doc/api.swagger.json"

openapi-generator-cli generate \
-i $url \
-g go \
-o . \
-p "packageName=nginxproxymanager" \
--git-user-id "sander0542" \
--git-repo-id "nginxproxymanager-go" \
--server-variables "protocol=http,host=localhost,port=81" \
--type-mappings "integer=int64,bool=boolAsInt"

rm .travis.yml
rm git_push.sh

rm -r test/
