GOPKG ?=	moul.io/openfiles

.PHONY: test-archs
test-archs:
	GOOS=darwin    GOARCH=386      go test -v -c -o /dev/null
	GOOS=darwin    GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=linux     GOARCH=386      go test -v -c -o /dev/null
	GOOS=linux     GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=linux     GOARCH=arm      go test -v -c -o /dev/null
	GOOS=linux     GOARCH=arm64    go test -v -c -o /dev/null
	GOOS=linux     GOARCH=ppc64    go test -v -c -o /dev/null
	GOOS=linux     GOARCH=ppc64le  go test -v -c -o /dev/null
	GOOS=linux     GOARCH=mips     go test -v -c -o /dev/null
	GOOS=linux     GOARCH=mipsle   go test -v -c -o /dev/null
	GOOS=linux     GOARCH=mips64   go test -v -c -o /dev/null
	GOOS=linux     GOARCH=mips64le go test -v -c -o /dev/null
	GOOS=linux     GOARCH=s390x    go test -v -c -o /dev/null
	GOOS=js        GOARCH=wasm     go test -v -c -o /dev/null
	GOOS=dragonfly GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=freebsd   GOARCH=386      go test -v -c -o /dev/null
	GOOS=freebsd   GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=freebsd   GOARCH=arm      go test -v -c -o /dev/null
	GOOS=netbsd    GOARCH=386      go test -v -c -o /dev/null
	GOOS=netbsd    GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=netbsd    GOARCH=arm      go test -v -c -o /dev/null
	GOOS=openbsd   GOARCH=386      go test -v -c -o /dev/null
	GOOS=openbsd   GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=openbsd   GOARCH=arm      go test -v -c -o /dev/null
	GOOS=plan9     GOARCH=386      go test -v -c -o /dev/null
	GOOS=plan9     GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=plan9     GOARCH=arm      go test -v -c -o /dev/null
	GOOS=solaris   GOARCH=amd64    go test -v -c -o /dev/null
	GOOS=windows   GOARCH=386      go test -v -c -o /dev/null
	GOOS=windows   GOARCH=amd64    go test -v -c -o /dev/null

include rules.mk
