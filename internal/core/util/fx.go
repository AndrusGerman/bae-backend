package util

import (
	"bae-backend/internal/core/domain"
	"fmt"
)

func NewParamTag(name string, tagType string) string {
	return fmt.Sprintf(`%s:"%s_%s"`, tagType, name, domain.NewId().Hex())
}
