package pidfile_test

import (
	"testing"
  "."
  "os"
  "path"
)

func TestPid(t *testing.T) {

  pidFile := path.Join( os.TempDir() , "foo.pid")
  err := pidfile.SetPidfile(pidFile)
  if(err != nil){
		t.Fatal("setting pid filed, actual", err)
  }

  running, err := pidfile.IsPidfileRunning(pidFile)
  if(err != nil){
		t.Fatal("IsPidfileRunning Fialed:", err)
  }

  if(running != true){
		t.Fatal("Should be True:", running)
  }


  pidfile.ExitIfRunning(pidFile)
	t.Fatal("Should not make it here")



}

