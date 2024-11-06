package dto

import "base-gin/domain/dao"

type PublisherCreateReq struct {
	Name string `json:"name" binding:"required,min=2,max=48"`
	City string `json:"city" binding:"required,max=32"`
}

func (o *PublisherCreateReq) ToEntity() dao.Publisher {
	var item dao.Publisher
	item.Name = o.Name
	item.City = o.City

	return item
}
