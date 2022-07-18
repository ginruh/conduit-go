package article

type ArticleService interface {
	List()
	Feed()
	Get(id string)
	Create()
	Update()
	Delete()
	Favorite()
	Unfavorite()
	ListTags()
}
