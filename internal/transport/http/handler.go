package http

import (
	"MailBeacon/cmd/web"
	"MailBeacon/internal/newsletter"
	"MailBeacon/internal/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type Handler struct {
	Router  *chi.Mux
	Server  *http.Server
	Service newsletter.NewsletterSevice
}

func NewHandler(service newsletter.NewsletterSevice) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)

	h.SetupRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	h.Server = &http.Server{
		Addr:         "0.0.0.0:" + strconv.Itoa(port),
		Handler:      h.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return h
}

func (h *Handler) SetupRoutes() {
	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJson(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	newsletterRouter := chi.NewRouter()
	newsletterRouter.Post("/signup", h.NewsletterSignup)
	newsletterRouter.Get("/confirm-email", h.ConfirmNewsletterSignup)

	newsletterRouter.Post("/send-confirmation-email", h.SendConfirmationEmail)
	newsletterRouter.Post("/send-welcome-email", h.SendWelcomeEmail)

	h.Router.Mount("/api/v1", apiRouter)
	h.Router.Mount("/api/v1/newsletter", newsletterRouter)

	// * Serve static files
	fileServer := http.FileServer(http.FS(web.Files))
	h.Router.Handle("/assets/*", fileServer)
	// h.Router.Get("/", templ.Handler(web.Newsletter()).ServeHTTP)
	// h.Router.Post("/hello", web.HelloWebHandler)

	h.Router.Group(func(r chi.Router) {
		r.Get("/", templ.Handler(web.SignUp()).ServeHTTP)
		// r.Post("/hello", templ.Handler(web.HelloPost()).ServeHTTP)
	})

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("shut down gracefully")
	return nil
}
