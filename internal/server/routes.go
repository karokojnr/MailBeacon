package server

import (
	"net/http"

	"MailBeacon/cmd/web"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	newsletterRouter := chi.NewRouter()
	newsletterRouter.Post("/signup", s.NewsletterSignup)
	newsletterRouter.Post("/confirm-email", s.ConfirmNewsletterSignup)
	
	newsletterRouter.Post("/send-confirmation-email", s.SendConfirmationEmail)
	newsletterRouter.Post("/send-welcome-email", s.SendWelcomeEmail)

	r.Mount("/api/newsletter", newsletterRouter)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)

	return r
}
