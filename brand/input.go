package brand

type GetBrandDetailInput struct {
	BrdId int `uri:"id" binding:"required"`
}

type CreateBrandInput struct {
	BrdName string `json:"brd_name" binding:"required"`
}
