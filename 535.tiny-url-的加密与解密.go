type Codec struct {
	id int
	m  map[int]string
}

func Constructor() Codec {
	return Codec{id: 0, m: make(map[int]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.id++
	this.m[this.id] = longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	id, _ := strconv.Atoi(shortUrl)
	return this.m[id] || ''
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */