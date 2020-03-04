package domain

type StoreAvailability struct {
	Status string `json:"status"`
	Items  []struct {
		Store struct {
			RetailPrice             float64    `json:"retailPrice"`
			SupplyStatusDescription string `json:"supplyStatusDescription"`
			Stock                   struct {
				OnHand float64 `json:"onHand"`
			} `json:"stock"`
		} `json:"store"`
		Sku         int    `json:"sku"`
		Description string `json:"description"`
		Images      []struct {
			URI string `json:"uri"`
		} `json:"images"`
		SubcategoryID int    `json:"subcategoryId"`
		Display       bool   `json:"display"`
		Status        string `json:"status"`
	} `json:"items"`
	Store struct {
		StoreCode int     `json:"storeCode"`
		Name      string  `json:"name"`
		Telephone string  `json:"telephone"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"store"`
}

type Result struct {
	Product   string
	Available float64
	NotFound  bool
}