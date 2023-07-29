package mdmon

import (
	"context"
	"fmt"
	"os/exec"

	"strings"

	"github.com/nicjohnson145/nota/internal/report"
	"github.com/nicjohnson145/nota/internal/util"
	"github.com/reugn/go-quartz/quartz"
	"github.com/rs/zerolog"
)

type JobConfig struct {
	Logger   zerolog.Logger
	Reporter report.Reporter
}

func NewJob(conf JobConfig) *Job {
	return &Job{
		log:      conf.Logger,
		reporter: conf.Reporter,
	}
}

var _ quartz.Job = (*Job)(nil)

type Job struct {
	log      zerolog.Logger
	reporter report.Reporter
}

type Disk struct {
	ID               string // 0
	Status           string // 1
	Name             string // 2
	State            string // 3
	PowerStatus      string // 4
	FailurePredicted bool   // 9
}

func (j *Job) Execute(_ context.Context) {
	j.log.Info().Msg("beginning disk check")

	output, err := j.executeCommand()
	if err != nil {
		j.log.Err(err).Msg("error executing command")
		j.reporter.Report(fmt.Sprintf("error executing command: %v", err))
		return
	}

	disks, err := j.parseCommandOutput(output)
	if err != nil {
		j.log.Err(err).Msg("error parsing command output")
		j.reporter.Report(fmt.Sprintf("error parsing command output: %v", err))
		return
	}

	if !j.disksOk(disks) {
		j.log.Error().Msg("some disks reported not ok")
		j.reporter.Report("disks reported not ok")
		return
	}

	j.log.Info().Msg("disks report OK")
}

func (j *Job) Description() string {
	return "Check the health of disks in the MD1000"
}

func (j *Job) Key() int {
	return util.Hash(fmt.Sprint("MD-MON"))
}

func (j *Job) executeCommand() (string, error) {
	var stdOut strings.Builder

	cmd := exec.Command("/bin/sh", "-c", "sudo /opt/dell/srvadmin/bin/omreport storage pdisk controller=0 -fmt ssv")
	cmd.Stdout = &stdOut

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running omreport command: %w", err)
	}

	return stdOut.String(), nil
}

func (j *Job) parseCommandOutput(input string) ([]Disk, error) {
	lines := strings.Split(input, "\n")

	start := -1
	for i, l := range lines {
		if strings.HasPrefix(l, "ID;Status;Name") {
			start = i + 1
			break
		}
	}

	if start == -1 {
		return nil, fmt.Errorf("unable to find header line in command ouput")
	}

	disks := []Disk{}

	for i := start; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, ";")
		disks = append(disks, Disk{
			ID:               parts[0],
			Status:           parts[1],
			Name:             parts[2],
			State:            parts[3],
			PowerStatus:      parts[4],
			FailurePredicted: parts[9] != "No",
		})
	}

	return disks, nil
}

func (j *Job) disksOk(disks []Disk) bool {
	ok := true
	for _, d := range disks {
		if d.Status != "Ok" {
			j.log.Error().Str("id", d.ID).Msg("disk not in status 'Ok'")
			ok = false
		}
		if d.State != "Online" && d.State != "Ready" {
			j.log.Error().Str("id", d.ID).Msg("disk state not Online/Ready")
			ok = false
		}
		if d.FailurePredicted {
			j.log.Error().Str("id", d.ID).Msg("disk has predicted failure")
			ok = false
		}
	}

	return ok
}
