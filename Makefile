all: ./rtl8720dn/rpc.go

./rtl8720dn/rpc.go: ./testdata/wioterminal-rtl8720dn.idl
	go run . -I $^ > $@
	gofmt -w ./rtl8720dn/rpc.go

./testdata/wioterminal-rtl8720dn.idl: ./testdata/erpc_idl/rpc_system.erpc FORCE
	perl ./testdata/concat_idl.pl ./testdata/erpc_idl/rpc_system.erpc > $@

smoketest:
	go build ./examples/01_check_firmware_version/
	go build ./examples/02_scanning_wifi_network/
	go build ./examples/03_connecting_to_specified_network/
	go build ./examples/04_wifi_client/
	go build ./examples/05_wifi_https_connection/
	go build ./examples/06_udp_client/
	go build ./examples/tlsclient/
	go build ./examples/webclient/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/01_check_firmware_version/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/02_scanning_wifi_network/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/03_connecting_to_specified_network/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/04_wifi_client/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/05_wifi_https_connection/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/06_udp_client/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/tlsclient/
	tinygo build -o /tmp/test.hex --target wioterminal --size short ./examples/webclient/

.PHONY: FORCE
FORCE:
