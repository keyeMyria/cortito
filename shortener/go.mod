module shortener

go 1.12

require (
	common v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0 // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.4
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/stretchr/testify v1.3.0
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace common => ../common
