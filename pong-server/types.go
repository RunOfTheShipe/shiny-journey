package main

type AircraftRequest struct {
	Now      float64    `json:"now"`
	Messages int        `json:"messages"`
	Aircraft []Aircraft `json:"aircraft"`
}

type Aircraft struct {
	Hex string `json:"hex"`
	// AltBaro int    `json:"alt_baro,omitempty"`
	// Version int    `json:"version,omitempty"`
	// NacP    int    `json:"nac_p,omitempty"`
	// Sil            int           `json:"sil,omitempty"`
	// SilType        string        `json:"sil_type,omitempty"`
	// Mlat           []interface{} `json:"mlat"`
	// Tisb           []interface{} `json:"tisb"`
	// Messages       int           `json:"messages"`
	// Seen           float64       `json:"seen"`
	// Rssi           float64       `json:"rssi"`
	// Flight         string        `json:"flight,omitempty"`
	// AltGeom        int           `json:"alt_geom,omitempty"`
	// Gs             float64       `json:"gs,omitempty"`
	// Track          float64       `json:"track,omitempty"`
	// BaroRate       int           `json:"baro_rate,omitempty"`
	// Squawk         string        `json:"squawk,omitempty"`
	// Emergency      string        `json:"emergency,omitempty"`
	// Category       string        `json:"category,omitempty"`
	// NavQnh         float64       `json:"nav_qnh,omitempty"`
	// NavAltitudeMcp int           `json:"nav_altitude_mcp,omitempty"`
	// NavModes       []string      `json:"nav_modes,omitempty"`
	// Lat            float64       `json:"lat,omitempty"`
	// Lon            float64       `json:"lon,omitempty"`
	// Nic            int           `json:"nic,omitempty"`
	// Rc             int           `json:"rc,omitempty"`
	// SeenPos        float64       `json:"seen_pos,omitempty"`
	// NicBaro        int           `json:"nic_baro,omitempty"`
	// NacV           int           `json:"nac_v,omitempty"`
	// Gva            int           `json:"gva,omitempty"`
	// Sda            int           `json:"sda,omitempty"`
	// GeomRate       int           `json:"geom_rate,omitempty"`
	// NavHeading     float64       `json:"nav_heading,omitempty"`
	// Type           string        `json:"type,omitempty"`
	// TrueHeading    float64       `json:"true_heading,omitempty"`
}
