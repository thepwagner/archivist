package index

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
	archivist "github.com/thepwagner/archivist/proto"
)

type DriveID string

type Drives map[DriveID]*archivist.Drive

func NewDrives(idx *archivist.Index) Drives {
	d := make(Drives)
	for _, drive := range idx.Drives {
		driveID := DriveID(drive.GetId())
		d[driveID] = drive
	}
	return d
}

func NewDrive(device string) (*archivist.Drive, error) {
	cmd := exec.Command("smartctl", "-a", device)
	out, err := cmd.CombinedOutput()
	if smartctlError(err) {
		return nil, fmt.Errorf("executing smartctl: %w", err)
	}
	return ParseSmartctl(string(out))
}

func smartctlError(err error) bool {
	if err == nil {
		return false
	}
	var exitError *exec.ExitError
	if !errors.As(err, &exitError) {
		return true

	}
	return exitError.ExitCode() != 4
}

var (
	deviceModelRe     = regexp.MustCompile("Device Model:[ ]*(.*)")
	modelNumberRe     = regexp.MustCompile("Model Number:[ ]*(.*)")
	serialNumberRe    = regexp.MustCompile("Serial Number:[ ]*(.*)")
	userCapacityRe    = regexp.MustCompile("User Capacity:[ ]*([0-9,]*)")
	powerCycleCountRe = regexp.MustCompile("Power_Cycle_Count[^\n]*([0-9]+)")
	powerCyclesRe     = regexp.MustCompile("Power Cycles:[ ]*([0-9,]*)")
	powerOnHoursRe    = regexp.MustCompile("Power On Hours:[ ]*([0-9,]*)")
	powerOnHoursRawRe = regexp.MustCompile("Power_On_Hours[^\n]* ([0-9]+) \\(")
)

func ParseSmartctl(smartctl string) (*archivist.Drive, error) {
	var d archivist.Drive
	modelNumberMatch := deviceModelRe.FindStringSubmatch(smartctl)
	if len(modelNumberMatch) > 0 {
		d.ModelNumber = modelNumberMatch[1]
	} else {
		modelNumberMatch := modelNumberRe.FindStringSubmatch(smartctl)
		if len(modelNumberMatch) > 0 {
			d.ModelNumber = modelNumberMatch[1]
		}
	}

	serialNumberMatch := serialNumberRe.FindStringSubmatch(smartctl)
	if len(serialNumberMatch) > 0 {
		d.SerialNumber = serialNumberMatch[1]
	}
	userCapacityMatch := userCapacityRe.FindStringSubmatch(smartctl)
	if len(userCapacityMatch) > 0 {
		if capacity, err := strconv.Atoi(strings.ReplaceAll(userCapacityMatch[1], ",", "")); err == nil {
			d.Capacity = uint64(capacity)
		}
	}
	powerCycleMatch := powerCycleCountRe.FindStringSubmatch(smartctl)
	if len(powerCycleMatch) > 0 {
		if cycles, err := strconv.Atoi(powerCycleMatch[1]); err == nil {
			d.PowerCycleCount = uint64(cycles)
		}
	} else {
		powerCycleMatch := powerCyclesRe.FindStringSubmatch(smartctl)
		if len(powerCycleMatch) > 0 {
			if cycles, err := strconv.Atoi(strings.ReplaceAll(powerCycleMatch[1], ",", "")); err == nil {
				d.PowerCycleCount = uint64(cycles)
			}
		}
	}
	powerOnHoursMatch := powerOnHoursRawRe.FindStringSubmatch(smartctl)
	if len(powerOnHoursMatch) > 0 {
		if hours, err := strconv.Atoi(powerOnHoursMatch[1]); err == nil {
			d.PowerOnHours = uint64(hours)
		}
	} else {
		powerOnHoursMatch := powerOnHoursRe.FindStringSubmatch(smartctl)
		if len(powerOnHoursMatch) > 0 {
			if hours, err := strconv.Atoi(strings.ReplaceAll(powerOnHoursMatch[1], ",", "")); err == nil {
				d.PowerOnHours = uint64(hours)
			}
		}
	}
	d.Id = uuid.NewV4().String()
	return &d, nil
}
