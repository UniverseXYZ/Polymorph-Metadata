package handlers

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/constants"
	"github.com/polymorph-metadata/db"
	"github.com/polymorph-metadata/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRarityById endpoints accepts id of a single polymorph and information for a single polymorph.
//
// If no polymorph is found returns empty response
func GetRarityById(id int) structs.RarityServiceResponse {
	godotenv.Load()
	polymorphDBName := os.Getenv("POLYMORPH_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return structs.RarityServiceResponse{}
	}

	findOptions := options.FindOneOptions{}
	removePrivateFieldsSingle(&findOptions)

	var filter bson.M = bson.M{}
	filter = bson.M{constants.MorphFieldNames.TokenId: id}

	var result = structs.RarityServiceResponse{}
	curr := collection.FindOne(context.Background(), filter, &findOptions)

	curr.Decode(&result)

	return result
}

// removePrivateFieldsSingle removes internal fields that are of no interest to the users of the API.
//
// Configuration of these fields can be found in helpers.apiConfig.go
func removePrivateFieldsSingle(findOptions *options.FindOneOptions) {
	noProjectionFields := bson.M{}
	for _, field := range config.MORPHS_NO_PROJECTION_FIELDS {
		noProjectionFields[field] = 0
	}
	findOptions.SetProjection(noProjectionFields)
}
