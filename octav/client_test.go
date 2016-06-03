package octav_test

import (
	"fmt"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"github.com/builderscon/octav/octav"
	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/client"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	"github.com/builderscon/octav/octav/validator"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type TestCtx struct {
	*testing.T
	APIClient *db.Client
	HTTPClient *client.Client
}

func NewTestCtx(t *testing.T) (*TestCtx, error) {
	vdb := db.Client{
		EID: tools.UUID(),
		Secret: tools.UUID(), // Todo
		Name: "Test Client",
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "failed to start DB transaction")
	}
	defer tx.AutoRollback()

	if vdb.Create(tx); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "failed to commit changes to DB")
	}

	ctx := &TestCtx{
		T: t,
		APIClient: &vdb,
	}

	return ctx, nil
}

func (ctx *TestCtx) SetAPIServer(ts *httptest.Server) {
	ctx.HTTPClient = client.New(ts.URL)
	ctx.HTTPClient.BasicAuth.Username = ctx.APIClient.EID
	ctx.HTTPClient.BasicAuth.Password = ctx.APIClient.Secret
}

func bigsight() *model.CreateVenueRequest {
	lf := tools.LocalizedFields{}
	lf.Set("ja", "name", `東京ビッグサイト`)
	lf.Set("ja", "address", `〒135-0063 東京都江東区有明３丁目１０−１`)

	r := model.CreateVenueRequest{}
	r.L10N = lf
	r.Name.Set("Tokyo Bigsight")
	r.Address.Set("Ariake 3-10-1, Koto-ku, Tokyo")
	r.Longitude.Set(35.6320326)
	r.Latitude.Set(139.7976891)

	return &r
}

func intlConferenceRoom(venueID string) *model.CreateRoomRequest {
	lf := tools.LocalizedFields{}
	lf.Set("ja", "name", `国際会議場`)

	r := model.CreateRoomRequest{}
	r.L10N = lf
	r.Capacity.Set(1000)
	r.Name.Set("International Conference Hall")
	r.VenueID.Set(venueID)

	return &r
}

func testCreateRoom(ctx *TestCtx, r *model.CreateRoomRequest) (*model.Room, error) {
	res, err := ctx.HTTPClient.CreateRoom(r)
	if !assert.NoError(ctx.T, err, "CreateRoom should succeed") {
		return nil, err
	}
	return res, nil
}

func testCreateVenue(ctx *TestCtx, v *model.CreateVenueRequest) (*model.Venue, error) {
	res, err := ctx.HTTPClient.CreateVenue(v)
	if !assert.NoError(ctx.T, err, "CreateVenue should succeed") {
		return nil, err
	}
	return res, nil
}

func testLookupRoom(ctx *TestCtx, id, lang string) (*model.Room, error) {
	r := &model.LookupRoomRequest{ID: id}
	if lang != "" {
		r.Lang.Set(lang)
	}
	venue, err := ctx.HTTPClient.LookupRoom(r)
	if !assert.NoError(ctx.T, err, "LookupRoom succeeds") {
		return nil, err
	}
	return venue, nil
}

func testCreateUser(ctx *TestCtx, in *model.CreateUserRequest) (*model.User, error) {
	res, err := ctx.HTTPClient.CreateUser(in)
	if !assert.NoError(ctx.T, err, "CreateUser should succeed") {
		return nil, err
	}
	return res, nil
}

func testLookupUser(ctx *TestCtx, id, lang string) (*model.User, error) {
	r := &model.LookupUserRequest{ID: id}
	if lang != "" {
		r.Lang.Set(lang)
	}
	user, err := ctx.HTTPClient.LookupUser(r)
	if !assert.NoError(ctx.T, err, "LookupUser succeeds") {
		return nil, err
	}
	return user, nil
}

func testDeleteUser(ctx *TestCtx, id string) error {
	err := ctx.HTTPClient.DeleteUser(&model.DeleteUserRequest{ID: id})
	if !assert.NoError(ctx.T, err, "DeleteUser should succeed") {
		return err
	}
	return nil
}

func testLookupVenue(ctx *TestCtx, id, lang string) (*model.Venue, error) {
	r := &model.LookupVenueRequest{ID: id}
	if lang != "" {
		r.Lang.Set(lang)
	}
	venue, err := ctx.HTTPClient.LookupVenue(r)
	if !assert.NoError(ctx.T, err, "LookupVenue succeeds") {
		return nil, err
	}
	return venue, nil
}

