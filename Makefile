.PHONY: clean build build-be build-fe
	
build: build-fe build-be
	
build-be:
	go build -o ./build/ ./cmd/srvplane
	
build-fe:
	cd web && \
	bun install && \
	bun run build && \
	cp -r ./dist ../cmd/srvplane 
	
clean: 
	rm -rf ./build ./web/dist
	
dev-be:
	SRVPLANE_DEBUG=true go run cmd/srvplane/main.go
	
dev-fe:
	cd web && bun dev