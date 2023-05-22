package main

import (
	"fmt"
	"time"

	"github.com/Akilan1999/p2p-rendering-computation/p2p/frp"
)

// Trigger the FRP implementation
func main() {
	natssh, err := EscapeNATSSH()
	if err != nil {
                fmt.Println(err)
		return
	}

	fmt.Println("ssh akilan@164.90.177.167 -p" + natssh)

	natMonitor, err := EscapeNATServerMonitor()
	if err != nil {
		return
	}

	fmt.Println("http://164.90.177.167:" + natMonitor)

	for {

	}
}

func EscapeNATSSH() (SSHport string, err error) {
	// Get free port from P2PRC server node
	serverPort, err := frp.GetFRPServerPort("http://164.90.177.167:8089")

	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)

	//port for the barrierKVM server
	SSHport, err = frp.StartFRPClientForServer("164.90.177.167", serverPort, "22", "8080")
	if err != nil {
		return
	}

	return
}

func EscapeNATServerMonitor() (MonitorPort string, err error) {
	// Get free port from P2PRC server node
	serverPort, err := frp.GetFRPServerPort("http://164.90.177.167:8089")

	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)

	//port for the barrierKVM server
	MonitorPort, err = frp.StartFRPClientForServer("164.90.177.167", serverPort, "19999", "8888")
	if err != nil {
		return
	}

	return
}
