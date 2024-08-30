
download-yaml:
	./scripts/fetch_api.sh

generate-client:
	oapi-codegen -package=generate -generate=types -o=generate/generated-client.go   ./api.yaml

generate-types:
	ogen --target target/dir -package api --clean openapi.yaml

	# swagger2openapi --yaml --outfile openapi.yaml https://dev.api.gateway.tech/swagger/doc.json


