package rpcdto

import "gophkeeper/models"

type (
	LoginDto struct {
		Login    string
		Password string
	}

	RegisterDto struct {
		Login    string
		Password string
	}
)

type (
	SaveLoginPasswordDto struct {
		User     models.GophUser
		Login    string
		Password string
	}

	LoadLoginPasswordDto struct {
		User  models.GophUser
		Login string
	}
)
