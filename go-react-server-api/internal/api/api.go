package api

import (
	"net/http"

	"github.com/RaulMS03/go-react-server-api/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.HandleCreateRoom)
			r.Get("/", a.HandleGetRoom)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.HandleCreateRoomMessage)
				r.Get("/", a.HandleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/" a.HandleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) HandleCreateRoom(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) HandleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {}   

func (h apiHandler) HandleGetRoom(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) HandleGetRoomMessage(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) HandleGetRoomMessages(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) HandleReactToMessage(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request) {}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {}
