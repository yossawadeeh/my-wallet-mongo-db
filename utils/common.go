package utils

func GetOffsetPage(page int, perPage int) int64 {
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 2
	}

	return int64(page-1) * int64(perPage)
}
