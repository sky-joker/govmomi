/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import (
	"github.com/vmware/govmomi/vim25/tasks"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualDiskManager struct {
	c *Client
}

// CopyVirtualDisk copies a virtual disk, performing conversions as specified in the spec.
func (m VirtualDiskManager) CopyVirtualDisk(c *Client,
	sourceName string, sourceDatacenter *Datacenter,
	destName string, destDatacenter *Datacenter,
	destSpec *types.VirtualDiskSpec, force bool) error {

	req := types.CopyVirtualDisk_Task{
		This:       *c.ServiceContent.VirtualDiskManager,
		SourceName: sourceName,
		DestName:   destName,
		DestSpec:   destSpec,
		Force:      force,
	}

	if sourceDatacenter != nil {
		ref := sourceDatacenter.Reference()
		req.SourceDatacenter = &ref
	}

	if destDatacenter != nil {
		ref := destDatacenter.Reference()
		req.DestDatacenter = &ref
	}

	task, err := tasks.CopyVirtualDisk(c, &req)
	if err != nil {
		return err
	}

	return m.c.waitForTask(task)
}
