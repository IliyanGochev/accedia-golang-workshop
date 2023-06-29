go get github.com/google/uuid
go get github.com/pborman/uuid
go mod tidy 
go build
go build .\main.go
$Env:GOOS = "linux"
