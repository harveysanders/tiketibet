package mongo_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	db "github.com/harveysanders/tiketibet/auth/mongo"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testStore store
var testDBName = "auth-test"
var testDBClient *mongo.Client

// TestMain will only run if the build tags include "integration". This TestMain will setup a mongoDB instance with dockertest and tear the container down when finished. This setup/teardown process can take 40 seconds or more.
func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pull mongodb docker image for version 7.0.2
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "7",
		Env: []string{
			// username and password for mongodb superuser
			"MONGO_INITDB_ROOT_USERNAME=root",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		testDBClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://root:password@localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return testDBClient.Ping(context.TODO(), nil)
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	testStore = db.NewStore(testDBClient, testDBName)

	// run tests
	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err = testDBClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	os.Exit(code)
}
