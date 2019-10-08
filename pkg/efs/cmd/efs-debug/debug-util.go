package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/amazon-ecs-agent/pkg/efs"
)

func main() {
	log.Println("Starting")
	defer log.Println("Stopping")
	tls := flag.Bool("tls", false, "use tls")
	ro := flag.Bool("ro", false, "readonly")
	t := flag.Int("t", 10, "number of seconds to delay unmount")
	net := flag.Int("netns", 0, "the netns pid to swap to")
	flag.Parse()
	src := flag.Arg(0)
	mnt := flag.Arg(1)

	if !efs.HasEFSSupport() {
		log.Fatal("No efs bin found")
	}

	if path, err := exec.LookPath("mount"); err != nil {
		log.Fatal("error looking for mount: %s", err)
	} else {
		log.Printf("mount path is %s", path)
	}

	infos, _ := ioutil.ReadDir("/mnt")
	str := strings.Builder{}
	for i, info := range infos {
		if i > 0 {
			str.WriteString(", ")
		}
		str.WriteString(info.Name())
	}

	log.Printf("ls: %s", str.String())

	log.Printf("src: %s, mnt: %s", src, mnt)
	log.Printf("tls: %t, ro: %t", *tls, *ro)

	if src == "" || mnt == "" {
		log.Fatal("source and mount must both be set")
	}
	efsThing := &efs.EFSMounter{
		Filesystem:        src,
		TransitEncryption: *tls,
		ReadOnly:          *ro,
		LocalTarget:       mnt,
		NetNSPid:          *net,
	}

	var err error
	err = efsThing.Mount()
	if err != nil {
		goto handleError
	}
	log.Println("successfully mounted")
	time.Sleep(time.Duration(*t) * time.Second)
	err = efsThing.Unmount()
	if err != nil {
		goto handleError
	}
	log.Printf("successfully unmounted")

handleError:
	if err != nil {
		log.Printf("Error: %s", err)
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Printf(string(exiterr.Stderr))
			os.Exit(1)
		}
	}
}
