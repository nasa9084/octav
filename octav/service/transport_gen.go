// Automatically generated by gentransport utility. DO NOT EDIT!
package service

import (
	"encoding/json"
	"errors"

	"github.com/builderscon/octav/octav/tools"

	"github.com/lestrrat/go-urlenc"
)

func (r CreateConferenceRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["title"] = r.Title
	if r.SubTitle.Valid() {
		m["sub_title"] = r.SubTitle.Value()
	}
	m["slug"] = r.Slug
	return m
}

func (r CreateConferenceRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalJSONWithL10N(buf, r.L10N)
}

func (r CreateConferenceRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalURLWithL10N(buf, r.L10N)
}

func (r *CreateConferenceRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["title"]; ok {
		switch jv.(type) {
		case string:
			r.Title = jv.(string)
			delete(m, "title")
		default:
			return ErrInvalidJSONFieldType{Field: "title"}
		}
	}
	if jv, ok := m["sub_title"]; ok {
		if err := r.SubTitle.Set(jv); err != nil {
			return errors.New("set field SubTitle failed: " + err.Error())
		}
		delete(m, "sub_title")
	}
	if jv, ok := m["slug"]; ok {
		switch jv.(type) {
		case string:
			r.Slug = jv.(string)
			delete(m, "slug")
		default:
			return ErrInvalidJSONFieldType{Field: "slug"}
		}
	}
	if err := tools.ExtractL10NFields(m, &r.L10N, []string{"title", "sub_title", "slug"}); err != nil {
		return err
	}
	return nil
}

func (r *CreateConferenceRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "title", "sub_title", "slug"), nil
}

func (r *CreateConferenceRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "title":
		if jv, ok := v.(string); ok {
			r.Title = jv
			return nil
		}
	case "sub_title":
		return r.SubTitle.Set(v)
	case "slug":
		if jv, ok := v.(string); ok {
			r.Slug = jv
			return nil
		}
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r UpdateConferenceRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = r.ID
	if r.Title.Valid() {
		m["title"] = r.Title.Value()
	}
	if r.SubTitle.Valid() {
		m["sub_title"] = r.SubTitle.Value()
	}
	if r.Slug.Valid() {
		m["slug"] = r.Slug.Value()
	}
	return m
}

func (r UpdateConferenceRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalJSONWithL10N(buf, r.L10N)
}

func (r UpdateConferenceRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalURLWithL10N(buf, r.L10N)
}

func (r *UpdateConferenceRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			r.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidJSONFieldType{Field: "id"}
		}
	}
	if jv, ok := m["title"]; ok {
		if err := r.Title.Set(jv); err != nil {
			return errors.New("set field Title failed: " + err.Error())
		}
		delete(m, "title")
	}
	if jv, ok := m["sub_title"]; ok {
		if err := r.SubTitle.Set(jv); err != nil {
			return errors.New("set field SubTitle failed: " + err.Error())
		}
		delete(m, "sub_title")
	}
	if jv, ok := m["slug"]; ok {
		if err := r.Slug.Set(jv); err != nil {
			return errors.New("set field Slug failed: " + err.Error())
		}
		delete(m, "slug")
	}
	if err := tools.ExtractL10NFields(m, &r.L10N, []string{"id", "title", "sub_title", "slug"}); err != nil {
		return err
	}
	return nil
}

func (r *UpdateConferenceRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "id", "title", "sub_title", "slug"), nil
}

