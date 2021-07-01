package bo

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	IsOwner    int
	List       interface{}
}

func PageUtil(totalCount int, pageNo int, pageSize int, list interface{}, isOwner int) Page {
	totalPage := totalCount / pageSize
	if totalCount%pageSize > 0 {
		totalPage = totalCount/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: totalPage, TotalCount: totalCount,
		FirstPage: pageNo == 1, LastPage: pageNo == totalPage, IsOwner: isOwner, List: list}
}
