package main

type SiriusConfig struct {
	AppID    string `json:"appId"`
	BaseID   string `json:"baseId"`
	Name     string `json:"name"`
	Selflink string `json:"selflink"`
	Integsvc struct {
		USPartnerAppRefreshComposite string `json:"USPartnerAppRefreshComposite"`
		USUpdateDeviceRefreshForCC   string `json:"USUpdateDeviceRefreshForCC"`
		USUpdateDeviceSATRefresh     string `json:"USUpdateDeviceSATRefresh"`
		USDemoConsumptionRules       string `json:"USDemoConsumptionRules"`
		VehicleDataRestService       string `json:"VehicleDataRestService"`
		ESBDeactivateRentalTrial     string `json:"ESBDeactivateRentalTrial"`
		USNewGetProperties           string `json:"USNewGetProperties"`
		USNewReportsGeneration       string `json:"USNewReportsGeneration"`
		USReportsGeneration          string `json:"USReportsGeneration"`
		DemoConsumptionRules         string `json:"DemoConsumptionRules"`
		DealerAppService7            string `json:"DealerAppService7"`
		DealerAppService6            string `json:"DealerAppService6"`
		DealerAppService5            string `json:"DealerAppService5"`
		DealerAppService4            string `json:"DealerAppService4"`
		DealerAppService1            string `json:"DealerAppService1"`
		DealerAppService0            string `json:"DealerAppService0"`
		DealerAppService2            string `json:"DealerAppService2"`
		DBChecking                   string `json:"DBChecking"`
		USESBRefresh                 string `json:"USESBRefresh"`
		DBSuccessUpdate              string `json:"DBSuccessUpdate"`
		DealerAppService3            string `json:"DealerAppService3"`
		USBlockListDevice            string `json:"USBlockListDevice"`
		PartnerAppRefresh            string `json:"PartnerAppRefresh"`
		VehicleData                  string `json:"VehicleData"`
		USGetServiceInfo             string `json:"USGetServiceInfo"`
		ReportsGeneration            string `json:"ReportsGeneration"`
	} `json:"integsvc"`
	Reportingsvc struct {
		Custom  string `json:"custom"`
		Session string `json:"session"`
	} `json:"reportingsvc"`
	ServicesMeta struct {
		USPartnerAppRefreshComposite struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USPartnerAppRefreshComposite"`
		USUpdateDeviceRefreshForCC struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USUpdateDeviceRefreshForCC"`
		USUpdateDeviceSATRefresh struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USUpdateDeviceSATRefresh"`
		USDemoConsumptionRules struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USDemoConsumptionRules"`
		VehicleDataRestService struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"VehicleDataRestService"`
		ESBDeactivateRentalTrial struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"ESBDeactivateRentalTrial"`
		USNewGetProperties struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USNewGetProperties"`
		USNewReportsGeneration struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USNewReportsGeneration"`
		USReportsGeneration struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USReportsGeneration"`
		DemoConsumptionRules struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DemoConsumptionRules"`
		DealerAppService7 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService7"`
		DealerAppService6 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService6"`
		DealerAppService5 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService5"`
		DealerAppService4 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService4"`
		DealerAppService1 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService1"`
		DealerAppService0 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService0"`
		DealerAppService2 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService2"`
		DBChecking struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DBChecking"`
		USESBRefresh struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USESBRefresh"`
		DBSuccessUpdate struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DBSuccessUpdate"`
		DealerAppService3 struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"DealerAppService3"`
		USBlockListDevice struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USBlockListDevice"`
		PartnerAppRefresh struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"PartnerAppRefresh"`
		VehicleData struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"VehicleData"`
		USGetServiceInfo struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"USGetServiceInfo"`
		ReportsGeneration struct {
			Version string `json:"version"`
			URL     string `json:"url"`
			Type    string `json:"type"`
		} `json:"ReportsGeneration"`
	} `json:"services_meta"`
}
