package mongodb

import (
	"context"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	log "github.com/sirupsen/logrus"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewMongoDbUserRepository(c *mongo.Collection) repo.UserRepository {
	return &userRepository{
		collection: c,
	}
}

func (r *userRepository) FindAll() ([]*models.User, error) {

	// setup filter
	findOptions := options.Find()
	findOptions.SetLimit(100)

	// Here's an array in which you can store the decoded documents
	var results []*models.User

	log.Debug("collection ", r.collection)

	// Passing nil as the filter matches all documents in the collection
	// cur, err  := r.collection.Find(context.TODO(), nil, findOptions)
	cur, err  := r.collection.Find(context.TODO(), bson.D{{}}, findOptions)

	log.Debug("cursor ", cur)

	if err != nil {
		log.Error("failed to execute the statement " , err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Error(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Error(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	log.Debug(cur)

	return results, nil
}

func (r *userRepository) Save(user *models.User) *models.User {
	insertResult, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(insertResult)
	log.Debug("Saved user to mongo db")

	return user
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) Delete(id int) error {
	return nil
}

func (r *userRepository) FindOne(id string) (*models.User, error) {

	// setup filter
	filter := bson.D{{"_id", id}}

	// Here's an array in which you can store the decoded documents
	var result *models.User

	// Passing nil as the filter matches all documents in the collection
	cur := r.collection.FindOne(context.TODO(), filter).Decode(&result)

	log.Debug(cur)

	return result, nil
}