func testUpdateRoom(ctx *TestCtx, in *model.UpdateRoomRequest) error {
	err := ctx.HTTPClient.UpdateRoom(in)
	if !assert.NoError(ctx.T, err, "UpdateRoom succeeds") {
		return err
	}
	return nil
}

func testDeleteRoom(ctx *TestCtx, id string) error {
	err := ctx.HTTPClient.DeleteRoom(&model.DeleteRoomRequest{ID: id})
	if !assert.NoError(ctx.T, err, "DeleteRoom should be successful") {
		return err
	}
	return err
}

func testUpdateVenue(ctx *TestCtx, in *model.UpdateVenueRequest) error {
	err := ctx.HTTPClient.UpdateVenue(in)
	if !assert.NoError(ctx.T, err, "UpdateVenue succeeds") {
		return err
	}
	return nil
}

func testDeleteVenue(ctx *TestCtx, id string) error {
	err := ctx.HTTPClient.DeleteVenue(&model.DeleteVenueRequest{ID: id})
	if !assert.NoError(ctx.T, err, "DeleteVenue should be successful") {
		return err
	}
	return err
}

func yapcasia(userID string) *model.CreateConferenceRequest {
	return &model.CreateConferenceRequest{
		Title:  "YAPC::Asia Tokyo",
		Slug:   "yapcasia",
		UserID: userID,
	}
}

func testCreateConference(ctx *TestCtx, in *model.CreateConferenceRequest) (*model.Conference, error) {
	res, err := ctx.HTTPClient.CreateConference(in)
	if !assert.NoError(ctx.T, err, "CreateConference should succeed") {
		return nil, err
	}
	return res, nil
}

func testLookupConference(ctx *TestCtx, id, lang string) (*model.Conference, error) {
	r := &model.LookupConferenceRequest{ID: id}
	if lang != "" {
		r.Lang.Set(lang)
	}
	conference, err := ctx.HTTPClient.LookupConference(r)
	if !assert.NoError(ctx.T, err, "LookupConference succeeds") {
		return nil, err
	}
	return conference, nil
}

func testUpdateConference(ctx *TestCtx, in *model.UpdateConferenceRequest) error {
	err := ctx.HTTPClient.UpdateConference(in)
	if !assert.NoError(ctx.T, err, "UpdateConference succeeds") {
		return err
	}
	return nil
}

func testDeleteConference(ctx *TestCtx, id string) error {
	err := ctx.HTTPClient.DeleteConference(&model.DeleteConferenceRequest{ID: id})
	if !assert.NoError(ctx.T, err, "DeleteConference should be successful") {
		return err
	}
	return err
}

func TestConferenceCRUD(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)

	res, err := testCreateConference(ctx, yapcasia(user.ID))
	if err != nil {
		return
	}
	defer testDeleteConference(ctx, res.ID)

	if !assert.NoError(ctx.T, validator.HTTPCreateConferenceResponse.Validate(res), "Validation should succeed") {
		return
	}

	res2, err := testLookupConference(ctx, res.ID, "")
	if err != nil {
		return
	}

	// The result from LookupConference contains the administrator field
	// Remove that (and make sure it's populated), then do the comparison
	admins := res2.Administrators
	res2.Administrators = model.UserList(nil)
	if !assert.Equal(ctx.T, res2, res, "LookupConference is the same as the conference created") {
		return
	}

	if !assert.Len(ctx.T, admins, 1, "There should be 1 administrator") {
		return
	}

	in := model.UpdateConferenceRequest{ID: res.ID}
	in.SubTitle.Set("Big Bang!")
	in.L10N.Set("ja", "title", "ヤップシー エイジア")
	if err := testUpdateConference(ctx, &in); err != nil {
		return
	}

	res3, err := testLookupConference(ctx, res.ID, "ja")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, res3.SubTitle, "Big Bang!", "Conference.SubTitle is the same as the conference updated") {
		return
	}

	if !assert.Equal(ctx.T, "ヤップシー エイジア", res3.Title, "Conference.title#ja is the same as the conference updated") {
		return
	}
}

