// DO NOT EDIT. Automatically generated by hsup
package octav

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/builderscon/octav/octav/service"
	"github.com/builderscon/octav/octav/validator"
	"github.com/gorilla/mux"
	"github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-urlenc"
	"golang.org/x/net/context"
)

type Server struct {
	*mux.Router
}

func Run(l string) error {
	return http.ListenAndServe(l, New())
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.SetupRoutes()
	return s
}

func httpError(w http.ResponseWriter, message string, st int, err error) {
	if pdebug.Enabled {
		if err == nil {
			pdebug.Printf("HTTP Error %s", message)
		} else {
			pdebug.Printf("HTTP Error %s: %s", message, err)
		}
	}
	http.Error(w, http.StatusText(st), st)
}

func getInteger(v url.Values, f string) ([]int64, error) {
	x, ok := v[f]
	if !ok {
		return nil, nil
	}

	ret := make([]int64, len(x))
	for i, e := range x {
		p, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil, err
		}
		ret[i] = p
	}

	return ret, nil
}

func httpCreateConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.CreateConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doCreateConference(context.Background(), w, r, payload)
}

func httpCreateRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doCreateRoom(context.Background(), w, r, payload)
}

func httpCreateSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.CreateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doCreateSession(context.Background(), w, r, payload)
}

func httpCreateUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doCreateUser(context.Background(), w, r, payload)
}

func httpCreateVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpCreateVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.CreateVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPCreateVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doCreateVenue(context.Background(), w, r, payload)
}

func httpDeleteConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.DeleteConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doDeleteConference(context.Background(), w, r, payload)
}

func httpDeleteRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.DeleteRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doDeleteRoom(context.Background(), w, r, payload)
}

func httpDeleteSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.DeleteSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doDeleteSession(context.Background(), w, r, payload)
}

func httpDeleteUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doDeleteUser(context.Background(), w, r, payload)
}

func httpDeleteVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpDeleteVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.DeleteVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPDeleteVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doDeleteVenue(context.Background(), w, r, payload)
}

func httpListRooms(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListRooms")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.ListRoomRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListRoomsRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doListRooms(context.Background(), w, r, payload)
}

func httpListSessionsByConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListSessionsByConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.ListSessionsByConferenceRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListSessionsByConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doListSessionsByConference(context.Background(), w, r, payload)
}

func httpListVenues(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpListVenues")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.ListVenueRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPListVenuesRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doListVenues(context.Background(), w, r, payload)
}

func httpLookupConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.LookupConferenceRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doLookupConference(context.Background(), w, r, payload)
}

func httpLookupRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.LookupRoomRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doLookupRoom(context.Background(), w, r, payload)
}

func httpLookupSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.LookupSessionRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doLookupSession(context.Background(), w, r, payload)
}

func httpLookupUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.LookupUserRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doLookupUser(context.Background(), w, r, payload)
}

func httpLookupVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpLookupVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `get` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.LookupVenueRequest
	if err := urlenc.Unmarshal([]byte(r.URL.RawQuery), &payload); err != nil {
		httpError(w, `Failed to parse url query string`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPLookupVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doLookupVenue(context.Background(), w, r, payload)
}

func httpUpdateConference(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateConference")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.UpdateConferenceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateConferenceRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doUpdateConference(context.Background(), w, r, payload)
}

func httpUpdateRoom(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateRoom")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateRoomRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doUpdateRoom(context.Background(), w, r, payload)
}

func httpUpdateSession(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateSession")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.UpdateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateSessionRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doUpdateSession(context.Background(), w, r, payload)
}

func httpUpdateUser(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateUser")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateUserRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doUpdateUser(context.Background(), w, r, payload)
}

func httpUpdateVenue(w http.ResponseWriter, r *http.Request) {
	if pdebug.Enabled {
		g := pdebug.Marker("httpUpdateVenue")
		defer g.End()
	}
	if strings.ToLower(r.Method) != `post` {
		httpError(w, `Method was `+r.Method, http.StatusNotFound, nil)
	}

	var payload service.UpdateVenueRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}

	if err := validator.HTTPUpdateVenueRequest.Validate(&payload); err != nil {
		httpError(w, `Invalid input`, http.StatusInternalServerError, err)
		return
	}
	doUpdateVenue(context.Background(), w, r, payload)
}

func (s *Server) SetupRoutes() {
	r := s.Router
	r.HandleFunc(`/v1/conference/create`, httpCreateConference)
	r.HandleFunc(`/v1/conference/delete`, httpDeleteConference)
	r.HandleFunc(`/v1/conference/lookup`, httpLookupConference)
	r.HandleFunc(`/v1/conference/update`, httpUpdateConference)
	r.HandleFunc(`/v1/room/create`, httpCreateRoom)
	r.HandleFunc(`/v1/room/delete`, httpDeleteRoom)
	r.HandleFunc(`/v1/room/list`, httpListRooms)
	r.HandleFunc(`/v1/room/lookup`, httpLookupRoom)
	r.HandleFunc(`/v1/room/update`, httpUpdateRoom)
	r.HandleFunc(`/v1/schedule/list`, httpListSessionsByConference)
	r.HandleFunc(`/v1/session/create`, httpCreateSession)
	r.HandleFunc(`/v1/session/delete`, httpDeleteSession)
	r.HandleFunc(`/v1/session/lookup`, httpLookupSession)
	r.HandleFunc(`/v1/session/update`, httpUpdateSession)
	r.HandleFunc(`/v1/user/create`, httpCreateUser)
	r.HandleFunc(`/v1/user/delete`, httpDeleteUser)
	r.HandleFunc(`/v1/user/lookup`, httpLookupUser)
	r.HandleFunc(`/v1/user/update`, httpUpdateUser)
	r.HandleFunc(`/v1/venue/create`, httpCreateVenue)
	r.HandleFunc(`/v1/venue/delete`, httpDeleteVenue)
	r.HandleFunc(`/v1/venue/list`, httpListVenues)
	r.HandleFunc(`/v1/venue/lookup`, httpLookupVenue)
	r.HandleFunc(`/v1/venue/update`, httpUpdateVenue)
}
