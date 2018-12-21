package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
	"github.com/fatih/color"
)

func PowerOn(address string, pooled bool) *nerr.E {
	log.L.Infof("Setting power of %v to on", address)
	command := "power \"on\""

	return sendCommand(command, address, pooled)
}

func PowerStandby(address string, pooled bool) *nerr.E {
	log.L.Infof("Seting power of %v to off", address)
	command := "power \"off\""

	return sendCommand(command, address, pooled)
}

func GetPower(address string, pooled bool) (status.Power, *nerr.E) {

	log.L.Infof("%s", color.HiCyanString("[helpers] querying power state of %v", address))

	response, err := queryState("power_status ?", address, pooled)
	if err != nil {
		return status.Power{}, err
	}

	var status status.Power
	responseString := string(response)

	if strings.Contains(responseString, "standby") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "startup") {
		status.Power = "on"
	} else if strings.Contains(responseString, "on") {
		status.Power = "on"
	} else if strings.Contains(responseString, "cooling1") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "cooling2") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "saving_standby") {
		status.Power = "standby"
	}

	return status, nil
}
