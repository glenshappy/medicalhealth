module medicalhealth

go 1.12

require (
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v0.0.0-20190318194812-d3c38a4eb497
	github.com/json-iterator/go v1.1.6
	github.com/mattn/go-isatty v0.0.7
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/ugorji/go v0.0.0-20190320090025-2dc34c0b8780
	golang.org/x/sys v0.0.0
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	golang.org/x/sys v0.0.0 => golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190710143415-6ec70d6a5542
)
