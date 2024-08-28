package validator

import (
	"fmt"
	"log"
)

type Validatable interface {
	Validate() error
}

func Validate(v Validatable) error {
	err := v.Validate()
	if err != nil {
		// ログ出力
		log.Printf("Validation failed: %v", err)
		// エラーメッセージのカスタマイズ
		return fmt.Errorf("validation error: %v", err)
	}
	return nil
}
