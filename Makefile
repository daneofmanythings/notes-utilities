install:
	@ cd ./cmd/nutils/ && go install
	@ echo ...installation successful!

test:
	@ go test ./...
