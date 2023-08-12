package zdbparser

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_parseDisksFromZdbOutput(t *testing.T) {
	file, err := os.Open("fixtures/zdb.in")
	require.NoError(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	disks := parseDisksFromZdbOutput(scanner)
	require.EqualValues(t, 23, len(disks))
	assert.EqualValues(t, "disk", disks[0].Type)
	assert.EqualValues(t, 0, disks[0].ID)
	assert.EqualValues(t, "11804555628199413764", disks[0].GUID)
	assert.EqualValues(t, "/dev/disk/by-id/wwn-0x5000cca242daae59-part1", disks[0].Path)
	assert.EqualValues(t, "id1,sd@n5000cca242daae59/a", disks[0].Devid)
	assert.EqualValues(t, "/scsi_vhci/disk@g5000cca242daae59:a", disks[0].PhysPath)
	assert.EqualValues(t, 1, disks[0].WholeDisk)
	assert.EqualValues(t, 192593, disks[0].DTL)
	assert.EqualValues(t, 4, disks[0].CreateTXG)

	assert.EqualValues(t, "disk", disks[1].Type)
	assert.EqualValues(t, 1, disks[1].ID)
	assert.EqualValues(t, "14811255447288370791", disks[1].GUID)
	assert.EqualValues(t, "/dev/disk/by-id/wwn-0x5000cca242da3d8c-part1", disks[1].Path)
}
