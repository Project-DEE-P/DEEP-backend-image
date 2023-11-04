package model

// In 접미사는 client에게 요청된 모델임을 뜻합니다.
// Out 접미사는 client에게 반환하는 모델임을 뜻합니다.
type (
	// OutGeneral 구조체는 response할 때 사용되는 구조체 입니다.
	OutGeneral struct {
		Status  uint   `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
)
