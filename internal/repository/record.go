package repository

import (
	"AI-Insurance-Agent/internal/model"
	"gorm.io/gorm"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(record *model.AnalysisRecord) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) FindByID(id int64) (*model.AnalysisRecord, error) {
	var record model.AnalysisRecord
	err := r.db.First(&record, id).Error
	return &record, err
}

func (r *RecordRepository) ListByUserID(userID int64, page, pageSize int) ([]model.AnalysisRecord, int64, error) {
	var records []model.AnalysisRecord
	var total int64

	offset := (page - 1) * pageSize
	err := r.db.Model(&model.AnalysisRecord{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&records).Error
	return records, total, err
}

func (r *RecordRepository) Delete(id, userID int64) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.AnalysisRecord{}).Error
}

func (r *RecordRepository) Update(record *model.AnalysisRecord) error {
	return r.db.Save(record).Error
}
