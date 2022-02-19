GOPATH := $(shell GOPATH)

default: build

build:
	GOPATH=${GOPATH} go build

install: build
	GOBIN=/usr/local/bin/ go install

install-daemon: install
	mkdir -p /Users/Shared/bin
	cp run/desklight.sh /Users/Shared/bin/
	cp run/custom.desklight.plist ~/Library/LaunchAgents/
	launchctl load -w ~/Library/LaunchAgents/custom.desklight.plist
	launchctl start custom.desklight

uninstall:
	sudo rm /usr/local/bin/lifx-cli
	rm /Users/Shared/bin/desklight.sh
	rm ~/Library/LaunchAgents/custom.desklight.plist
	launchctl unload -w ~/Library/LaunchAgents/custom.desklight.plist
	launchctl stop custom.desklight