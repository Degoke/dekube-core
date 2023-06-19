.PHONY: make-vendor verify-codegen
setup:
	make-vendor
	verify-codegen

verify-codegen:
	chmod +x ./update-codegen.sh
	./update-codegen.sh

make-vendor:
	go mod vendor