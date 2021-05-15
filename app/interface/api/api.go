package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type APIRoute http.Handler

type API struct {
	r *chi.Mux
}

type APIResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (api *API) AddRouter(route string, router APIRoute) {
	api.r.Mount(route, router)
}

func (api *API) Start(port string) error {
	log.Infof("[API] Listening for REST API Requests on port %v\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), api.r)
}

func (api *API) StartTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, api.r)
}

func NewAPI() *API {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:  []string{"*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
	})

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.AllowContentType("application/json"),
		middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StandardLogger(), NoColor: true}),
		middleware.Compress(6, "gzip"),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.NoCache,
		middleware.Timeout(60*time.Second),
		c.Handler,
	)

	return &API{r: r}

}
