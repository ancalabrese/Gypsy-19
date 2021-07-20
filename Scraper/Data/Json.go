package Data

import (
	"encoding/json"
	"io"
)

//ToJson encodes the interface i in Json format and writes it to w
func ToPrettyJson(i interface{}, w io.Writer) error {
	bytes, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func ToJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func FromJson(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
