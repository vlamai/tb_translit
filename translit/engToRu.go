package translit

import "unicode"

var keyboard = map[rune]rune{
	'q':  'й',
	'w':  'ц',
	'e':  'у',
	'r':  'к',
	't':  'е',
	'a':  'ф',
	's':  'ы',
	'd':  'в',
	'f':  'а',
	'g':  'п',
	'z':  'я',
	'x':  'ч',
	'c':  'с',
	'v':  'м',
	'b':  'и',
	'y':  'н',
	'u':  'г',
	'i':  'ш',
	'o':  'щ',
	'p':  'з',
	'[':  'х',
	']':  'ъ',
	'h':  'р',
	'j':  'о',
	'k':  'л',
	'l':  'д',
	';':  'ж',
	'\'': 'э',
	'n':  'т',
	'm':  'ь',
	',':  'б',
	'.':  'ю',
	'/':  '.',
}

func EngToRu(inp string) string {
	result := make([]rune, len(inp))
	for i, c := range inp {
		isCapital := unicode.IsUpper(c)
		if isCapital {
			c = unicode.ToLower(c)
		}
		let, ok := keyboard[c]
		if !ok {
			result[i] = c
		} else {
			if isCapital {
				let = unicode.ToUpper(let)
			}
			result[i] = let
		}

	}
	return string(result)
}
