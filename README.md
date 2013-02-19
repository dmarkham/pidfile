pidfile
==========

`go get github.com/dmarkham/pidfile`

```go
// if another process is running with the pid contained
// in your pid file your program will exit cleanly.
// This is useful for adding things into a crontab
// and knowing they will not fork bomb the server if something
// hangs or goes wrong.
pidfile.ExitIfRunning( "/tmp/mypid.pid")
```
A Proc::Pidfile like package for go.
Only works on Linux.