func (r *UpdateConferenceRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "id":
		if jv, ok := v.(string); ok {
			r.ID = jv
			return nil
		}
	case "title":
		return r.Title.Set(v)
	case "sub_title":
		return r.SubTitle.Set(v)
	case "slug":
		return r.Slug.Set(v)
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r CreateSessionRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	if r.ConferenceID.Valid() {
		m["conference_id"] = r.ConferenceID.Value()
	}
	if r.SpeakerID.Valid() {
		m["speaker_id"] = r.SpeakerID.Value()
	}
	if r.Title.Valid() {
		m["title"] = r.Title.Value()
	}
	if r.Abstract.Valid() {
		m["abstract"] = r.Abstract.Value()
	}
	if r.Memo.Valid() {
		m["memo"] = r.Memo.Value()
	}
	if r.Duration.Valid() {
		m["duration"] = r.Duration.Value()
	}
	if r.MaterialLevel.Valid() {
		m["material_level"] = r.MaterialLevel.Value()
	}
	if r.Tags.Valid() {
		m["tags"] = r.Tags.Value()
	}
	if r.Category.Valid() {
		m["category"] = r.Category.Value()
	}
	if r.SpokenLanguage.Valid() {
		m["spoken_language"] = r.SpokenLanguage.Value()
	}
	if r.SlideLanguage.Valid() {
		m["slide_language"] = r.SlideLanguage.Value()
	}
	if r.SlideSubtitles.Valid() {
		m["slide_subtitles"] = r.SlideSubtitles.Value()
	}
	if r.SlideURL.Valid() {
		m["slide_url"] = r.SlideURL.Value()
	}
	if r.VideoURL.Valid() {
		m["video_url"] = r.VideoURL.Value()
	}
	if r.PhotoPermission.Valid() {
		m["photo_permission"] = r.PhotoPermission.Value()
	}
	if r.VideoPermission.Valid() {
		m["video_permission"] = r.VideoPermission.Value()
	}
	return m
}

func (r CreateSessionRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalJSONWithL10N(buf, r.L10N)
}

func (r CreateSessionRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalURLWithL10N(buf, r.L10N)
}

func (r *CreateSessionRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["conference_id"]; ok {
		if err := r.ConferenceID.Set(jv); err != nil {
			return errors.New("set field ConferenceID failed: " + err.Error())
		}
		delete(m, "conference_id")
	}
	if jv, ok := m["speaker_id"]; ok {
		if err := r.SpeakerID.Set(jv); err != nil {
			return errors.New("set field SpeakerID failed: " + err.Error())
		}
		delete(m, "speaker_id")
	}
	if jv, ok := m["title"]; ok {
		if err := r.Title.Set(jv); err != nil {
			return errors.New("set field Title failed: " + err.Error())
		}
		delete(m, "title")
	}
	if jv, ok := m["abstract"]; ok {
		if err := r.Abstract.Set(jv); err != nil {
			return errors.New("set field Abstract failed: " + err.Error())
		}
		delete(m, "abstract")
	}
	if jv, ok := m["memo"]; ok {
		if err := r.Memo.Set(jv); err != nil {
			return errors.New("set field Memo failed: " + err.Error())
		}
		delete(m, "memo")
	}
	if jv, ok := m["duration"]; ok {
		if err := r.Duration.Set(jv); err != nil {
			return errors.New("set field Duration failed: " + err.Error())
		}
		delete(m, "duration")
	}
	if jv, ok := m["material_level"]; ok {
		if err := r.MaterialLevel.Set(jv); err != nil {
			return errors.New("set field MaterialLevel failed: " + err.Error())
		}
		delete(m, "material_level")
	}
	if jv, ok := m["tags"]; ok {
		if err := r.Tags.Set(jv); err != nil {
			return errors.New("set field Tags failed: " + err.Error())
		}
		delete(m, "tags")
	}
	if jv, ok := m["category"]; ok {
		if err := r.Category.Set(jv); err != nil {
			return errors.New("set field Category failed: " + err.Error())
		}
		delete(m, "category")
	}
	if jv, ok := m["spoken_language"]; ok {
		if err := r.SpokenLanguage.Set(jv); err != nil {
			return errors.New("set field SpokenLanguage failed: " + err.Error())
		}
		delete(m, "spoken_language")
	}
	if jv, ok := m["slide_language"]; ok {
		if err := r.SlideLanguage.Set(jv); err != nil {
			return errors.New("set field SlideLanguage failed: " + err.Error())
		}
		delete(m, "slide_language")
	}
	if jv, ok := m["slide_subtitles"]; ok {
		if err := r.SlideSubtitles.Set(jv); err != nil {
			return errors.New("set field SlideSubtitles failed: " + err.Error())
		}
		delete(m, "slide_subtitles")
	}
	if jv, ok := m["slide_url"]; ok {
		if err := r.SlideURL.Set(jv); err != nil {
			return errors.New("set field SlideURL failed: " + err.Error())
		}
		delete(m, "slide_url")
	}
	if jv, ok := m["video_url"]; ok {
		if err := r.VideoURL.Set(jv); err != nil {
			return errors.New("set field VideoURL failed: " + err.Error())
		}
		delete(m, "video_url")
	}
	if jv, ok := m["photo_permission"]; ok {
		if err := r.PhotoPermission.Set(jv); err != nil {
			return errors.New("set field PhotoPermission failed: " + err.Error())
		}
		delete(m, "photo_permission")
	}
	if jv, ok := m["video_permission"]; ok {
		if err := r.VideoPermission.Set(jv); err != nil {
			return errors.New("set field VideoPermission failed: " + err.Error())
		}
		delete(m, "video_permission")
	}
	if err := tools.ExtractL10NFields(m, &r.L10N, []string{"conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission"}); err != nil {
		return err
	}
	return nil
}

