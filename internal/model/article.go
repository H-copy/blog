package model


type Article struct{
	* Model
	Title string `json:"title"`
	Content string `json:"content"`
	Desc string `json:"desc"`
	ConverImageUrl string `json:conver_image_url`
	State uint8 `json:"state"`
}

func (a  Article) TableName() string{
	return "blog_article"
}