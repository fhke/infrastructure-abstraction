package test

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	dynamodblocal "github.com/abhirockzz/dynamodb-local-testcontainers-go"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/fhke/infrastructure-abstraction/server/controller"
	"github.com/fhke/infrastructure-abstraction/server/handler"
	dynamoModuleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository/dynamo"
	dynamoStackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository/dynamo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"k8s.io/utils/ptr"
)

func runServer(ctx context.Context, t *testing.T) {
	log := zap.NewNop().Sugar()

	ddbCl := initDynamoDB(ctx, t)
	ginE := gin.Default()
	handler.New(
		log,
		controller.New(
			log,
			dynamoStackRepository.New(ddbCl, dynamoTableNameStacks),
			dynamoModuleRepository.New(ddbCl, dynamoTableNameModules),
		),
	).Register(ginE)

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: ginE,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	for i := range 10 {
		time.Sleep(time.Second * time.Duration(i) / 2)
		if _, err := net.Dial("tcp4", serverAddr); err == nil {
			t.Log("Server is listening")
			return
		}
		t.Log("Waiting for server to start listening")
	}

	t.Fatal("Timed out waiting for server to start listening")
}

func initDynamoDB(ctx context.Context, t *testing.T) *dynamodb.Client {
	// set up a local dynamodb container
	ddb, err := dynamodblocal.RunContainer(ctx)
	require.NoError(t, err, "Setup of dynamodb container failed")

	// tidy up at end of test
	t.Cleanup(func() {
		if err := ddb.Terminate(context.Background()); err != nil {
			panic(err)
		}
	})

	// get client
	dynamoCl, err := ddb.GetDynamoDBClient(ctx)
	require.NoError(t, err, "Get DynamoDB client failed")

	// create tables
	for _, createTableInput := range []dynamodb.CreateTableInput{
		genCreateTableInput(dynamoTableNameStacks, "stackId"),
		genCreateTableInput(dynamoTableNameModules, "moduleName"),
	} {
		_, err = dynamoCl.CreateTable(ctx, &createTableInput)
		require.NoError(t, err, "Failed to create table")
	}

	return dynamoCl
}

func genCreateTableInput(tableName, primaryKey string) dynamodb.CreateTableInput {
	return dynamodb.CreateTableInput{
		TableName:   ptr.To(tableName),
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: ptr.To(primaryKey),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: ptr.To(primaryKey),
				KeyType:       types.KeyTypeHash,
			},
		},
	}

}
