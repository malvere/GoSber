package model

type StripResponse struct {
	Success bool   `json:"success,omitempty"`
	Total   string `json:"total,omitempty"`
	Items   []struct {
		Goods struct {
			GoodsID string `json:"goodsId,omitempty"`
			Title   string `json:"title,omitempty"`
			WebURL  string `json:"webUrl,omitempty"`
		} `json:"goods,omitempty"`
		Price        int     `json:"price,omitempty"`
		BonusPercent int     `json:"bonusPercent,omitempty"`
		BonusAmount  int     `json:"bonusAmount,omitempty"`
		Rating       float32 `json:"rating,omitempty"`
		ReviewCount  int     `json:"reviewCount,omitempty"`
		OfferCount   int     `json:"offerCount,omitempty"`
		FinalPrice   int     `json:"finalPrice,omitempty"`
		IsAvailable  bool    `json:"isAvailable,omitempty"`
	} `json:"items,omitempty"`
}
