package langreg

import (
	"errors"
	"fmt"
)

// LangCodeInfo returns the English and native language in it's script for a
// given string, and an error if any.  If there are more than one official
// names for the language (either English or native), they are separated by a
// semi-colon (;)
// Language codes should always be lowercase, and this is enforced.
func LangCodeInfo(s string) (english, native string, err error) {

	// codes have to be two characters long
	if len(s) != 2 {
		return "", "", errors.New("ISO 639-1 language codes must be 2 characters long")
	}

	switch s {
	case "aa":
		return "Afar", "Afaraf", nil
	case "ab":
		return "Abkhazian", "аҧсуа бызшәа; аҧсшәа", nil
	case "af":
		return "Afrikaans", "Afrikaans", nil
	case "ak":
		return "Akan", "Akan", nil
	case "sq":
		return "Albanian", "Shqip", nil
	case "am":
		return "Amharic", "አማርኛ", nil
	case "ar":
		return "Arabic", "العربية", nil
	case "an":
		return "Aragonese", "aragonés", nil
	case "hy":
		return "Armenian", "Հայերեն", nil
	case "as":
		return "Assamese", "অসমীয়া", nil
	case "av":
		return "Avaric", "авар мацӀ; магӀарул мацӀ", nil
	case "ae":
		return "Avestan", "avesta", nil
	case "ay":
		return "Aymara", "aymar aru", nil
	case "az":
		return "Azerbaijani", "azərbaycan dili", nil
	case "ba":
		return "Bashkir", "башҡорт теле", nil
	case "bm":
		return "Bambara", "bamanankan", nil
	case "eu":
		return "Basque", "euskara; euskera", nil
	case "be":
		return "Belarusian", "беларуская мова", nil
	case "bn":
		return "Bengali", "বাংলা", nil
	case "bh":
		return "Bihari languages", "भोजपुरी", nil
	case "bi":
		return "Bislama", "Bislama", nil
	case "bs":
		return "Bosnian", "bosanski jezik", nil
	case "br":
		return "Breton", "brezhoneg", nil
	case "bg":
		return "Bulgarian", "български език", nil
	case "my":
		return "Burmese", "ဗမာစာ", nil
	case "ca":
		return "Catalan; Valencian", "català; valencià", nil
	case "ch":
		return "Chamorro", "Chamoru", nil
	case "ce":
		return "Chechen", "нохчийн мотт", nil
	case "zh":
		return "Chinese", "中文(Zhōngwén); 汉语; 漢語", nil
	case "cu":
		return "Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic", "ѩзыкъ словѣньскъ", nil
	case "cv":
		return "Chuvash", "чӑваш чӗлхи", nil
	case "kw":
		return "Cornish", "Kernewek", nil
	case "co":
		return "Corsican", "corsu; lingua corsa", nil
	case "cr":
		return "Cree", "ᓀᐦᐃᔭᐍᐏᐣ", nil
	case "cs":
		return "Czech", "čeština; český jazyk", nil
	case "da":
		return "Danish", "dansk", nil
	case "dv":
		return "Divehi; Dhivehi; Maldivian", "ދިވެހި", nil
	case "nl":
		return "Dutch; Flemish", "Nederlands; Vlaams", nil
	case "dz":
		return "Dzongkha", "རྫོང་ཁ", nil
	case "en":
		return "English", "English", nil
	case "eo":
		return "Esperanto", "Esperanto", nil
	case "et":
		return "Estonian", "eesti; eesti keel", nil
	case "ee":
		return "Ewe", "Eʋegbe", nil
	case "fo":
		return "Faroese", "føroyskt", nil
	case "fj":
		return "Fijian", "vosa Vakaviti", nil
	case "fi":
		return "Finnish", "suomi; suomen kieli", nil
	case "fr":
		return "French", "français; langue française", nil
	case "fy":
		return "Western Frisian", "Frysk", nil
	case "ff":
		return "Fulah", "Fulfulde; Pulaar; Pular", nil
	case "ka":
		return "Georgian", "ქართული", nil
	case "de":
		return "German", "Deutsch", nil
	case "gd":
		return "Gaelic; Scottish Gaelic", "Gàidhlig", nil
	case "ga":
		return "Irish", "Gaeilge", nil
	case "gl":
		return "Galician", "galego", nil
	case "gv":
		return "Manx", "Gaelg; Gailck", nil
	case "el":
		return "Greek, Modern (1453-)", "ελληνικά", nil
	case "gn":
		return "Guarani", "Avañeẽ", nil
	case "gu":
		return "Gujarati", "ગુજરાતી", nil
	case "ht":
		return "Haitian; Haitian Creole", "Kreyòl ayisyen", nil
	case "ha":
		return "Hausa", "(Hausa) هَوُسَ", nil
	case "he":
		return "Hebrew", "עברית", nil
	case "hz":
		return "Herero", "Otjiherero", nil
	case "hi":
		return "Hindi", "हिन्दी; हिंदी", nil
	case "ho":
		return "Hiri Motu", "Hiri Motu", nil
	case "hr":
		return "Croatian", "hrvatski jezik", nil
	case "hu":
		return "Hungarian", "magyar", nil
	case "ig":
		return "Igbo", "Asụsụ Igbo", nil
	case "is":
		return "Icelandic", "Íslenska", nil
	case "io":
		return "Ido", "Ido", nil
	case "ii":
		return "Sichuan Yi; Nuosu", "ꆈꌠ꒿ Nuosuhxop", nil
	case "iu":
		return "Inuktitut", "ᐃᓄᒃᑎᑐᑦ", nil
	case "ie":
		return "Interlingue; Occidental", "Interlingue; Occidental", nil
	case "ia":
		return "Interlingua (International Auxiliary language Association)", "Interlingua", nil
	case "id":
		return "Indonesian", "Bahasa Indonesia", nil
	case "ik":
		return "Inupiaq", "Iñupiaq; Iñupiatun", nil
	case "it":
		return "Italian", "italiano", nil
	case "jv":
		return "Javanese", "basa Jawa", nil
	case "ja":
		return "Japanese", "日本語(にほんご)", nil
	case "kl":
		return "Kalaallisut; Greenlandic", "kalaallisut; kalaallit oqaasii", nil
	case "kn":
		return "Kannada", "ಕನ್ನಡ", nil
	case "ks":
		return "Kashmiri", "कश्मीरी; كشميري‎", nil
	case "kr":
		return "Kanuri", "Kanuri", nil
	case "kk":
		return "Kazakh", "қазақ тілі", nil
	case "km":
		return "Central Khmer", "ខ្មែរ; ខេមរភាសា; ភាសាខ្មែរ", nil
	case "ki":
		return "Kikuyu; Gikuyu", "Gĩkũyũ", nil
	case "rw":
		return "Kinyarwanda", "Ikinyarwanda", nil
	case "ky":
		return "Kirghiz; Kyrgyz", "Кыргызча; Кыргыз тили", nil
	case "kv":
		return "Komi", "коми кыв", nil
	case "kg":
		return "Kongo", "Kikongo", nil
	case "ko":
		return "Korean", "한국어; 조선어", nil
	case "kj":
		return "Kuanyama; Kwanyama", "Kuanyama", nil
	case "ku":
		return "Kurdish", "Kurdî; كوردی‎", nil
	case "lo":
		return "Lao", "ພາສາລາວ", nil
	case "la":
		return "Latin", "latine; lingua latina", nil
	case "lv":
		return "Latvian", "latviešu valoda", nil
	case "li":
		return "Limburgan; Limburger; Limburgish", "Limburgs", nil
	case "ln":
		return "Lingala", "Lingála", nil
	case "lt":
		return "Lithuanian", "lietuvių kalba", nil
	case "lb":
		return "Luxembourgish; Letzeburgesch", "Lëtzebuergesch", nil
	case "lu":
		return "Luba-Katanga", "Tshiluba", nil
	case "lg":
		return "Ganda", "Luganda", nil
	case "mk":
		return "Macedonian", "македонски јазик", nil
	case "mh":
		return "Marshallese", "Kajin M̧ajeļ", nil
	case "ml":
		return "Malayalam", "മലയാളം", nil
	case "mi":
		return "Maori", "te reo Māori", nil
	case "mr":
		return "Marathi", "मराठी", nil
	case "ms":
		return "Malay", "bahasa Melayu; بهاس ملايو‎", nil
	case "mg":
		return "Malagasy", "fiteny malagasy", nil
	case "mt":
		return "Maltese", "Malti", nil
	case "mn":
		return "Mongolian", "монгол", nil
	case "na":
		return "Nauru", "Ekakairũ Naoero", nil
	case "nv":
		return "Navajo; Navaho", "Diné bizaad; Dinékʼehǰí", nil
	case "nr":
		return "Ndebele, South; South Ndebele", "isiNdebele", nil
	case "nd":
		return "Ndebele, North; North Ndebele", "isiNdebele", nil
	case "ng":
		return "Ndonga", "Owambo", nil
	case "ne":
		return "Nepali", "नेपाली", nil
	case "nn":
		return "Norwegian Nynorsk; Nynorsk, Norwegian", "Norsk nynorsk", nil
	case "nb":
		return "Bokmål, Norwegian; Norwegian Bokmål", "Norsk bokmål", nil
	case "no":
		return "Norwegian", "Norsk", nil
	case "ny":
		return "Chichewa; Chewa; Nyanja", "chiCheŵa; chinyanja", nil
	case "oc":
		return "Occitan (post 1500); Provençal", "occitan; lenga dòc", nil
	case "oj":
		return "Ojibwa", "ᐊᓂᔑᓈᐯᒧᐎᓐ", nil
	case "or":
		return "Oriya", "ଓଡ଼ିଆ", nil
	case "om":
		return "Oromo", "Afaan Oromoo", nil
	case "os":
		return "Ossetian; Ossetic", "ирон æвзаг", nil
	case "pa":
		return "Panjabi; Punjabi", "ਪੰਜਾਬੀ; پنجابی‎", nil
	case "fa":
		return "Persian", "فارسی", nil
	case "pi":
		return "Pali", "पाऴि", nil
	case "pl":
		return "Polish", "język polski; polszczyzna", nil
	case "pt":
		return "Portuguese", "português", nil
	case "ps":
		return "Pushto; Pashto", "پښتو", nil
	case "qu":
		return "Quechua", "Runa Simi; Kichwa", nil
	case "rm":
		return "Romansh", "rumantsch grischun", nil
	case "ro":
		return "Romanian; Moldavian; Moldovan", "limba română", nil
	case "rn":
		return "Rundi", "Ikirundi", nil
	case "ru":
		return "Russian", "русский язык", nil
	case "sg":
		return "Sango", "yângâ tî sängö", nil
	case "sa":
		return "Sanskrit", "संस्कृतम्", nil
	case "si":
		return "Sinhala; Sinhalese", "සිංහල", nil
	case "sk":
		return "Slovak", "slovenčina; slovenský jazyk", nil
	case "sl":
		return "Slovenian", "slovenski jezik; slovenščina", nil
	case "se":
		return "Northern Sami", "Davvisámegiella", nil
	case "sm":
		return "Samoan", "gagana faa Samoa", nil
	case "sn":
		return "Shona", "chiShona", nil
	case "sd":
		return "Sindhi", "सिन्धी; سنڌي، سندھی‎", nil
	case "so":
		return "Somali", "Soomaaliga; af Soomaali", nil
	case "st":
		return "Sotho, Southern", "Sesotho", nil
	case "es":
		return "Spanish; Castilian", "español; castellano", nil
	case "sc":
		return "Sardinian", "sardu", nil
	case "sr":
		return "Serbian", "српски језик", nil
	case "ss":
		return "Swati", "SiSwati", nil
	case "su":
		return "Sundanese", "Basa Sunda", nil
	case "sw":
		return "Swahili", "Kiswahili", nil
	case "sv":
		return "Swedish", "Svenska", nil
	case "ty":
		return "Tahitian", "Reo Tahiti", nil
	case "ta":
		return "Tamil", "தமிழ்", nil
	case "tt":
		return "Tatar", "татар теле; tatar tele", nil
	case "te":
		return "Telugu", "తెలుగు", nil
	case "tg":
		return "Tajik", "тоҷикӣ;toğikī;تاجیکی‎", nil
	case "tl":
		return "Tagalog", "Wikang Tagalog; ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔", nil
	case "th":
		return "Thai", "ไทย", nil
	case "bo":
		return "Tibetan", "བོད་ཡིག", nil
	case "ti":
		return "Tigrinya", "ትግርኛ", nil
	case "to":
		return "Tonga (Tonga Islands)", "faka Tonga", nil
	case "tn":
		return "Tswana", "Setswana", nil
	case "ts":
		return "Tsonga", "Xitsonga", nil
	case "tk":
		return "Turkmen", "Türkmen; Түркмен", nil
	case "tr":
		return "Turkish", "Türkçe", nil
	case "tw":
		return "Twi", "Twi", nil
	case "ug":
		return "Uighur; Uyghur", "Uyƣurqə; ئۇيغۇرچە‎", nil
	case "uk":
		return "Ukrainian", "українська мова", nil
	case "ur":
		return "Urdu", "اردو", nil
	case "uz":
		return "Uzbek", "O‘zbek; Ўзбек; أۇزبېك‎", nil
	case "ve":
		return "Venda", "Tshivenḓa", nil
	case "vi":
		return "Vietnamese", "Tiếng Việt", nil
	case "vo":
		return "Volapük", "Volapük", nil
	case "cy":
		return "Welsh", "Cymraeg", nil
	case "wa":
		return "Walloon", "walon", nil
	case "wo":
		return "Wolof", "Wollof", nil
	case "xh":
		return "Xhosa", "isiXhosa", nil
	case "yi":
		return "Yiddish", "ייִדיש", nil
	case "yo":
		return "Yoruba", "Yorùbá", nil
	case "za":
		return "Zhuang; Chuang", "Saɯ cueŋƅ; Saw cuengh", nil
	case "zu":
		return "Zulu", "isiZulu", nil
	}
	return "", "", fmt.Errorf("\"%s\" is not a valid ISO-639-1 language code", s)
}

// IsValidLanguageCode returns true if s is a valid ISO 639-1 language code
func IsValidLanguageCode(s string) bool {
	_, _, err := LangCodeInfo(s)
	if err != nil {
		return false
	}
	return true
}

// LangEnglishName returns the English name(s) corresponding to the language code
// s.  If there are multiple names, they are separated by a `;`.
func LangEnglishName(s string) (string, error) {
	en, _, err := LangCodeInfo(s)
	return en, err
}

// LangNativeName returns the native name(s) corresponding to the language code s
// in the native script(s).  If there are multiple names, they are separated
// by a `;`.
func LangNativeName(s string) (string, error) {
	_, nat, err := LangCodeInfo(s)
	return nat, err
}
