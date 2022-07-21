package utils

const byteEmpty = 0x20

// Slugify convert string to slug
func Slugify(title string) string {
	var slug []byte

	for _, x := range title {
		switch {
		case ('a' <= x && x <= 'z') || ('0' <= x && x <= '9'):
			slug = append(slug, byte(x))
		case 'A' <= x && x <= 'Z':
			slug = append(slug, byte(x)+byteEmpty)
		case x == '-' || x == ' ':
			slug = append(slug, '-')
		}
	}

	return string(slug)
}
