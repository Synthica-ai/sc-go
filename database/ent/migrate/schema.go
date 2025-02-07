// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APITokensColumns holds the columns for the "api_tokens" table.
	APITokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "hashed_token", Type: field.TypeString, Size: 2147483647},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "short_string", Type: field.TypeString, Size: 2147483647},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "uses", Type: field.TypeInt, Default: 0},
		{Name: "credits_spent", Type: field.TypeInt, Default: 0},
		{Name: "last_used_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// APITokensTable holds the schema information for the "api_tokens" table.
	APITokensTable = &schema.Table{
		Name:       "api_tokens",
		Columns:    APITokensColumns,
		PrimaryKey: []*schema.Column{APITokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_tokens_users_api_tokens",
				Columns:    []*schema.Column{APITokensColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CreditsColumns holds the columns for the "credits" table.
	CreditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "remaining_amount", Type: field.TypeInt32},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "stripe_line_item_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "replenished_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "credit_type_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// CreditsTable holds the schema information for the "credits" table.
	CreditsTable = &schema.Table{
		Name:       "credits",
		Columns:    CreditsColumns,
		PrimaryKey: []*schema.Column{CreditsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credits_credit_types_credits",
				Columns:    []*schema.Column{CreditsColumns[7]},
				RefColumns: []*schema.Column{CreditTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "credits_users_credits",
				Columns:    []*schema.Column{CreditsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "credit_expires_at_user_id_remaining_amount",
				Unique:  false,
				Columns: []*schema.Column{CreditsColumns[2], CreditsColumns[8], CreditsColumns[1]},
			},
		},
	}
	// CreditTypesColumns holds the columns for the "credit_types" table.
	CreditTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "amount", Type: field.TypeInt32},
		{Name: "stripe_product_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"free", "subscription", "one_time"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CreditTypesTable holds the schema information for the "credit_types" table.
	CreditTypesTable = &schema.Table{
		Name:       "credit_types",
		Columns:    CreditTypesColumns,
		PrimaryKey: []*schema.Column{CreditTypesColumns[0]},
	}
	// DeviceInfoColumns holds the columns for the "device_info" table.
	DeviceInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "os", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "browser", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// DeviceInfoTable holds the schema information for the "device_info" table.
	DeviceInfoTable = &schema.Table{
		Name:       "device_info",
		Columns:    DeviceInfoColumns,
		PrimaryKey: []*schema.Column{DeviceInfoColumns[0]},
	}
	// DisposableEmailsColumns holds the columns for the "disposable_emails" table.
	DisposableEmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "domain", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// DisposableEmailsTable holds the schema information for the "disposable_emails" table.
	DisposableEmailsTable = &schema.Table{
		Name:       "disposable_emails",
		Columns:    DisposableEmailsColumns,
		PrimaryKey: []*schema.Column{DisposableEmailsColumns[0]},
	}
	// GenerationsColumns holds the columns for the "generations" table.
	GenerationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "width", Type: field.TypeInt32},
		{Name: "height", Type: field.TypeInt32},
		{Name: "inference_steps", Type: field.TypeInt32},
		{Name: "guidance_scale", Type: field.TypeFloat32},
		{Name: "num_outputs", Type: field.TypeInt32},
		{Name: "nsfw_count", Type: field.TypeInt32, Default: 0},
		{Name: "seed", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"queued", "started", "succeeded", "failed"}},
		{Name: "failure_reason", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "country_code", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "init_image_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "prompt_strength", Type: field.TypeFloat32, Nullable: true},
		{Name: "was_auto_submitted", Type: field.TypeBool, Default: false},
		{Name: "stripe_product_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "api_token_id", Type: field.TypeUUID, Nullable: true},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "model_id", Type: field.TypeUUID},
		{Name: "negative_prompt_id", Type: field.TypeUUID, Nullable: true},
		{Name: "prompt_id", Type: field.TypeUUID, Nullable: true},
		{Name: "scheduler_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// GenerationsTable holds the schema information for the "generations" table.
	GenerationsTable = &schema.Table{
		Name:       "generations",
		Columns:    GenerationsColumns,
		PrimaryKey: []*schema.Column{GenerationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "generations_api_tokens_generations",
				Columns:    []*schema.Column{GenerationsColumns[19]},
				RefColumns: []*schema.Column{APITokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_device_info_generations",
				Columns:    []*schema.Column{GenerationsColumns[20]},
				RefColumns: []*schema.Column{DeviceInfoColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_generation_models_generations",
				Columns:    []*schema.Column{GenerationsColumns[21]},
				RefColumns: []*schema.Column{GenerationModelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_negative_prompts_generations",
				Columns:    []*schema.Column{GenerationsColumns[22]},
				RefColumns: []*schema.Column{NegativePromptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_prompts_generations",
				Columns:    []*schema.Column{GenerationsColumns[23]},
				RefColumns: []*schema.Column{PromptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_schedulers_generations",
				Columns:    []*schema.Column{GenerationsColumns[24]},
				RefColumns: []*schema.Column{SchedulersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "generations_users_generations",
				Columns:    []*schema.Column{GenerationsColumns[25]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "generation_user_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationsColumns[25], GenerationsColumns[17]},
			},
			{
				Name:    "generation_user_id_status_created_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationsColumns[25], GenerationsColumns[8], GenerationsColumns[17]},
			},
			{
				Name:    "generation_user_id_status",
				Unique:  false,
				Columns: []*schema.Column{GenerationsColumns[25], GenerationsColumns[8]},
			},
			{
				Name:    "generation_created_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationsColumns[17]},
			},
			{
				Name:    "generation_updated_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationsColumns[18]},
			},
		},
	}
	// GenerationModelsColumns holds the columns for the "generation_models" table.
	GenerationModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name_in_worker", Type: field.TypeString, Size: 2147483647},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "is_hidden", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// GenerationModelsTable holds the schema information for the "generation_models" table.
	GenerationModelsTable = &schema.Table{
		Name:       "generation_models",
		Columns:    GenerationModelsColumns,
		PrimaryKey: []*schema.Column{GenerationModelsColumns[0]},
	}
	// GenerationOutputsColumns holds the columns for the "generation_outputs" table.
	GenerationOutputsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "image_path", Type: field.TypeString, Size: 2147483647},
		{Name: "upscaled_image_path", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "gallery_status", Type: field.TypeEnum, Enums: []string{"not_submitted", "submitted", "approved", "rejected"}, Default: "not_submitted"},
		{Name: "is_favorited", Type: field.TypeBool, Default: false},
		{Name: "has_embeddings", Type: field.TypeBool, Default: false},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "generation_id", Type: field.TypeUUID},
	}
	// GenerationOutputsTable holds the schema information for the "generation_outputs" table.
	GenerationOutputsTable = &schema.Table{
		Name:       "generation_outputs",
		Columns:    GenerationOutputsColumns,
		PrimaryKey: []*schema.Column{GenerationOutputsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "generation_outputs_generations_generation_outputs",
				Columns:    []*schema.Column{GenerationOutputsColumns[9]},
				RefColumns: []*schema.Column{GenerationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "generationoutput_id_gallery_status",
				Unique:  false,
				Columns: []*schema.Column{GenerationOutputsColumns[0], GenerationOutputsColumns[3]},
			},
			{
				Name:    "generationoutput_gallery_status",
				Unique:  false,
				Columns: []*schema.Column{GenerationOutputsColumns[3]},
			},
			{
				Name:    "generationoutput_created_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationOutputsColumns[7]},
			},
			{
				Name:    "generationoutput_updated_at",
				Unique:  false,
				Columns: []*schema.Column{GenerationOutputsColumns[8]},
			},
			{
				Name:    "generationoutput_generation_id",
				Unique:  false,
				Columns: []*schema.Column{GenerationOutputsColumns[9]},
			},
		},
	}
	// NegativePromptsColumns holds the columns for the "negative_prompts" table.
	NegativePromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NegativePromptsTable holds the schema information for the "negative_prompts" table.
	NegativePromptsTable = &schema.Table{
		Name:       "negative_prompts",
		Columns:    NegativePromptsColumns,
		PrimaryKey: []*schema.Column{NegativePromptsColumns[0]},
	}
	// PromptsColumns holds the columns for the "prompts" table.
	PromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PromptsTable holds the schema information for the "prompts" table.
	PromptsTable = &schema.Table{
		Name:       "prompts",
		Columns:    PromptsColumns,
		PrimaryKey: []*schema.Column{PromptsColumns[0]},
	}
	// SchedulersColumns holds the columns for the "schedulers" table.
	SchedulersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name_in_worker", Type: field.TypeString, Size: 2147483647},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "is_hidden", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SchedulersTable holds the schema information for the "schedulers" table.
	SchedulersTable = &schema.Table{
		Name:       "schedulers",
		Columns:    SchedulersColumns,
		PrimaryKey: []*schema.Column{SchedulersColumns[0]},
	}
	// UpscalesColumns holds the columns for the "upscales" table.
	UpscalesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "width", Type: field.TypeInt32},
		{Name: "height", Type: field.TypeInt32},
		{Name: "scale", Type: field.TypeInt32},
		{Name: "country_code", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"queued", "started", "succeeded", "failed"}},
		{Name: "failure_reason", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "stripe_product_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "system_generated", Type: field.TypeBool, Default: false},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "api_token_id", Type: field.TypeUUID, Nullable: true},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "model_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// UpscalesTable holds the schema information for the "upscales" table.
	UpscalesTable = &schema.Table{
		Name:       "upscales",
		Columns:    UpscalesColumns,
		PrimaryKey: []*schema.Column{UpscalesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upscales_api_tokens_upscales",
				Columns:    []*schema.Column{UpscalesColumns[13]},
				RefColumns: []*schema.Column{APITokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upscales_device_info_upscales",
				Columns:    []*schema.Column{UpscalesColumns[14]},
				RefColumns: []*schema.Column{DeviceInfoColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upscales_upscale_models_upscales",
				Columns:    []*schema.Column{UpscalesColumns[15]},
				RefColumns: []*schema.Column{UpscaleModelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upscales_users_upscales",
				Columns:    []*schema.Column{UpscalesColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UpscaleModelsColumns holds the columns for the "upscale_models" table.
	UpscaleModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name_in_worker", Type: field.TypeString, Size: 2147483647},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "is_hidden", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UpscaleModelsTable holds the schema information for the "upscale_models" table.
	UpscaleModelsTable = &schema.Table{
		Name:       "upscale_models",
		Columns:    UpscaleModelsColumns,
		PrimaryKey: []*schema.Column{UpscaleModelsColumns[0]},
	}
	// UpscaleOutputsColumns holds the columns for the "upscale_outputs" table.
	UpscaleOutputsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "image_path", Type: field.TypeString, Size: 2147483647},
		{Name: "input_image_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "generation_output_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "upscale_id", Type: field.TypeUUID},
	}
	// UpscaleOutputsTable holds the schema information for the "upscale_outputs" table.
	UpscaleOutputsTable = &schema.Table{
		Name:       "upscale_outputs",
		Columns:    UpscaleOutputsColumns,
		PrimaryKey: []*schema.Column{UpscaleOutputsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upscale_outputs_generation_outputs_upscale_outputs",
				Columns:    []*schema.Column{UpscaleOutputsColumns[6]},
				RefColumns: []*schema.Column{GenerationOutputsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "upscale_outputs_upscales_upscale_outputs",
				Columns:    []*schema.Column{UpscaleOutputsColumns[7]},
				RefColumns: []*schema.Column{UpscalesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "upscaleoutput_image_path",
				Unique:  false,
				Columns: []*schema.Column{UpscaleOutputsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Size: 2147483647},
		{Name: "stripe_customer_id", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "active_product_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "last_sign_in_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_seen_at", Type: field.TypeTime},
		{Name: "banned_at", Type: field.TypeTime, Nullable: true},
		{Name: "scheduled_for_deletion_on", Type: field.TypeTime, Nullable: true},
		{Name: "data_deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role_name", Type: field.TypeEnum, Enums: []string{"SUPER_ADMIN", "GALLERY_ADMIN"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_users_user_roles",
				Columns:    []*schema.Column{UserRolesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APITokensTable,
		CreditsTable,
		CreditTypesTable,
		DeviceInfoTable,
		DisposableEmailsTable,
		GenerationsTable,
		GenerationModelsTable,
		GenerationOutputsTable,
		NegativePromptsTable,
		PromptsTable,
		SchedulersTable,
		UpscalesTable,
		UpscaleModelsTable,
		UpscaleOutputsTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	APITokensTable.ForeignKeys[0].RefTable = UsersTable
	APITokensTable.Annotation = &entsql.Annotation{
		Table: "api_tokens",
	}
	CreditsTable.ForeignKeys[0].RefTable = CreditTypesTable
	CreditsTable.ForeignKeys[1].RefTable = UsersTable
	CreditsTable.Annotation = &entsql.Annotation{
		Table: "credits",
	}
	CreditTypesTable.Annotation = &entsql.Annotation{
		Table: "credit_types",
	}
	DeviceInfoTable.Annotation = &entsql.Annotation{
		Table: "device_info",
	}
	DisposableEmailsTable.Annotation = &entsql.Annotation{
		Table: "disposable_emails",
	}
	GenerationsTable.ForeignKeys[0].RefTable = APITokensTable
	GenerationsTable.ForeignKeys[1].RefTable = DeviceInfoTable
	GenerationsTable.ForeignKeys[2].RefTable = GenerationModelsTable
	GenerationsTable.ForeignKeys[3].RefTable = NegativePromptsTable
	GenerationsTable.ForeignKeys[4].RefTable = PromptsTable
	GenerationsTable.ForeignKeys[5].RefTable = SchedulersTable
	GenerationsTable.ForeignKeys[6].RefTable = UsersTable
	GenerationsTable.Annotation = &entsql.Annotation{
		Table: "generations",
	}
	GenerationModelsTable.Annotation = &entsql.Annotation{
		Table: "generation_models",
	}
	GenerationOutputsTable.ForeignKeys[0].RefTable = GenerationsTable
	GenerationOutputsTable.Annotation = &entsql.Annotation{
		Table: "generation_outputs",
	}
	NegativePromptsTable.Annotation = &entsql.Annotation{
		Table: "negative_prompts",
	}
	PromptsTable.Annotation = &entsql.Annotation{
		Table: "prompts",
	}
	SchedulersTable.Annotation = &entsql.Annotation{
		Table: "schedulers",
	}
	UpscalesTable.ForeignKeys[0].RefTable = APITokensTable
	UpscalesTable.ForeignKeys[1].RefTable = DeviceInfoTable
	UpscalesTable.ForeignKeys[2].RefTable = UpscaleModelsTable
	UpscalesTable.ForeignKeys[3].RefTable = UsersTable
	UpscalesTable.Annotation = &entsql.Annotation{
		Table: "upscales",
	}
	UpscaleModelsTable.Annotation = &entsql.Annotation{
		Table: "upscale_models",
	}
	UpscaleOutputsTable.ForeignKeys[0].RefTable = GenerationOutputsTable
	UpscaleOutputsTable.ForeignKeys[1].RefTable = UpscalesTable
	UpscaleOutputsTable.Annotation = &entsql.Annotation{
		Table: "upscale_outputs",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.Annotation = &entsql.Annotation{
		Table: "user_roles",
	}
}
