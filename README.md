# go-erpcgen

Generate Go / TinyGo code using eRPC idl as input.  

* https://github.com/EmbeddedRPC/eRPC
* https://github.com/Seeed-Studio/seeed-ambd-firmware

## examples/*

The examples are a rewrite of the code in the following URL to work with TinyGo.  
Right now, it simply enumerates the eRPC calls.  

* https://wiki.seeedstudio.com/Wio-Terminal-Wi-Fi/
* https://github.com/tinygo-org/drivers/tree/release/examples/wifinina

## how to run examples

You can use the wio terminal to try out the code in the examples.  

* https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/

### run with TinyGo

The examples can be flashed as follows.  
After flashing, open the console and check the result.  

```
$ tinygo flash --target wioterminal --size short ./examples/webclient/
```

### run with Go

First, flash `testdata/wioterminal_passthrough/wioterminal_passthrough.uf2`.  
After that, run the following.  
The first time you run wioterminal after powering it on, the `--init` option is required.  
After that, do not add the `--init` option.  

```
$ go run ./examples/webclient/ --port COM59 --debug --init
```

```
$ go run ./examples/webclient/ --port COM59 --debug
```

It is important to note that the passthrough part is used for relaying, so high-speed and large data transmissions will often fail.  
