# GoPort

## The port scanner that works :eyes:

### Features:
- Scans from any range (currently only for TCP)
- Saves logs to text files every time you scan something
- Uses [Go routines](https://golangbot.com/goroutines/) / tends to be concurrent (fast)

### To-do:
- [ ] Advanced multi-threading
- [ ] Vulnerability scanner
- [ ] Service detector (e.g: SSH, Apache, e.t.c.)
- [ ] OS detector

### Contributors
- [pasenidis](https://github.com/pasenidis)
- [ntiakakis](https://github.com/ntiakakis)

### Build
Linux build variables
`GOARCH=amd64`
`GOOS=linux`
```sh
git clone https://github.com/pasenidis/goport
cd goport
go build main.go
./main
```
