module github.com/ArtisanCloud/PowerWeChat

go 1.16

replace github.com/ArtisanCloud/PowerLibs => ../PowerLibs

replace github.com/ArtisanCloud/PowerSocialite => ../PowerSocialite

require (
	github.com/ArtisanCloud/PowerLibs v1.1.6
	github.com/ArtisanCloud/PowerSocialite v1.0.10
	github.com/ArtisanCloud/go-socialite v1.0.9 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.1
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
)
