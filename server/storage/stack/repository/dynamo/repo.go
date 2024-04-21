package dynamo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	storageErrors "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
	"k8s.io/utils/ptr"
)

const primaryKey = "stackId"

type (
	dynamoRepository struct {
		cl        *dynamodb.Client
		tableName string
	}
	dynamoModel struct {
		StackID string `dynamodbav:"stackId"`
		model.Stack
	}
)

func New(cl *dynamodb.Client, tableName string) repository.Repository {
	return &dynamoRepository{
		cl:        cl,
		tableName: tableName,
	}
}

func (d *dynamoRepository) GetStack(ctx context.Context, name, repository string) (model.Stack, error) {
	return d.getDocument(ctx, name, repository)
}

func (d *dynamoRepository) AddStack(ctx context.Context, stack model.Stack) error {
	itemMap, err := attributevalue.MarshalMap(dynamoModel{
		StackID: toStackId(stack.Name, stack.Repository),
		Stack:   stack,
	})
	if err != nil {
		return fmt.Errorf("error marshalling new value: %w", err)
	}

	_, err = d.cl.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      itemMap,
		TableName: ptr.To(d.tableName),
	})
	if err != nil {
		return fmt.Errorf("error putting item: %w", err)
	}
	return nil
}

func (d *dynamoRepository) UpdateStack(ctx context.Context, stack model.Stack) error {
	if _, err := d.getDocument(ctx, stack.Name, stack.Repository); err != nil {
		if errors.Is(err, storageErrors.ErrNotFound) {
			return err
		}
		return fmt.Errorf("get document error: %w", err)
	}
	return d.AddStack(ctx, stack)
}

func (d *dynamoRepository) getDocument(ctx context.Context, name, repo string) (model.Stack, error) {
	itemOut, err := d.cl.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			primaryKey: &types.AttributeValueMemberS{
				Value: toStackId(name, repo),
			},
		},
		TableName:      ptr.To(d.tableName),
		ConsistentRead: ptr.To(false),
	})

	if err != nil {
		return model.Stack{}, fmt.Errorf("error getting item from DynamoDB: %w", err)
	}
	if len(itemOut.Item) == 0 {
		return model.Stack{}, storageErrors.ErrNotFound
	}

	item := dynamoModel{
		Stack: model.Stack{
			Modules: make(map[string]model.Module),
		}}
	if err := attributevalue.UnmarshalMap(itemOut.Item, &item); err != nil {
		return model.Stack{}, fmt.Errorf("unmarshal error: %w", err)
	}
	return item.Stack, nil

}

func toStackId(name, repo string) string {
	return fmt.Sprintf("%s::%s", name, repo)
}
