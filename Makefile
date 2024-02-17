build-image:
	docker build -t bianavic/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

lint:
	golangci-lint run
	go fmt -n ./...

unit-test:
	go test ./...