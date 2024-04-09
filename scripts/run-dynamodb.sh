#!/usr/bin/env bash

docker run \
    --rm \
    -d \
    -p 8000:8000 \
    --name dynamo \
    amazon/dynamodb-local:latest \
    -jar DynamoDBLocal.jar -sharedDb

sleep 3

export AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY AWS_ENDPOINT_URL=http://localhost:8000

aws dynamodb create-table \
    --table-name stacks \
    --attribute-definitions \
        AttributeName=stackId,AttributeType=S \
    --key-schema \
        AttributeName=stackId,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST > /dev/null

aws dynamodb create-table \
    --table-name modules \
    --attribute-definitions \
        AttributeName=moduleName,AttributeType=S \
    --key-schema \
        AttributeName=moduleName,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST > /dev/null
