module github.com/ArtisanCloud/go-wechat

go 1.16

replace github.com/ArtisanCloud/go-libs => ../go-libs

replace github.com/ArtisanCloud/go-socialite => ../go-socialite

require (
	github.com/ArtisanCloud/go-libs v1.0.12
	github.com/ArtisanCloud/go-socialite v1.0.2
	github.com/clbanning/mxj/v2 v2.3.2 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-redis/redis/v8 v8.10.0 // indirect
	github.com/stretchr/testify v1.7.0
)
