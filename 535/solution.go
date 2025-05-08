package solution

import "crypto/sha256"

type Codec struct {
	urls map[string]string
	base string
}

func Constructor() Codec {
	return Codec{
		urls: map[string]string{},
		base: "https://tinyurl.com/",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	hasher := sha256.New()
	hasher.Write([]byte(longUrl))
	shortUrl := hasher.Sum(nil)[:8]

	this.urls[string(shortUrl)] = longUrl

	return string(shortUrl)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if longUrl, ok := this.urls[shortUrl]; ok {
		return longUrl
	}

	return ""
}
