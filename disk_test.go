package zdbparser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// test for ParseDisk function
func Test_ParseDisk(t *testing.T) {
	disk := ParseDisk(`type: 'disk'
                id: 0
                guid: 11804555628199413764
                path: '/dev/disk/by-id/wwn-0x5000cca242daae59-part1'
                devid: 'id1,sd@n5000cca242daae59/a'
                phys_path: '/scsi_vhci/disk@g5000cca242daae59:a'
                whole_disk: 1
                DTL: 192593
                create_txg: 4
                com.delphix:vdev_zap_leaf: 36`)

	require.EqualValues(t, "disk", disk.Type)
	assert.EqualValues(t, 0, disk.ID)
	assert.EqualValues(t, "11804555628199413764", disk.GUID)
	assert.EqualValues(t, "/dev/disk/by-id/wwn-0x5000cca242daae59-part1", disk.Path)
	assert.EqualValues(t, "id1,sd@n5000cca242daae59/a", disk.Devid)
	assert.EqualValues(t, "/scsi_vhci/disk@g5000cca242daae59:a", disk.PhysPath)
	assert.EqualValues(t, 1, disk.WholeDisk)
	assert.EqualValues(t, 192593, disk.DTL)
	assert.EqualValues(t, 4, disk.CreateTXG)
}
