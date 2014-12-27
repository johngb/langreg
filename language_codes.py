"""Generate LangCodeInfo()."""
import csv
import os

from collections import defaultdict


PACKAGE_HEAD = u"""
package langreg
import (
  "errors"
  "fmt"
)
"""


FUNCTION_HEAD = u"""
// LangCodeInfo returns the English and native language in it's script for a
// given string, and an error if any.  If there are more than one official
// names for the language (either English or native), they are separated by a
// semi-colon (;)
// Language codes should always be lowercase, and this is enforced.
func LangCodeInfo(s string) (english, native string, err error) {

  // codes have to be two characters long
  if len(s) != 2 {
    return "", "",
      errors.New("ISO 639-1 language codes must be 2 characters long")
  }
"""

FUNCTION_TAIL = u"""
  return "", "",
    fmt.Errorf("\\"%s\\" is not a valid ISO-639-1 language code", s)
}
"""


def switch_statement():
    """Generate the switch statement in LangCodeInfo."""

    yield u'switch s[0] {'

    by_first = defaultdict(dict)
    with open('language_codes.csv', 'rb') as src:
        for code, english, native in [line for line in csv.reader(src)]:
            by_first[code[0]][code] = english, native
    for first in sorted(by_first):
        yield u"case '{}':".format(first)
        yield u'switch s[1] {'

        for code in sorted(by_first[first]):
            english, native = by_first[first][code]
            yield u"case '{}':".format(code[1])
            yield u'return "{}", "{}", nil'.format(
                english.decode('utf-8'), native.decode('utf-8'))

        yield u'}'

    yield u'}'


def main():
    """Write LangCodeInfo to language_code_info.go."""
    with open('language_code_info.go', 'wb') as out:
        out.write(PACKAGE_HEAD)
        out.write(FUNCTION_HEAD)
        for line in switch_statement():
            out.write(line.encode('utf-8') + '\n')
        out.write(FUNCTION_TAIL)
    os.system('gofmt -w language_code_info.go')


if __name__ == '__main__':
    main()
