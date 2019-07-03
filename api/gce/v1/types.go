package main

// GCEInstanceList GCE Instance List 構造体
type GCEInstanceList struct {
	Kind  string `json:"kind"`
	ID    string `json:"id"`
	Items []struct {
		Kind              string `json:"kind"`
		ID                string `json:"id"`
		CreationTimestamp string `json:"creationTimestamp"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		Tags              struct {
			Items       []string `json:"items"`
			Fingerprint string   `json:"fingerprint"`
		} `json:"tags"`
		MachineType       string `json:"machineType"`
		Status            string `json:"status"`
		Zone              string `json:"zone"`
		CanIPForward      bool   `json:"canIpForward"`
		NetworkInterfaces []struct {
			Kind          string `json:"kind"`
			Network       string `json:"network"`
			Subnetwork    string `json:"subnetwork"`
			NetworkIP     string `json:"networkIP"`
			Name          string `json:"name"`
			AccessConfigs []struct {
				Kind        string `json:"kind"`
				Type        string `json:"type"`
				Name        string `json:"name"`
				NatIP       string `json:"natIP"`
				NetworkTier string `json:"networkTier"`
			} `json:"accessConfigs"`
			Fingerprint string `json:"fingerprint"`
		} `json:"networkInterfaces"`
		Disks []struct {
			Kind            string   `json:"kind"`
			Type            string   `json:"type"`
			Mode            string   `json:"mode"`
			Source          string   `json:"source"`
			DeviceName      string   `json:"deviceName"`
			Index           int      `json:"index"`
			Boot            bool     `json:"boot"`
			AutoDelete      bool     `json:"autoDelete"`
			Licenses        []string `json:"licenses"`
			Interface       string   `json:"interface"`
			GuestOsFeatures []struct {
				Type string `json:"type"`
			} `json:"guestOsFeatures"`
		} `json:"disks"`
		Metadata struct {
			Kind        string `json:"kind"`
			Fingerprint string `json:"fingerprint"`
		} `json:"metadata"`
		ServiceAccounts []struct {
			Email  string   `json:"email"`
			Scopes []string `json:"scopes"`
		} `json:"serviceAccounts"`
		SelfLink   string `json:"selfLink"`
		Scheduling struct {
			OnHostMaintenance string `json:"onHostMaintenance"`
			AutomaticRestart  bool   `json:"automaticRestart"`
			Preemptible       bool   `json:"preemptible"`
		} `json:"scheduling"`
		CPUPlatform        string `json:"cpuPlatform"`
		LabelFingerprint   string `json:"labelFingerprint"`
		StartRestricted    bool   `json:"startRestricted"`
		DeletionProtection bool   `json:"deletionProtection"`
	} `json:"items"`
	SelfLink string `json:"selfLink"`
}
