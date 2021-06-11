package kernel

import "github.com/ArtisanCloud/go-libs/object"

type Config struct {
	*object.Collection
}

func NewConfig(items *object.HashMap) *Config{
	return &Config{
		object.NewCollection(items),
	}
}