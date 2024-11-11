.PHONY: build
build:
	sam build --no-cached

.PHONY: run
run: build
	sam local start-api --docker-network lambda-local

.PHONY: invoke-%
invoke-%: build
	sam local invoke $(subst invoke-,,$@) --docker-network lambda-local --event ${EVENT_FILE}

.PHONY: test
test:
	go test ./handlers/...

.PHONY: clean
clean:
	rm -rf .aws-sam