func (r *CreateSessionRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission"), nil
}

func (r *CreateSessionRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "conference_id":
		return r.ConferenceID.Set(v)
	case "speaker_id":
		return r.SpeakerID.Set(v)
	case "title":
		return r.Title.Set(v)
	case "abstract":
		return r.Abstract.Set(v)
	case "memo":
		return r.Memo.Set(v)
	case "duration":
		return r.Duration.Set(v)
	case "material_level":
		return r.MaterialLevel.Set(v)
	case "tags":
		return r.Tags.Set(v)
	case "category":
		return r.Category.Set(v)
	case "spoken_language":
		return r.SpokenLanguage.Set(v)
	case "slide_language":
		return r.SlideLanguage.Set(v)
	case "slide_subtitles":
		return r.SlideSubtitles.Set(v)
	case "slide_url":
		return r.SlideURL.Set(v)
	case "video_url":
		return r.VideoURL.Set(v)
	case "photo_permission":
		return r.PhotoPermission.Set(v)
	case "video_permission":
		return r.VideoPermission.Set(v)
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r UpdateSessionRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = r.ID
	if r.ConferenceID.Valid() {
		m["conference_id"] = r.ConferenceID.Value()
	}
	if r.SpeakerID.Valid() {
		m["speaker_id"] = r.SpeakerID.Value()
	}
	if r.Title.Valid() {
		m["title"] = r.Title.Value()
	}
	if r.Abstract.Valid() {
		m["abstract"] = r.Abstract.Value()
	}
	if r.Memo.Valid() {
		m["memo"] = r.Memo.Value()
	}
	if r.Duration.Valid() {
		m["duration"] = r.Duration.Value()
	}
	if r.MaterialLevel.Valid() {
		m["material_level"] = r.MaterialLevel.Value()
	}
	if r.Tags.Valid() {
		m["tags"] = r.Tags.Value()
	}
	if r.Category.Valid() {
		m["category"] = r.Category.Value()
	}
	if r.SpokenLanguage.Valid() {
		m["spoken_language"] = r.SpokenLanguage.Value()
	}
	if r.SlideLanguage.Valid() {
		m["slide_language"] = r.SlideLanguage.Value()
	}
	if r.SlideSubtitles.Valid() {
		m["slide_subtitles"] = r.SlideSubtitles.Value()
	}
	if r.SlideURL.Valid() {
		m["slide_url"] = r.SlideURL.Value()
	}
	if r.VideoURL.Valid() {
		m["video_url"] = r.VideoURL.Value()
	}
	if r.PhotoPermission.Valid() {
		m["photo_permission"] = r.PhotoPermission.Value()
	}
	if r.VideoPermission.Valid() {
		m["video_permission"] = r.VideoPermission.Value()
	}
	if r.SortOrder.Valid() {
		m["sort_order"] = r.SortOrder.Value()
	}
	if r.HasInterpretation.Valid() {
		m["has_interpretation"] = r.HasInterpretation.Value()
	}
	if r.Status.Valid() {
		m["status"] = r.Status.Value()
	}
	if r.Confirmed.Valid() {
		m["confirmed"] = r.Confirmed.Value()
	}
	return m
}

func (r UpdateSessionRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalJSONWithL10N(buf, r.L10N)
}

func (r UpdateSessionRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return tools.MarshalURLWithL10N(buf, r.L10N)
}

