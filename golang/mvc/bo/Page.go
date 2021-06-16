package bo

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(totalCount int, pageNo int, pageSize int, list interface{}) Page {
	totalPage := totalCount / pageSize
	if totalCount%pageSize > 0 {
		totalPage = totalCount/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: totalPage, TotalCount: totalCount,
		FirstPage: pageNo == 1, LastPage: pageNo == totalPage, List: list}
}
