package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/shared"
)

var PsqlBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (r *Repository) CreateUser(id uuid.UUID, email string, stripeCustomerId string, lastSignIn *time.Time, db *ent.Client) (*ent.User, error) {
	if db == nil {
		db = r.DB
	}
	cq := db.User.Create().SetID(id).SetStripeCustomerID(stripeCustomerId).SetEmail(email)
	if lastSignIn != nil {
		cq.SetLastSignInAt(*lastSignIn)
	}
	return cq.Save(r.Ctx)
}

func (r *Repository) SetActiveProductID(id uuid.UUID, stripeProductID string, db *ent.Client) error {
	if db == nil {
		db = r.DB
	}
	return db.User.UpdateOneID(id).SetActiveProductID(stripeProductID).Exec(r.Ctx)
}

// Only unset if the active product ID matches the stripe product ID given
func (r *Repository) UnsetActiveProductID(id uuid.UUID, stripeProductId string, db *ent.Client) (int, error) {
	if db == nil {
		db = r.DB
	}
	return db.User.Update().Where(user.IDEQ(id), user.ActiveProductIDEQ(stripeProductId)).ClearActiveProductID().Save(r.Ctx)
}

// Update last_seen_at
func (r *Repository) UpdateLastSeenAt(id uuid.UUID) error {
	return r.DB.User.UpdateOneID(id).SetLastSeenAt(time.Now()).Exec(r.Ctx)
}

// Update last_seen_at
func (r *Repository) UpdateUserSettings(id uuid.UUID, data map[string]interface{}, ctx context.Context) error {

	sqlBuilder := PsqlBuilder.Update("user_settings").Where(sq.Eq{"user_id": id})

	_, sqlBuilder = SetJsonbMapAndReturnColumns(sqlBuilder, data)

	query, args, _ := sqlBuilder.ToSql()

	_, err := r.DB.ExecContext(ctx, query, args...)
	return err
}

// Update last_seen_at
func (r *Repository) UpdateUser(id uuid.UUID, data map[string]interface{}, ctx context.Context) error {
	sqlBuilder := PsqlBuilder.Update("users").Where(sq.Eq{"id": id})

	_, sqlBuilder = SetJsonbMapAndReturnColumns(sqlBuilder, data)

	query, args, _ := sqlBuilder.ToSql()

	_, err := r.DB.ExecContext(ctx, query, args...)
	return err
}

type UserSettings struct {
	AspectRatio     string `json:"aspect_ratio"`
	InitialImageURL string `json:"initial_image_url"`
	ModelID         string `json:"model_id"`
	InferenceSteps  int    `json:"inference_steps"`
	SchedulerID     string `json:"scheduler_id"`
	GuidanceScale   int    `json:"guidance_scale"`
	PublicMode      bool   `json:"public_mode"`
}

// Update last_seen_at
func (r *Repository) GetUserSettings(id uuid.UUID, ctx context.Context) (UserSettings, error) {
	var res UserSettings

	rows, err := r.DB.QueryContext(ctx, `
		select
			aspect_ratio,
			initial_image_url,
			model_id,
			inference_steps,
			scheduler_id,
			guidance_scale,
			public_mode
		from user_settings
		where user_id=$1;
	`, id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(
			&res.AspectRatio,
			&res.InitialImageURL,
			&res.ModelID,
			&res.InferenceSteps,
			&res.SchedulerID,
			&res.GuidanceScale,
			&res.PublicMode,
		)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

// Sync stripe product IDs
func (r *Repository) SyncStripeProductIDs(productCustomerIDMap map[string][]string) error {
	if err := r.WithTx(func(tx *ent.Tx) error {
		allCustomersWithProducts := make([]string, 0)
		for productID, customerIDs := range productCustomerIDMap {
			allCustomersWithProducts = append(allCustomersWithProducts, customerIDs...)
			_, err := tx.User.Update().Where(user.StripeCustomerIDIn(customerIDs...)).SetActiveProductID(productID).Save(r.Ctx)
			if err != nil {
				return err
			}
		}
		err := tx.User.Update().Where(user.StripeCustomerIDNotIn(allCustomersWithProducts...)).ClearActiveProductID().Exec(r.Ctx)
		return err
	}); err != nil {
		return err
	}
	return nil
}

// Ban users
func (r *Repository) BanUsers(userIDs []uuid.UUID) (int, error) {
	return r.DB.User.Update().Where(user.IDIn(userIDs...)).SetBannedAt(time.Now()).SetScheduledForDeletionOn(time.Now().Add(shared.DELETE_BANNED_USER_DATA_AFTER)).Save(r.Ctx)
}

// Unban users
func (r *Repository) UnbanUsers(userIDs []uuid.UUID) (int, error) {
	return r.DB.User.Update().Where(user.IDIn(userIDs...)).ClearBannedAt().ClearScheduledForDeletionOn().Save(r.Ctx)
}

func SetJsonbMapAndReturnColumns(sqlBuilder sq.UpdateBuilder, data map[string]interface{}) ([]string, sq.UpdateBuilder) {
	keys := make([]string, 0, len(data))

	for field, value := range data {
		keys = append(keys, field)

		switch vData := value.(type) {
		case string:
			sqlBuilder = sqlBuilder.Set(field, vData)
		case map[string]interface{}:
			if len(vData) == 0 {
				return keys, sqlBuilder
			}

			sqlBuilder = setNestedUpdateField(sqlBuilder, field, vData)
		default:
			sqlBuilder = sqlBuilder.Set(field, vData)
		}
	}

	return keys, sqlBuilder
}

func setNestedUpdateField(sqlBuilder sq.UpdateBuilder, parentFieldName string, data map[string]interface{}) sq.UpdateBuilder {
	for k, v := range data {
		updateField := fmt.Sprintf("%s['%s']", parentFieldName, k)

		switch v.(type) {
		case string:
			sqlBuilder = sqlBuilder.Set(updateField, sq.Expr("to_jsonb(?::text)", v))
		case map[string]interface{}:
			sqlBuilder = setNestedUpdateField(sqlBuilder, updateField, v.(map[string]interface{}))
		default:
			sqlBuilder = sqlBuilder.Set(updateField, v)
		}

	}

	return sqlBuilder
}
