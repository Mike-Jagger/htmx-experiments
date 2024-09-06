# all: build test - if I had tests ready 

run: build
	@./bin/gowiki.exe

compile:
	@go install

build:
	@cd ./handlers && go build
	@cd ./shared && go build
	@cd ./util && go build