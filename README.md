# golang-mock-testing
# steps
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@v1.6.0
mockgen -destination=mocks/mockRunner.go -package=mocks github.com/an1l4/mocktest/iuser IUserInterface
go test -v github.com/an1l4/mocktest/user