func TestRoomCRUD(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	venue, err := testCreateVenue(ctx, bigsight())
	if err != nil {
		return
	}

	res, err := testCreateRoom(ctx, intlConferenceRoom(venue.ID))
	if err != nil {
		return
	}

	if !assert.NotEmpty(ctx.T, res.ID, "Returned structure has ID") {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPCreateRoomResponse.Validate(res), "Validation should succeed") {
		return
	}

	res2, err := testLookupRoom(ctx, res.ID, "")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, res2, res, "LookupRoom is the same as the room created") {
		return
	}

	in := model.UpdateRoomRequest{ID: res.ID}
	in.L10N.Set("ja", "name", "国際会議場")
	if err := testUpdateRoom(ctx, &in); err != nil {
		return
	}

	res3, err := testLookupRoom(ctx, res.ID, "ja")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, "国際会議場", res3.Name, "Room.name#ja is the same as the conference updated") {
		return
	}

	if err := testDeleteRoom(ctx, res.ID); err != nil {
		return
	}

	if err := testDeleteVenue(ctx, venue.ID); err != nil {
		return
	}
}

func testCreateSession(ctx *TestCtx, in *model.CreateSessionRequest) (*model.Session, error) {
	res, err := ctx.HTTPClient.CreateSession(in)
	if !assert.NoError(ctx.T, err, "CreateSession should succeed") {
		return nil, err
	}
	return res, nil
}

func testLookupSession(ctx *TestCtx, id, lang string) (*model.Session, error) {
	r := &model.LookupSessionRequest{ID: id}
	if lang != "" {
		r.Lang.Set(lang)
	}
	v, err := ctx.HTTPClient.LookupSession(r)
	if !assert.NoError(ctx.T, err, "LookupSession succeeds") {
		return nil, err
	}
	return v, nil
}

func testUpdateSession(ctx *TestCtx, in *model.UpdateSessionRequest) error {
	err := ctx.HTTPClient.UpdateSession(in)
	if !assert.NoError(ctx.T, err, "UpdateSession succeeds") {
		return err
	}
	return nil
}

func testDeleteSession(ctx *TestCtx, id string) error {
	err := ctx.HTTPClient.DeleteSession(&model.DeleteSessionRequest{ID: id})
	if !assert.NoError(ctx.T, err, "DeleteSession should be successful") {
		return err
	}
	return err
}

func bconsession(cid, uid string) *model.CreateSessionRequest {
	in := model.CreateSessionRequest{}
	in.ConferenceID.Set(cid)
	in.SpeakerID.Set(uid)
	in.Title.Set("How To Write A Conference Backend")
	in.Duration.Set(60)
	in.Abstract.Set("Use lots of reflection and generate lots of code")
	return &in
}

func TestSessionCRUD(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)

	conference, err := testCreateConference(ctx, yapcasia(user.ID))
	if err != nil {
		return
	}

	res, err := testCreateSession(ctx, bconsession(conference.ID, user.ID))
	if err != nil {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPCreateSessionResponse.Validate(res), "Validation should succeed") {
		return
	}

	res2, err := testLookupSession(ctx, res.ID, "")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, res2, res, "LookupSession is the same as the room created") {
		return
	}

	in := model.UpdateSessionRequest{ID: res.ID}
	in.L10N.Set("ja", "title", "カンファレンス用ソフトウェアの作り方")
	if err := testUpdateSession(ctx, &in); err != nil {
		return
	}

	res3, err := testLookupSession(ctx, res.ID, "ja")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, "カンファレンス用ソフトウェアの作り方", res3.Title, "Session.title#ja is the same as the conference updated") {
		return
	}

	if err := testDeleteSession(ctx, res.ID); err != nil {
		return
	}

	if err := testDeleteConference(ctx, conference.ID); err != nil {
		return
	}
}

var ghidL = sync.Mutex{}
var ghid = 0

func johndoe() *model.CreateUserRequest {
	ghidL.Lock()
	defer ghidL.Unlock()

	ghid++
	lf := tools.LocalizedFields{}
	lf.Set("ja", "first_name", "ジョン")
	lf.Set("ja", "last_name", "ドー")

	r := model.CreateUserRequest{
		Nickname: tools.UUID(),
	}
	r.AuthVia.Set("github")
	r.AuthUserID.Set(strconv.Itoa(ghid))
	r.FirstName.Set("John")
	r.LastName.Set("Doe")
	r.Email.Set("john.doe@example.com")
	r.TshirtSize.Set("XL")
	r.L10N = lf
	return &r
}

