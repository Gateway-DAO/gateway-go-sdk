#!/bin/bash

URL="https://dev.api.gateway.tech/swagger/doc.json"
FILE="api.json"

# Download the OpenAPI spec
curl -o $FILE $URL

