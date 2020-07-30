// Retrieve summary property for all machines
// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
// Print summary per vm (see also: govc/vm/info.go)
//
// TODO: Iterate multiple vcenter
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pathcl/elese/client"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

func main() {
	// start time

	var vcenter = []string{"https://somevcenter.domain.tld/sdk", "https://somevc.domain.tld/sdk"}

	// Connect and login to ESX or vCenter
	for _, j := range vcenter {
		ctx := context.Background()
		c, err := client.NewClient(ctx, j)
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
	}
}
