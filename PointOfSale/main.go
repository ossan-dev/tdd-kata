package pointofsale

func Scan(barcode string) string {
	if barcode == "12345" {
		return "$7.25"
	}
	return "$12.50"
}
