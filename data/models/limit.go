package models

import (
	"github.com/kamva/mgm/v3"
)

// limit represents a limit model
type Limit struct {
	mgm.DefaultModel    `bson:",inline"`
	UserSlug            string `json:"user_slug" bson:"userSlug"`
	ServiceName         string `json:"service_name" bson:"serviceName"`
	DailyLimit          int64  `json:"daily_limit" bson:"dailyLimit"`
	DailyCountLimit     int64  `json:"daily_count_limit" bson:"dailyCountLimit"`
	MonthlyLimit        int64  `json:"monthly_limit" bson:"monthlyLimit"`
	MonthlyCountLimit   int64  `json:"monthly_count_limit" bson:"monthlyCountLimit"`
	PerTransactionLimit int64  `json:"per_transaction_limit" bson:"perTransactionLimit"`
}
