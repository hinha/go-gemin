package blog

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"errors"

	"go-gemin/db/mongo"
)

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Articles struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title    string             `json:"title"`
	Post     string             `json:"post"`
	Create_At time.Time         `json:"createdAt"`
}

var (
	ctx = context.Background()
	ErrInvalidName = errors.New("invalid post name")
	ErrInvalidTitle = errors.New("invalid title name")
)

func InsertArticles(title, post string) (article *Articles, err error)  {
	if (title == "") || (post == "") {
		fmt.Println(err)
		return
	}
	article = &Articles{
		ID:       primitive.NewObjectID(),
		Title:    title,
		Post:     post,
		Create_At: time.Now().Local(),
	}

	//result := article1.Insert
	collection := mongo.DB.Collection("blog")
	result, err := collection.InsertOne(ctx, article)

	if err != nil {
		err = ErrInvalidName
		fmt.Println(err)
		return
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	fmt.Println("\n\n",err, "ops")

	return
}