package responses

import (
	"time"

	"github.com/stablecog/sc-go/database/ent/userrole"
)

type GetUserResponse struct {
	TotalRemainingCredits int        `json:"total_remaining_credits"`
	HasNonfreeCredits     bool       `json:"has_nonfree_credits"`
	ProductID             string     `json:"product_id,omitempty"`
	PriceID               string     `json:"price_id,omitempty"`
	CancelsAt             *time.Time `json:"cancels_at,omitempty"`
	RenewsAt              *time.Time `json:"renews_at,omitempty"`
	MoreCreditsAt         *time.Time `json:"more_credits_at,omitempty"`
	// The current amoount of free credits server offers
	FreeCreditAmount *int                `json:"free_credit_amount,omitempty"`
	StripeHadError   bool                `json:"stripe_had_error"`
	Roles            []userrole.RoleName `json:"roles,omitempty"`
	Role             string              `json:"role"`
	Username         string              `json:"username"`
	CompleteProfile  bool                `json:"complete_profile"`
}
