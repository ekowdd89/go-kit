package author

import "io"



type AuthorContract interface {
	io.Close
	FetchAll()([]Author)
}

type Author struct {
	Id int64
	Name string
}