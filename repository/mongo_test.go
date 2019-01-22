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
	testColName := "recipes"

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

func (suite *MongoTestSuite) TestFindOne_ok() {
	t := suite.T()

	foundRecipe, e := suite.testee.FindOne("123456789")

	assert.NoError(t, e)
	assert.Equal(t, testresources.SampleRecipes()[0], foundRecipe)
}

func (suite *MongoTestSuite) TestInsert_ok() {
	t := suite.T()

	rid, e1 := suite.testee.Insert(testresources.NewRecipeToBeInserted())

	assert.NoError(t, e1)
	assert.Len(t, rid, 36)

	foundRecipe, e2 := suite.testee.FindOne(rid)
	assert.NoError(t, e2)
	expected := testresources.NewRecipeToBeInserted()
	expected.Rid = rid
	assert.Equal(t, foundRecipe, &expected)
}

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
