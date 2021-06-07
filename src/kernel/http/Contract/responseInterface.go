package contract

import "github.com/ArtisanCloud/go-libs/object"

type ResponseContract interface {
	GetStatusCode() int
	WithStatus() object.HashMap
	GetReasonPhrase() object.HashMap
}

type PromiseInterface interface {

}