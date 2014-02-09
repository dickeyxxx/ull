package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os/exec"
)

func rmdir(dir string) {
	cmd := exec.Command("rm", "-rf", dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		panic(err)
	}
	log.Println("Removed dir:", dir)
}

func cp(here string, there string) {
	cmd := exec.Command("cp", here, there)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		panic(err)
	}
	log.Println("Updated keys")
}

func gitClone(url string, path string) {
	cmd := exec.Command("git", "clone", url, path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		panic(err)
	}
	log.Println("Cloned repo", url, "to", path)
}

func watch() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	pubsub := redis.PubSubConn{c}
	pubsub.Subscribe("authorized_keys")
	for {
		switch v := pubsub.Receive().(type) {
		case redis.Message:
			rmdir("/var/ull/kvasir")
			gitClone("git@github.com:dickeyxxx/kvasir", "/var/ull/kvasir")
			cp("/var/ull/kvasir/authorized_keys", "authorized_keys")
		case error:
			panic(v)
		}
	}
}

func main() {
	watch()
}
