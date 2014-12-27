package langreg

import (
	"errors"
	"fmt"
)

// RegionCodeInfo returns the English regional name, and ccTLD that
// corresponds to the ISO 3166-1 alpha-2 region codes.
// Region codes should always be uppercase, and this is enforced.
// E.g. "US" is valid, but "us" is not.
func RegionCodeInfo(s string) (region string, err error) {

	// codes have to be two characters long
	if len(s) != 2 {
		return "",
			errors.New("ISO 3166-1 alpha-2 region codes must be 2 characters long")
	}
	switch s[0] {
	case 'A':
		switch s[1] {
		case 'D':
			return "Andorra", nil
		case 'E':
			return "United Arab Emirates", nil
		case 'F':
			return "Afghanistan", nil
		case 'G':
			return "Antigua and Barbuda", nil
		case 'I':
			return "Anguilla", nil
		case 'L':
			return "Albania", nil
		case 'M':
			return "Armenia", nil
		case 'O':
			return "Angola", nil
		case 'Q':
			return "Antarctica", nil
		case 'R':
			return "Argentina", nil
		case 'S':
			return "American Samoa", nil
		case 'T':
			return "Austria", nil
		case 'U':
			return "Australia", nil
		case 'W':
			return "Aruba", nil
		case 'X':
			return "Aland Islands !Åland Islands", nil
		case 'Z':
			return "Azerbaijan", nil
		}
	case 'B':
		switch s[1] {
		case 'A':
			return "Bosnia and Herzegovina", nil
		case 'B':
			return "Barbados", nil
		case 'D':
			return "Bangladesh", nil
		case 'E':
			return "Belgium", nil
		case 'F':
			return "Burkina Faso", nil
		case 'G':
			return "Bulgaria", nil
		case 'H':
			return "Bahrain", nil
		case 'I':
			return "Burundi", nil
		case 'J':
			return "Benin", nil
		case 'L':
			return "Saint Barthélemy", nil
		case 'M':
			return "Bermuda", nil
		case 'N':
			return "Brunei Darussalam", nil
		case 'O':
			return "Bolivia, Plurinational State of", nil
		case 'Q':
			return "Bonaire, Sint Eustatius and Saba", nil
		case 'R':
			return "Brazil", nil
		case 'S':
			return "Bahamas", nil
		case 'T':
			return "Bhutan", nil
		case 'V':
			return "Bouvet Island", nil
		case 'W':
			return "Botswana", nil
		case 'Y':
			return "Belarus", nil
		case 'Z':
			return "Belize", nil
		}
	case 'C':
		switch s[1] {
		case 'A':
			return "Canada", nil
		case 'C':
			return "Cocos (Keeling) Islands", nil
		case 'D':
			return "Congo, the Democratic Republic of the", nil
		case 'F':
			return "Central African Republic", nil
		case 'G':
			return "Congo", nil
		case 'H':
			return "Switzerland", nil
		case 'I':
			return "Cote dIvoire !Côte dIvoire", nil
		case 'K':
			return "Cook Islands", nil
		case 'L':
			return "Chile", nil
		case 'M':
			return "Cameroon", nil
		case 'N':
			return "China", nil
		case 'O':
			return "Colombia", nil
		case 'R':
			return "Costa Rica", nil
		case 'U':
			return "Cuba", nil
		case 'V':
			return "Cabo Verde", nil
		case 'W':
			return "Curaçao", nil
		case 'X':
			return "Christmas Island", nil
		case 'Y':
			return "Cyprus", nil
		case 'Z':
			return "Czech Republic", nil
		}
	case 'D':
		switch s[1] {
		case 'E':
			return "Germany", nil
		case 'J':
			return "Djibouti", nil
		case 'K':
			return "Denmark", nil
		case 'M':
			return "Dominica", nil
		case 'O':
			return "Dominican Republic", nil
		case 'Z':
			return "Algeria", nil
		}
	case 'E':
		switch s[1] {
		case 'C':
			return "Ecuador", nil
		case 'E':
			return "Estonia", nil
		case 'G':
			return "Egypt", nil
		case 'H':
			return "Western Sahara", nil
		case 'R':
			return "Eritrea", nil
		case 'S':
			return "Spain", nil
		case 'T':
			return "Ethiopia", nil
		}
	case 'F':
		switch s[1] {
		case 'I':
			return "Finland", nil
		case 'J':
			return "Fiji", nil
		case 'K':
			return "Falkland Islands (Malvinas)", nil
		case 'M':
			return "Micronesia, Federated States of", nil
		case 'O':
			return "Faroe Islands", nil
		case 'R':
			return "France", nil
		}
	case 'G':
		switch s[1] {
		case 'A':
			return "Gabon", nil
		case 'B':
			return "United Kingdom", nil
		case 'D':
			return "Grenada", nil
		case 'E':
			return "Georgia", nil
		case 'F':
			return "French Guiana", nil
		case 'G':
			return "Guernsey", nil
		case 'H':
			return "Ghana", nil
		case 'I':
			return "Gibraltar", nil
		case 'L':
			return "Greenland", nil
		case 'M':
			return "Gambia", nil
		case 'N':
			return "Guinea", nil
		case 'P':
			return "Guadeloupe", nil
		case 'Q':
			return "Equatorial Guinea", nil
		case 'R':
			return "Greece", nil
		case 'S':
			return "South Georgia and the South Sandwich Islands", nil
		case 'T':
			return "Guatemala", nil
		case 'U':
			return "Guam", nil
		case 'W':
			return "Guinea-Bissau", nil
		case 'Y':
			return "Guyana", nil
		}
	case 'H':
		switch s[1] {
		case 'K':
			return "Hong Kong", nil
		case 'M':
			return "Heard Island and McDonald Islands", nil
		case 'N':
			return "Honduras", nil
		case 'R':
			return "Croatia", nil
		case 'T':
			return "Haiti", nil
		case 'U':
			return "Hungary", nil
		}
	case 'I':
		switch s[1] {
		case 'D':
			return "Indonesia", nil
		case 'E':
			return "Ireland", nil
		case 'L':
			return "Israel", nil
		case 'M':
			return "Isle of Man", nil
		case 'N':
			return "India", nil
		case 'O':
			return "British Indian Ocean Territory", nil
		case 'Q':
			return "Iraq", nil
		case 'R':
			return "Iran, Islamic Republic of", nil
		case 'S':
			return "Iceland", nil
		case 'T':
			return "Italy", nil
		}
	case 'J':
		switch s[1] {
		case 'E':
			return "Jersey", nil
		case 'M':
			return "Jamaica", nil
		case 'O':
			return "Jordan", nil
		case 'P':
			return "Japan", nil
		}
	case 'K':
		switch s[1] {
		case 'E':
			return "Kenya", nil
		case 'G':
			return "Kyrgyzstan", nil
		case 'H':
			return "Cambodia", nil
		case 'I':
			return "Kiribati", nil
		case 'M':
			return "Comoros", nil
		case 'N':
			return "Saint Kitts and Nevis", nil
		case 'P':
			return "Korea, Democratic Peoples Republic of", nil
		case 'R':
			return "Korea, Republic of", nil
		case 'W':
			return "Kuwait", nil
		case 'Y':
			return "Cayman Islands", nil
		case 'Z':
			return "Kazakhstan", nil
		}
	case 'L':
		switch s[1] {
		case 'A':
			return "Lao Peoples Democratic Republic", nil
		case 'B':
			return "Lebanon", nil
		case 'C':
			return "Saint Lucia", nil
		case 'I':
			return "Liechtenstein", nil
		case 'K':
			return "Sri Lanka", nil
		case 'R':
			return "Liberia", nil
		case 'S':
			return "Lesotho", nil
		case 'T':
			return "Lithuania", nil
		case 'U':
			return "Luxembourg", nil
		case 'V':
			return "Latvia", nil
		case 'Y':
			return "Libya", nil
		}
	case 'M':
		switch s[1] {
		case 'A':
			return "Morocco", nil
		case 'C':
			return "Monaco", nil
		case 'D':
			return "Moldova, Republic of", nil
		case 'E':
			return "Montenegro", nil
		case 'F':
			return "Saint Martin (French part)", nil
		case 'G':
			return "Madagascar", nil
		case 'H':
			return "Marshall Islands", nil
		case 'K':
			return "Macedonia, the former Yugoslav Republic of", nil
		case 'L':
			return "Mali", nil
		case 'M':
			return "Myanmar", nil
		case 'N':
			return "Mongolia", nil
		case 'O':
			return "Macao", nil
		case 'P':
			return "Northern Mariana Islands", nil
		case 'Q':
			return "Martinique", nil
		case 'R':
			return "Mauritania", nil
		case 'S':
			return "Montserrat", nil
		case 'T':
			return "Malta", nil
		case 'U':
			return "Mauritius", nil
		case 'V':
			return "Maldives", nil
		case 'W':
			return "Malawi", nil
		case 'X':
			return "Mexico", nil
		case 'Y':
			return "Malaysia", nil
		case 'Z':
			return "Mozambique", nil
		}
	case 'N':
		switch s[1] {
		case 'A':
			return "Namibia", nil
		case 'C':
			return "New Caledonia", nil
		case 'E':
			return "Niger", nil
		case 'F':
			return "Norfolk Island", nil
		case 'G':
			return "Nigeria", nil
		case 'I':
			return "Nicaragua", nil
		case 'L':
			return "Netherlands", nil
		case 'O':
			return "Norway", nil
		case 'P':
			return "Nepal", nil
		case 'R':
			return "Nauru", nil
		case 'U':
			return "Niue", nil
		case 'Z':
			return "New Zealand", nil
		}
	case 'O':
		switch s[1] {
		case 'M':
			return "Oman", nil
		}
	case 'P':
		switch s[1] {
		case 'A':
			return "Panama", nil
		case 'E':
			return "Peru", nil
		case 'F':
			return "French Polynesia", nil
		case 'G':
			return "Papua New Guinea", nil
		case 'H':
			return "Philippines", nil
		case 'K':
			return "Pakistan", nil
		case 'L':
			return "Poland", nil
		case 'M':
			return "Saint Pierre and Miquelon", nil
		case 'N':
			return "Pitcairn", nil
		case 'R':
			return "Puerto Rico", nil
		case 'S':
			return "Palestine, State of", nil
		case 'T':
			return "Portugal", nil
		case 'W':
			return "Palau", nil
		case 'Y':
			return "Paraguay", nil
		}
	case 'Q':
		switch s[1] {
		case 'A':
			return "Qatar", nil
		}
	case 'R':
		switch s[1] {
		case 'E':
			return "Reunion !Réunion", nil
		case 'O':
			return "Romania", nil
		case 'S':
			return "Serbia", nil
		case 'U':
			return "Russian Federation", nil
		case 'W':
			return "Rwanda", nil
		}
	case 'S':
		switch s[1] {
		case 'A':
			return "Saudi Arabia", nil
		case 'B':
			return "Solomon Islands", nil
		case 'C':
			return "Seychelles", nil
		case 'D':
			return "Sudan", nil
		case 'E':
			return "Sweden", nil
		case 'G':
			return "Singapore", nil
		case 'H':
			return "Saint Helena, Ascension and Tristan da Cunha", nil
		case 'I':
			return "Slovenia", nil
		case 'J':
			return "Svalbard and Jan Mayen", nil
		case 'K':
			return "Slovakia", nil
		case 'L':
			return "Sierra Leone", nil
		case 'M':
			return "San Marino", nil
		case 'N':
			return "Senegal", nil
		case 'O':
			return "Somalia", nil
		case 'R':
			return "Suriname", nil
		case 'S':
			return "South Sudan", nil
		case 'T':
			return "Sao Tome and Principe", nil
		case 'V':
			return "El Salvador", nil
		case 'X':
			return "Sint Maarten (Dutch part)", nil
		case 'Y':
			return "Syrian Arab Republic", nil
		case 'Z':
			return "Swaziland", nil
		}
	case 'T':
		switch s[1] {
		case 'C':
			return "Turks and Caicos Islands", nil
		case 'D':
			return "Chad", nil
		case 'F':
			return "French Southern Territories", nil
		case 'G':
			return "Togo", nil
		case 'H':
			return "Thailand", nil
		case 'J':
			return "Tajikistan", nil
		case 'K':
			return "Tokelau", nil
		case 'L':
			return "Timor-Leste", nil
		case 'M':
			return "Turkmenistan", nil
		case 'N':
			return "Tunisia", nil
		case 'O':
			return "Tonga", nil
		case 'R':
			return "Turkey", nil
		case 'T':
			return "Trinidad and Tobago", nil
		case 'V':
			return "Tuvalu", nil
		case 'W':
			return "Taiwan, Province of China", nil
		case 'Z':
			return "Tanzania, United Republic of", nil
		}
	case 'U':
		switch s[1] {
		case 'A':
			return "Ukraine", nil
		case 'G':
			return "Uganda", nil
		case 'M':
			return "United States Minor Outlying Islands", nil
		case 'S':
			return "United States", nil
		case 'Y':
			return "Uruguay", nil
		case 'Z':
			return "Uzbekistan", nil
		}
	case 'V':
		switch s[1] {
		case 'A':
			return "Holy See (Vatican City State)", nil
		case 'C':
			return "Saint Vincent and the Grenadines", nil
		case 'E':
			return "Venezuela, Bolivarian Republic of", nil
		case 'G':
			return "Virgin Islands, British", nil
		case 'I':
			return "Virgin Islands, U.S.", nil
		case 'N':
			return "Viet Nam", nil
		case 'U':
			return "Vanuatu", nil
		}
	case 'W':
		switch s[1] {
		case 'F':
			return "Wallis and Futuna", nil
		case 'S':
			return "Samoa", nil
		}
	case 'Y':
		switch s[1] {
		case 'E':
			return "Yemen", nil
		case 'T':
			return "Mayotte", nil
		}
	case 'Z':
		switch s[1] {
		case 'A':
			return "South Africa", nil
		case 'M':
			return "Zambia", nil
		case 'W':
			return "Zimbabwe", nil
		}
	}

	return "",
		fmt.Errorf("\"%s\" is not a valid ISO 3166-1 alpha-2 region code", s)
}
