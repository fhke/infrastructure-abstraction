# POC of managed infra interface

## Quickstart

1. Run local DynamoDB: `./scripts/run-dynamodb.sh`
1. Set AWS environment variables: `export AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY AWS_ENDPOINT_URL=http://localhost:8000`
1. Start configuration management server: `go run ./server -debug -storage dynamo`
1. Configure modules: `go run ./client/examples/common/setup`
1. Build example stack: `go run ./client/examples/cdk/`
