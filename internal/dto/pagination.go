package dto

type PageResult[T any] struct {
	Items      []T    `json:"items"`
	NextOffset *int32 `json:"nextOffset"`
}
