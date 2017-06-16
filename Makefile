default: test

testdeps:
	@go get github.com/smartystreets/goconvey

test: testdeps
	@go test ./...


testrace: testdeps
	@go test ./... -race

testall: test testrace

bench:
	@go test ./... -run=NONE -bench=.
