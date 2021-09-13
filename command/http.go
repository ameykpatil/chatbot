package command

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ameykpatil/chatbot/client"
	"github.com/ameykpatil/chatbot/handler"
	"github.com/ameykpatil/chatbot/service"
	"github.com/ameykpatil/chatbot/storage"
)

// HTTPCommand holds the necessary data for http server
type HTTPCommand struct {
	BaseCommand
}

// NewHTTPCommand creates an instance of HTTPCommand
func NewHTTPCommand(baseCommand BaseCommand) *HTTPCommand {
	return &HTTPCommand{baseCommand}
}

// Run starts up the http
func (cmd *HTTPCommand) Run(cliCtx *cli.Context) error {

	ctx, cancelCtx := context.WithCancel(context.TODO())
	defer cancelCtx()

	// setup basic stuff
	mongoClient := cmd.setup(ctx)

	// init router
	router := cmd.newRouter(ctx, mongoClient)

	log.Printf("starting http server on %d port", cmd.config.HTTP.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cmd.config.HTTP.Port), router)
}

// creates a router with all the api endpoints
func (cmd *HTTPCommand) newRouter(ctx context.Context, mongoClient *mongo.Client) *chi.Mux {
	r := chi.NewRouter()

	replyReader := storage.NewReplyReader(mongoClient)
	intentClient := cmd.newIntentClient()
	replyService := service.NewReplyService(*replyReader, *intentClient)

	getReplyHandler := handler.NewGetReplyHandler(*replyService)

	// endpoint to check if the server is running
	r.Get("/ping", handler.PingHandler)

	// endpoint to serve a reply
	r.Method(http.MethodGet, "/replies", getReplyHandler)

	return r
}

// creates a client for intent-service
func (cmd *HTTPCommand) newIntentClient() *client.IntentClient {
	return client.NewIntentClient(
		cmd.newHTTPClient("intent-service", time.Second*10),
		cmd.config.IntentService.URL,
		cmd.config.IntentService.APIKey,
	)
}