func TestCreateUser(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	res, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}

	if !assert.NotEmpty(ctx.T, res.ID, "Returned structure has ID") {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPCreateUserResponse.Validate(res), "Validation should succeed") {
		return
	}

	res2, err := testLookupUser(ctx, res.ID, "")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, res2, res, "LookupUser is the same as the user created") {
		return
	}

	res3, err := testLookupUser(ctx, res.ID, "ja")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, "ジョン", res3.FirstName, "User.first_name#ja is localized") {
		return
	}

	if !assert.Equal(ctx.T, "ドー", res3.LastName, "User.last_name#ja is localized") {
		return
	}

	if err := testDeleteUser(ctx, res.ID); err != nil {
		return
	}
}

func TestVenueCRUD(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	res, err := testCreateVenue(ctx, bigsight())
	if err != nil {
		return
	}

	if !assert.NotEmpty(ctx.T, res.ID, "Returned structure has ID") {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPCreateVenueResponse.Validate(res), "Validation should succeed") {
		return
	}

	res2, err := testLookupVenue(ctx, res.ID, "")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, res2, res, "LookupVenue is the same as the venue created") {
		return
	}

	in := model.UpdateVenueRequest{ID: res.ID}
	in.L10N.Set("ja", "name", "東京ビッグサイト")
	if err := testUpdateVenue(ctx, &in); err != nil {
		return
	}

	res3, err := testLookupVenue(ctx, res.ID, "ja")
	if err != nil {
		return
	}

	if !assert.Equal(ctx.T, "東京ビッグサイト", res3.Name, "Venue.name#ja is the same as the object updated") {
		return
	}

	if err := testDeleteVenue(ctx, res.ID); err != nil {
		return
	}
}

func TestDeleteConferenceDates(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)

	conf, err := testCreateConference(ctx, &model.CreateConferenceRequest{
		UserID: user.ID,
	})
	if err != nil {
		return
	}
	defer testDeleteConference(ctx, conf.ID)

	err = ctx.HTTPClient.AddConferenceDates(&model.AddConferenceDatesRequest{
		ConferenceID: conf.ID,
		Dates: []model.ConferenceDate{
			model.ConferenceDate{
				Date:  model.NewDate(2016, 3, 22),
				Open:  model.NewWallClock(10, 0),
				Close: model.NewWallClock(19, 0),
			},
		},
	})
	if !assert.NoError(ctx.T, err, "AddConferenceDates works") {
		return
	}

	err = ctx.HTTPClient.DeleteConferenceDates(&model.DeleteConferenceDatesRequest{
		ConferenceID: conf.ID,
		Dates:        []model.Date{model.NewDate(2016, 3, 22)},
	})
	if !assert.NoError(ctx.T, err, "DeleteConferenceDates works") {
		return
	}

	conf2, err := testLookupConference(ctx, conf.ID, "")
	if err != nil {
		return
	}

	if !assert.Len(ctx.T, conf2.Dates, 0, "There should be no dates set") {
		return
	}
}

func TestConferenceAdmins(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)

	conf, err := testCreateConference(ctx, &model.CreateConferenceRequest{
		UserID: user.ID,
	})
	if err != nil {
		return
	}
	defer testDeleteConference(ctx, conf.ID)

	for i := 0; i < 9; i++ {
		extraAdmin, err := testCreateUser(ctx, johndoe())
		if err != nil {
			return
		}
		defer testDeleteUser(ctx, extraAdmin.ID)

		err = ctx.HTTPClient.AddConferenceAdmin(&model.AddConferenceAdminRequest{
			ConferenceID: conf.ID,
			UserID:       extraAdmin.ID,
		})
		if !assert.NoError(ctx.T, err, "AddConferenceAdmin should succeed") {
			return
		}
	}

	conf2, err := testLookupConference(ctx, conf.ID, "")
	if err != nil {
		return
	}

	if !assert.Len(ctx.T, conf2.Administrators, 10, "There should be 10 admins") {
		return
	}

	for _, admin := range conf2.Administrators {
		err = ctx.HTTPClient.DeleteConferenceAdmin(&model.DeleteConferenceAdminRequest{
			ConferenceID: conf.ID,
			UserID:       admin.ID,
		})
		if !assert.NoError(ctx.T, err, "DeleteConferenceAdmin should succeed") {
			return
		}
	}

	conf3, err := testLookupConference(ctx, conf.ID, "")
	if err != nil {
		return
	}

	if !assert.Len(ctx.T, conf3.Administrators, 0, "There should be 0 admins") {
		return
	}

}

