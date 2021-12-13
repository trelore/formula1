package graph

import (
	"net/http"
	"time"
)

type Resolver struct {
	client *http.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
