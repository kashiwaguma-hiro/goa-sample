GOA_DESIGN=github.com/kashiwaguma-hiro/goa-sample/design

.PHONY:gen
gen:
	goa gen ${GOA_DESIGN}
	cp ./gen/http/openapi3.yaml ./docs/

.PHONY:clean
clean:
	rm -rf cmd/* admin.go

.PHONY:example
example:
	goa example ${GOA_DESIGN}

