package rpcdto

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
	SaveRecordRequestDto struct {
		Token string
		Key   string
		Data  string
	}

	LoadRecordRequestDto struct {
		Token string
		Key   string
	}
)
