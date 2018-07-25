package helpers

import (
	"fmt"
	"strings"

	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

var validADCPInputs = []string{
	"video1",
	"svideo1",
	"rgb1",
	"rgb2",
	"dvi1",
	"hdmi1",
	"hdmi2",
	"network",
	"usb_a",
	"usb_b",
	"hdbaset1",
	"option1",
}

func SetInput(address, port string, pooled bool) *nerr.E {
	log.L.Debugf("Setting input on %s to %s", address, port)

	validInput := false
	for _, input := range validADCPInputs {
		if strings.EqualFold(port, input) {
			validInput = true
			break
		}
	}

	if !validInput {
		return nerr.Create(fmt.Sprintf("error: %s is not a valid ADCP input.", port), "invalid port")
	}

	command := fmt.Sprintf("input \"%s\"", port)
	return sendCommand(command, address, pooled)
}

func GetInputStatus(address string, pooled bool) (status.Input, *nerr.E) {
	log.L.Debugf("Querying input status of %s", address)

	response, err := queryState("input ?", address, pooled)
	if err != nil {
		return status.Input{}, err.Add("Couldn't query input status")
	}

	status := status.Input{
		Input: strings.Trim(string(response), "\""),
	}
	return status, nil
}
