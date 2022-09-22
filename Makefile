all : proto brain body
.PHONY : all proto brain body

proto : thesocialrobot/thesocialrobot.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative thesocialrobot/thesocialrobot.proto

# is there a better way to run the go build 'go build core' does not work
brain : proto
	cd core && go build .

body : proto
	cd fake-body && go build .