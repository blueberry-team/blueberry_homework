package req

type NameRequest struct {
    Name string `json:"name"`
}

type DeleteByIndexRequest struct {
    Index int `json:"index"`
}

type DeleteByNameRequest struct {
    Name string `json:"name"`
}
