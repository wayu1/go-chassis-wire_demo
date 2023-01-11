package types

type InfoParam struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

type InfoRsp struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}
