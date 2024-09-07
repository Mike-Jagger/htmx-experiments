all: build test

run: 
	build
	compile
	exec

exec: @cd ./bin && .\gowiki.exe

compile:
	@go install

build:
	@cd ./shared && go build
	@cd ./util && go build
	@cd ./handlers && go build