package command

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ameykpatil/chatbot/config"
	"github.com/ameykpatil/chatbot/storage"
)

type (
	// BaseCommand hold common command properties
	BaseCommand struct {
		ctx    context.Context
		config *config.Specification
	}
)

// NewBaseCommand creates a structure with common shared properties of the commands
func NewBaseCommand(ctx context.Context, config *config.Specification) BaseCommand {
	return BaseCommand{
		ctx:    ctx,
		config: config,
	}
}

// setup initialises basic common stuff
func (cmd *BaseCommand) setup(ctx context.Context) *mongo.Client {

	// database connection
	log.Println("initializing the database connection")
	mongoClient, err := cmd.newDatabaseConnection(ctx)
	if err != nil {
		log.Fatalf("failed initializing database : %s", err)
	}

	// this step is to load the database with some replies
	// ideal way would be to give a facility of POST api
	err = storage.LoadRepliesToDB(ctx, mongoClient, storage.DefaultReplies)
	if err != nil {
		log.Fatalf("failed to load replies to database : %s", err)
	}

	// setup or initialize more stuff here such as
	// prometheus exporter for metrics
	// register tracer

	return mongoClient
}

// creates new database connection & returns corresponding client
func (cmd *BaseCommand) newDatabaseConnection(ctx context.Context) (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(cmd.config.MongoDB.URI))
	if err != nil {
		log.Fatal(err)
	}

	ctxTimeout, _ := context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(ctxTimeout)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

// creates new http with given timeout to use for making outgoing calls
func (cmd *BaseCommand) newHTTPClient(name string, timeout time.Duration) *http.Client {

	c := &http.Client{
		Timeout: timeout,
	}

	// add common stuff required for http client of outgoing calls
	// circuit breaker
	// passing tracing info
	// common logging
	// etc.

	return c
}
