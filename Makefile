
all:
	go build -o cat cat.go
	go build -o echo echo.go
	go build -o exit exit.go
	go build -o flase false.go
	go build -o logname logname.go
	go build -o mv mv.go
	go build -o pwd pwd.go
	go build -o rm rm.go
	go build -o sleep sleep.go
	go build -o uptime uptime.go
	go build -o whoami whoami.go

gobuild:
	go build -o gocat cat.go
	go build -o goecho echo.go
	go build -o goexit exit.go
	go build -o gofalse false.go
	go build -o gologname logname.go
	go build -o gomv mv.go
	go build -o gopwd pwd.go
	go build -o gorm rm.go
	go build -o gosleep sleep.go
	go build -o gouptime uptime.go
	go build -o gowhoami whoami.go

fmt:
	go fmt cat.go
	go fmt echo.go
	go fmt exit.go
	go fmt false.go
	go fmt logname.go
	go fmt mv.go
	go fmt pwd.go
	go fmt rm.go
	go fmt sleep.go
	go fmt uptime.go
	go fmt whoami.go

clean:
	rm -f cat echo exit false logname mv pwd rm sleep uptime whoami
