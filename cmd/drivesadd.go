package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thepwagner/archivist/index"
	archivist "github.com/thepwagner/archivist/proto"
)

var driveAddCmd = &cobra.Command{
	Use:   "add [device]",
	Short: "Add drive",
	Args:  cobra.MinimumNArgs(1),
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		device := args[0]
		newDrive, err := index.NewDrive(device)
		if err != nil {
			return err
		}
		logrus.WithFields(logrus.Fields{
			"dev":    device,
			"model":  newDrive.GetModelNumber(),
			"serial": newDrive.GetSerialNumber(),
		}).Debug("Parsed drive")

		drives := index.NewDrives(idx)
		for _, drive := range drives {
			if drive.GetSerialNumber() == newDrive.GetSerialNumber() {
				return fmt.Errorf("drive already exists: %s", drive.GetId())
			}
		}

		idx.Drives = append(idx.Drives, newDrive)
		return nil
	}),
}

func init() {
	drivesCmd.AddCommand(driveAddCmd)
}
