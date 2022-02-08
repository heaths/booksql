package model

import "github.com/Azure/azure-sdk-for-go/sdk/data/aztables"

type Entity interface {
	Entity() *aztables.EDMEntity
}

func (e *Author) Entity() *aztables.EDMEntity {
	properties := map[string]interface{}{
		"id":       e.ID,
		"lastName": e.LastName,
	}
	if e.FirstName != nil {
		properties["firstName"] = *e.FirstName
	}
	if e.Birthday != nil {
		properties["birthday"] = aztables.EDMDateTime(*e.Birthday)
	}
	return &aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: "author",
			RowKey:       e.ID,
		},
		Properties: properties,
	}
}

func (e *Book) Entity() *aztables.EDMEntity {
	properties := map[string]interface{}{
		"isbn":  e.Isbn,
		"title": e.Title,
	}
	if e.Year != nil {
		properties["year"] = *e.Year
	}
	if len(e.Authors) != 0 {
		authors := make([]string, len(e.Authors))
		for i, author := range e.Authors {
			authors[i] = author.ID
		}
		properties["authorIds"] = authors
	}
	return &aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: "book",
			RowKey:       e.Isbn,
		},
		Properties: properties,
	}
}
