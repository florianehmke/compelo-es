package query

import (
	"compelo/event"
	"log"
)

type RatingHandler struct {
	*Compelo
}

func (r *RatingHandler) on(e event.Event) {
	r.Lock()
	defer r.Unlock()

	log.Println("[Rating Handler] Handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.MatchCreated:
		r.handleMatchCreated(e)
	}
}

func (r *RatingHandler) handleMatchCreated(e *event.MatchCreated) {
	log.Println("==> Calculate the rating..")
}
