package main

type (
	Repo struct {
		Clone string
	}

	Build struct {
		Path   string
		Event  string
		Commit string
	}

	Netrc struct {
		Machine  string
		Login    string
		Password string
	}

	Share struct {
		Pool     string
	}
)
