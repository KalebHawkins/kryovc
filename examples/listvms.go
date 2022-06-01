package examples

import (
	"fmt"
	"os"
)

const (
	username = "user"
	password = "pass"
)

// Get creates a new vcenter session and lists all virtual machines.
func ListVms() {
	// Create a session.
	if _, err := kvc.CreateSession(username, password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vms, err := kvc.ListVm()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range vms {
		fmt.Println(v)
	}
}
