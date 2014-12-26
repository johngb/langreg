Langreg
=====

langreg is a lightweight Go package to handle ISO region and language names for general computing use. This only uses ISO 639-1 language codes, and ISO 3166-1 alpha-2 country codes.  These are commonly used for language and region settings.  E.g. `en_US`, `en_GB`, `zu_ZA`

## Source of Data

The data used for the language codes and English names, come from the [official source](http://loc.gov/standards/iso639-2/ISO-639-2_utf-8.txt) at the Library of Congress.  As the official specification doesn't include native names (What were they thinking?!), I've scraped the native names from Wikipedia's [list of ISO 639-1 codes](http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).

As this is intended for general use cases, rather than distinguishing between ancient and modern languages, some of the names from the official source have been simplified where I felt they would be clear.  For example "Greek, Modern (1453-)" - the official ISO name - is less clear than "Greek", and so I used "Greek".  There are a few other minor cases like this, but nothing significant.

The data used for the region codes and data has been scraped from Wikipedia's [ISO 3166-1 page](http://en.wikipedia.org/wiki/ISO_3166-1). I couldn't find any single download from ISO without paying for it or scraping a page per letter of the alphabet, so I stuck with Wikipedia - because lazy.

If you find any errors, please post an issue, and I'll get to it as soon as I can.