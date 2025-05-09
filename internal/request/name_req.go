package request

type NameRequest struct {
    Name string `json:"name"`
}

type ChangeNameRequest struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

type DeleteByIndexRequest struct {
    Index int `json:"index"`
}

type DeleteByNameRequest struct {
    Name string `json:"name"`
}
