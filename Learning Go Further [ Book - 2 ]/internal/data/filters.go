package data

type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

func (f Filters) Limit() int {
	if f.PageSize <= 0 || f.PageSize > 100 {
		return 10
	}
	return f.PageSize
}

func (f Filters) Offset() int {
	if f.Page <= 0 {
		return 0
	}
	return (f.Page - 1) * f.Limit()
}
