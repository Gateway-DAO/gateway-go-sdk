#!/bin/bash

URL="https://dev.api.gateway.tech/docs/openapi3_full.json"
FILE="api.json"

curl -o $FILE $URL

