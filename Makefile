build:
	@go build -v .

block:
	sudo ./webblock-cli block

unblock:
	sudo ./webblock-cli unblock