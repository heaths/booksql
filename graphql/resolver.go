//go:generate go run github.com/99designs/gqlgen generate
package graphql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/heaths/booksql/graphql/model"
)

type Resolver struct {
	authors *aztables.Client
	books   *aztables.Client
}

func NewResolver(ctx context.Context, connectionString string, populate bool) (*Resolver, error) {
	svc, err := aztables.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}

	authors, err := svc.CreateTable(ctx, "authors", nil)
	if err != nil {
		return nil, err
	}

	if populate {
		err = seed(ctx, authors, seedAuthors)
		if err != nil {
			return nil, err
		}
	}

	books, err := svc.CreateTable(ctx, "books", nil)
	if err != nil {
		return nil, err
	}

	if populate {
		err = seed(ctx, books, seedBooks)
		if err != nil {
			return nil, err
		}
	}

	resolver := &Resolver{
		authors: authors,
		books:   books,
	}

	return resolver, nil
}

func seed(ctx context.Context, client *aztables.Client, fn func() []model.Entity) error {
	options := &aztables.InsertEntityOptions{
		UpdateMode: aztables.ReplaceEntity,
	}
	for _, entity := range fn() {
		buffer, err := json.Marshal(entity.Entity())
		if err != nil {
			return err
		}
		_, err = client.InsertEntity(ctx, buffer, options)
		if err != nil {
			return err
		}
	}
	return nil
}

func seedAuthors() []model.Entity {
	return []model.Entity{
		&model.Author{
			ID:        "1",
			FirstName: to.StringPtr("Frank"),
			LastName:  "Herbert",
		},
		&model.Author{
			ID:        "2",
			FirstName: to.StringPtr("Douglas"),
			LastName:  "Adams",
		},
	}
}

func seedBooks() []model.Entity {
	return []model.Entity{
		&model.Book{
			Isbn:  "978-0593099322",
			Title: "Dune",
			Authors: []*model.Author{
				{
					ID: "1",
				},
			},
		},
		&model.Book{
			Isbn:  "978-0593098233",
			Title: "Dune Messiah",
			Authors: []*model.Author{
				{
					ID: "1",
				},
			},
		},
		&model.Book{
			Isbn:  "978-0593098240",
			Title: "Children of Dune",
			Authors: []*model.Author{
				{
					ID: "1",
				},
			},
		},
		&model.Book{
			Isbn:  "978-0345418913",
			Title: "Hitchhiker's Guide to the Galaxy",
			Authors: []*model.Author{
				{
					ID: "2",
				},
			},
		},
		&model.Book{
			Isbn:  "978-1529034530",
			Title: "Restaurant at the End of the Universe",
			Authors: []*model.Author{
				{
					ID: "2",
				},
			},
		},
		&model.Book{
			Isbn:  "978-1529034547",
			Title: "Restaurant at the End of the Univers",
			Authors: []*model.Author{
				{
					ID: "2",
				},
			},
		},
	}
}
