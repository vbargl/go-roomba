package status

import (
	"reflect"
)

type Status struct {
	Audio              any `json:"audio"`
	BatPct             any `json:"batPct"`
	BatteryType        any `json:"batteryType"`
	Bbchg              any `json:"bbchg"`
	Bbchg3             any `json:"bbchg3"`
	Bbmssn             any `json:"bbmssn"`
	Bbnav              any `json:"bbnav"`
	Bbpanic            any `json:"bbpanic"`
	Bbpause            any `json:"bbpause"`
	Bbrstinfo          any `json:"bbrstinfo"`
	Bbrun              any `json:"bbrun"`
	Bbswitch           any `json:"bbswitch"`
	Bbsys              any `json:"bbsys"`
	Bin                any `json:"bin"`
	BinPause           any `json:"binPause"`
	BootloaderVer      any `json:"bootloaderVer"`
	Cap                any `json:"cap"`
	CarpetBoost        any `json:"carpetBoost"`
	CleanMissionStatus any `json:"cleanMissionStatus"`
	CleanSchedule      any `json:"cleanSchedule"`
	CloudEnv           any `json:"cloudEnv"`
	Country            any `json:"country"`
	Dock               any `json:"dock"`
	EcoCharge          any `json:"ecoCharge"`
	HardwareRev        any `json:"hardwareRev"`
	Langs              any `json:"langs"`
	Language           any `json:"language"`
	LastCommand        any `json:"lastCommand"`
	Localtimeoffset    any `json:"localtimeoffset"`
	Mac                any `json:"mac"`
	MapUploadAllowed   any `json:"mapUploadAllowed"`
	MobilityVer        any `json:"mobilityVer"`
	Name               any `json:"name"`
	NavSwVer           any `json:"navSwVer"`
	Netinfo            any `json:"netinfo"`
	NoAutoPasses       any `json:"noAutoPasses"`
	NoPP               any `json:"noPP"`
	OpenOnly           any `json:"openOnly"`
	Pose               any `json:"pose"`
	SchedHold          any `json:"schedHold"`
	Signal             any `json:"signal"`
	Sku                any `json:"sku"`
	SoftwareVer        any `json:"softwareVer"`
	SoundVer           any `json:"soundVer"`
	SvcEndpoints       any `json:"svcEndpoints"`
	Timezone           any `json:"timezone"`
	TwoPass            any `json:"twoPass"`
	Tz                 any `json:"tz"`
	UiSwVer            any `json:"uiSwVer"`
	UmiVer             any `json:"umiVer"`
	Utctime            any `json:"utctime"`
	VacHigh            any `json:"vacHigh"`
	Wifistat           any `json:"wifistat"`
	WifiSwVer          any `json:"wifiSwVer"`
	Wlcfg              any `json:"wlcfg"`
}

func (s *Status) IsAllValuesPresent() bool {
	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			return false
		}
	}

	return true
}
