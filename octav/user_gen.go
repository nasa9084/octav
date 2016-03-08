// Automatically generated by genmodel utility. DO NOT EDIT!
package octav

import "encoding/json"

func (v User) GetPropNames() ([]string, error) {
	l, _ := v.L10N.GetPropNames()
	return append(l, "id", "first_name", "last_name", "nickname", "email", "tshirt_size"), nil
}

func (v User) GetPropValue(s string) (interface{}, error) {
	switch s {
	case "id":
		return v.ID, nil
	case "first_name":
		return v.FirstName, nil
	case "last_name":
		return v.LastName, nil
	case "nickname":
		return v.Nickname, nil
	case "email":
		return v.Email, nil
	case "tshirt_size":
		return v.TshirtSize, nil
	default:
		return v.L10N.GetPropValue(s)
	}
}

func (v User) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["first_name"] = v.FirstName
	m["last_name"] = v.LastName
	m["nickname"] = v.Nickname
	m["email"] = v.Email
	m["tshirt_size"] = v.TshirtSize
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalJSONWithL10N(buf, v.L10N)
}
func (v *User) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			v.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["first_name"]; ok {
		switch jv.(type) {
		case string:
			v.FirstName = jv.(string)
			delete(m, "first_name")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["last_name"]; ok {
		switch jv.(type) {
		case string:
			v.LastName = jv.(string)
			delete(m, "last_name")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["nickname"]; ok {
		switch jv.(type) {
		case string:
			v.Nickname = jv.(string)
			delete(m, "nickname")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["email"]; ok {
		switch jv.(type) {
		case string:
			v.Email = jv.(string)
			delete(m, "email")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["tshirt_size"]; ok {
		switch jv.(type) {
		case string:
			v.TshirtSize = jv.(string)
			delete(m, "tshirt_size")
		default:
			return ErrInvalidFieldType
		}
	}
	return nil
}
