package model
//文章标签
//记录文章和标签之间的1:N的关联关系
type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}
func(a ArticleTag ) TableName() string{
	return "blog_article_tag"
}