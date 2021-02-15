package structs

import (
	"math"
)

type Pagination struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

func (pagination Pagination) GetPage() int64 {
	return int64(math.Max(1, float64(pagination.Page)))
}

func (pagination Pagination) GetSize() int64 {
	return int64(math.Max(1, float64(pagination.Size)))
}

func (pagination Pagination) GetSkip() int64 {
	return pagination.GetSize() * (pagination.GetPage() - 1)
}