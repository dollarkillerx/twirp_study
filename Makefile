build_twirp:
	protoc --twirp_out=. --go_out=pb pb/*.proto