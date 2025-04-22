package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome     string `json:"nome" validate:"nonzero"`
	Document string `json:"document" validate:"len=9, regexp=^[0-9]*$"`
}

func AlunoValidate(aluno *Aluno) error {

	err := validator.Validate(aluno)

	if err != nil {
		return err
	}

	return nil
}
