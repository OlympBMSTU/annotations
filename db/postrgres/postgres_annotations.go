package postgres

import (
	"github.com/jackc/pgx"
)

type AnnotationsService struct {
	pool *pgx.ConnPool
}

func (service *AnnotationsService) CreateAnnotation(annotation model.AnnoatationModel) result.DbResult {
	return OKResult()
}

func (service *AnnotationsService) DeleteAnnotation(authorID int) result.DbResult {
	return result.DbResult{}
}

func (service *AnnotationsService) GetAnnotation(authorID int) result.DbResult {
	return result.DbResult{}
}
	// TODO we cant search by userID, also by fileName, but subject, maybe tags
func (service *AnnotationsService) GetAnnotationsList(subject string, limit int, offset int) result.DbResult {
	return result.DbResult{}
}