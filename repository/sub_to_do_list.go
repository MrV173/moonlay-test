package repository

import (
	"math"
	"moonlay-test/models"

	"gorm.io/gorm"
)

type SublistRepository interface {
	FindSublistsByFilter(ListID, page, limit int, title, description string) (PaginationSublistResult, error)
	GetSubList(ID int) (models.Sublist, error)
	CreateSublist(sublist models.Sublist) (models.Sublist, error)
	UpdateSublist(sublist models.Sublist) (models.Sublist, error)
	DeleteSublist(sublist models.Sublist) (models.Sublist, error)
}

type PaginationSublistResult struct {
	Sublists    []models.SublistResponse `json:"sublists"`
	TotalItems  int64                    `json:"total_items"`
	TotalPages  int                      `json:"total_pages"`
	CurrentPage int                      `json:"current_page"`
	Limit       int                      `json:"limit"`
}

func RepositorySublist(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSublistsByFilter(listID, page, limit int, title, description string) (PaginationSublistResult, error) {
	var result PaginationSublistResult

	query := r.db.Model(&models.Sublist{}).Where("list_id = ?", listID)

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	var totalItems int64
	errCount := query.Count(&totalItems).Error
	if errCount != nil {
		return PaginationSublistResult{}, errCount
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	var sublists []models.SublistResponse
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&sublists).Error
	if err != nil {
		return PaginationSublistResult{}, err
	}

	result.Sublists = sublists
	result.TotalItems = totalItems
	result.TotalPages = totalPages
	result.CurrentPage = page
	result.Limit = limit

	return result, nil
}

func (r *repository) GetSubList(ID int) (models.Sublist, error) {
	var sublist models.Sublist
	err := r.db.Preload("List").First(&sublist, ID).Error
	return sublist, err
}

func (r *repository) CreateSublist(sublist models.Sublist) (models.Sublist, error) {
	err := r.db.Create(&sublist).Error
	return sublist, err
}

func (r *repository) UpdateSublist(sublist models.Sublist) (models.Sublist, error) {
	err := r.db.Save(&sublist).Error
	return sublist, err
}

func (r *repository) DeleteSublist(sublist models.Sublist) (models.Sublist, error) {
	err := r.db.Delete(&sublist).Error
	return sublist, err
}
