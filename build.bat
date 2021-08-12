@echo off
go generate
go build -ldflags "-H windowsgui" -o sprite_viewer.exe 
