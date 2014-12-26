package langreg

import "errors"

// RegionCodeInfo returns the English regional name, and ccTLD that
// corresponds to the ISO 3166-1 alpha-2 region codes.
// Region codes should always be uppercase, and this is enforced.
// E.g. "US" is valid, but "us" is not.
func RegionCodeInfo(s string) (region, ccTLD string, err error) {
	// codes have to be two characters long
	if len(s) != 2 {
		return "", "", errors.New("ISO 3166-1 alpha-2 region codes must be 2 characters long")
	}
	switch s {
	case "AD":
		return "Andorra", ".ad", nil
	case "AE":
		return "United Arab Emirates", ".ae", nil
	case "AF":
		return "Afghanistan", ".af", nil
	case "AG":
		return "Antigua and Barbuda", ".ag", nil
	case "AI":
		return "Anguilla", ".ai", nil
	case "AL":
		return "Albania", ".al", nil
	case "AM":
		return "Armenia", ".am", nil
	case "AO":
		return "Angola", ".ao", nil
	case "AQ":
		return "Antarctica", ".aq", nil
	case "AR":
		return "Argentina", ".ar", nil
	case "AS":
		return "American Samoa", ".as", nil
	case "AT":
		return "Austria", ".at", nil
	case "AU":
		return "Australia", ".au", nil
	case "AW":
		return "Aruba", ".aw", nil
	case "AX":
		return "Aland Islands !Åland Islands", ".ax", nil
	case "AZ":
		return "Azerbaijan", ".az", nil
	case "BA":
		return "Bosnia and Herzegovina", ".ba", nil
	case "BB":
		return "Barbados", ".bb", nil
	case "BD":
		return "Bangladesh", ".bd", nil
	case "BE":
		return "Belgium", ".be", nil
	case "BF":
		return "Burkina Faso", ".bf", nil
	case "BG":
		return "Bulgaria", ".bg", nil
	case "BH":
		return "Bahrain", ".bh", nil
	case "BI":
		return "Burundi", ".bi", nil
	case "BJ":
		return "Benin", ".bj", nil
	case "BL":
		return "Saint Barthélemy", ".bl", nil
	case "BM":
		return "Bermuda", ".bm", nil
	case "BN":
		return "Brunei Darussalam", ".bn", nil
	case "BO":
		return "Bolivia, Plurinational State of", ".bo", nil
	case "BQ":
		return "Bonaire, Sint Eustatius and Saba", ".bq", nil
	case "BR":
		return "Brazil", ".br", nil
	case "BS":
		return "Bahamas", ".bs", nil
	case "BT":
		return "Bhutan", ".bt", nil
	case "BV":
		return "Bouvet Island", ".bv", nil
	case "BW":
		return "Botswana", ".bw", nil
	case "BY":
		return "Belarus", ".by", nil
	case "BZ":
		return "Belize", ".bz", nil
	case "CA":
		return "Canada", ".ca", nil
	case "CC":
		return "Cocos (Keeling) Islands", ".cc", nil
	case "CD":
		return "Congo, the Democratic Republic of the", ".cd", nil
	case "CF":
		return "Central African Republic", ".cf", nil
	case "CG":
		return "Congo", ".cg", nil
	case "CH":
		return "Switzerland", ".ch", nil
	case "CI":
		return "Cote dIvoire !Côte dIvoire", ".ci", nil
	case "CK":
		return "Cook Islands", ".ck", nil
	case "CL":
		return "Chile", ".cl", nil
	case "CM":
		return "Cameroon", ".cm", nil
	case "CN":
		return "China", ".cn", nil
	case "CO":
		return "Colombia", ".co", nil
	case "CR":
		return "Costa Rica", ".cr", nil
	case "CU":
		return "Cuba", ".cu", nil
	case "CV":
		return "Cabo Verde", ".cv", nil
	case "CW":
		return "Curaçao", ".cw", nil
	case "CX":
		return "Christmas Island", ".cx", nil
	case "CY":
		return "Cyprus", ".cy", nil
	case "CZ":
		return "Czech Republic", ".cz", nil
	case "DE":
		return "Germany", ".de", nil
	case "DJ":
		return "Djibouti", ".dj", nil
	case "DK":
		return "Denmark", ".dk", nil
	case "DM":
		return "Dominica", ".dm", nil
	case "DO":
		return "Dominican Republic", ".do", nil
	case "DZ":
		return "Algeria", ".dz", nil
	case "EC":
		return "Ecuador", ".ec", nil
	case "EE":
		return "Estonia", ".ee", nil
	case "EG":
		return "Egypt", ".eg", nil
	case "EH":
		return "Western Sahara", ".eh", nil
	case "ER":
		return "Eritrea", ".er", nil
	case "ES":
		return "Spain", ".es", nil
	case "ET":
		return "Ethiopia", ".et", nil
	case "FI":
		return "Finland", ".fi", nil
	case "FJ":
		return "Fiji", ".fj", nil
	case "FK":
		return "Falkland Islands (Malvinas)", ".fk", nil
	case "FM":
		return "Micronesia, Federated States of", ".fm", nil
	case "FO":
		return "Faroe Islands", ".fo", nil
	case "FR":
		return "France", ".fr", nil
	case "GA":
		return "Gabon", ".ga", nil
	case "GB":
		return "United Kingdom", ".uk", nil
	case "GD":
		return "Grenada", ".gd", nil
	case "GE":
		return "Georgia", ".ge", nil
	case "GF":
		return "French Guiana", ".gf", nil
	case "GG":
		return "Guernsey", ".gg", nil
	case "GH":
		return "Ghana", ".gh", nil
	case "GI":
		return "Gibraltar", ".gi", nil
	case "GL":
		return "Greenland", ".gl", nil
	case "GM":
		return "Gambia", ".gm", nil
	case "GN":
		return "Guinea", ".gn", nil
	case "GP":
		return "Guadeloupe", ".gp", nil
	case "GQ":
		return "Equatorial Guinea", ".gq", nil
	case "GR":
		return "Greece", ".gr", nil
	case "GS":
		return "South Georgia and the South Sandwich Islands", ".gs", nil
	case "GT":
		return "Guatemala", ".gt", nil
	case "GU":
		return "Guam", ".gu", nil
	case "GW":
		return "Guinea-Bissau", ".gw", nil
	case "GY":
		return "Guyana", ".gy", nil
	case "HK":
		return "Hong Kong", ".hk", nil
	case "HM":
		return "Heard Island and McDonald Islands", ".hm", nil
	case "HN":
		return "Honduras", ".hn", nil
	case "HR":
		return "Croatia", ".hr", nil
	case "HT":
		return "Haiti", ".ht", nil
	case "HU":
		return "Hungary", ".hu", nil
	case "ID":
		return "Indonesia", ".id", nil
	case "IE":
		return "Ireland", ".ie", nil
	case "IL":
		return "Israel", ".il", nil
	case "IM":
		return "Isle of Man", ".im", nil
	case "IN":
		return "India", ".in", nil
	case "IO":
		return "British Indian Ocean Territory", ".io", nil
	case "IQ":
		return "Iraq", ".iq", nil
	case "IR":
		return "Iran, Islamic Republic of", ".ir", nil
	case "IS":
		return "Iceland", ".is", nil
	case "IT":
		return "Italy", ".it", nil
	case "JE":
		return "Jersey", ".je", nil
	case "JM":
		return "Jamaica", ".jm", nil
	case "JO":
		return "Jordan", ".jo", nil
	case "JP":
		return "Japan", ".jp", nil
	case "KE":
		return "Kenya", ".ke", nil
	case "KG":
		return "Kyrgyzstan", ".kg", nil
	case "KH":
		return "Cambodia", ".kh", nil
	case "KI":
		return "Kiribati", ".ki", nil
	case "KM":
		return "Comoros", ".km", nil
	case "KN":
		return "Saint Kitts and Nevis", ".kn", nil
	case "KP":
		return "Korea, Democratic Peoples Republic of", ".kp", nil
	case "KR":
		return "Korea, Republic of", ".kr", nil
	case "KW":
		return "Kuwait", ".kw", nil
	case "KY":
		return "Cayman Islands", ".ky", nil
	case "KZ":
		return "Kazakhstan", ".kz", nil
	case "LA":
		return "Lao Peoples Democratic Republic", ".la", nil
	case "LB":
		return "Lebanon", ".lb", nil
	case "LC":
		return "Saint Lucia", ".lc", nil
	case "LI":
		return "Liechtenstein", ".li", nil
	case "LK":
		return "Sri Lanka", ".lk", nil
	case "LR":
		return "Liberia", ".lr", nil
	case "LS":
		return "Lesotho", ".ls", nil
	case "LT":
		return "Lithuania", ".lt", nil
	case "LU":
		return "Luxembourg", ".lu", nil
	case "LV":
		return "Latvia", ".lv", nil
	case "LY":
		return "Libya", ".ly", nil
	case "MA":
		return "Morocco", ".ma", nil
	case "MC":
		return "Monaco", ".mc", nil
	case "MD":
		return "Moldova, Republic of", ".md", nil
	case "ME":
		return "Montenegro", ".me", nil
	case "MF":
		return "Saint Martin (French part)", ".mf", nil
	case "MG":
		return "Madagascar", ".mg", nil
	case "MH":
		return "Marshall Islands", ".mh", nil
	case "MK":
		return "Macedonia, the former Yugoslav Republic of", ".mk", nil
	case "ML":
		return "Mali", ".ml", nil
	case "MM":
		return "Myanmar", ".mm", nil
	case "MN":
		return "Mongolia", ".mn", nil
	case "MO":
		return "Macao", ".mo", nil
	case "MP":
		return "Northern Mariana Islands", ".mp", nil
	case "MQ":
		return "Martinique", ".mq", nil
	case "MR":
		return "Mauritania", ".mr", nil
	case "MS":
		return "Montserrat", ".ms", nil
	case "MT":
		return "Malta", ".mt", nil
	case "MU":
		return "Mauritius", ".mu", nil
	case "MV":
		return "Maldives", ".mv", nil
	case "MW":
		return "Malawi", ".mw", nil
	case "MX":
		return "Mexico", ".mx", nil
	case "MY":
		return "Malaysia", ".my", nil
	case "MZ":
		return "Mozambique", ".mz", nil
	case "NA":
		return "Namibia", ".na", nil
	case "NC":
		return "New Caledonia", ".nc", nil
	case "NE":
		return "Niger", ".ne", nil
	case "NF":
		return "Norfolk Island", ".nf", nil
	case "NG":
		return "Nigeria", ".ng", nil
	case "NI":
		return "Nicaragua", ".ni", nil
	case "NL":
		return "Netherlands", ".nl", nil
	case "NO":
		return "Norway", ".no", nil
	case "NP":
		return "Nepal", ".np", nil
	case "NR":
		return "Nauru", ".nr", nil
	case "NU":
		return "Niue", ".nu", nil
	case "NZ":
		return "New Zealand", ".nz", nil
	case "OM":
		return "Oman", ".om", nil
	case "PA":
		return "Panama", ".pa", nil
	case "PE":
		return "Peru", ".pe", nil
	case "PF":
		return "French Polynesia", ".pf", nil
	case "PG":
		return "Papua New Guinea", ".pg", nil
	case "PH":
		return "Philippines", ".ph", nil
	case "PK":
		return "Pakistan", ".pk", nil
	case "PL":
		return "Poland", ".pl", nil
	case "PM":
		return "Saint Pierre and Miquelon", ".pm", nil
	case "PN":
		return "Pitcairn", ".pn", nil
	case "PR":
		return "Puerto Rico", ".pr", nil
	case "PS":
		return "Palestine, State of", ".ps", nil
	case "PT":
		return "Portugal", ".pt", nil
	case "PW":
		return "Palau", ".pw", nil
	case "PY":
		return "Paraguay", ".py", nil
	case "QA":
		return "Qatar", ".qa", nil
	case "RE":
		return "Reunion !Réunion", ".re", nil
	case "RO":
		return "Romania", ".ro", nil
	case "RS":
		return "Serbia", ".rs", nil
	case "RU":
		return "Russian Federation", ".ru", nil
	case "RW":
		return "Rwanda", ".rw", nil
	case "SA":
		return "Saudi Arabia", ".sa", nil
	case "SB":
		return "Solomon Islands", ".sb", nil
	case "SC":
		return "Seychelles", ".sc", nil
	case "SD":
		return "Sudan", ".sd", nil
	case "SE":
		return "Sweden", ".se", nil
	case "SG":
		return "Singapore", ".sg", nil
	case "SH":
		return "Saint Helena, Ascension and Tristan da Cunha", ".sh", nil
	case "SI":
		return "Slovenia", ".si", nil
	case "SJ":
		return "Svalbard and Jan Mayen", ".sj", nil
	case "SK":
		return "Slovakia", ".sk", nil
	case "SL":
		return "Sierra Leone", ".sl", nil
	case "SM":
		return "San Marino", ".sm", nil
	case "SN":
		return "Senegal", ".sn", nil
	case "SO":
		return "Somalia", ".so", nil
	case "SR":
		return "Suriname", ".sr", nil
	case "SS":
		return "South Sudan", ".ss", nil
	case "ST":
		return "Sao Tome and Principe", ".st", nil
	case "SV":
		return "El Salvador", ".sv", nil
	case "SX":
		return "Sint Maarten (Dutch part)", ".sx", nil
	case "SY":
		return "Syrian Arab Republic", ".sy", nil
	case "SZ":
		return "Swaziland", ".sz", nil
	case "TC":
		return "Turks and Caicos Islands", ".tc", nil
	case "TD":
		return "Chad", ".td", nil
	case "TF":
		return "French Southern Territories", ".tf", nil
	case "TG":
		return "Togo", ".tg", nil
	case "TH":
		return "Thailand", ".th", nil
	case "TJ":
		return "Tajikistan", ".tj", nil
	case "TK":
		return "Tokelau", ".tk", nil
	case "TL":
		return "Timor-Leste", ".tl", nil
	case "TM":
		return "Turkmenistan", ".tm", nil
	case "TN":
		return "Tunisia", ".tn", nil
	case "TO":
		return "Tonga", ".to", nil
	case "TR":
		return "Turkey", ".tr", nil
	case "TT":
		return "Trinidad and Tobago", ".tt", nil
	case "TV":
		return "Tuvalu", ".tv", nil
	case "TW":
		return "Taiwan, Province of China", ".tw", nil
	case "TZ":
		return "Tanzania, United Republic of", ".tz", nil
	case "UA":
		return "Ukraine", ".ua", nil
	case "UG":
		return "Uganda", ".ug", nil
	case "UM":
		return "United States Minor Outlying Islands", ".um", nil
	case "US":
		return "United States", ".us", nil
	case "UY":
		return "Uruguay", ".uy", nil
	case "UZ":
		return "Uzbekistan", ".uz", nil
	case "VA":
		return "Holy See (Vatican City State)", ".va", nil
	case "VC":
		return "Saint Vincent and the Grenadines", ".vc", nil
	case "VE":
		return "Venezuela, Bolivarian Republic of", ".ve", nil
	case "VG":
		return "Virgin Islands, British", ".vg", nil
	case "VI":
		return "Virgin Islands, U.S.", ".vi", nil
	case "VN":
		return "Viet Nam", ".vn", nil
	case "VU":
		return "Vanuatu", ".vu", nil
	case "WF":
		return "Wallis and Futuna", ".wf", nil
	case "WS":
		return "Samoa", ".ws", nil
	case "YE":
		return "Yemen", ".ye", nil
	case "YT":
		return "Mayotte", ".yt", nil
	case "ZA":
		return "South Africa", ".za", nil
	case "ZM":
		return "Zambia", ".zm", nil
	case "ZW":
		return "Zimbabwe", ".zw", nil
	}
	return "", "", errors.New("\"%s\" is not a valid ISO 3166-1 alpha-2 region code")
}

// IsValidRegionCode returns true if s is a valid ISO1366-1_alpa-2 region code.
func IsValidRegionCode(s string) bool {
	_, _, err := RegionCodeInfo(s)
	if err != nil {
		return false
	}
	return true
}

// RegionName returns the English name of the ISO1366-1_alpa-2 region code s.
func RegionName(s string) (string, error) {
	name, _, err := RegionCodeInfo(s)
	return name, err
}

// RegionCCTLD returns the name ccTLD  of the ISO1366-1_alpa-2 region code s.
func RegionCCTLD(s string) (string, error) {
	_, ccTLD, err := RegionCodeInfo(s)
	return ccTLD, err
}
