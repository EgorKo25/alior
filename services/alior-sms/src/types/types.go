package types

type Service struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

type PagineBody struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
