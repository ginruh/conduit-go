package tag

type ListTagsResponse struct {
	Tags []string `json:"tags"`
}

func (s TagServiceImpl) List() (*ListTagsResponse, error) {
	tags, err := s.q.ListArticleTags()
	if err != nil {
		return nil, err
	}
	return &ListTagsResponse{Tags: tags}, nil
}
