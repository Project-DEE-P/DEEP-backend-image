package cerrors

type (
	validationErr struct {
		errMessage string
	}

	parsingErr struct {
		errMessage string
	}

	authorizationErr struct {
		errMessage string
	}
)

func (v validationErr) Error() string {
	return v.errMessage
}

func ValidationErr(errMessage string) {
	panic(validationErr{errMessage: errMessage})
}

func (p parsingErr) Error() string {
	return p.errMessage
}

func ParsingErr(errMessage string) {
	panic(parsingErr{errMessage: errMessage})
}

func (a authorizationErr) Error() string {
	return a.errMessage
}

func AuthorizationErr(errMessage string) {
	panic(authorizationErr{errMessage: errMessage})
}
