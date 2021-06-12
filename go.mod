module github.com/ArtisanCloud/go-wechat

go 1.16

replace github.com/ArtisanCloud/go-libs => ../go-libs

replace github.com/ArtisanCloud/go-socialite => ../go-socialite

require (
	github.com/ArtisanCloud/go-libs v1.0.11
	github.com/ArtisanCloud/go-socialite v1.0.0
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-redis/redis/v8 v8.10.0 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
