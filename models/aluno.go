package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func ValidaAluno(a *Aluno) error {
	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}
