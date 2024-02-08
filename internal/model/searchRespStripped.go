package model

type StripResponse struct {
	Success bool `json:"success"`
	Items   []struct {
		Goods struct {
			GoodsID string `json:"goodsId"`
			Title   string `json:"title"`
			WebURL  string `json:"webUrl"`
		} `json:"goods"`
		Price        int     `json:"price"`
		BonusPercent int     `json:"bonusPercent"`
		BonusAmount  int     `json:"bonusAmount"`
		Rating       float32 `json:"rating"`
		ReviewCount  int     `json:"reviewCount"`
		OfferCount   int     `json:"offerCount"`
		FinalPrice   int     `json:"finalPrice"`
		IsAvailable  bool    `json:"isAvailable"`
	} `json:"items"`
}
