package webcustomer

import "mime/multipart"

type CustomerCreateRequest struct {
	Name    string                `json:"name" form:"name"`
	NoHp    int64                 `json:"no_hp" form:"no_hp"`
	Address string                `json:"address" form:"address"`
	Photo   *multipart.FileHeader `form:"photo" `
	Balance int64                 `json:"balance" form:"balance"`
}

type CustomerPhotoRequest struct {
	Photo *multipart.FileHeader `form:"photo_customer"`
}
