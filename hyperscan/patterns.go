package hyperscan

const (
	// FloatNumber for matching floating point numbers
	FloatNumber = `(?:` +
		`[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?.)`

	// IPv4Address for matching IPv4 address
	IPv4Address = `(?:` +
		`(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3}` +
		`(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))`

	// EmailAddress for matching email address
	EmailAddress = `(?:` +
		`^[A-Za-z0-9](([_\.\-]?[a-zA-Z0-9]+)*)@` +
		`([A-Za-z0-9]+)(([\.\-]?[a-zA-Z0-9]+)*)\.([A-Za-z]{2,})$)`

	// CreditCard for matching credit card number
	CreditCard = `(?:` +
		`4[0-9]{12}(?:[0-9]{3})?|` + // Visa
		`5[1-5][0-9]{14}|` + // MasterCard
		`3[47][0-9]{13}|` + // American Express
		`3(?:0[0-5]|[68][0-9])[0-9]{11}|` + // Diners Club
		`6(?:011|5[0-9]{2})[0-9]{12}|` + // Discover
		`(?:2131|1800|35\d{3})\d{11})` // JCB
)
