package dynamo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	storageErrors "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
	"k8s.io/utils/ptr"
)

const primaryKey = "moduleName"

type dynamoRepository struct {
	cl        *dynamodb.Client
	tableName string
}

type dynamoModel struct {
	Name string `dynamodbav:"moduleName"`

	Versions []model.ModuleVersion `dynamodbav:"versions"`
}

func New(cl *dynamodb.Client, tableName string) repository.Repository {
	return &dynamoRepository{
		cl:        cl,
		tableName: tableName,
	}
}

func (d *dynamoRepository) GetVersions(name string) ([]model.ModuleVersion, error) {
	item, err := d.getDocument(context.TODO(), name)
	if err != nil {
		return nil, err
	}
	return item.Versions, nil
}

func (d *dynamoRepository) AddVersion(mv model.ModuleVersion) error {
	item, err := d.getDocument(context.TODO(), mv.Name)
	if err != nil {
		if errors.Is(err, storageErrors.ErrNotFound) {
			item = dynamoModel{
				Name:     mv.Name,
				Versions: make([]model.ModuleVersion, 0),
			}
		} else {
			return fmt.Errorf("error getting data: %w", err)
		}
	}

	item.Versions = append(item.Versions, mv)
	itemMap, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("error marshalling new value: %w", err)
	}

	_, err = d.cl.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item:      itemMap,
		TableName: ptr.To(d.tableName),
	})
	if err != nil {
		return fmt.Errorf("error putting item: %w", err)
	}
	return nil
}

func (d *dynamoRepository) getDocument(ctx context.Context, name string) (dynamoModel, error) {
	itemOut, err := d.cl.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			primaryKey: &types.AttributeValueMemberS{
				Value: name,
			},
		},
		TableName:      ptr.To(d.tableName),
		ConsistentRead: ptr.To(false),
	})
	if err != nil {
		return dynamoModel{}, fmt.Errorf("error getting item from DynamoDB: %w", err)
	}
	if len(itemOut.Item) == 0 {
		return dynamoModel{}, storageErrors.ErrNotFound
	}

	item := dynamoModel{
		Versions: make([]model.ModuleVersion, 0),
	}
	if err := attributevalue.UnmarshalMap(itemOut.Item, &item); err != nil {
		return dynamoModel{}, fmt.Errorf("unmarshal error: %w", err)
	}
	return item, nil

}
