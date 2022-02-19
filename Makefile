default: build

clean:
	rm -rf bin/

setup:
	mkdir -p /Users/Shared/bin

build:
	GOPATH=${GOPATH} go build -o bin/

install:
	GOBIN=/usr/local/bin/ go install

install-handler:
	gcc -framework Foundation -o bin/xpc_set_event_stream_handler stream/xpc_set_event_stream_handler.m
	cp bin/xpc_set_event_stream_handler /Users/Shared/bin/

install-agent: install-handler
	cp run/desklight.sh /Users/Shared/bin/
	cp run/custom.desklight.plist ~/Library/LaunchAgents/
	launchctl load -w ~/Library/LaunchAgents/custom.desklight.plist
	launchctl start custom.desklight

install-all: install install-agent
	

uninstall:
	rm /usr/local/bin/lifx-cli

uninstall-agent:
	launchctl unload -w ~/Library/LaunchAgents/custom.desklight.plist
	launchctl stop custom.desklight
	rm /Users/Shared/bin/desklight.sh
	rm ~/Library/LaunchAgents/custom.desklight.plist


	# rm /Users/Shared/bin/xpc_set_event_stream_handler
