package repository

import (
	"math"
	"moonlay-test/models"

	"gorm.io/gorm"
)

type ListRepository interface {
	FindListsByFilter(title, description string, page, limit int) (PaginationResult, error)
	GetList(ID int) (models.List, error)
	CreateList(list models.List) (models.List, error)
	UpdateList(list models.List) (models.List, error)
	DeleteList(list models.List) (models.List, error)
}

type PaginationResult struct {
	Lists       []models.List `json:"lists"`
	TotalItems  int64         `json:"total_items"`
	TotalPages  int           `json:"total_pages"`
	CurrentPage int           `json:"current_page"`
	Limit       int           `json:"limit"`
}

func RepositoryList(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindListsByFilter(title, description string, page, limit int) (PaginationResult, error) {
	var lists []models.List
	var totalItems int64
	query := r.db.Model(&models.List{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	offset := (page - 1) * limit

	query = query.Offset(offset).Limit(limit)

	err := query.Preload("Sublists").Find(&lists).Error

	if err != nil {
		return PaginationResult{}, err
	}

	errCount := r.db.Model(&models.List{}).Where(query).Count(&totalItems).Error

	if errCount != nil {
		return PaginationResult{}, errCount
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	paginationResult := PaginationResult{
		Lists:       lists,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
	}

	return paginationResult, nil
}

func (r *repository) GetList(ID int) (models.List, error) {
	var list models.List
	err := r.db.Preload("Sublists").First(&list, ID).Error

	return list, err
}

func (r *repository) CreateList(list models.List) (models.List, error) {
	err := r.db.Create(&list).Error
	return list, err
}

func (r *repository) UpdateList(list models.List) (models.List, error) {
	err := r.db.Save(&list).Error
	return list, err
}

func (r *repository) DeleteList(list models.List) (models.List, error) {
	err := r.db.Delete(&list).Error
	return list, err
}
