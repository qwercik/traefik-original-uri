package traefik_original_uri

import (
    "context"
    "net/http"
)

type Config struct {
    Header string
}

func CreateConfig() *Config {
    return &Config {
        Header: "X-Original-Uri",
    }
}

type OriginalUri struct {
    next http.Handler
    header string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    return &OriginalUri {
        header: config.Header,
        next: next,
    }, nil
}

func (self *OriginalUri) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    req.Header.Set(self.header, req.Host + req.URL.RequestURI())
    self.next.ServeHTTP(rw, req)
}
