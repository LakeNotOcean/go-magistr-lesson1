module github.com/LakeNotOcean/go-magistr-lesson1

go 1.25.1

replace (
github.com/LakeNotOcean/go-magistr-lesson1/handlers => ./config
github.com/LakeNotOcean/go-magistr-lesson1/config => ./handlers
)

require (
	github.com/go-resty/resty/v2 v2.16.5
	github.com/joho/godotenv v1.5.1
)

require golang.org/x/net v0.33.0 // indirect
