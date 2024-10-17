package request

//type Pagination struct {
//	PageIndex int `form:"pageIndex"` //第几页
//	PageSize  int `form:"pageSize"`  //每页返回条数
//}
//
//func (m *Pagination) GetPage() int {
//	if m.PageIndex <= 0 {
//		m.PageIndex = 1
//	}
//	return m.PageIndex
//}
//
//func (m *Pagination) GetSize() int {
//	if m.PageSize <= 0 {
//		m.PageSize = 10
//	}
//	return m.PageSize
//}

type Pagination struct {
	Page     int `json:"pageIndex" form:"pageIndex" query:"-"` // 页码
	PageSize int `json:"pageSize" form:"pageSize" query:"-"`   // 每页大小
	//	Keyword string `json:"keyword" form:"keyword"` //关键字
}

func (e *Pagination) GetPage() int {
	if e.Page < 1 {
		return 1
	}
	return e.Page
}

func (e *Pagination) GetSize() int {
	if e.PageSize < 1 {
		return 10
	}
	return e.PageSize
}

func (e *Pagination) GetOffset() int {
	return (e.GetPage() - 1) * e.GetSize()
}
