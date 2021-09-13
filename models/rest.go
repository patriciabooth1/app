package models

type TestRequest struct {
	Field1 string `json:"field_1,omitempty"`
	Field2 string `json:"field_2,omitempty"`
	Field3 string `json:"field_3,omitempty"`
}

func (t *TestRequest) ToSanitizedJSONString() string {

	out := `{`

	if t.Field1 != "" {
		out += `"field_1": "` + t.Field1 + `"`
	}

	if t.Field2 != "" {
		if len(out) > 1 {
			out += `, `
		}
		out += `"field_2": "` + t.Field2 + `"`
	}

	if t.Field3 != "" {
		if len(out) > 1 {
			out += `, `
		}
		out += `"field_3": "` + t.Field3 + `"`
	}

	out += `}`

	return out
}