func (r *UpdateSessionRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			r.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidJSONFieldType{Field: "id"}
		}
	}
	if jv, ok := m["conference_id"]; ok {
		if err := r.ConferenceID.Set(jv); err != nil {
			return errors.New("set field ConferenceID failed: " + err.Error())
		}
		delete(m, "conference_id")
	}
	if jv, ok := m["speaker_id"]; ok {
		if err := r.SpeakerID.Set(jv); err != nil {
			return errors.New("set field SpeakerID failed: " + err.Error())
		}
		delete(m, "speaker_id")
	}
	if jv, ok := m["title"]; ok {
		if err := r.Title.Set(jv); err != nil {
			return errors.New("set field Title failed: " + err.Error())
		}
		delete(m, "title")
	}
	if jv, ok := m["abstract"]; ok {
		if err := r.Abstract.Set(jv); err != nil {
			return errors.New("set field Abstract failed: " + err.Error())
		}
		delete(m, "abstract")
	}
	if jv, ok := m["memo"]; ok {
		if err := r.Memo.Set(jv); err != nil {
			return errors.New("set field Memo failed: " + err.Error())
		}
		delete(m, "memo")
	}
	if jv, ok := m["duration"]; ok {
		if err := r.Duration.Set(jv); err != nil {
			return errors.New("set field Duration failed: " + err.Error())
		}
		delete(m, "duration")
	}
	if jv, ok := m["material_level"]; ok {
		if err := r.MaterialLevel.Set(jv); err != nil {
			return errors.New("set field MaterialLevel failed: " + err.Error())
		}
		delete(m, "material_level")
	}
	if jv, ok := m["tags"]; ok {
		if err := r.Tags.Set(jv); err != nil {
			return errors.New("set field Tags failed: " + err.Error())
		}
		delete(m, "tags")
	}
	if jv, ok := m["category"]; ok {
		if err := r.Category.Set(jv); err != nil {
			return errors.New("set field Category failed: " + err.Error())
		}
		delete(m, "category")
	}
	if jv, ok := m["spoken_language"]; ok {
		if err := r.SpokenLanguage.Set(jv); err != nil {
			return errors.New("set field SpokenLanguage failed: " + err.Error())
		}
		delete(m, "spoken_language")
	}
	if jv, ok := m["slide_language"]; ok {
		if err := r.SlideLanguage.Set(jv); err != nil {
			return errors.New("set field SlideLanguage failed: " + err.Error())
		}
		delete(m, "slide_language")
	}
	if jv, ok := m["slide_subtitles"]; ok {
		if err := r.SlideSubtitles.Set(jv); err != nil {
			return errors.New("set field SlideSubtitles failed: " + err.Error())
		}
		delete(m, "slide_subtitles")
	}
	if jv, ok := m["slide_url"]; ok {
		if err := r.SlideURL.Set(jv); err != nil {
			return errors.New("set field SlideURL failed: " + err.Error())
		}
		delete(m, "slide_url")
	}
	if jv, ok := m["video_url"]; ok {
		if err := r.VideoURL.Set(jv); err != nil {
			return errors.New("set field VideoURL failed: " + err.Error())
		}
		delete(m, "video_url")
	}
	if jv, ok := m["photo_permission"]; ok {
		if err := r.PhotoPermission.Set(jv); err != nil {
			return errors.New("set field PhotoPermission failed: " + err.Error())
		}
		delete(m, "photo_permission")
	}
	if jv, ok := m["video_permission"]; ok {
		if err := r.VideoPermission.Set(jv); err != nil {
			return errors.New("set field VideoPermission failed: " + err.Error())
		}
		delete(m, "video_permission")
	}
	if jv, ok := m["sort_order"]; ok {
		if err := r.SortOrder.Set(jv); err != nil {
			return errors.New("set field SortOrder failed: " + err.Error())
		}
		delete(m, "sort_order")
	}
	if jv, ok := m["has_interpretation"]; ok {
		if err := r.HasInterpretation.Set(jv); err != nil {
			return errors.New("set field HasInterpretation failed: " + err.Error())
		}
		delete(m, "has_interpretation")
	}
	if jv, ok := m["status"]; ok {
		if err := r.Status.Set(jv); err != nil {
			return errors.New("set field Status failed: " + err.Error())
		}
		delete(m, "status")
	}
	if jv, ok := m["confirmed"]; ok {
		if err := r.Confirmed.Set(jv); err != nil {
			return errors.New("set field Confirmed failed: " + err.Error())
		}
		delete(m, "confirmed")
	}
	if err := tools.ExtractL10NFields(m, &r.L10N, []string{"id", "conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission", "sort_order", "has_interpretation", "status", "confirmed"}); err != nil {
		return err
	}
	return nil
}

