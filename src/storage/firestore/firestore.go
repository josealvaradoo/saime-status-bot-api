package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/josealvaradoo/saime-status-bot/src/utils"
	"google.golang.org/api/option"
)

var (
	Client *firestore.Client
	Ctx    = context.Background()
)

func newFirestoreConnection() *firestore.Client {
	keys := option.WithCredentialsFile(".gcp/credentials.json")
	client, err := firestore.NewClient(Ctx, utils.Env(utils.GCP_PROJECT_ID), keys)

	if err != nil {
		fmt.Println("Could not create a firestore connection")
	}

	return client
}

func New() {
	Client = newFirestoreConnection()
}
