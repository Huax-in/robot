#Using Walk
go get github.com/lxn/walk
Create Manifest test.manifest
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso

#build (无命令窗)
go build -ldflags="-H windowsgui"