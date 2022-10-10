package exceptions

type EmailAlreadyExist struct {
	Error string
}

func NewEmailAlreadyExist(error string) EmailAlreadyExist {
	return EmailAlreadyExist{Error: error}
}
