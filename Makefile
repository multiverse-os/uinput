# uinput Makefile
###############################################################################


default: install-deps build 

#install:
#	probably put it in ~/.local/share/bin but need it in path too
#	or to not need to modify path there is the uglier /usr/local/bin/
#	but that requires root

build:
	mkdir -p bin
	go build cmd/parse-usb-ids/ 
	mv cmd/parse-usb-ids/parse-usb-ids bin/

install-deps:
	sudo apt-get install libusb-dev
	go mod tidy

clean:
	rm -rf bin
	rm -rf cmd/parse-usb-ids/parse-usb-ids
	rm -rf go.sum
