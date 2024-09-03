package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName 方法指定了 Tag 结构体对应的数据库表名为 "blog_tag"
func (t Tag) TableName() string {
	return "blog_tag"
}
