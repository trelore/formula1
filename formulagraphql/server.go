package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/trelore/formula1/formulagraphql/graph"
	"github.com/trelore/formula1/formulagraphql/graph/generated"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
)

const defaultPort = "8080"

var (
	ocagentHost = os.Getenv("OC_AGENT_HOST")
)

func handle(path string, h func(w http.ResponseWriter, r *http.Request)) {
	http.Handle(path, &ochttp.Handler{
		Handler: http.HandlerFunc(h),
	})
}

func logging(log *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			// Loop over header names
			for name, values := range r.Header {
				// Loop over all values for the name.
				for _, value := range values {
					log.Infow("headers", name, value)
				}
			}
		})
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(5*time.Second),
		ocagent.WithAddress(ocagentHost),
		ocagent.WithServiceName("formulagraphql"))
	if err != nil {
		log.Fatalf("Failed to create ocagent-exporter: %v", err)
	}
	trace.RegisterExporter(oce)

	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver()}))
	loggingMiddleware := logging(sugar)

	handle("/", playground.Handler("GraphQL playground", "/query"))
	handle("/query", loggingMiddleware(srv).ServeHTTP)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
