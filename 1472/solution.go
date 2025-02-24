package solution

type BrowserHistory struct {
	pages       []string
	currentPage int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		pages:       []string{homepage},
		currentPage: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	if this.currentPage < len(this.pages)-1 {
		this.pages = this.pages[:this.currentPage+1]
		this.pages = append(this.pages, url)
		this.currentPage++
	} else {
		this.pages = append(this.pages, url)
		this.currentPage = len(this.pages) - 1
	}
}

func (this *BrowserHistory) Back(steps int) string {
	if this.currentPage-steps >= 0 {
		this.currentPage -= steps
		return this.pages[this.currentPage]
	} else {
		this.currentPage = 0
		return this.pages[this.currentPage]
	}
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.currentPage+steps <= len(this.pages)-1 {
		this.currentPage += steps
		return this.pages[this.currentPage]
	} else {
		this.currentPage = len(this.pages) - 1
		return this.pages[this.currentPage]
	}
}
