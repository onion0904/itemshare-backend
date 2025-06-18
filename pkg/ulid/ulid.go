package ulid

import "github.com/oklog/ulid/v2"

func NewUlid() string {
	return ulid.Make().String()
}

func IsValid(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}
