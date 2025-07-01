package pokeapi

import (
	"net/http"
	"time"

	"github.com/hehacz/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
