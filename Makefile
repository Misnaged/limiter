APP:=bolter
APP_ENTRY_POINT:=./cmd/limiter.go

fire:
	MallocNanoZone=0 go run -race $(APP_ENTRY_POINT) serve
