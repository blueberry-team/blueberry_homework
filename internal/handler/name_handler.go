package handler

import (
	"encoding/json"
	"net/http"

	"blueberry_homework/internal/repository"
)

type NameRequest struct {
    Name string `json:"name"`
}

// NameHandler는 이름 관련 HTTP 요청을 처리하는 핸들러입니다.
type NameHandler struct {
    repo repository.NameRepository
}

// NewNameHandler는 새로운 NameHandler 인스턴스를 생성합니다.
// repository.NameRepository를 의존성으로 주입받습니다.
func NewNameHandler(r repository.NameRepository) *NameHandler {
    return &NameHandler{repo: r}
}

// CreateName은 새로운 이름을 생성하는 HTTP 엔드포인트 핸들러입니다.
// POST 요청의 body에서 JSON 형식의 이름을 받아 저장합니다.
// 성공 시 201 Created 상태 코드를 반환합니다.
// TODO: success message 반환
func (h *NameHandler) CreateName(w http.ResponseWriter, r *http.Request) {
    var req NameRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil || req.Name == "" {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    h.repo.CreateName(req.Name)
    w.WriteHeader(http.StatusCreated)
}

// GetNames는 저장된 모든 이름을 조회하는 HTTP 엔드포인트 핸들러입니다.
// 저장된 이름 목록을 JSON 형식으로 반환합니다.
// TODO: success message 반환
// TODO: 공백 문자열 반환
func (h *NameHandler) GetNames(w http.ResponseWriter, r *http.Request) {
    names := h.repo.GetNames()
    json.NewEncoder(w).Encode(names)
}
