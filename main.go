package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"./webhook"
)

var (
	listenPort = ":8080"
	secret     = "metadata"
	// commandStr = []string{"--git-dir=/Users/andy/Desktop/github/ContainerStation/.git", "pull"}

	deployTrigger = runDeployer()
)

func main() {

	webhook.SetSecret([]byte(secret))

	http.HandleFunc("/", webhook.HandlePush(func(ev *webhook.Event) {

		push := ev.PushEvent()
		log.Printf("push=", push)
		if push == nil {
			return
		}
		log.Printf("push: verified=%v %#v", ev.Verified, push)

		resultChan := make(chan error)
		select {
		case deployTrigger <- resultChan:
			err := <-resultChan
			if err != nil {
				log.Println("Deploy error:", err)
				return
			}

		}
	}))

	log.Fatal(http.ListenAndServe(listenPort, nil))
}

func runDeployer() chan chan error {
	trigger := make(chan chan error)
	go func() {
		for resultChan := range trigger {
			resultChan <- performDeploy()
		}
	}()
	return trigger
}

func performDeploy() error {
	// arv := commandStr
	// c := exec.Command("git", arv...)
	c := exec.Command("./pull.sh")
	d, err := c.CombinedOutput()
	if err != nil {
		//fmt.Println(err)
		return fmt.Errorf("%s\n---\n%s", err.Error(), d)
	}
	fmt.Println(string(d))
	return nil
}
