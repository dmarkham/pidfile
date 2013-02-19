package pidfile

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

func init() {

	if runtime.GOOS != "linux" {
		panic("pidfile: Your OS doesnt work yet. Sry..")
	}

}

func IsPidfileRunning(file string) (running bool, err error) {
	f, err := os.Open(file)
	if err != nil {
		return false, fmt.Errorf("Can not Open (%s): %s", file, err)
	}

	buf := make([]byte, 30)
	read, err := f.Read(buf)
	if err != nil {
		return false, fmt.Errorf("Can not Read (%s): %s", file, err)
	}

	pid, err := strconv.ParseInt(string(buf[0:read]), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Can not Parse (%s): %s", string(buf[0:read]), err)
	}

	_, err = os.Stat(path.Join("/proc", strconv.Itoa(int(pid))))
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}

	panic("Should not Reach")
}

// Sets a pid file without checking anything
// will over write a file if it exists
func SetPidfile(file string) (err error) {

	pid := os.Getpid()
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Can not create (%s): %s", file, err)
	}
	_, err = f.WriteString(strconv.Itoa(pid))
	if err != nil {
		return fmt.Errorf("Can not write (%s): %s", file, err)
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("Can not close (%s): %s", file, err)
	}

	return nil

}
