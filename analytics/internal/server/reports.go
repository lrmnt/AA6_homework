package server

import (
	"net/http"
	"time"
)

func (s *Server) getLastDaySum(w http.ResponseWriter, r *http.Request) {
	out, err := s.service.GetLastDaySum(r.Context())
	if err != nil {
		s.s.Respond500(w, "can not load last day sum", err)
		return
	}

	s.s.RespondJSON(w, out)
}

func (s *Server) getLastDayPopugDebtsCount(w http.ResponseWriter, r *http.Request) {
	out, err := s.service.GetLastDayPopugDebtsCount(r.Context())
	if err != nil {
		s.s.Respond500(w, "can not load last day popugs debts", err)
		return
	}

	s.s.RespondJSON(w, out)
}

func (s *Server) getHighestPriceForPeriod(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.s.Respond500(w, "can not parse dorm", err)
		return

	}

	from := r.Form.Get("from")
	to := r.Form.Get("to")
	fromTime, err := time.Parse("2006-01-02", from)
	if err != nil {
		s.s.Respond400(w, "can not parse from time", err)
		return
	}
	toTime, err := time.Parse("2006-01-02", to)
	if err != nil {
		s.s.Respond400(w, "can not parse to time", err)
		return
	}

	out, err := s.service.GetHighestPriceForPeriod(r.Context(), fromTime, toTime)
	if err != nil {
		s.s.Respond500(w, "can not load highest price", err)
		return
	}

	s.s.RespondJSON(w, out)
}
