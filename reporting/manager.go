//go:generate mockgen -source manager.go -destination mock/manager_mock.go -package mock
package reporting

import (
	"github.com/TestardR/reporting/reportstore"
	"github.com/TestardR/reporting/uuid"
)

type ReportManager struct {
	uuidGenerator uuid.Generator
	store         reportstore.Store
}

func NewReportManager(gen uuid.Generator, store reportstore.Store) *ReportManager {
	return &ReportManager{
		uuidGenerator: gen,
		store:         store,
	}
}

func (m *ReportManager) CreateReport(request CreateReportRequest) (response CreateReportResponse, err error) {
	reportID := m.uuidGenerator.Generate()
	r := reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   request.UserID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    request.Title,
	}
	err = m.store.CreatReport(r)

	response.ReportID = r.ReportID
	return
}
