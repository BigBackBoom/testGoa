
REPO:=github.com/testGoa

init: depend bootstrap
gen: clean generate
rerun: clean generate run

depend:
	@which goagen || go get -v github.com/goadesign/goa/goagen

bootstrap:
	@goagen bootstrap -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
	@go build -o build

run:
	@./build
