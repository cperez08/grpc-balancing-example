gen-go-from-proto:
	protoc -I . --go_out ./ --go_opt plugins=grpc --go_opt paths=source_relative ./protos/post.proto

call-service:
	./call.sh localhost:$(PORT) ok $(CALLS)

call-service-fail:
	./call.sh localhost:$(PORT) fail $(CALLS)

## TAG = docker tag (my-image:v1), DEMO = demo #? (demo1, demo2, demo3), TARGET = (client, server)
build-image:
	docker build -t $(TAG) --build-arg DEMO=$(DEMO) -f Dockerfile.$(TARGET) .
