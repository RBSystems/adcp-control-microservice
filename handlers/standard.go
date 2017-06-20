package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/byuoitav/adcp-control-microservice/helpers"
	"github.com/labstack/echo"
)

func SetVolume(context echo.Context) error {

	address := context.Param("address")
	volumeLevel := context.Param("level")

	level, err := strconv.Atoi(volumeLevel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid volume level %s. Must be in range 0-100. %s", volumeLevel, err.Error()))
	}

	err = helpers.SetVolume(address, level)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, "ok")
}

func PowerOn(context echo.Context) error {
	address := context.Param("address")

	err := helpers.PowerOn(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	time.Sleep(3 * time.Second)
	return context.JSON(http.StatusOK, "ok")

}

func PowerStandby(context echo.Context) error {
	address := context.Param("address")

	err := helpers.PowerStandby(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	time.Sleep(2 * time.Second)
	return context.JSON(http.StatusOK, "ok")

}

func VolumeLevel(context echo.Context) error {

	address := context.Param("address")

	level, err := helpers.GetVolumeLevel(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, level)
}

func Mute(context echo.Context) error {

	address := context.Param("address")

	err := helpers.SetMute(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "ok")
}

func UnMute(context echo.Context) error {

	address := context.Param("address")

	err := helpers.SetMute(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "ok")
}

func MuteStatus(context echo.Context) error {

	address := context.Param("address")

	status, err := helpers.GetMuteStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func PowerStatus(context echo.Context) error {

	address := context.Param("address")

	status, err := helpers.GetPowerStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}
