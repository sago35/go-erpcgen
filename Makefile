all: ./rtl8720dn/rpc.go

./rtl8720dn/rpc.go: ./testdata/wioterminal-rtl8720dn.idl
	go run . -I $^ > $@

./testdata/wioterminal-rtl8720dn.idl: ./testdata/erpc_idl/rpc_system.erpc
	perl ./testdata/concat_idl.pl ./testdata/erpc_idl/rpc_system.erpc > $@
