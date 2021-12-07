package repository

import (
	"../entity"
)


type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FIndAll() ([]entity.Post, error)
}

type repo struct{}

// NePostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId string = "pragmatic-reviews" // programmatic-reviews/
	collectionName string ="posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	 ctx := context.Background()
	 client, err := firestore.NewClient(ctx, projectId)
	 if err != nill {
		 log.Fatalf("Failed to create a Firestore Client: %v", err)
		 return nill, err
	 }
	 _, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		 "ID":    post.ID,
		 "Title": post.Title,
		 "Text":  post.Text,
	 })

	 if err != nil {
		 log.Fatal("Failed adding a new post : %v", err)
		 return nil, err
	 }
}