package kernel

import "github.com/ArtisanCloud/PowerLibs/v2/object"

type Config struct {
	*object.Collection
}

func NewConfig(items *object.HashMap) *Config {

	return &Config{
		object.NewCollection(items),
	}
}
