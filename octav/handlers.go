package octav

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"context"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/service"
	"github.com/lestrrat/go-apache-logformat"
	ical "github.com/lestrrat/go-ical"
	"github.com/lestrrat/go-pdebug"
)

var mwset middlewareSet

type middlewareSet struct{}

func (m middlewareSet) Wrap(h http.Handler) http.Handler {
	return apachelog.CombinedLog.Wrap(h, os.Stdout)
}

const trustedCall = "octav.api.trustedCall"

func isTrustedCall(ctx context.Context) bool {
	allow, ok := ctx.Value(trustedCall).(bool)
	return ok && allow
}

func init() {
	httpError = httpErrorAsJSON
	mwset = middlewareSet{}
}

type httpCoder interface {
	HTTPCode() int
}

func httpCodeFromError(err error) int {
	if v, ok := err.(httpCoder); ok {
		return v.HTTPCode()
	}
	return http.StatusInternalServerError
}

func httpWithOptionalBasicAuth(h HandlerWithContext) HandlerWithContext {
	return wrapBasicAuth(h, true)
}
func httpWithBasicAuth(h HandlerWithContext) HandlerWithContext {
	return wrapBasicAuth(h, false)
}

func wrapBasicAuth(h HandlerWithContext, authIsOptional bool) HandlerWithContext {
	return HandlerWithContext(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		// Verify access token in the Basic-Auth
		clientID, clientSecret, ok := r.BasicAuth()
		if !ok {
			if pdebug.Enabled {
				pdebug.Printf("clientID and/or clientSecret not provided")
			}

			if authIsOptional {
				// if the authentication is optional, then we can just proceed
				if pdebug.Enabled {
					pdebug.Printf("authentication is optional, allowing regular access")
				}
				h(ctx, w, r)
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="octav"`)
			httpError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, nil)
			return
		}

		s := service.Client()
		if err := s.Authenticate(clientID, clientSecret); err != nil {
			if pdebug.Enabled {
				pdebug.Printf("Failed to authenticate client: %s", err)
			}
			code := httpCodeFromError(err)
			httpError(w, http.StatusText(code), code, err)
			return
		}

		if pdebug.Enabled {
			pdebug.Printf("Authentication succeeded, proceeding to call handler")
		}
		ctx = context.WithValue(ctx, trustedCall, true)
		h(ctx, w, r)
	})
}

func httpJSONWithStatus(w http.ResponseWriter, v interface{}, st int) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		httpError(w, `encode json`, http.StatusInternalServerError, err)
		return
	}

	if pdebug.Enabled {
		pdebug.Printf("response buffer: %s", buf.String())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(st)
	buf.WriteTo(w)
}

type jsonerr struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func httpErrorAsJSON(w http.ResponseWriter, message string, st int, err error) {
	v := jsonerr{
		Message: message,
	}
	if err != nil {
		v.Error = err.Error()
	}
	httpJSONWithStatus(w, v, st)
}

func httpJSON(w http.ResponseWriter, v interface{}) {
	httpJSONWithStatus(w, v, http.StatusOK)
}

func doHealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	httpJSON(w, map[string]interface{}{
		"message": "Hello, World!",
	})
}

func doCreateConferenceSeries(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateConferenceSeriesRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doCreateConferenceSeries")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateConferenceSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.ConferenceSeries()
	var c model.ConferenceSeries
	if err := s.CreateFromPayload(tx, &c, payload); err != nil {
		httpError(w, `CreateConferenceSeries`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateConferenceSeries`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doLookupConferenceSeries(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupConferenceSeriesRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupConferenceSeries")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupConferenceSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.ConferenceSeries()
	var c model.ConferenceSeries
	if err := s.LookupFromPayload(tx, &c, payload); err != nil {
		httpError(w, `LookupConferenceSeries`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doDeleteConferenceSeries(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteConferenceSeriesRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteConferenceSeries")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteConferenceSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.ConferenceSeries()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteConferenceSeries`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteConferenceSeries`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doListConferenceSeries(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListConferenceSeriesRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferencesSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.ConferenceSeries()
	l := []model.ConferenceSeries{}
	if err := s.LoadByRange(tx, &l, payload.Since.String, int(payload.Limit.Int)); err != nil {
		httpError(w, `ListConferenceSeries`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, l)
}

func doAddConferenceSeriesAdmin(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddConferenceSeriesAdminRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddConferenceSeriesAdmin`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.ConferenceSeries()
	if err := s.AddAdministratorFromPayload(tx, payload); err != nil {
		httpError(w, `AddConferenceSeriesAdmin`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `AddConferenceSeriesAdmin`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doCreateConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateConferenceRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doCreateConference")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var c model.Conference
	if err := s.CreateFromPayload(tx, payload, &c); err != nil {
		httpError(w, `CreateConference`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateConference`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, model.ObjectID{ID: c.ID, Type: "conference"})
}

func doLookupConferenceBySlug(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupConferenceBySlugRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupConferenceBySlug")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupConferenceBySlug`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var c model.Conference
	if err := s.LookupBySlug(tx, &c, payload); err != nil {
		httpError(w, `LookupConferenceBySlug`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doLookupConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupConferenceRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupConference")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var c model.Conference
	if err := s.LookupFromPayload(tx, &c, payload); err != nil {
		httpError(w, `LookupConference`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doUpdateConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateConferenceRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateConference")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var updateErr error
	if updateErr = s.UpdateFromPayload(ctx, tx, payload); !errors.IsIgnorable(updateErr) {
		httpError(w, `UpdateConference`, http.StatusInternalServerError, updateErr)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateConference`, http.StatusInternalServerError, err)
		return
	}

	// This extra bit is for finalizing the image upload
	if cb, ok := errors.IsFinalizationRequired(updateErr); ok {
		if err := cb(); err != nil {
			httpError(w, `Failed to finalize image uploads`, http.StatusInternalServerError, err)
			return
		}
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteConferenceRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteConference")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.Delete(tx, payload.ID); err != nil {
		httpError(w, `DeleteConference`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteConference`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteConferenceDate(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteConferenceDateRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteConferenceDate`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.DeleteDateFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteConferenceDates`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteConferenceDates`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doListConferenceDate(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListConferenceDateRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferenceDate`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var cdl model.ConferenceDateList
	if err := s.LoadDates(tx, &cdl, payload.ConferenceID); err != nil {
		httpError(w, `ListConferenceDate`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, cdl)
}

func doAddConferenceDate(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateConferenceDateRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddConferenceDates`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	var v model.ConferenceDate
	s := service.ConferenceDate()
	if err := s.CreateFromPayload(tx, payload, &v); err != nil {
		httpError(w, `AddConferenceDates`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `AddConferenceDates`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doDeleteConferenceAdmin(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteConferenceAdminRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.DeleteAdministratorFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doListConferenceAdmin(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListConferenceAdminRequest) {
	trustedCall := isTrustedCall(ctx)

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var cdl model.UserList
	if err := s.LoadAdmins(tx, &cdl, trustedCall, payload.ConferenceID, payload.Lang.String); err != nil {
		httpError(w, `ListConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, cdl)
}

func doAddConferenceAdmin(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddConferenceAdminRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.AddAdministratorFromPayload(tx, payload); err != nil {
		httpError(w, `AddConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `AddConferenceAdmin`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteTrack(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteTrackRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteTrack`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Track()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteTrack`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteTrack`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doCreateTrack(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateTrackRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateTrack`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Track()
	if err := s.CreateFromPayload(tx, payload, nil); err != nil {
		httpError(w, `CreateTrack`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateTrack`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doUpdateTrack(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateTrackRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateTrack`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Track()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateTrack`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateTrack`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupTrack(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupTrackRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupTrack")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupTrack`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Track()
	var v model.Track
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupTrack`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doDeleteConferenceVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteConferenceVenueRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteConferenceVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.DeleteVenueFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteConferenceVenue`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteConferenceVenue`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doAddConferenceVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddConferenceVenueRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddConferenceVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.AddVenueFromPayload(tx, payload); err != nil {
		httpError(w, `AddConferenceVenue`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `AddConferenceVenue`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doListConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListConferenceRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferences`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var v model.ConferenceList
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListConference`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doCreateRoom(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateRoomRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateRoom`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Room()
	var v model.Room
	if err := s.CreateFromPayload(tx, &v, payload); err != nil {
		httpError(w, `CreateRoom`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateRoom`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doUpdateRoom(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateRoomRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateRoom")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateRoom`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Room()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateRoom`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateRoom`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doCreateSession(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateSessionRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateSession`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Session()
	var v model.Session
	if err := s.CreateFromPayload(tx, &v, payload); err != nil {
		httpError(w, `CreateSession`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateSession`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)

	go s.PostSocialServices(&v)
}

func doUpdateSession(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateSessionRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateSession`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Session()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateConference`, http.StatusNotFound, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateSession`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteSession(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteSessionRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteSession")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteSession`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Session()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteSession`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteSession`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doCreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateUserRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateUser`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	var v model.User
	if err := s.CreateFromPayload(tx, &v, payload); err != nil {
		httpError(w, `CreateUser`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateUser`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doDeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteUserRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteUser")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteUser`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteUser`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteUser`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doListUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListUserRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListUsers`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	var v model.UserList
	payload.TrustedCall = isTrustedCall(ctx)
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListUsers`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doLookupUserByAuthUserID(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupUserByAuthUserIDRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupUserByAuthUserID")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupUserByAuthUserID`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	var v model.User
	payload.TrustedCall = isTrustedCall(ctx)
	if err := s.LookupUserByAuthUserIDFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupUserByAuthUserID`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doLookupUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupUserRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupUser")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupUser`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	var v model.User
	payload.TrustedCall = isTrustedCall(ctx)
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupUser`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doCreateVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateVenueRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doCreateVenue")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Venue()
	var v model.Venue
	if err := s.CreateFromPayload(tx, &v, payload); err != nil {
		httpError(w, `CreateVenue`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateVenue`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doUpdateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateUserRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateUser")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateUser`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateUser`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateUser`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doListRoom(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListRoomRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListRoom`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Room()
	var v model.RoomList
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListRoom`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doLookupRoom(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupRoomRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupRoom")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupRoom`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Room()
	var v model.Room
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupRoom`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doDeleteRoom(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteRoomRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteRoom")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteRoom`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Room()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteRoom`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteRoom`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteVenueRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteVenue")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Venue()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteVenue`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteVenue`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupVenueRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupVenue")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Venue()
	var v model.Venue
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupVenue`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doUpdateVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateVenueRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateVenue`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Venue()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateVenue`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateVenue`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doListVenue(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListVenueRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListVenues`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Venue()
	var v model.VenueList
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListVenues`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doLookupSession(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupSessionRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupSession")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupSession`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Session()
	var v model.Session
	payload.TrustedCall = isTrustedCall(ctx)
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupSession`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

type listSessionsCacheEntry struct {
	Expires time.Time
	List    model.SessionList
}

func doListSessions(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListSessionsRequest) {
	var v model.SessionList

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListSessions`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Session()
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListSessions`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doCreateQuestion(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateQuestionRequest) {
}

func doDeleteQuestion(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteQuestionRequest) {
}

func doListQuestion(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListQuestionRequest) {
}

func doCreateSessionSurveyResponse(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateSessionSurveyResponseRequest) {

}

func doAddFeaturedSpeaker(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddFeaturedSpeakerRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doAddFeaturedSpeaker")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.FeaturedSpeaker()
	var c model.FeaturedSpeaker
	if err := s.CreateFromPayload(tx, payload, &c); err != nil {
		httpError(w, `AddFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateConference`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doDeleteFeaturedSpeaker(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteFeaturedSpeakerRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteFeaturedSpeaker")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.FeaturedSpeaker()
	if err := s.DeleteFromPayload(tx, payload); err != nil {
		httpError(w, `DeleteFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doUpdateFeaturedSpeaker(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateFeaturedSpeakerRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateFeaturedSpeaker")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.FeaturedSpeaker()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupFeaturedSpeaker(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupFeaturedSpeakerRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupFeaturedSpeaker")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.FeaturedSpeaker()
	var c model.FeaturedSpeaker
	if err := s.LookupFromPayload(tx, &c, payload); err != nil {
		httpError(w, `LookupFeaturedSpeaker`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doListFeaturedSpeakers(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListFeaturedSpeakersRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferencesSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.FeaturedSpeaker()
	var l model.FeaturedSpeakerList
	if err := s.ListFromPayload(tx, &l, payload); err != nil {
		httpError(w, `ListFeaturedSpeakers`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, l)
}

func doAddSponsor(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddSponsorRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doAddSponsor")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddSponsor`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Sponsor()
	var c model.Sponsor

	if err := s.CreateFromPayload(ctx, tx, payload, &c); err != nil {
		httpError(w, `Faild to create sponsor from payload`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit transaction`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doDeleteSponsor(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteSponsorRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteSponsor")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteSponsor`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Sponsor()
	if err := s.DeleteFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `DeleteSponsor`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteSponsor`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doUpdateSponsor(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateSponsorRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateSponsor")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateSponsor`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Sponsor()
	updateErr := s.UpdateFromPayload(ctx, tx, payload)
	if !errors.IsIgnorable(updateErr) {
		httpError(w, `Failed to update data from payload`, http.StatusInternalServerError, updateErr)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}

	// This extra bit is for finalizing the image upload
	if cb, ok := errors.IsFinalizationRequired(updateErr); ok {
		if err := cb(); err != nil {
			httpError(w, `Failed to finalize image uploads`, http.StatusInternalServerError, err)
			return
		}
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupSponsor(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupSponsorRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupSponsor")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupSponsor`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Sponsor()
	var c model.Sponsor
	if err := s.LookupFromPayload(tx, &c, payload); err != nil {
		httpError(w, `LookupSponsor`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doListSponsors(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListSponsorsRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListConferencesSeries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Sponsor()
	var l model.SponsorList
	if err := s.ListFromPayload(tx, &l, payload); err != nil {
		httpError(w, `ListSponsors`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, l)
}

func doAddSessionType(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddSessionTypeRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doAddSessionType")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `AddSessionType`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.SessionType()
	var c model.SessionType

	if err := s.CreateFromPayload(ctx, tx, payload, &c); err != nil {
		httpError(w, `Faild to create sponsor from payload`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit transaction`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doDeleteSessionType(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteSessionTypeRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doDeleteSessionType")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteSessionType`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.SessionType()
	if err := s.DeleteFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `DeleteSessionType`, http.StatusInternalServerError, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteSessionType`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doUpdateSessionType(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateSessionTypeRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doUpdateSessionType")
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateSessionType`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.SessionType()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `Failed to update data from payload`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupSessionType(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupSessionTypeRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupSessionType")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupSessionType`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.SessionType()
	var c model.SessionType
	if err := s.LookupFromPayload(tx, &c, payload); err != nil {
		httpError(w, `LookupSessionType`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, c)
}

func doListSessionTypesByConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListSessionTypesByConferenceRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doListSessionTypesbyConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.SessionType()
	var l model.SessionTypeList
	if err := s.ListFromPayload(tx, &l, payload); err != nil {
		httpError(w, `doListSessionTypesbyConference`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, l)
}

func doListConferencesByOrganizer(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListConferencesByOrganizerRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doListConferencesByOrganizer`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	var l model.ConferenceList
	if err := s.ListByOrganizerFromPayload(tx, &l, payload); err != nil {
		httpError(w, `doListConferencesByOrganizer`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, l)
}

func doCreateTemporaryEmail(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateTemporaryEmailRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doCreateTemporaryEmail`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	var res model.CreateTemporaryEmailResponse
	s := service.User()
	if err := s.CreateTemporaryEmailFromPayload(tx, &res.ConfirmationKey, payload); err != nil {
		httpError(w, `doCreateTemporaryEmail`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, res)
}

func doConfirmTemporaryEmail(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ConfirmTemporaryEmailRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doConfirmTemporaryEmail`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.User()
	if err := s.ConfirmTemporaryEmailFromPayload(tx, payload); err != nil {
		httpError(w, `doConfirmTemporaryEmail`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}
}

func doAddConferenceCredential(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.AddConferenceCredentialRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doAddConferenceCredential`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.AddCredentialFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `doAddConferenceCredential`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}
}

func doTweetAsConference(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.TweetAsConferenceRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doTweetAsConference`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.Conference()
	if err := s.TweetFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `doTweetAsConference`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}
}

func doGetConferenceSchedule(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.GetConferenceScheduleRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `GetConferenceSchedule`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	var conf model.Conference
	sc := service.Conference()
	if err := sc.Lookup(tx, &conf, payload.ConferenceID); err != nil {
		httpError(w, `GetConferenceSchedule`, http.StatusInternalServerError, err)
		return
	}

	var series model.ConferenceSeries
	ss := service.ConferenceSeries()
	if err := ss.Lookup(tx, &series, conf.SeriesID); err != nil {
		httpError(w, `GetConferenceSchedule`, http.StatusInternalServerError, err)
		return
	}

	// Create a fake payload
	var lp model.ListSessionsRequest
	lp.ConferenceID.Set(payload.ConferenceID)
	lp.Status = []string{"accepted"}
	lp.Lang.Set("all")

	preferredLang := "en"
	if payload.Lang.Valid() {
		preferredLang = payload.Lang.String
	}
	s := service.Session()
	var v model.SessionList
	if err := s.ListFromPayload(tx, &v, &lp); err != nil {
		httpError(w, `GetConferenceSchedule`, http.StatusInternalServerError, err)
		return
	}

	sts := service.SessionType()
	stm := map[string]model.SessionType{}
	c, _ := ical.New()
	c.AddProperty("x-wr-calname", conf.Title)
	c.AddProperty("x-wr-timezone", conf.Timezone)
	c.AddProperty("calscale", "GREGORIAN")

	c.AddEntry(ical.NewTimezone(conf.Timezone))

	tz, err := time.LoadLocation(conf.Timezone)
	if err != nil {
		tz = time.UTC
	}

	// supported languages
	// first, try the preferred language, then try the other supported
	// languages
	languages := []string{preferredLang, "ja", "en"}
	for _, session := range v {
		e := ical.NewEvent()
		e.AddProperty("url", fmt.Sprintf("https://builderscon.io/%s/%s/session/%s", series.Slug, conf.Slug, session.ID))

		var abstract string
		var abstractLang string
		var title string
		var titleLang string
		for _, lang := range languages {
			if len(abstract) == 0 {
				if lang == "en" {
					if v := strings.TrimSpace(session.Abstract); len(v) > 0 {
						abstract = v
						abstractLang = lang
					}
				} else {
					v, _ := session.LocalizedFields.Get(lang, "abstract")
					v = strings.TrimSpace(v)
					if len(v) > 0 {
						abstract = v
						abstractLang = lang
					}
				}
			}

			if len(title) == 0 {
				if lang == "en" {
					if v := strings.TrimSpace(session.Title); len(v) > 0 {
						title = v
						titleLang = lang
					}
				} else {
					v, _ := session.LocalizedFields.Get(lang, "title")
					v = strings.TrimSpace(v)
					if len(v) > 0 {
						title = v
						titleLang = lang
					}
				}
			}
			if len(title) > 0 && len(abstract) > 0 {
				break
			}
		}
		e.AddProperty("description", abstract, ical.WithParameters(ical.Parameters{
			"language": []string{abstractLang},
		}))
		e.AddProperty("summary", title, ical.WithParameters(ical.Parameters{
			"language": []string{titleLang},
		}))
		if !session.StartsOn.IsZero() {
			e.AddProperty(
				"dtstart",
				session.StartsOn.In(tz).Format("20060102T150405"),
				ical.WithParameters(ical.Parameters{
					"tzid": []string{conf.Timezone},
				}),
			)

			// Grr, this is silly. We should implement this in go-ics
			st, ok := stm[session.SessionTypeID]
			if !ok {
				if err := sts.Lookup(tx, &st, session.SessionTypeID); err == nil {
					ok = true
					stm[session.SessionTypeID] = st
				}
			}

			if ok {
				dur := st.Duration
				var durbuf bytes.Buffer
				durbuf.WriteByte('P')
				if hour := int(math.Floor(float64(dur) / 3600.0)); hour > 0 {
					durbuf.WriteString(strconv.Itoa(hour))
					durbuf.WriteByte('H')
					dur = dur - 3600*hour
				}

				if min := int(math.Floor(float64(dur) / 60.0)); min > 0 {
					durbuf.WriteString(strconv.Itoa(min))
					durbuf.WriteByte('M')
					dur = dur - 60*min
				}

				if dur > 0 {
					durbuf.WriteString(strconv.Itoa(dur))
					durbuf.WriteByte('S')
				}
				e.AddProperty("duration", durbuf.String())
			}
		}
		c.AddEntry(e)
	}

	w.Header().Set("Content-Type", "text/calendar")
	w.WriteHeader(http.StatusOK)
	c.WriteTo(w)
}

func doVerifyUser(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.VerifyUserRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doVerify`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	su := service.User()
	if payload.UserID != payload.ID {
		if err := su.IsAdministrator(tx, payload.UserID); err != nil {
			httpError(w, `doVerify`, http.StatusInternalServerError, err)
			return
		}
	}

	var m model.User
	if err = su.Lookup(tx, &m, payload.ID); err != nil {
		httpError(w, `doVerify`, http.StatusInternalServerError, err)
		return
	}

	go su.Verify(&m)

	httpJSON(w, map[string]interface{}{
		"message": "Verify scheduled",
	})
}

func doSendSelectionResultNotification(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.SendSelectionResultNotificationRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doSendSelectionResultNotification`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	payload.TrustedCall = isTrustedCall(ctx)

	s := service.Session()
	if err := s.SendSelectionResultNotificationFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `doSendSelectionResultNotification`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `Failed to commit data`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, map[string]interface{}{
		"message": "Notification scheduled",
	})
}

func doSendAllSelectionResultNotification(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.SendAllSelectionResultNotificationRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `doSendAllSelectionResultNotification`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	var vdbl db.SessionList
	if err := vdbl.LoadByConference(tx, payload.ConferenceID, "", time.Time{}, time.Time{}, []string{model.StatusAccepted, model.StatusRejected}, nil); err != nil {
		httpError(w, `doSendAllSelectionResultNotification`, http.StatusInternalServerError, err)
		return
	}

	trustedCall := isTrustedCall(ctx)

	// Do this asynchronously
	go func() {
		s := service.Session()
		var req model.SendSelectionResultNotificationRequest
		for _, vdb := range vdbl {
			tx, err := db.Begin()
			if err != nil {
				tx.Rollback()
				return
			}

			req.SessionID = vdb.EID
			req.Force = payload.Force
			req.UserID = payload.UserID
			req.TrustedCall = trustedCall

			if err := s.SendSelectionResultNotificationFromPayload(ctx, tx, &req); err != nil {
				tx.Rollback()
				continue
			}

			if err := tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
	}()

	httpJSON(w, map[string]interface{}{
		"message": "Notification scheduled",
	})
}

func doCreateBlogEntry(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.CreateBlogEntryRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `CreateBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.BlogEntry()
	var m model.BlogEntry
	if err := s.CreateFromPayload(ctx, tx, &m, payload); err != nil {
		httpError(w, `CreateBlogEntry`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `CreateBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, &m)
}

func doUpdateBlogEntry(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.UpdateBlogEntryRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `UpdateBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.BlogEntry()
	if err := s.UpdateFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `UpdateBlogEntry`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `UpdateBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doDeleteBlogEntry(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.DeleteBlogEntryRequest) {
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `DeleteBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.BlogEntry()
	if err := s.DeleteFromPayload(ctx, tx, payload); err != nil {
		httpError(w, `DeleteBlogEntry`, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Commit(); err != nil {
		httpError(w, `DeleteBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	httpJSON(w, map[string]string{"status": "success"})
}

func doLookupBlogEntry(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.LookupBlogEntryRequest) {
	if pdebug.Enabled {
		g := pdebug.Marker("doLookupBlogEntry")
		defer g.End()
	}
	tx, err := db.Begin()
	if err != nil {
		httpError(w, `LookupBlogEntry`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.BlogEntry()
	var v model.BlogEntry
	if err := s.LookupFromPayload(tx, &v, payload); err != nil {
		httpError(w, `LookupBlogEntry`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

func doListBlogEntries(ctx context.Context, w http.ResponseWriter, r *http.Request, payload *model.ListBlogEntriesRequest) {
	var v model.BlogEntryList

	tx, err := db.Begin()
	if err != nil {
		httpError(w, `ListBlogEntries`, http.StatusInternalServerError, err)
		return
	}
	defer tx.AutoRollback()

	s := service.BlogEntry()
	if err := s.ListFromPayload(tx, &v, payload); err != nil {
		httpError(w, `ListBlogEntries`, http.StatusInternalServerError, err)
		return
	}

	httpJSON(w, v)
}

