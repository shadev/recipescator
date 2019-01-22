package repository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/shadev/recipescator/testresources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MongoTestSuite struct {
	suite.Suite
	testee MongoRepo
}

func (suite *MongoTestSuite) SetupTest() {
	testDbName := "recipescator-test-db"
	testColName := "recipes-find"

	client, _ := mongo.Connect(context.Background(), "mongodb://localhost")
	err := client.
		Database(testDbName).
		Collection(testColName).
		Drop(context.Background())
	_, err = client.
		Database(testDbName).
		Collection(testColName).
		InsertMany(context.Background(), testresources.SampleRecipesAsInterface())
	assert.Nil(suite.T(), err)

	suite.testee = MongoRepo{Client: client, Db: testDbName, Collection: testColName}
}

func (suite *MongoTestSuite) TestFindAll_ok() {
	t := suite.T()

	actualRecipes, e := suite.testee.FindAll()

	assert.NoError(t, e)
	assert.Len(t, actualRecipes, 2)
	assert.ElementsMatch(t, testresources.SampleRecipes(), actualRecipes)
}

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
