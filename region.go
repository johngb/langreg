package langreg

import "errors"

// RegionCodeInfo returns the English regional name, and ccTLD that
// corresponds to the ISO 3166-1 alpha-2 region codes.
// Region codes should always be uppercase, and this is enforced.
// E.g. "US" is valid, but "us" is not.
func RegionCodeInfo(s string) (region string, err error) {
	// codes have to be two characters long
	if len(s) != 2 {
		return "", errors.New("ISO 3166-1 alpha-2 region codes must be 2 characters long")
	}
	switch s {
	case "AD":
		return "Andorra", nil
	case "AE":
		return "United Arab Emirates", nil
	case "AF":
		return "Afghanistan", nil
	case "AG":
		return "Antigua and Barbuda", nil
	case "AI":
		return "Anguilla", nil
	case "AL":
		return "Albania", nil
	case "AM":
		return "Armenia", nil
	case "AO":
		return "Angola", nil
	case "AQ":
		return "Antarctica", nil
	case "AR":
		return "Argentina", nil
	case "AS":
		return "American Samoa", nil
	case "AT":
		return "Austria", nil
	case "AU":
		return "Australia", nil
	case "AW":
		return "Aruba", nil
	case "AX":
		return "Aland Islands !Åland Islands", nil
	case "AZ":
		return "Azerbaijan", nil
	case "BA":
		return "Bosnia and Herzegovina", nil
	case "BB":
		return "Barbados", nil
	case "BD":
		return "Bangladesh", nil
	case "BE":
		return "Belgium", nil
	case "BF":
		return "Burkina Faso", nil
	case "BG":
		return "Bulgaria", nil
	case "BH":
		return "Bahrain", nil
	case "BI":
		return "Burundi", nil
	case "BJ":
		return "Benin", nil
	case "BL":
		return "Saint Barthélemy", nil
	case "BM":
		return "Bermuda", nil
	case "BN":
		return "Brunei Darussalam", nil
	case "BO":
		return "Bolivia, Plurinational State of", nil
	case "BQ":
		return "Bonaire, Sint Eustatius and Saba", nil
	case "BR":
		return "Brazil", nil
	case "BS":
		return "Bahamas", nil
	case "BT":
		return "Bhutan", nil
	case "BV":
		return "Bouvet Island", nil
	case "BW":
		return "Botswana", nil
	case "BY":
		return "Belarus", nil
	case "BZ":
		return "Belize", nil
	case "CA":
		return "Canada", nil
	case "CC":
		return "Cocos (Keeling) Islands", nil
	case "CD":
		return "Congo, the Democratic Republic of the", nil
	case "CF":
		return "Central African Republic", nil
	case "CG":
		return "Congo", nil
	case "CH":
		return "Switzerland", nil
	case "CI":
		return "Cote dIvoire !Côte dIvoire", nil
	case "CK":
		return "Cook Islands", nil
	case "CL":
		return "Chile", nil
	case "CM":
		return "Cameroon", nil
	case "CN":
		return "China", nil
	case "CO":
		return "Colombia", nil
	case "CR":
		return "Costa Rica", nil
	case "CU":
		return "Cuba", nil
	case "CV":
		return "Cabo Verde", nil
	case "CW":
		return "Curaçao", nil
	case "CX":
		return "Christmas Island", nil
	case "CY":
		return "Cyprus", nil
	case "CZ":
		return "Czech Republic", nil
	case "DE":
		return "Germany", nil
	case "DJ":
		return "Djibouti", nil
	case "DK":
		return "Denmark", nil
	case "DM":
		return "Dominica", nil
	case "DO":
		return "Dominican Republic", nil
	case "DZ":
		return "Algeria", nil
	case "EC":
		return "Ecuador", nil
	case "EE":
		return "Estonia", nil
	case "EG":
		return "Egypt", nil
	case "EH":
		return "Western Sahara", nil
	case "ER":
		return "Eritrea", nil
	case "ES":
		return "Spain", nil
	case "ET":
		return "Ethiopia", nil
	case "FI":
		return "Finland", nil
	case "FJ":
		return "Fiji", nil
	case "FK":
		return "Falkland Islands (Malvinas)", nil
	case "FM":
		return "Micronesia, Federated States of", nil
	case "FO":
		return "Faroe Islands", nil
	case "FR":
		return "France", nil
	case "GA":
		return "Gabon", nil
	case "GB":
		return "United Kingdom", nil
	case "GD":
		return "Grenada", nil
	case "GE":
		return "Georgia", nil
	case "GF":
		return "French Guiana", nil
	case "GG":
		return "Guernsey", nil
	case "GH":
		return "Ghana", nil
	case "GI":
		return "Gibraltar", nil
	case "GL":
		return "Greenland", nil
	case "GM":
		return "Gambia", nil
	case "GN":
		return "Guinea", nil
	case "GP":
		return "Guadeloupe", nil
	case "GQ":
		return "Equatorial Guinea", nil
	case "GR":
		return "Greece", nil
	case "GS":
		return "South Georgia and the South Sandwich Islands", nil
	case "GT":
		return "Guatemala", nil
	case "GU":
		return "Guam", nil
	case "GW":
		return "Guinea-Bissau", nil
	case "GY":
		return "Guyana", nil
	case "HK":
		return "Hong Kong", nil
	case "HM":
		return "Heard Island and McDonald Islands", nil
	case "HN":
		return "Honduras", nil
	case "HR":
		return "Croatia", nil
	case "HT":
		return "Haiti", nil
	case "HU":
		return "Hungary", nil
	case "ID":
		return "Indonesia", nil
	case "IE":
		return "Ireland", nil
	case "IL":
		return "Israel", nil
	case "IM":
		return "Isle of Man", nil
	case "IN":
		return "India", nil
	case "IO":
		return "British Indian Ocean Territory", nil
	case "IQ":
		return "Iraq", nil
	case "IR":
		return "Iran, Islamic Republic of", nil
	case "IS":
		return "Iceland", nil
	case "IT":
		return "Italy", nil
	case "JE":
		return "Jersey", nil
	case "JM":
		return "Jamaica", nil
	case "JO":
		return "Jordan", nil
	case "JP":
		return "Japan", nil
	case "KE":
		return "Kenya", nil
	case "KG":
		return "Kyrgyzstan", nil
	case "KH":
		return "Cambodia", nil
	case "KI":
		return "Kiribati", nil
	case "KM":
		return "Comoros", nil
	case "KN":
		return "Saint Kitts and Nevis", nil
	case "KP":
		return "Korea, Democratic Peoples Republic of", nil
	case "KR":
		return "Korea, Republic of", nil
	case "KW":
		return "Kuwait", nil
	case "KY":
		return "Cayman Islands", nil
	case "KZ":
		return "Kazakhstan", nil
	case "LA":
		return "Lao Peoples Democratic Republic", nil
	case "LB":
		return "Lebanon", nil
	case "LC":
		return "Saint Lucia", nil
	case "LI":
		return "Liechtenstein", nil
	case "LK":
		return "Sri Lanka", nil
	case "LR":
		return "Liberia", nil
	case "LS":
		return "Lesotho", nil
	case "LT":
		return "Lithuania", nil
	case "LU":
		return "Luxembourg", nil
	case "LV":
		return "Latvia", nil
	case "LY":
		return "Libya", nil
	case "MA":
		return "Morocco", nil
	case "MC":
		return "Monaco", nil
	case "MD":
		return "Moldova, Republic of", nil
	case "ME":
		return "Montenegro", nil
	case "MF":
		return "Saint Martin (French part)", nil
	case "MG":
		return "Madagascar", nil
	case "MH":
		return "Marshall Islands", nil
	case "MK":
		return "Macedonia, the former Yugoslav Republic of", nil
	case "ML":
		return "Mali", nil
	case "MM":
		return "Myanmar", nil
	case "MN":
		return "Mongolia", nil
	case "MO":
		return "Macao", nil
	case "MP":
		return "Northern Mariana Islands", nil
	case "MQ":
		return "Martinique", nil
	case "MR":
		return "Mauritania", nil
	case "MS":
		return "Montserrat", nil
	case "MT":
		return "Malta", nil
	case "MU":
		return "Mauritius", nil
	case "MV":
		return "Maldives", nil
	case "MW":
		return "Malawi", nil
	case "MX":
		return "Mexico", nil
	case "MY":
		return "Malaysia", nil
	case "MZ":
		return "Mozambique", nil
	case "NA":
		return "Namibia", nil
	case "NC":
		return "New Caledonia", nil
	case "NE":
		return "Niger", nil
	case "NF":
		return "Norfolk Island", nil
	case "NG":
		return "Nigeria", nil
	case "NI":
		return "Nicaragua", nil
	case "NL":
		return "Netherlands", nil
	case "NO":
		return "Norway", nil
	case "NP":
		return "Nepal", nil
	case "NR":
		return "Nauru", nil
	case "NU":
		return "Niue", nil
	case "NZ":
		return "New Zealand", nil
	case "OM":
		return "Oman", nil
	case "PA":
		return "Panama", nil
	case "PE":
		return "Peru", nil
	case "PF":
		return "French Polynesia", nil
	case "PG":
		return "Papua New Guinea", nil
	case "PH":
		return "Philippines", nil
	case "PK":
		return "Pakistan", nil
	case "PL":
		return "Poland", nil
	case "PM":
		return "Saint Pierre and Miquelon", nil
	case "PN":
		return "Pitcairn", nil
	case "PR":
		return "Puerto Rico", nil
	case "PS":
		return "Palestine, State of", nil
	case "PT":
		return "Portugal", nil
	case "PW":
		return "Palau", nil
	case "PY":
		return "Paraguay", nil
	case "QA":
		return "Qatar", nil
	case "RE":
		return "Reunion !Réunion", nil
	case "RO":
		return "Romania", nil
	case "RS":
		return "Serbia", nil
	case "RU":
		return "Russian Federation", nil
	case "RW":
		return "Rwanda", nil
	case "SA":
		return "Saudi Arabia", nil
	case "SB":
		return "Solomon Islands", nil
	case "SC":
		return "Seychelles", nil
	case "SD":
		return "Sudan", nil
	case "SE":
		return "Sweden", nil
	case "SG":
		return "Singapore", nil
	case "SH":
		return "Saint Helena, Ascension and Tristan da Cunha", nil
	case "SI":
		return "Slovenia", nil
	case "SJ":
		return "Svalbard and Jan Mayen", nil
	case "SK":
		return "Slovakia", nil
	case "SL":
		return "Sierra Leone", nil
	case "SM":
		return "San Marino", nil
	case "SN":
		return "Senegal", nil
	case "SO":
		return "Somalia", nil
	case "SR":
		return "Suriname", nil
	case "SS":
		return "South Sudan", nil
	case "ST":
		return "Sao Tome and Principe", nil
	case "SV":
		return "El Salvador", nil
	case "SX":
		return "Sint Maarten (Dutch part)", nil
	case "SY":
		return "Syrian Arab Republic", nil
	case "SZ":
		return "Swaziland", nil
	case "TC":
		return "Turks and Caicos Islands", nil
	case "TD":
		return "Chad", nil
	case "TF":
		return "French Southern Territories", nil
	case "TG":
		return "Togo", nil
	case "TH":
		return "Thailand", nil
	case "TJ":
		return "Tajikistan", nil
	case "TK":
		return "Tokelau", nil
	case "TL":
		return "Timor-Leste", nil
	case "TM":
		return "Turkmenistan", nil
	case "TN":
		return "Tunisia", nil
	case "TO":
		return "Tonga", nil
	case "TR":
		return "Turkey", nil
	case "TT":
		return "Trinidad and Tobago", nil
	case "TV":
		return "Tuvalu", nil
	case "TW":
		return "Taiwan, Province of China", nil
	case "TZ":
		return "Tanzania, United Republic of", nil
	case "UA":
		return "Ukraine", nil
	case "UG":
		return "Uganda", nil
	case "UM":
		return "United States Minor Outlying Islands", nil
	case "US":
		return "United States", nil
	case "UY":
		return "Uruguay", nil
	case "UZ":
		return "Uzbekistan", nil
	case "VA":
		return "Holy See (Vatican City State)", nil
	case "VC":
		return "Saint Vincent and the Grenadines", nil
	case "VE":
		return "Venezuela, Bolivarian Republic of", nil
	case "VG":
		return "Virgin Islands, British", nil
	case "VI":
		return "Virgin Islands, U.S.", nil
	case "VN":
		return "Viet Nam", nil
	case "VU":
		return "Vanuatu", nil
	case "WF":
		return "Wallis and Futuna", nil
	case "WS":
		return "Samoa", nil
	case "YE":
		return "Yemen", nil
	case "YT":
		return "Mayotte", nil
	case "ZA":
		return "South Africa", nil
	case "ZM":
		return "Zambia", nil
	case "ZW":
		return "Zimbabwe", nil
	}
	return "", errors.New("\"%s\" is not a valid ISO 3166-1 alpha-2 region code")
}

// IsValidRegionCode returns true if s is a valid ISO1366-1_alpa-2 region code.
func IsValidRegionCode(s string) bool {
	_, err := RegionCodeInfo(s)
	if err != nil {
		return false
	}
	return true
}

// RegionName returns the English name of the ISO1366-1_alpa-2 region code s.
func RegionName(s string) (string, error) {
	name, err := RegionCodeInfo(s)
	return name, err
}
