package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"rest/pkg/offers"
	"rest/pkg/rest"
	"strconv"
)

type Server struct {
	offersSvc *offers.Service
	router    chi.Router
}

func NewServer(offersSvc *offers.Service, router chi.Router) *Server {
	return &Server{offersSvc: offersSvc, router: router}
}

func (s *Server) Init() error {
	s.router.Get("/offers", s.handleGetOffers)
	s.router.Get("/offers/{id}", s.handleGetOfferByID)
	s.router.Post("/offers", s.handleSaveOffer)
	s.router.Delete("/offers/{id}", s.handleRemoveOfferByID)

	return nil
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func (s *Server) handleGetOffers(writer http.ResponseWriter, request *http.Request) {
	items, err := s.offersSvc.All(request.Context())
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusInternalServerError)
		return
	}

	rest.WriteAsJSON(writer, items, http.StatusOK)
}

func (s *Server) handleGetOfferByID(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusBadRequest)
		return
	}

	item, err := s.offersSvc.ByID(request.Context(), id)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusInternalServerError)
		return
	}

	rest.WriteAsJSON(writer, item, http.StatusOK)
}

func (s *Server) handleSaveOffer(writer http.ResponseWriter, request *http.Request) {
	itemToSave := &offers.Offer{}
	err := json.NewDecoder(request.Body).Decode(&itemToSave)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusBadRequest)
		return
	}

	item, err := s.offersSvc.Save(request.Context(), itemToSave)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusInternalServerError)
		return
	}

	rest.WriteAsJSON(writer, item, http.StatusOK)
}

func (s *Server) handleRemoveOfferByID(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusBadRequest)
		return
	}

	item, err := s.offersSvc.DeleteByID(request.Context(), id)
	if err != nil {
		rest.WriteAsJSONErr(writer, err, http.StatusInternalServerError)
		return
	}

	rest.WriteAsJSON(writer, item, http.StatusOK)
}
