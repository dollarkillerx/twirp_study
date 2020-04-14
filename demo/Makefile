build_twirp:
	protoc --twirp_out=. --go_out=pb pb/*.proto

build_grpc:
	protoc --go_out=plugins=grpc:. *.proto