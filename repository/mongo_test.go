package repository

import (
	"github.com/shadev/recipescator/testresources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAll_ok(t *testing.T) {
	testee := MongoRepo{}

	actualRecipes, e := testee.FindAll()

	assert.NoError(t, e)
	assert.Len(t, actualRecipes, 2)
	assert.ElementsMatch(t, testresources.SampleRecipes(), actualRecipes)
}
