package chargePoint

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/models"
)

// BootNotificationRequest represents a CP -> CSMS boot message
type BootNotificationReq struct {
	ChargePointModel        models.CiString20Type  `json:"chargePointModel"`
	ChargePointVendor       models.CiString20Type  `json:"chargePointVendor"`
	ChargeBoxSerialNumber   *models.CiString25Type `json:"chargeBoxSerialNumber,omitempty"`
	ChargePointSerialNumber *models.CiString25Type `json:"chargePointSerialNumber,omitempty"`
	FirmwareVersion         *models.CiString50Type `json:"firmwareVersion,omitempty"`
	Iccid                   *models.CiString20Type `json:"iccid,omitempty"`
	Imsi                    *models.CiString20Type `json:"imsi,omitempty"`
	MeterSerialNumber       *models.CiString25Type `json:"meterSerialNumber,omitempty"`
	MeterType               *models.CiString25Type `json:"meterType,omitempty"`
}

// BootNotificationConf represents a CSMS -> CP response
type BootNotificationConf struct {
	CurrentTime time.Time                `json:"currentTime"`
	Interval    int                      `json:"interval"`
	Status      enums.RegistrationStatus `json:"status"`
}
