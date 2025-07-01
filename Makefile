.PHONY: build update

build:
	-xargs -d '\n' rm < .openapi-generator/FILES
	openapi-generator-cli generate -i openapi/openapi.yaml -g go --git-user-id DFMailbox --git-repo-id go-client --git-host github.com
	echo ".direnv/" >> .gitignore

update:
	git submodule update --remote
