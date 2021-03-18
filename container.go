package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run container.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("unknown")
	}
}

func run() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/home/rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"

	for _, ns := range []string{"pids", "cpu", "memory"} {
		nspath := fmt.Sprintf("%s/%s/asr/", cgroups, ns)
		os.Mkdir(nspath, 0755)

		switch ns {
		case "pids":
			//pid namespace
			must(ioutil.WriteFile(filepath.Join(nspath, "pids.max"), []byte("20"), 0700))
		case "cpu":
			//10% of all cpu time
			must(ioutil.WriteFile(filepath.Join(nspath, "cpu.cfs_period_us"), []byte("1000000"), 0700))
			must(ioutil.WriteFile(filepath.Join(nspath, "cpu.cfs_quota_us"), []byte("100000"), 0700))
		case "memory":
			//10MB
			must(ioutil.WriteFile(filepath.Join(nspath, "memory.limit_in_bytes"), []byte("10000000"), 0700))
			must(ioutil.WriteFile(filepath.Join(nspath, "memory.memsw.limit_in_bytes"), []byte("10000000"), 0700))
		}

		// Removes the new cgroup in place after the container exits
		must(ioutil.WriteFile(filepath.Join(nspath, "notify_on_release"), []byte("1"), 0700))
		must(ioutil.WriteFile(filepath.Join(nspath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
