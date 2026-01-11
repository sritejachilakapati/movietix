package service

func calcNextOffset(limit int32, offset int32, resultsCount int) *int32 {
	if int32(resultsCount) < limit {
		return nil
	}
	newOffset := offset + limit
	return &newOffset
}
