package model

import "time"

type Response struct {
	Success bool `json:"success"`
	Meta    struct {
		Time       time.Time `json:"time"`
		TraceID    string    `json:"traceId"`
		RequestID  string    `json:"requestId"`
		AppVersion string    `json:"appVersion"`
	} `json:"meta"`
	Errors []any  `json:"errors"`
	Total  string `json:"total"`
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
	Items  []struct {
		Goods struct {
			GoodsID    string `json:"goodsId"`
			Title      string `json:"title"`
			TitleImage string `json:"titleImage"`
			Attributes []struct {
				Title string `json:"title"`
				Slug  string `json:"slug"`
				Value string `json:"value"`
				Group struct {
					Title string `json:"title"`
					Slug  string `json:"slug"`
				} `json:"group"`
				IsWebShort         bool   `json:"isWebShort"`
				IsWebListing       bool   `json:"isWebListing"`
				Sequence           int    `json:"sequence"`
				FeatureDescription string `json:"featureDescription"`
			} `json:"attributes"`
			WebURL string `json:"webUrl"`
			Slug   string `json:"slug"`
			Boxes  []struct {
				Box           string  `json:"box"`
				PackagingUnit string  `json:"packagingUnit"`
				Width         float64 `json:"width"`
				Height        float64 `json:"height"`
				Length        float64 `json:"length"`
				WeightUnit    string  `json:"weightUnit"`
				Weight        float64 `json:"weight"`
			} `json:"boxes"`
			CategoryID       string   `json:"categoryId"`
			Brand            string   `json:"brand"`
			ContentFlags     []string `json:"contentFlags"`
			ContentFlagsStr  []string `json:"contentFlagsStr"`
			Stocks           int      `json:"stocks"`
			PhotosCount      int      `json:"photosCount"`
			OffersCount      int      `json:"offersCount"`
			LogisticTags     []string `json:"logisticTags"`
			Images           []string `json:"images"`
			Documents        []any    `json:"documents"`
			Description      string   `json:"description"`
			Videos           []any    `json:"videos"`
			MainCollectionID string   `json:"mainCollectionId"`
			Package          struct {
				PackageType string `json:"packageType"`
				MinQuantity string `json:"minQuantity"`
				WeightUnit  string `json:"weightUnit"`
			} `json:"package"`
		} `json:"goods"`
		Price         int     `json:"price"`
		PriceFrom     int     `json:"priceFrom"`
		PriceTo       int     `json:"priceTo"`
		BonusPercent  int     `json:"bonusPercent"`
		BonusAmount   int     `json:"bonusAmount"`
		Rating        float32 `json:"rating"`
		ReviewCount   int     `json:"reviewCount"`
		OfferCount    int     `json:"offerCount"`
		FinalPrice    int     `json:"finalPrice"`
		FavoriteOffer struct {
			ID                    string `json:"id"`
			Price                 int    `json:"price"`
			Score                 int    `json:"score"`
			IsFavorite            bool   `json:"isFavorite"`
			MerchantID            string `json:"merchantId"`
			DeliveryPossibilities []struct {
				Code                 string `json:"code"`
				Date                 string `json:"date"`
				Amount               int    `json:"amount"`
				IsDefault            bool   `json:"isDefault"`
				MaxDeliveryDays      any    `json:"maxDeliveryDays"`
				AppliedPromotionTags []any  `json:"appliedPromotionTags"`
				IsDbm                bool   `json:"isDbm"`
				DeliveryPriceType    string `json:"deliveryPriceType"`
				DisplayName          string `json:"displayName"`
				DisplayDeliveryDate  string `json:"displayDeliveryDate"`
				DeliveryOptions      []any  `json:"deliveryOptions"`
			} `json:"deliveryPossibilities"`
			FinalPrice               int    `json:"finalPrice"`
			BonusPercent             int    `json:"bonusPercent"`
			BonusAmount              int    `json:"bonusAmount"`
			AvailableQuantity        int    `json:"availableQuantity"`
			BonusAmountFinalPrice    int    `json:"bonusAmountFinalPrice"`
			Discounts                []any  `json:"discounts"`
			PriceAdjustments         []any  `json:"priceAdjustments"`
			DeliveryDate             string `json:"deliveryDate"`
			PickupDate               string `json:"pickupDate"`
			MerchantOfferID          string `json:"merchantOfferId"`
			MerchantName             string `json:"merchantName"`
			MerchantLogoURL          string `json:"merchantLogoUrl"`
			MerchantURL              string `json:"merchantUrl"`
			MerchantSummaryRating    int    `json:"merchantSummaryRating"`
			IsBpgByMerchant          bool   `json:"isBpgByMerchant"`
			Nds                      int    `json:"nds"`
			OldPrice                 int    `json:"oldPrice"`
			OldPriceChangePercentage int    `json:"oldPriceChangePercentage"`
			MaxDeliveryDays          any    `json:"maxDeliveryDays"`
			BpgType                  string `json:"bpgType"`
			CreditPaymentAmount      int    `json:"creditPaymentAmount"`
			InstallmentPaymentAmount int    `json:"installmentPaymentAmount"`
			ShowMerchant             any    `json:"showMerchant"`
			SalesBlockInfo           any    `json:"salesBlockInfo"`
			DueDate                  string `json:"dueDate"`
			DueDateText              string `json:"dueDateText"`
			LocationID               string `json:"locationId"`
			SpasiboIsAvailable       bool   `json:"spasiboIsAvailable"`
			IsShowcase               bool   `json:"isShowcase"`
			LoyaltyPromotionFlags    []any  `json:"loyaltyPromotionFlags"`
			PricesPerMeasurement     []any  `json:"pricesPerMeasurement"`
			AvailablePaymentMethods  []any  `json:"availablePaymentMethods"`
			SuperPrice               int    `json:"superPrice"`
			WarehouseID              string `json:"warehouseId"`
			BnplPaymentParams        any    `json:"bnplPaymentParams"`
			InstallmentPaymentParams any    `json:"installmentPaymentParams"`
			BonusInfoGroups          []any  `json:"bonusInfoGroups"`
		} `json:"favoriteOffer"`
		RelatedItems          []any `json:"relatedItems"`
		ProductSelectors      []any `json:"productSelectors"`
		ExtraOptions          []any `json:"extraOptions"`
		DeliveryPossibilities []struct {
			Code                 string `json:"code"`
			Date                 string `json:"date"`
			Amount               int    `json:"amount"`
			IsDefault            bool   `json:"isDefault"`
			MaxDeliveryDays      any    `json:"maxDeliveryDays"`
			AppliedPromotionTags []any  `json:"appliedPromotionTags"`
			IsDbm                bool   `json:"isDbm"`
			DeliveryPriceType    string `json:"deliveryPriceType"`
			DisplayName          string `json:"displayName"`
			DisplayDeliveryDate  string `json:"displayDeliveryDate"`
			DeliveryOptions      []any  `json:"deliveryOptions"`
		} `json:"deliveryPossibilities"`
		Discounts                    []any    `json:"discounts"`
		ContentFlagsStr              []string `json:"contentFlagsStr"`
		ContentFlags                 []string `json:"contentFlags"`
		Badges                       []any    `json:"badges"`
		CrossedPrice                 int      `json:"crossedPrice"`
		CrossedPricePeriod           string   `json:"crossedPricePeriod"`
		LastPrice                    int      `json:"lastPrice"`
		IsAvailable                  bool     `json:"isAvailable"`
		CrossedPriceChangePercentage int      `json:"crossedPriceChangePercentage"`
		Collections                  []string `json:"collections"`
		HasOtherOffers               bool     `json:"hasOtherOffers"`
	} `json:"items"`
	Tags                []any `json:"tags"`
	Categories          []any `json:"categories"`
	MerchantIds         []any `json:"merchantIds"`
	PriceRange          any   `json:"priceRange"`
	Filters             []any `json:"filters"`
	SelectedFilterCount int   `json:"selectedFilterCount"`
	CollectionSelector  struct {
		Collections []struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			Slug             string `json:"slug"`
			WebURL           string `json:"webUrl"`
			IsSelected       bool   `json:"isSelected"`
			ChildCollections []struct {
				ID               string `json:"id"`
				Name             string `json:"name"`
				Slug             string `json:"slug"`
				WebURL           string `json:"webUrl"`
				IsSelected       bool   `json:"isSelected"`
				ChildCollections []struct {
					ID               string `json:"id"`
					Name             string `json:"name"`
					Slug             string `json:"slug"`
					WebURL           string `json:"webUrl"`
					IsSelected       bool   `json:"isSelected"`
					ChildCollections []any  `json:"childCollections"`
					IsActive         bool   `json:"isActive"`
					RelativeURL      string `json:"relativeUrl"`
					Type             string `json:"type"`
				} `json:"childCollections"`
				IsActive    bool   `json:"isActive"`
				RelativeURL string `json:"relativeUrl"`
				Type        string `json:"type"`
			} `json:"childCollections"`
			IsActive    bool   `json:"isActive"`
			RelativeURL string `json:"relativeUrl"`
			Type        string `json:"type"`
		} `json:"collections"`
	} `json:"collectionSelector"`
	Processor struct {
		Type            string `json:"type"`
		GoodsID         string `json:"goodsId"`
		CollectionID    string `json:"collectionId"`
		MenuNodeID      string `json:"menuNodeId"`
		MerchantID      string `json:"merchantId"`
		MerchantSlug    string `json:"merchantSlug"`
		URL             string `json:"url"`
		BrandSlug       string `json:"brandSlug"`
		MerchantLogoURL string `json:"merchantLogoUrl"`
		MenuNodeTitle   string `json:"menuNodeTitle"`
	} `json:"processor"`
	HasPlus18             bool     `json:"hasPlus18"`
	Navigation            any      `json:"navigation"`
	Flags                 []any    `json:"flags"`
	Keywords              []string `json:"keywords"`
	View                  string   `json:"view"`
	AllowedServiceSchemes []any    `json:"allowedServiceSchemes"`
	Sorting               []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"sorting"`
	ListingSize       int `json:"listingSize"`
	AssumedCollection struct {
		CollectionID   string `json:"collectionId"`
		ParentID       string `json:"parentId"`
		CollectionType string `json:"collectionType"`
		Title          string `json:"title"`
		IsDepartment   bool   `json:"isDepartment"`
		Hierarchy      []any  `json:"hierarchy"`
		URL            string `json:"url"`
		Images         struct {
			Mid10 string `json:"mid10"`
		} `json:"images"`
		Description             string `json:"description"`
		Slug                    string `json:"slug"`
		NavigationMode          string `json:"navigationMode"`
		AllowedServiceSchemes   []any  `json:"allowedServiceSchemes"`
		Code                    string `json:"code"`
		MainListingCollectionID string `json:"mainListingCollectionId"`
		Attributes              any    `json:"attributes"`
		Rating                  struct {
		} `json:"rating"`
		ShortTitle                       string `json:"shortTitle"`
		MainListingCollectionRelativeURL string `json:"mainListingCollectionRelativeUrl"`
		Name                             string `json:"name"`
		DisplayName                      string `json:"displayName"`
	} `json:"assumedCollection"`
	AlternativeAssumedCollections []any  `json:"alternativeAssumedCollections"`
	QueryChangesInfo              any    `json:"queryChangesInfo"`
	SearchTextContext             string `json:"searchTextContext"`
	GoodsURL                      string `json:"goodsURL"`
	Version                       string `json:"version"`
}
