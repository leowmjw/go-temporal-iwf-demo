run:
	@cd ./cmd/collection && go run *.go

runSub:
	@cd ./cmd/subscription && go run *.go

start:
	@curl http://localhost:8803/sim/start

test:
	@gotest ./...

tools:
	@echo "Download all the tools ... "

deps:
	@echo "Installing the deps .."
	@brew install ariga/tap/atlas
