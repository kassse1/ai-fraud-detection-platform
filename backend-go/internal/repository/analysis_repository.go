package repository

import (
	"context"
	"database/sql"

	"github.com/kassse1/ai-fraud-backend/internal/domain"
)

type AnalysisRepository struct {
	db *sql.DB
}

func NewAnalysisRepository(db *sql.DB) *AnalysisRepository {
	return &AnalysisRepository{db: db}
}

func (r *AnalysisRepository) Save(ctx context.Context, res domain.AnalysisResult, text string) error {
	query := `
		INSERT INTO analysis_results
		(text, risk_score, is_scam, is_ai_generated, explanation)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		text,
		res.RiskScore,
		res.IsScam,
		res.IsAIGenerated,
		res.Explanation,
	)
	return err
}
func (r *AnalysisRepository) List(
	ctx context.Context,
	limit int,
	offset int,
) ([]domain.AnalysisResultWithText, error) {

	rows, err := r.db.QueryContext(ctx, `
		SELECT text, risk_score, is_scam, is_ai_generated, explanation, created_at
		FROM analysis_results
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.AnalysisResultWithText

	for rows.Next() {
		var rlt domain.AnalysisResultWithText
		err := rows.Scan(
			&rlt.Text,
			&rlt.RiskScore,
			&rlt.IsScam,
			&rlt.IsAIGenerated,
			&rlt.Explanation,
			&rlt.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, rlt)
	}

	return results, nil
}
