package mdmon

import (
	"testing"
	"github.com/stretchr/testify/require"
	"os"
)

func TestParseCommandOutput(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		content, err := os.ReadFile("./testdata/omreport-sample-output")
		require.NoError(t, err)

		job := Job{}
		disks, err := job.parseCommandOutput(string(content))
		require.NoError(t, err)

		require.Equal(
			t,
			[]Disk{
				{
					ID: "0:0:0",
					Status: "Ok",
					Name: "Physical Disk 0:0:0",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:1",
					Status: "Ok",
					Name: "Physical Disk 0:0:1",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:2",
					Status: "Ok",
					Name: "Physical Disk 0:0:2",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:3",
					Status: "Ok",
					Name: "Physical Disk 0:0:3",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:4",
					Status: "Ok",
					Name: "Physical Disk 0:0:4",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:5",
					Status: "Ok",
					Name: "Physical Disk 0:0:5",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:6",
					Status: "Ok",
					Name: "Physical Disk 0:0:6",
					State: "Ready",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:7",
					Status: "Ok",
					Name: "Physical Disk 0:0:7",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:8",
					Status: "Ok",
					Name: "Physical Disk 0:0:8",
					State: "Ready",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:11",
					Status: "Ok",
					Name: "Physical Disk 0:0:11",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:12",
					Status: "Ok",
					Name: "Physical Disk 0:0:12",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:13",
					Status: "Ok",
					Name: "Physical Disk 0:0:13",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
				{
					ID: "0:0:14",
					Status: "Ok",
					Name: "Physical Disk 0:0:14",
					State: "Online",
					PowerStatus: "Spun Up",
					FailurePredicted: false,
				},
			},
			disks,
		)
	})
}

func TestDisksOk(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		content, err := os.ReadFile("./testdata/omreport-sample-output")
		require.NoError(t, err)

		job := Job{}
		disks, err := job.parseCommandOutput(string(content))
		require.NoError(t, err)

		require.True(t, job.disksOk(disks))
	})
}