func TestListConference(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)

	for i := 0; i < 10; i++ {
		conf, err := testCreateConference(ctx, &model.CreateConferenceRequest{
			UserID: user.ID,
		})
		if err != nil {
			return
		}

		err = ctx.HTTPClient.AddConferenceDates(&model.AddConferenceDatesRequest{
			ConferenceID: conf.ID,
			Dates: []model.ConferenceDate{
				model.ConferenceDate{
					Date:  model.NewDate(2016, 3, 22),
					Open:  model.NewWallClock(10, 0),
					Close: model.NewWallClock(19, 0),
				},
			},
		})
		if !assert.NoError(ctx.T, err, "AddConferenceDates works") {
			return
		}

		defer testDeleteConference(ctx, conf.ID)
	}

	in := model.ListConferenceRequest{}
	in.RangeStart.Set("2016-03-21")
	in.RangeEnd.Set("2016-03-23")
	res, err := ctx.HTTPClient.ListConference(&in)
	if !assert.NoError(ctx.T, err, "ListConference should succeed") {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPListConferenceResponse.Validate(res), "Validation should succeed") {
		return
	}

	if !assert.Len(ctx.T, res, 10, "ListConference returns 10 rooms") {
		return
	}
}

func TestListRoom(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	venue, err := testCreateVenue(ctx, bigsight())
	if err != nil {
		return
	}

	_, err = testCreateRoom(ctx, intlConferenceRoom(venue.ID))
	if err != nil {
		return
	}

	in := model.ListRoomRequest{
		VenueID: venue.ID,
	}
	res, err := ctx.HTTPClient.ListRoom(&in)
	if !assert.NoError(ctx.T, err, "ListRoom should succeed") {
		return
	}

	if !assert.NoError(ctx.T, validator.HTTPListRoomResponse.Validate(res), "Validation should succeed") {
		return
	}

	if !assert.Len(ctx.T, res, 1, "ListRoom returns 1 rooms") {
		return
	}
}

func TestListSessionByConference(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	user, err := testCreateUser(ctx, johndoe())
	if err != nil {
		return
	}
	defer testDeleteUser(ctx, user.ID)
	conference, err := testCreateConference(ctx, yapcasia(user.ID))
	if err != nil {
		return
	}

	for i := 0; i < 10; i++ {
		sin := model.CreateSessionRequest{}
		sin.ConferenceID.Set(conference.ID)
		sin.SpeakerID.Set(user.ID)
		sin.Title.Set(fmt.Sprintf("Title %d", i))
		sin.Duration.Set(60)
		sin.Abstract.Set("Use lots of reflection and generate lots of code")
		_, err := testCreateSession(ctx, &sin)
		if err != nil {
			return
		}
	}

	in := model.ListSessionByConferenceRequest{
		ConferenceID: conference.ID,
	}
	res, err := ctx.HTTPClient.ListSessionByConference(&in)
	if !assert.NoError(ctx.T, err, "ListSessionByConference should succeed") {
		return
	}
	if !assert.NoError(ctx.T, validator.HTTPListSessionByConferenceResponse.Validate(res), "Validation should succeed") {
		return
	}

	if !assert.Len(ctx.T, res, 10, "There should be 10 sessions") {
		return
	}
}

func TestListVenue(t *testing.T) {
	ctx, err := NewTestCtx(t)
	if !assert.NoError(ctx, err, "failed to create test ctx") {
		return
	}

	ts := httptest.NewServer(octav.New())
	defer ts.Close()

	ctx.SetAPIServer(ts)

	in := model.ListVenueRequest{}
	res, err := ctx.HTTPClient.ListVenue(&in)
	if !assert.NoError(ctx.T, err, "ListVenue should succeed") {
		return
	}
	if !assert.NoError(ctx.T, validator.HTTPListVenueResponse.Validate(res), "Validation should succeed") {
		return
	}
}