func (r *UpdateSessionRequest) GetPropNames() ([]string, error) {
	l, _ := r.L10N.GetPropNames()
	return append(l, "id", "conference_id", "speaker_id", "title", "abstract", "memo", "duration", "material_level", "tags", "category", "spoken_language", "slide_language", "slide_subtitles", "slide_url", "video_url", "photo_permission", "video_permission", "sort_order", "has_interpretation", "status", "confirmed"), nil
}

func (r *UpdateSessionRequest) SetPropValue(s string, v interface{}) error {
	switch s {
	case "id":
		if jv, ok := v.(string); ok {
			r.ID = jv
			return nil
		}
	case "conference_id":
		return r.ConferenceID.Set(v)
	case "speaker_id":
		return r.SpeakerID.Set(v)
	case "title":
		return r.Title.Set(v)
	case "abstract":
		return r.Abstract.Set(v)
	case "memo":
		return r.Memo.Set(v)
	case "duration":
		return r.Duration.Set(v)
	case "material_level":
		return r.MaterialLevel.Set(v)
	case "tags":
		return r.Tags.Set(v)
	case "category":
		return r.Category.Set(v)
	case "spoken_language":
		return r.SpokenLanguage.Set(v)
	case "slide_language":
		return r.SlideLanguage.Set(v)
	case "slide_subtitles":
		return r.SlideSubtitles.Set(v)
	case "slide_url":
		return r.SlideURL.Set(v)
	case "video_url":
		return r.VideoURL.Set(v)
	case "photo_permission":
		return r.PhotoPermission.Set(v)
	case "video_permission":
		return r.VideoPermission.Set(v)
	case "sort_order":
		return r.SortOrder.Set(v)
	case "has_interpretation":
		return r.HasInterpretation.Set(v)
	case "status":
		return r.Status.Set(v)
	case "confirmed":
		return r.Confirmed.Set(v)
	default:
		return errors.New("unknown column '" + s + "'")
	}
	return ErrInvalidFieldType{Field: s}
}

func (r ListVenueRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Since.Valid() {
		m["since"] = r.Since.Value()
	}
	if r.Lang.Valid() {
		m["lang"] = r.Lang.Value()
	}
	if r.Limit.Valid() {
		m["limit"] = r.Limit.Value()
	}
	return m
}

func (r ListVenueRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r ListVenueRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *ListVenueRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["since"]; ok {
		if err := r.Since.Set(jv); err != nil {
			return errors.New("set field Since failed: " + err.Error())
		}
		delete(m, "since")
	}
	if jv, ok := m["lang"]; ok {
		if err := r.Lang.Set(jv); err != nil {
			return errors.New("set field Lang failed: " + err.Error())
		}
		delete(m, "lang")
	}
	if jv, ok := m["limit"]; ok {
		if err := r.Limit.Set(jv); err != nil {
			return errors.New("set field Limit failed: " + err.Error())
		}
		delete(m, "limit")
	}
	return nil
}

func (r ListSessionsByConferenceRequest) collectMarshalData() map[string]interface{} {
	m := make(map[string]interface{})
	m["conference_id"] = r.ConferenceID
	if r.Date.Valid() {
		m["date"] = r.Date.Value()
	}
	return m
}

func (r ListSessionsByConferenceRequest) MarshalJSON() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r ListSessionsByConferenceRequest) MarshalURL() ([]byte, error) {
	m := r.collectMarshalData()
	buf, err := urlenc.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *ListSessionsByConferenceRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["conference_id"]; ok {
		switch jv.(type) {
		case string:
			r.ConferenceID = jv.(string)
			delete(m, "conference_id")
		default:
			return ErrInvalidJSONFieldType{Field: "conference_id"}
		}
	}
	if jv, ok := m["date"]; ok {
		if err := r.Date.Set(jv); err != nil {
			return errors.New("set field Date failed: " + err.Error())
		}
		delete(m, "date")
	}
	return nil
}
