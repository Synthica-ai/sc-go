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

func (r *Repository) UpdateAccount(userID uuid.UUID, data map[string]interface{}, ctx context.Context) error {
	sqlBuilder := PsqlBuilder.Update("account_info").Where(sq.Eq{"user_id": userID})

	_, sqlBuilder = SetJsonbMapAndReturnColumns(sqlBuilder, data)

	query, args, _ := sqlBuilder.ToSql()

	_, err := r.DB.ExecContext(ctx, query, args...)
	return err
}

type Account struct {
	Firstame           string `json:"first_name"`
	Lastname           string `json:"last_name"`
	WritingFor         string `json:"writing_for"`
	Presentation       string `json:"presentation"`
	Industry           string `json:"industry"`
	Description        string `json:"description"`
	CompanyName        string `json:"company_name"`
	CompanyIndustry    string `json:"company_industry"`
	CompanyDescription string `json:"сompany_description"`
	WritingDNA         string `json:"writing_dna"`
}

type UserSettings struct {
	Width          int32    `json:"width"`
	Height         int32    `json:"height"`
	Seed           int      `json:"seed"`
	PromptStrength *float32 `json:"prompt_strength"`
	Prompt         *string  `json:"prompt"`
	NegativePrompt *string  `json:"negative_prompt"`

	AspectRatio     string  `json:"aspect_ratio"`
	InitialImageURL *string `json:"initial_image_url"`
	ModelID         string  `json:"model_id"`
	InferenceSteps  int32   `json:"inference_steps"`
	SchedulerID     string  `json:"scheduler_id"`
	GuidanceScale   float32 `json:"guidance_scale"`
	PublicMode      bool    `json:"public_mode"`
}

type AIFriends struct {
	ID           int     `json:"id"`
	Name         *string `json:"name"`
	Age          *int    `json:"age"`
	Eyes         *string `json:"eyes"`
	Hair         *string `json:"hair"`
	Height       *string `json:"height"`
	Occupation   *string `json:"occupation"`
	SiteProfile  *string `json:"site_profile"`
	ProfileImage *string `json:"profile_image"`
	Languages    *string `json:"languages"`
	VoiceID      *string `json:"voice_id"`
}

