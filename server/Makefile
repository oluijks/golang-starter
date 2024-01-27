# Makefile

.PHONY: help tidy audit test mock docs server air

.DEFAULT_GOAL := help

##@
##@ Quality Control
##@

tidy: ##@ format and tidy
	go fmt ./...
	go mod tidy -v
	
audit: ##@ audit: run quality control checks
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...
	
##@
##@ Documentation 
##@	
docs: ##@ docs: run generated docs in the browser ()
	pkgsite -http=:6060

##@
##@ Tests 
##@

test: ##@ run tests with coverage
	go test -v -cover ./...

##@
##@ Api Server 
##@

server: ##@ run the program
	go run cmd/server/main.go

air: ##@ run air
	~/go/bin/air

##@
##@ Misc 
##@

help: ##@ (Default) Print listing of key targets with their descriptions
	@printf "\nUsage: make <command>\n"
	@grep -F -h "##@" $(MAKEFILE_LIST) | grep -F -v grep -F | sed -e 's/\\$$//' | awk 'BEGIN {FS = ":*[[:space:]]*##@[[:space:]]*"}; \
	{ \
		if($$2 == "") \
			pass; \
		else if($$0 ~ /^#/) \
			printf "\n%s\n", $$2; \
		else if($$1 == "") \
			printf "     %-20s%s\n", "", $$2; \
		else \
			printf "\n    \033[34m%-20s\033[0m %s\n", $$1, $$2; \
	}'	