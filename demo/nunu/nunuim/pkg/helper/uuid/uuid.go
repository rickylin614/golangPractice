package uuid

import uuid "github.com/gofrs/uuid"

func GenUUID() string {
	id, _ := uuid.NewV7()
	return id.String()
}
