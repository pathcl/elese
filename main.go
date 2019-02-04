// Retrieve summary property for all machines
// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
// Print summary per vm (see also: govc/vm/info.go)
//
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pathcl/elese/client"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("\n")
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	// start time
	defer timeTrack(time.Now(), "main")
	ctx := context.Background()

	// Connect and login to ESX or vCenter
	c, err := client.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Logout(ctx)
	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Fatal(err)
	}

	defer v.Destroy(ctx)

	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		log.Fatal(err)
	}

	for _, vm := range vms {
		fmt.Printf("%s: %s %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName, vm.Summary.Guest.IpAddress)
	}

	log.Printf(" ")
	log.Println("Total vms", len(vms))
}
