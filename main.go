package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	opw "github.com/1password/onepassword-sdk-go"
)

var token = os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")
var vaultID = os.Getenv("OP_VAULT_ID")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := opw.NewClient(
		ctx,
		opw.WithServiceAccountToken(token),
		opw.WithIntegrationInfo("Test Integration", "v0.1.0"),
	)
	if err != nil {
		log.Fatal(err)
	}
	items, err := list(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		log.Printf("found: %s", item.Title)
	}
	err = createSecureNote(ctx, c, "test B", "some data")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok")
}

func list(ctx context.Context, c *opw.Client) ([]*opw.ItemOverview, error) {
	// Note!
	//
	// Use Items.Get on the overview to get the meat
	it, err := c.Items.ListAll(ctx, vaultID)
	if err != nil {
		return nil, err
	}
	var out []*opw.ItemOverview
	for {
		itemOverview, err := it.Next()
		if err != nil {
			if errors.Is(err, opw.ErrorIteratorDone) {
				break
			}
			return nil, err
		}
		out = append(out, itemOverview)
	}
	return out, nil
}

func createSecureNote(ctx context.Context, c *opw.Client, title string, body string) error {
	// Note!
	//
	// Secret notes require a section with a body:value field
	sectionID := "body"

	_, err := c.Items.Create(ctx, opw.ItemCreateParams{
		VaultID:  vaultID,
		Category: opw.ItemCategorySecureNote,
		Title:    title,
		Fields: []opw.ItemField{
			{
				ID:        "notesPlain",
				Title:     "Text",
				FieldType: opw.ItemFieldTypeText,
				Value:     body,
				// Details: &opw.ItemFieldDetails{
				// 	Type: "STRING",
				// },
				SectionID: &sectionID,
			},
		},
		Sections: []opw.ItemSection{
			{
				ID:    sectionID,
				Title: "Body",
			},
		},
	})
	return err
}

func createPassword(ctx context.Context, c *opw.Client, title string, body string) error {
	_, err := c.Items.Create(ctx, opw.ItemCreateParams{
		VaultID:  vaultID,
		Category: opw.ItemCategoryPassword,
		Title:    title,
		Fields: []opw.ItemField{
			{
				ID:        "password",
				FieldType: opw.ItemFieldTypeConcealed,
				Value:     body,
				// Details: &opw.ItemFieldDetails{
				// 	Type: "STRING",
				// },
			},
		},
	})
	return err
}
