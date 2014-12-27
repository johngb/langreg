"""Generate RegionCodeInfo()."""
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
"""

FUNCTION_TAIL = u"""
  return "",
    fmt.Errorf("\\"%s\\" is not a valid ISO 3166-1 alpha-2 region code", s)
}
"""


def switch_statement():
    """Generate the switch statement in RegionCodeInfo."""

    yield u'switch s[0] {'

    by_first = defaultdict(dict)
    with open('region_codes.csv', 'rb') as src:
        for code, region in [line for line in csv.reader(src)]:
            by_first[code[0]][code] = region
    for first in sorted(by_first):
        yield u"case '{}':".format(first)
        yield u'switch s[1] {'

        for code in sorted(by_first[first]):
            region = by_first[first][code]
            yield u"case '{}':".format(code[1])
            yield u'return "{}", nil'.format(region.decode('utf-8'))

        yield u'}'

    yield u'}'


def main():
    """Write RegionCodeInfo to region_code_info.go."""
    with open('region_code_info.go', 'wb') as out:
        out.write(PACKAGE_HEAD)
        out.write(FUNCTION_HEAD)
        for line in switch_statement():
            out.write(line.encode('utf-8') + '\n')
        out.write(FUNCTION_TAIL)
    os.system('gofmt -w region_code_info.go')


if __name__ == '__main__':
    main()
