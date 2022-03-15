package uuid

import (
	"github.com/google/uuid"
)

func FromString(inputParam string) uuid.UUID {
	const defaultUUID = "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	return uuid.NewSHA1(uuid.MustParse(defaultUUID), []byte(inputParam))
}
