package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/heaths/booksql/graphql/generated"
	"github.com/heaths/booksql/graphql/model"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, input model.AuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, isbn string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Authors(ctx context.Context, name *string) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, title *string) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Book(ctx context.Context, isbn string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
