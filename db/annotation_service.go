package db

import (
	"github.com/OlympBMSTU/annotations/db/result"
	"github.com/OlympBMSTU/annotations/entities"

)

type IAnnotationService interface {
	CreateAnnotation(annotation entities.AnnotationModel) result.DbResult
	DeleteAnnotation(authorID int) result.DbResult
	GetAnnotation(authorID int) result.DbResult
	// TODO we cant search by userID, also by fileName, but subject, maybe tags
	GetAnnotationsList(subject string, limit int, offset int) result.DbResult
}

