package models

type Cabinet struct {
	ID             uint   `gorm:"primaryKey"`
	Department     string `json:"department"`
	Feature        string `json:"feature"`
	Number         string `json:"number"`
	Name           string `json:"name"`
}

type Equipment struct {
	ID                uint   `gorm:"primaryKey"`
	CabinetID         uint   `gorm:"index"`
	ItemName          string `json:"item_name"`
	StandardName      string `json:"standard_name"`
	QtyToOrder        int    `json:"qty_to_order"`
	Manufacturer      string `json:"manufacturer"`
	Model             string `json:"model"`
	Article           string `json:"article"`
	Specialist        string `json:"specialist"`
	TMCType           string `json:"tmc_type"`
}

type SpecItem struct {
	ItemName     string `json:"item_name"`
	StandardName string `json:"standard_name"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Article      string `json:"article"`
	TotalQty     int    `json:"total_qty"`
	Specialist   string `json:"specialist"`
	TMCType      string `json:"tmc_type"`
}
