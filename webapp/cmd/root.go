package cmd

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const port = "8081"

//go:embed static/*
var staticFiles embed.FS

var driversStandingsPage = template.Must(template.ParseFS(staticFiles, "static/drivers.html"))
var constructorsStandingsPage = template.Must(template.ParseFS(staticFiles, "static/constructors.html"))

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "webapp",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			log.Fatal(err)
		}
	},
}

func run() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	hostAddress := "http://localhost:8080/query"
	if host, ok := os.LookupEnv("GRAPHQL_EP"); ok {
		hostAddress = host
	}
	c := client{host: hostAddress, log: sugar}
	http.HandleFunc("/drivers", c.driversStandings)
	http.HandleFunc("/constructors", c.constructorsStandings)
	addr := fmt.Sprintf(":%s", port)
	sugar.Infof("running on address: %s", addr)
	return http.ListenAndServe(addr, nil)
}

type client struct {
	host string
	log  *zap.SugaredLogger
}

var driversQuery struct {
	DriverStandings struct {
		Drivers []struct {
			Points graphql.String `graphql:"points"`
			Driver struct {
				GivenName  graphql.String
				FamilyName graphql.String
			} `graphql:"Driver"`
		} `graphql:"drivers"`
	} `graphql:"DriverStandings(filter: {top: 5})"`
}

func (c *client) driversStandings(w http.ResponseWriter, r *http.Request) {
	client := graphql.NewClient(c.host, &http.Client{Timeout: 10 * time.Second})
	err := client.Query(context.Background(), &driversQuery, nil)
	if err != nil {
		c.log.Warnw("getting drivers standings",
			"err", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	driversStandingsPage.Execute(w, driversQuery)
}

var constructorsQuery struct {
	ConstructorStandings struct {
		Teams []struct {
			Position graphql.String `graphql:"position"`
			Points   graphql.String `graphql:"points"`
			Team     struct {
				Name graphql.String
				URL  graphql.String
			} `graphql:"team"`
		} `graphql:"teams"`
	} `graphql:"ConstructorStandings"`
}

func (c *client) constructorsStandings(w http.ResponseWriter, r *http.Request) {
	client := graphql.NewClient(c.host, &http.Client{Timeout: 10 * time.Second})
	err := client.Query(context.Background(), &constructorsQuery, nil)
	if err != nil {
		c.log.Warnw("getting drivers standings",
			"err", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	constructorsStandingsPage.Execute(w, constructorsQuery)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
