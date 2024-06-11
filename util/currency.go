package util

const (
	USD     = "USD"
	RUPEES  = "RUPEES"
	RUPEE   = "RUPEE"
	DOLLARS = "DOLLARS"
	EUROS   = "EUROS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, RUPEES, RUPEE, DOLLARS, EUROS:
		{
			return true
		}
		
	}
	return false
}
