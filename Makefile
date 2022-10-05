APP:=limiter
APP_ENTRY_POINT:=./cmd/limiter.go

run:
	MallocNanoZone=0 go run -race $(APP_ENTRY_POINT) serve