func (r *Repository) GetAIFriendContext(id string, ctx context.Context) (string, error) {
	var res string

	rows, err := r.DB.QueryContext(ctx, `
	select
		CONCAT(
			(CASE WHEN name IS NULL OR name = '' THEN '' ELSE CONCAT('Your name is: ', name, E'\n') END),
			(CASE WHEN age IS NULL THEN '' ELSE CONCAT('You are: ', age, ' years old', E'\n') END),
			(CASE WHEN gender IS NULL OR gender = '' THEN '' ELSE CONCAT('You are: ', gender, E'\n') END),
			(CASE WHEN occupation IS NULL OR occupation = '' THEN '' ELSE CONCAT('Your occupation is: ', occupation, E'\n') END),
			(CASE WHEN education IS NULL OR education = '' THEN '' ELSE CONCAT('Your education is: ', education, E'\n') END),
			(CASE WHEN hobbies IS NULL OR hobbies = '' THEN '' ELSE CONCAT('Your favorite hobbies are ', hobbies, E'\n') END),
			(CASE WHEN business_interests IS NULL OR business_interests = '' THEN '' ELSE CONCAT('Your business interests include ', business_interests, E'\n') END),
			(CASE WHEN religion IS NULL OR religion = '' THEN '' ELSE CONCAT('You practice the ', religion,  ' religion. ', E'\n') END),
			(CASE WHEN relationship IS NULL OR relationship = '' THEN '' ELSE CONCAT('Your current relationship status is ', relationship, '. ', E'\n') END),
			(CASE WHEN zodiac IS NULL OR zodiac = '' THEN '' ELSE CONCAT('You are a ', zodiac, ', according to your zodiac sign. ', E'\n') END),
			(CASE WHEN music IS NULL OR music = '' THEN '' ELSE CONCAT('You enjoy listening to ', music, E'\n') END),
			(CASE WHEN movie IS NULL OR movie = '' THEN '' ELSE CONCAT('You enjoy watching ', movie,  ' movies. ', E'\n') END),
			(CASE WHEN cooking IS NULL OR cooking = '' THEN '' ELSE CONCAT('In cooking, you love ', cooking, '. ', E'\n') END),
			(CASE WHEN social_account IS NULL OR social_account = '' THEN '' ELSE CONCAT('You love using ', social_account, '. ', E'\n') END),
			(CASE WHEN hair IS NULL OR eyes IS NULL THEN '' ELSE CONCAT('You have ', hair,  ' hair and ', eyes,  ' eyes. ', E'\n') END),
			(CASE WHEN body_type IS NULL OR body_type = '' THEN '' ELSE CONCAT('Your body type is ', body_type, '. ', E'\n') END),
			(CASE WHEN height_ft_in IS NULL OR height_ft_in = '' THEN '' ELSE CONCAT('You are ', height_ft_in,  ' in height. ', E'\n') END),
			(CASE WHEN smoking IS NULL OR smoking = '' THEN '' ELSE CONCAT('You ', smoking,  ' smoke. ', E'\n') END),
			(CASE WHEN drinking IS NULL OR drinking = '' THEN '' ELSE CONCAT('You ', drinking,  ' drink alcohol. ', E'\n') END),
			(CASE WHEN home_town IS NULL OR home_town = '' THEN '' ELSE CONCAT('You live in ', home_town, E'\n') END),
			(CASE WHEN business_personality IS NULL OR business_personality = '' THEN '' ELSE CONCAT('Your business personality is ', business_personality, '. ', E'\n') END),
			context_you_are,
			E'\n',
			business_background,
			E'\n',
			(select context from ai_friends_policy limit 1)
		)
	from ai_friends where id=$1;
	`, id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(
			&res,
		)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

// Update last_seen_at
func (r *Repository) GetAIFriends(ctx context.Context) ([]AIFriends, error) {
	res := make([]AIFriends, 0)

	rows, err := r.DB.QueryContext(ctx, `
		select
			id,
			name,
			age,
			eyes,
			hair,
			height_ft_in,
			occupation,
			site_profile,
			profile_image,
			languages,
			voice_id
		from ai_friends;
	`)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var fr AIFriends
		err := rows.Scan(
			&fr.ID,
			&fr.Name,
			&fr.Age,
			&fr.Eyes,
			&fr.Hair,
			&fr.Height,
			&fr.Occupation,
			&fr.SiteProfile,
			&fr.ProfileImage,
			&fr.Languages,
			&fr.VoiceID,
		)
		if err != nil {
			return res, err
		}

		res = append(res, fr)
	}

	return res, nil
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

// Update last_seen_at
func (r *Repository) GetAccount(userID uuid.UUID, ctx context.Context) (Account, error) {
	var res Account

	rows, err := r.DB.QueryContext(ctx, `
		select
			first_name,
			last_name,
			writing_for,
			presentation,
			industry,
			description,
			company_name,
			company_industry,
			сompany_description,
			writing_dna
		from account_info
		where user_id=$1;
	`, userID)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(
			&res.Firstame,
			&res.Lastname,
			&res.WritingFor,
			&res.Presentation,
			&res.Industry,
			&res.Description,
			&res.CompanyName,
			&res.CompanyIndustry,
			&res.CompanyDescription,
			&res.WritingDNA,
		)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

// Update last_seen_at
func (r *Repository) GetImageSettings(imageUrl string, ctx context.Context) (UserSettings, error) {
	var res UserSettings

	rows, err := r.DB.QueryContext(ctx, `
		select
			width,
			height,
			inference_steps,
			guidance_scale,
			seed,
			init_image_url,
			prompt_strength,
			model_id,
			prn.text as negative_prompt,
			pr.text as prompt,
			scheduler_id   
		from generation_outputs go
		JOIN generations gen ON gen.id = go.generation_id
		JOIN prompts pr ON pr.id = gen.prompt_id
		LEFT JOIN negative_prompts prn ON prn.id = gen.negative_prompt_id
		where go.image_path=$1;
	`, imageUrl)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(
			&res.Width,
			&res.Height,
			&res.InferenceSteps,
			&res.GuidanceScale,
			&res.Seed,
			&res.InitialImageURL,
			&res.PromptStrength,
			&res.ModelID,
			&res.NegativePrompt,
			&res.Prompt,
			&res.SchedulerID,
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
