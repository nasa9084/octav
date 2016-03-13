// DO NOT EDIT. Automatically generated by hsup
package validator

import (
	"github.com/lestrrat/go-jsval"
)

var HTTPCreateConferenceRequest *jsval.JSVal
var HTTPCreateConferenceResponse *jsval.JSVal
var HTTPCreateRoomRequest *jsval.JSVal
var HTTPCreateRoomResponse *jsval.JSVal
var HTTPCreateSessionRequest *jsval.JSVal
var HTTPCreateSessionResponse *jsval.JSVal
var HTTPCreateUserRequest *jsval.JSVal
var HTTPCreateUserResponse *jsval.JSVal
var HTTPCreateVenueRequest *jsval.JSVal
var HTTPCreateVenueResponse *jsval.JSVal
var HTTPDeleteConferenceRequest *jsval.JSVal
var HTTPDeleteRoomRequest *jsval.JSVal
var HTTPDeleteUserRequest *jsval.JSVal
var HTTPDeleteVenueRequest *jsval.JSVal
var HTTPListRoomsRequest *jsval.JSVal
var HTTPListRoomsResponse *jsval.JSVal
var HTTPListSessionsByConferenceRequest *jsval.JSVal
var HTTPListSessionsByConferenceResponse *jsval.JSVal
var HTTPListVenuesRequest *jsval.JSVal
var HTTPListVenuesResponse *jsval.JSVal
var HTTPLookupConferenceRequest *jsval.JSVal
var HTTPLookupConferenceResponse *jsval.JSVal
var HTTPLookupRoomRequest *jsval.JSVal
var HTTPLookupRoomResponse *jsval.JSVal
var HTTPLookupSessionRequest *jsval.JSVal
var HTTPLookupSessionResponse *jsval.JSVal
var HTTPLookupUserRequest *jsval.JSVal
var HTTPLookupUserResponse *jsval.JSVal
var HTTPLookupVenueRequest *jsval.JSVal
var HTTPLookupVenueResponse *jsval.JSVal
var HTTPUpdateConferenceRequest *jsval.JSVal
var HTTPUpdateSessionRequest *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint
var R4 jsval.Constraint
var R5 jsval.Constraint
var R6 jsval.Constraint
var R7 jsval.Constraint
var R8 jsval.Constraint
var R9 jsval.Constraint
var R10 jsval.Constraint
var R11 jsval.Constraint
var R12 jsval.Constraint
var R13 jsval.Constraint
var R14 jsval.Constraint
var R15 jsval.Constraint
var R16 jsval.Constraint
var R17 jsval.Constraint
var R18 jsval.Constraint
var R19 jsval.Constraint
var R20 jsval.Constraint
var R21 jsval.Constraint
var R22 jsval.Constraint
var R23 jsval.Constraint
var R24 jsval.Constraint
var R25 jsval.Constraint
var R26 jsval.Constraint
var R27 jsval.Constraint
var R28 jsval.Constraint
var R29 jsval.Constraint
var R30 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().Enum("pending", "accepted", "rejected").Default("pending")
	R1 = jsval.String().Enum("allow", "disallow").Default("allow")
	R2 = jsval.Boolean().Default(false)
	R3 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			"dates",
			jsval.OneOf().
				Add(
					jsval.NullConstraint,
				).
				Add(
					jsval.Reference(M).RefersTo("#/definitions/date_array"),
				),
		).
		AddProp(
			"description",
			jsval.Reference(M).RefersTo("#/definitions/string_en"),
		).
		AddProp(
			"id",
			jsval.Reference(M).RefersTo("#/definitions/uuid"),
		).
		AddProp(
			"name",
			jsval.Reference(M).RefersTo("#/definitions/string_en"),
		).
		AddProp(
			"slug",
			jsval.Reference(M).RefersTo("#/definitions/string_en"),
		).
		AddProp(
			"venue",
			jsval.Reference(M).RefersTo("#/definitions/venue"),
		).
		PatternPropertiesString(
			"description#[a-z]+",
			jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
		).
		PatternPropertiesString(
			"title#[a-z]+",
			jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
		)
	R4 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		)
	R5 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo("#/definitions/date"),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R6 = jsval.String().RegexpString("^\\d+-\\d+-\\d+(,\\d+-\\d+-\\d+-)*$")
	R7 = jsval.OneOf().
		Add(
			jsval.String().Format("date-time"),
		).
		Add(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				),
		)
	R8 = jsval.Integer()
	R9 = jsval.String().Format("email")
	R10 = jsval.String().Default("en")
	R11 = jsval.Number()
	R12 = jsval.Number()
	R13 = jsval.String()
	R14 = jsval.String()
	R15 = jsval.String().Enum("beginner", "intermediate", "advanced").Default("beginner")
	R16 = jsval.Integer().Minimum(0)
	R17 = jsval.Integer().Minimum(0).Default(10)
	R18 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			"capcity",
			jsval.Reference(M).RefersTo("#/definitions/positiveInteger"),
		).
		AddProp(
			"id",
			jsval.Reference(M).RefersTo("#/definitions/uuid"),
		).
		AddProp(
			"name",
			jsval.String(),
		).
		AddProp(
			"venue_id",
			jsval.Reference(M).RefersTo("#/definitions/uuid"),
		).
		PatternPropertiesString(
			"name#[a-z]+",
			jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
		)
	R19 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			"abstract",
			jsval.String(),
		).
		AddProp(
			"category",
			jsval.String(),
		).
		AddProp(
			"conference",
			jsval.OneOf().
				Add(
					jsval.NullConstraint,
				).
				Add(
					jsval.Reference(M).RefersTo("#/definitions/conference"),
				),
		).
		AddProp(
			"confirmed",
			jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
		).
		AddProp(
			"duration",
			jsval.Reference(M).RefersTo("#/definitions/duration"),
		).
		AddProp(
			"has_interpretation",
			jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
		).
		AddProp(
			"material_level",
			jsval.Reference(M).RefersTo("#/definitions/material_level"),
		).
		AddProp(
			"memo",
			jsval.String(),
		).
		AddProp(
			"photo_permission",
			jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
		).
		AddProp(
			"room",
			jsval.OneOf().
				Add(
					jsval.NullConstraint,
				).
				Add(
					jsval.Reference(M).RefersTo("#/definitions/room"),
				),
		).
		AddProp(
			"slide_language",
			jsval.Reference(M).RefersTo("#/definitions/language"),
		).
		AddProp(
			"slide_subtitles",
			jsval.Reference(M).RefersTo("#/definitions/language"),
		).
		AddProp(
			"slide_url",
			jsval.Reference(M).RefersTo("#/definitions/url"),
		).
		AddProp(
			"speaker",
			jsval.OneOf().
				Add(
					jsval.NullConstraint,
				).
				Add(
					jsval.Object().
						AdditionalProperties(
							jsval.EmptyConstraint,
						),
				).
				Add(
					jsval.Reference(M).RefersTo("#/definitions/speaker_array"),
				),
		).
		AddProp(
			"spoken_language",
			jsval.Reference(M).RefersTo("#/definitions/language"),
		).
		AddProp(
			"starts_on",
			jsval.Reference(M).RefersTo("#/definitions/datetime"),
		).
		AddProp(
			"status",
			jsval.Reference(M).RefersTo("#/definitions/acceptance_status"),
		).
		AddProp(
			"tags",
			jsval.OneOf().
				Add(
					jsval.String(),
				).
				Add(
					jsval.Reference(M).RefersTo("#/definitions/tag_array"),
				),
		).
		AddProp(
			"title",
			jsval.String(),
		).
		AddProp(
			"video_permission",
			jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
		).
		AddProp(
			"video_url",
			jsval.Reference(M).RefersTo("#/definitions/url"),
		)
	R20 = jsval.Object().
		AdditionalProperties(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				PatternPropertiesString(
					"^[a-z0-9-]+-profile",
					jsval.Object().
						AdditionalProperties(
							jsval.EmptyConstraint,
						),
				),
		).
		AddProp(
			"email",
			jsval.Reference(M).RefersTo("#/definitions/email"),
		).
		AddProp(
			"id",
			jsval.Reference(M).RefersTo("#/definitions/uuid"),
		).
		AddProp(
			"name",
			jsval.String(),
		)
	R21 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo("#/definitions/speaker"),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R22 = jsval.String()
	R23 = jsval.String()
	R24 = jsval.String()
	R25 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo("#/definitions/tag"),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R26 = jsval.String().Enum("XXXL", "XXL", "XL", "L", "M", "S", "XS")
	R27 = jsval.String().Format("url")
	R28 = jsval.String().RegexpString("^[a-fA-F0-9-]+$")
	R29 = jsval.String().RegexpString("^[a-fA-F0-9-]+$").Default("")
	R30 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			"id",
			jsval.Reference(M).RefersTo("#/definitions/uuid"),
		).
		AddProp(
			"name",
			jsval.String(),
		).
		AddProp(
			"rooms",
			jsval.Reference(M).RefersTo("#/definitions/room"),
		)
	M.SetReference("#/definitions/acceptance_status", R0)
	M.SetReference("#/definitions/binary_permission_default_allow", R1)
	M.SetReference("#/definitions/boolean_default_false", R2)
	M.SetReference("#/definitions/conference", R3)
	M.SetReference("#/definitions/date", R4)
	M.SetReference("#/definitions/date_array", R5)
	M.SetReference("#/definitions/datestr", R6)
	M.SetReference("#/definitions/datetime", R7)
	M.SetReference("#/definitions/duration", R8)
	M.SetReference("#/definitions/email", R9)
	M.SetReference("#/definitions/language", R10)
	M.SetReference("#/definitions/latitude", R11)
	M.SetReference("#/definitions/longitude", R12)
	M.SetReference("#/definitions/markdown_en", R13)
	M.SetReference("#/definitions/markdown_i18n", R14)
	M.SetReference("#/definitions/material_level", R15)
	M.SetReference("#/definitions/positiveInteger", R16)
	M.SetReference("#/definitions/positiveIntegerDefault10", R17)
	M.SetReference("#/definitions/room", R18)
	M.SetReference("#/definitions/session", R19)
	M.SetReference("#/definitions/speaker", R20)
	M.SetReference("#/definitions/speaker_array", R21)
	M.SetReference("#/definitions/string_en", R22)
	M.SetReference("#/definitions/string_i18n", R23)
	M.SetReference("#/definitions/tag", R24)
	M.SetReference("#/definitions/tag_array", R25)
	M.SetReference("#/definitions/tshirt_size", R26)
	M.SetReference("#/definitions/url", R27)
	M.SetReference("#/definitions/uuid", R28)
	M.SetReference("#/definitions/uuidDefaultEmpty", R29)
	M.SetReference("#/definitions/venue", R30)
	HTTPCreateConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("slug", "title").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"description",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"slug",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"starts_on",
					jsval.Reference(M).RefersTo("#/definitions/datetime"),
				).
				AddProp(
					"sub_title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"dates",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/date_array"),
						),
				).
				AddProp(
					"description",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"slug",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"venue",
					jsval.Reference(M).RefersTo("#/definitions/venue"),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("name", "venue_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"capacity",
					jsval.Reference(M).RefersTo("#/definitions/positiveInteger"),
				).
				AddProp(
					"name",
					jsval.String(),
				).
				AddProp(
					"venue_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateRoomResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"capcity",
					jsval.Reference(M).RefersTo("#/definitions/positiveInteger"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.String(),
				).
				AddProp(
					"venue_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateSessionRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("abstract", "conference_id", "duration", "speaker_id", "title").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"abstract",
					jsval.Reference(M).RefersTo("#/definitions/markdown_en"),
				).
				AddProp(
					"category",
					jsval.String(),
				).
				AddProp(
					"conference_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"duration",
					jsval.Reference(M).RefersTo("#/definitions/duration"),
				).
				AddProp(
					"material_level",
					jsval.Reference(M).RefersTo("#/definitions/material_level"),
				).
				AddProp(
					"memo",
					jsval.String(),
				).
				AddProp(
					"photo_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"slide_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_subtitles",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				AddProp(
					"speaker_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"spoken_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"tags",
					jsval.String(),
				).
				AddProp(
					"title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"video_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"video_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				PatternPropertiesString(
					"abstract#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/markdown_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateSessionResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"abstract",
					jsval.String(),
				).
				AddProp(
					"category",
					jsval.String(),
				).
				AddProp(
					"conference",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/conference"),
						),
				).
				AddProp(
					"confirmed",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"duration",
					jsval.Reference(M).RefersTo("#/definitions/duration"),
				).
				AddProp(
					"has_interpretation",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"material_level",
					jsval.Reference(M).RefersTo("#/definitions/material_level"),
				).
				AddProp(
					"memo",
					jsval.String(),
				).
				AddProp(
					"photo_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"room",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/room"),
						),
				).
				AddProp(
					"slide_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_subtitles",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				AddProp(
					"speaker",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Object().
								AdditionalProperties(
									jsval.EmptyConstraint,
								),
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/speaker_array"),
						),
				).
				AddProp(
					"spoken_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"starts_on",
					jsval.Reference(M).RefersTo("#/definitions/datetime"),
				).
				AddProp(
					"status",
					jsval.Reference(M).RefersTo("#/definitions/acceptance_status"),
				).
				AddProp(
					"tags",
					jsval.OneOf().
						Add(
							jsval.String(),
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/tag_array"),
						),
				).
				AddProp(
					"title",
					jsval.String(),
				).
				AddProp(
					"video_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"video_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				),
		)

	HTTPCreateUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("email", "first_name", "last_name", "nickname", "tshirt_size").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/email"),
				).
				AddProp(
					"first_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"last_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"nickname",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"tshirt_size",
					jsval.Reference(M).RefersTo("#/definitions/tshirt_size"),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateUserResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/email"),
				).
				AddProp(
					"first_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"last_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"nickname",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"tshirt_size",
					jsval.Reference(M).RefersTo("#/definitions/tshirt_size"),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("address", "name").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"address",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"latitude",
					jsval.Reference(M).RefersTo("#/definitions/latitude"),
				).
				AddProp(
					"longitude",
					jsval.Reference(M).RefersTo("#/definitions/longitude"),
				).
				AddProp(
					"name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				PatternPropertiesString(
					"address#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPCreateVenueResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.String(),
				).
				AddProp(
					"rooms",
					jsval.Reference(M).RefersTo("#/definitions/room"),
				),
		)

	HTTPDeleteConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPDeleteRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPDeleteUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPDeleteVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPListRoomsRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("venue_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"lang",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"limit",
					jsval.Reference(M).RefersTo("#/definitions/positiveIntegerDefault10"),
				).
				AddProp(
					"venue_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPListRoomsResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo("#/definitions/room"),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPListSessionsByConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("conference_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"conference_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"date",
					jsval.Reference(M).RefersTo("#/definitions/datestr"),
				),
		)

	HTTPListSessionsByConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo("#/definitions/session"),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPListVenuesRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"lang",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"since",
					jsval.Reference(M).RefersTo("#/definitions/uuidDefaultEmpty"),
				),
		)

	HTTPListVenuesResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo("#/definitions/venue"),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPLookupConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPLookupConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"dates",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/date_array"),
						),
				).
				AddProp(
					"description",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"slug",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"venue",
					jsval.Reference(M).RefersTo("#/definitions/venue"),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPLookupRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPLookupRoomResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"capcity",
					jsval.Reference(M).RefersTo("#/definitions/positiveInteger"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.String(),
				).
				AddProp(
					"venue_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPLookupSessionRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPLookupSessionResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"abstract",
					jsval.String(),
				).
				AddProp(
					"category",
					jsval.String(),
				).
				AddProp(
					"conference",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/conference"),
						),
				).
				AddProp(
					"confirmed",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"duration",
					jsval.Reference(M).RefersTo("#/definitions/duration"),
				).
				AddProp(
					"has_interpretation",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"material_level",
					jsval.Reference(M).RefersTo("#/definitions/material_level"),
				).
				AddProp(
					"memo",
					jsval.String(),
				).
				AddProp(
					"photo_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"room",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/room"),
						),
				).
				AddProp(
					"slide_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_subtitles",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				AddProp(
					"speaker",
					jsval.OneOf().
						Add(
							jsval.NullConstraint,
						).
						Add(
							jsval.Object().
								AdditionalProperties(
									jsval.EmptyConstraint,
								),
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/speaker_array"),
						),
				).
				AddProp(
					"spoken_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"starts_on",
					jsval.Reference(M).RefersTo("#/definitions/datetime"),
				).
				AddProp(
					"status",
					jsval.Reference(M).RefersTo("#/definitions/acceptance_status"),
				).
				AddProp(
					"tags",
					jsval.OneOf().
						Add(
							jsval.String(),
						).
						Add(
							jsval.Reference(M).RefersTo("#/definitions/tag_array"),
						),
				).
				AddProp(
					"title",
					jsval.String(),
				).
				AddProp(
					"video_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"video_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				),
		)

	HTTPLookupUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPLookupUserResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/email"),
				).
				AddProp(
					"first_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"last_name",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"nickname",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"tshirt_size",
					jsval.Reference(M).RefersTo("#/definitions/tshirt_size"),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPLookupVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				),
		)

	HTTPLookupVenueResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"name",
					jsval.String(),
				).
				AddProp(
					"rooms",
					jsval.Reference(M).RefersTo("#/definitions/room"),
				),
		)

	HTTPUpdateConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"description",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"slug",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"starts_on",
					jsval.Reference(M).RefersTo("#/definitions/datetime"),
				).
				AddProp(
					"sub_title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

	HTTPUpdateSessionRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"abstract",
					jsval.Reference(M).RefersTo("#/definitions/markdown_en"),
				).
				AddProp(
					"category",
					jsval.String(),
				).
				AddProp(
					"conference_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"confirmed",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"duration",
					jsval.Reference(M).RefersTo("#/definitions/duration"),
				).
				AddProp(
					"has_interpretation",
					jsval.Reference(M).RefersTo("#/definitions/boolean_default_false"),
				).
				AddProp(
					"id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"material_level",
					jsval.Reference(M).RefersTo("#/definitions/material_level"),
				).
				AddProp(
					"memo",
					jsval.String(),
				).
				AddProp(
					"photo_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"slide_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_subtitles",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"slide_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				AddProp(
					"sort_order",
					jsval.Reference(M).RefersTo("#/definitions/positiveInteger"),
				).
				AddProp(
					"speaker_id",
					jsval.Reference(M).RefersTo("#/definitions/uuid"),
				).
				AddProp(
					"spoken_language",
					jsval.Reference(M).RefersTo("#/definitions/language"),
				).
				AddProp(
					"status",
					jsval.Reference(M).RefersTo("#/definitions/acceptance_status"),
				).
				AddProp(
					"tags",
					jsval.String(),
				).
				AddProp(
					"title",
					jsval.Reference(M).RefersTo("#/definitions/string_en"),
				).
				AddProp(
					"video_permission",
					jsval.Reference(M).RefersTo("#/definitions/binary_permission_default_allow"),
				).
				AddProp(
					"video_url",
					jsval.Reference(M).RefersTo("#/definitions/url"),
				).
				PatternPropertiesString(
					"abstract#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/markdown_i18n"),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo("#/definitions/string_i18n"),
				),
		)

}
