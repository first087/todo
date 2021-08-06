package serialize

import (
	"encoding/json"
	"io"

	msgpack "github.com/vmihailenco/msgpack/v5"
)

type JSONSerializer struct{}

func (j JSONSerializer) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j JSONSerializer) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

type MessagePackSerializer struct{}

func (MessagePackSerializer) Decode(r io.Reader, v interface{}) error {
	return msgpack.NewDecoder(r).Decode(v)
}

func (MessagePackSerializer) Encode(w io.Writer, v interface{}) error {
	return msgpack.NewEncoder(w).Encode(v)
}
