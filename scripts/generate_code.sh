#!/bin/sh

`go env GOPATH`/bin/swagger generate client --skip-validation -f ./swagger.json  -A avatax -P models.Principal --default-scheme https 