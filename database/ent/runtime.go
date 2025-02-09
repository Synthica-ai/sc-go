// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/credit"
	"github.com/stablecog/sc-go/database/ent/credittype"
	"github.com/stablecog/sc-go/database/ent/deviceinfo"
	"github.com/stablecog/sc-go/database/ent/disposableemail"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/generationmodel"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/negativeprompt"
	"github.com/stablecog/sc-go/database/ent/prompt"
	"github.com/stablecog/sc-go/database/ent/scheduler"
	"github.com/stablecog/sc-go/database/ent/schema"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/upscalemodel"
	"github.com/stablecog/sc-go/database/ent/upscaleoutput"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/database/ent/userrole"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apitokenFields := schema.ApiToken{}.Fields()
	_ = apitokenFields
	// apitokenDescIsActive is the schema descriptor for is_active field.
	apitokenDescIsActive := apitokenFields[4].Descriptor()
	// apitoken.DefaultIsActive holds the default value on creation for the is_active field.
	apitoken.DefaultIsActive = apitokenDescIsActive.Default.(bool)
	// apitokenDescUses is the schema descriptor for uses field.
	apitokenDescUses := apitokenFields[5].Descriptor()
	// apitoken.DefaultUses holds the default value on creation for the uses field.
	apitoken.DefaultUses = apitokenDescUses.Default.(int)
	// apitokenDescCreditsSpent is the schema descriptor for credits_spent field.
	apitokenDescCreditsSpent := apitokenFields[6].Descriptor()
	// apitoken.DefaultCreditsSpent holds the default value on creation for the credits_spent field.
	apitoken.DefaultCreditsSpent = apitokenDescCreditsSpent.Default.(int)
	// apitokenDescCreatedAt is the schema descriptor for created_at field.
	apitokenDescCreatedAt := apitokenFields[9].Descriptor()
	// apitoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	apitoken.DefaultCreatedAt = apitokenDescCreatedAt.Default.(func() time.Time)
	// apitokenDescUpdatedAt is the schema descriptor for updated_at field.
	apitokenDescUpdatedAt := apitokenFields[10].Descriptor()
	// apitoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apitoken.DefaultUpdatedAt = apitokenDescUpdatedAt.Default.(func() time.Time)
	// apitoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apitoken.UpdateDefaultUpdatedAt = apitokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apitokenDescID is the schema descriptor for id field.
	apitokenDescID := apitokenFields[0].Descriptor()
	// apitoken.DefaultID holds the default value on creation for the id field.
	apitoken.DefaultID = apitokenDescID.Default.(func() uuid.UUID)
	creditFields := schema.Credit{}.Fields()
	_ = creditFields
	// creditDescReplenishedAt is the schema descriptor for replenished_at field.
	creditDescReplenishedAt := creditFields[4].Descriptor()
	// credit.DefaultReplenishedAt holds the default value on creation for the replenished_at field.
	credit.DefaultReplenishedAt = creditDescReplenishedAt.Default.(func() time.Time)
	// creditDescCreatedAt is the schema descriptor for created_at field.
	creditDescCreatedAt := creditFields[7].Descriptor()
	// credit.DefaultCreatedAt holds the default value on creation for the created_at field.
	credit.DefaultCreatedAt = creditDescCreatedAt.Default.(func() time.Time)
	// creditDescUpdatedAt is the schema descriptor for updated_at field.
	creditDescUpdatedAt := creditFields[8].Descriptor()
	// credit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	credit.DefaultUpdatedAt = creditDescUpdatedAt.Default.(func() time.Time)
	// credit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	credit.UpdateDefaultUpdatedAt = creditDescUpdatedAt.UpdateDefault.(func() time.Time)
	// creditDescID is the schema descriptor for id field.
	creditDescID := creditFields[0].Descriptor()
	// credit.DefaultID holds the default value on creation for the id field.
	credit.DefaultID = creditDescID.Default.(func() uuid.UUID)
	credittypeFields := schema.CreditType{}.Fields()
	_ = credittypeFields
	// credittypeDescCreatedAt is the schema descriptor for created_at field.
	credittypeDescCreatedAt := credittypeFields[6].Descriptor()
	// credittype.DefaultCreatedAt holds the default value on creation for the created_at field.
	credittype.DefaultCreatedAt = credittypeDescCreatedAt.Default.(func() time.Time)
	// credittypeDescUpdatedAt is the schema descriptor for updated_at field.
	credittypeDescUpdatedAt := credittypeFields[7].Descriptor()
	// credittype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	credittype.DefaultUpdatedAt = credittypeDescUpdatedAt.Default.(func() time.Time)
	// credittype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	credittype.UpdateDefaultUpdatedAt = credittypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// credittypeDescID is the schema descriptor for id field.
	credittypeDescID := credittypeFields[0].Descriptor()
	// credittype.DefaultID holds the default value on creation for the id field.
	credittype.DefaultID = credittypeDescID.Default.(func() uuid.UUID)
	deviceinfoFields := schema.DeviceInfo{}.Fields()
	_ = deviceinfoFields
	// deviceinfoDescCreatedAt is the schema descriptor for created_at field.
	deviceinfoDescCreatedAt := deviceinfoFields[4].Descriptor()
	// deviceinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	deviceinfo.DefaultCreatedAt = deviceinfoDescCreatedAt.Default.(func() time.Time)
	// deviceinfoDescUpdatedAt is the schema descriptor for updated_at field.
	deviceinfoDescUpdatedAt := deviceinfoFields[5].Descriptor()
	// deviceinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deviceinfo.DefaultUpdatedAt = deviceinfoDescUpdatedAt.Default.(func() time.Time)
	// deviceinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deviceinfo.UpdateDefaultUpdatedAt = deviceinfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deviceinfoDescID is the schema descriptor for id field.
	deviceinfoDescID := deviceinfoFields[0].Descriptor()
	// deviceinfo.DefaultID holds the default value on creation for the id field.
	deviceinfo.DefaultID = deviceinfoDescID.Default.(func() uuid.UUID)
	disposableemailFields := schema.DisposableEmail{}.Fields()
	_ = disposableemailFields
	// disposableemailDescCreatedAt is the schema descriptor for created_at field.
	disposableemailDescCreatedAt := disposableemailFields[2].Descriptor()
	// disposableemail.DefaultCreatedAt holds the default value on creation for the created_at field.
	disposableemail.DefaultCreatedAt = disposableemailDescCreatedAt.Default.(func() time.Time)
	// disposableemailDescUpdatedAt is the schema descriptor for updated_at field.
	disposableemailDescUpdatedAt := disposableemailFields[3].Descriptor()
	// disposableemail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	disposableemail.DefaultUpdatedAt = disposableemailDescUpdatedAt.Default.(func() time.Time)
	// disposableemail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	disposableemail.UpdateDefaultUpdatedAt = disposableemailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// disposableemailDescID is the schema descriptor for id field.
	disposableemailDescID := disposableemailFields[0].Descriptor()
	// disposableemail.DefaultID holds the default value on creation for the id field.
	disposableemail.DefaultID = disposableemailDescID.Default.(func() uuid.UUID)
	generationFields := schema.Generation{}.Fields()
	_ = generationFields
	// generationDescNsfwCount is the schema descriptor for nsfw_count field.
	generationDescNsfwCount := generationFields[6].Descriptor()
	// generation.DefaultNsfwCount holds the default value on creation for the nsfw_count field.
	generation.DefaultNsfwCount = generationDescNsfwCount.Default.(int32)
	// generationDescWasAutoSubmitted is the schema descriptor for was_auto_submitted field.
	generationDescWasAutoSubmitted := generationFields[13].Descriptor()
	// generation.DefaultWasAutoSubmitted holds the default value on creation for the was_auto_submitted field.
	generation.DefaultWasAutoSubmitted = generationDescWasAutoSubmitted.Default.(bool)
	// generationDescCreatedAt is the schema descriptor for created_at field.
	generationDescCreatedAt := generationFields[24].Descriptor()
	// generation.DefaultCreatedAt holds the default value on creation for the created_at field.
	generation.DefaultCreatedAt = generationDescCreatedAt.Default.(func() time.Time)
	// generationDescUpdatedAt is the schema descriptor for updated_at field.
	generationDescUpdatedAt := generationFields[25].Descriptor()
	// generation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	generation.DefaultUpdatedAt = generationDescUpdatedAt.Default.(func() time.Time)
	// generation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	generation.UpdateDefaultUpdatedAt = generationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// generationDescID is the schema descriptor for id field.
	generationDescID := generationFields[0].Descriptor()
	// generation.DefaultID holds the default value on creation for the id field.
	generation.DefaultID = generationDescID.Default.(func() uuid.UUID)
	generationmodelFields := schema.GenerationModel{}.Fields()
	_ = generationmodelFields
	// generationmodelDescIsActive is the schema descriptor for is_active field.
	generationmodelDescIsActive := generationmodelFields[2].Descriptor()
	// generationmodel.DefaultIsActive holds the default value on creation for the is_active field.
	generationmodel.DefaultIsActive = generationmodelDescIsActive.Default.(bool)
	// generationmodelDescIsDefault is the schema descriptor for is_default field.
	generationmodelDescIsDefault := generationmodelFields[3].Descriptor()
	// generationmodel.DefaultIsDefault holds the default value on creation for the is_default field.
	generationmodel.DefaultIsDefault = generationmodelDescIsDefault.Default.(bool)
	// generationmodelDescIsHidden is the schema descriptor for is_hidden field.
	generationmodelDescIsHidden := generationmodelFields[4].Descriptor()
	// generationmodel.DefaultIsHidden holds the default value on creation for the is_hidden field.
	generationmodel.DefaultIsHidden = generationmodelDescIsHidden.Default.(bool)
	// generationmodelDescCreatedAt is the schema descriptor for created_at field.
	generationmodelDescCreatedAt := generationmodelFields[5].Descriptor()
	// generationmodel.DefaultCreatedAt holds the default value on creation for the created_at field.
	generationmodel.DefaultCreatedAt = generationmodelDescCreatedAt.Default.(func() time.Time)
	// generationmodelDescUpdatedAt is the schema descriptor for updated_at field.
	generationmodelDescUpdatedAt := generationmodelFields[6].Descriptor()
	// generationmodel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	generationmodel.DefaultUpdatedAt = generationmodelDescUpdatedAt.Default.(func() time.Time)
	// generationmodel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	generationmodel.UpdateDefaultUpdatedAt = generationmodelDescUpdatedAt.UpdateDefault.(func() time.Time)
	// generationmodelDescID is the schema descriptor for id field.
	generationmodelDescID := generationmodelFields[0].Descriptor()
	// generationmodel.DefaultID holds the default value on creation for the id field.
	generationmodel.DefaultID = generationmodelDescID.Default.(func() uuid.UUID)
	generationoutputFields := schema.GenerationOutput{}.Fields()
	_ = generationoutputFields
	// generationoutputDescIsFavorited is the schema descriptor for is_favorited field.
	generationoutputDescIsFavorited := generationoutputFields[4].Descriptor()
	// generationoutput.DefaultIsFavorited holds the default value on creation for the is_favorited field.
	generationoutput.DefaultIsFavorited = generationoutputDescIsFavorited.Default.(bool)
	// generationoutputDescHasEmbeddings is the schema descriptor for has_embeddings field.
	generationoutputDescHasEmbeddings := generationoutputFields[5].Descriptor()
	// generationoutput.DefaultHasEmbeddings holds the default value on creation for the has_embeddings field.
	generationoutput.DefaultHasEmbeddings = generationoutputDescHasEmbeddings.Default.(bool)
	// generationoutputDescCreatedAt is the schema descriptor for created_at field.
	generationoutputDescCreatedAt := generationoutputFields[8].Descriptor()
	// generationoutput.DefaultCreatedAt holds the default value on creation for the created_at field.
	generationoutput.DefaultCreatedAt = generationoutputDescCreatedAt.Default.(func() time.Time)
	// generationoutputDescUpdatedAt is the schema descriptor for updated_at field.
	generationoutputDescUpdatedAt := generationoutputFields[9].Descriptor()
	// generationoutput.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	generationoutput.DefaultUpdatedAt = generationoutputDescUpdatedAt.Default.(func() time.Time)
	// generationoutput.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	generationoutput.UpdateDefaultUpdatedAt = generationoutputDescUpdatedAt.UpdateDefault.(func() time.Time)
	// generationoutputDescID is the schema descriptor for id field.
	generationoutputDescID := generationoutputFields[0].Descriptor()
	// generationoutput.DefaultID holds the default value on creation for the id field.
	generationoutput.DefaultID = generationoutputDescID.Default.(func() uuid.UUID)
	negativepromptFields := schema.NegativePrompt{}.Fields()
	_ = negativepromptFields
	// negativepromptDescCreatedAt is the schema descriptor for created_at field.
	negativepromptDescCreatedAt := negativepromptFields[2].Descriptor()
	// negativeprompt.DefaultCreatedAt holds the default value on creation for the created_at field.
	negativeprompt.DefaultCreatedAt = negativepromptDescCreatedAt.Default.(func() time.Time)
	// negativepromptDescUpdatedAt is the schema descriptor for updated_at field.
	negativepromptDescUpdatedAt := negativepromptFields[3].Descriptor()
	// negativeprompt.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	negativeprompt.DefaultUpdatedAt = negativepromptDescUpdatedAt.Default.(func() time.Time)
	// negativeprompt.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	negativeprompt.UpdateDefaultUpdatedAt = negativepromptDescUpdatedAt.UpdateDefault.(func() time.Time)
	// negativepromptDescID is the schema descriptor for id field.
	negativepromptDescID := negativepromptFields[0].Descriptor()
	// negativeprompt.DefaultID holds the default value on creation for the id field.
	negativeprompt.DefaultID = negativepromptDescID.Default.(func() uuid.UUID)
	promptFields := schema.Prompt{}.Fields()
	_ = promptFields
	// promptDescCreatedAt is the schema descriptor for created_at field.
	promptDescCreatedAt := promptFields[2].Descriptor()
	// prompt.DefaultCreatedAt holds the default value on creation for the created_at field.
	prompt.DefaultCreatedAt = promptDescCreatedAt.Default.(func() time.Time)
	// promptDescUpdatedAt is the schema descriptor for updated_at field.
	promptDescUpdatedAt := promptFields[3].Descriptor()
	// prompt.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	prompt.DefaultUpdatedAt = promptDescUpdatedAt.Default.(func() time.Time)
	// prompt.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	prompt.UpdateDefaultUpdatedAt = promptDescUpdatedAt.UpdateDefault.(func() time.Time)
	// promptDescID is the schema descriptor for id field.
	promptDescID := promptFields[0].Descriptor()
	// prompt.DefaultID holds the default value on creation for the id field.
	prompt.DefaultID = promptDescID.Default.(func() uuid.UUID)
	schedulerFields := schema.Scheduler{}.Fields()
	_ = schedulerFields
	// schedulerDescIsActive is the schema descriptor for is_active field.
	schedulerDescIsActive := schedulerFields[2].Descriptor()
	// scheduler.DefaultIsActive holds the default value on creation for the is_active field.
	scheduler.DefaultIsActive = schedulerDescIsActive.Default.(bool)
	// schedulerDescIsDefault is the schema descriptor for is_default field.
	schedulerDescIsDefault := schedulerFields[3].Descriptor()
	// scheduler.DefaultIsDefault holds the default value on creation for the is_default field.
	scheduler.DefaultIsDefault = schedulerDescIsDefault.Default.(bool)
	// schedulerDescIsHidden is the schema descriptor for is_hidden field.
	schedulerDescIsHidden := schedulerFields[4].Descriptor()
	// scheduler.DefaultIsHidden holds the default value on creation for the is_hidden field.
	scheduler.DefaultIsHidden = schedulerDescIsHidden.Default.(bool)
	// schedulerDescCreatedAt is the schema descriptor for created_at field.
	schedulerDescCreatedAt := schedulerFields[5].Descriptor()
	// scheduler.DefaultCreatedAt holds the default value on creation for the created_at field.
	scheduler.DefaultCreatedAt = schedulerDescCreatedAt.Default.(func() time.Time)
	// schedulerDescUpdatedAt is the schema descriptor for updated_at field.
	schedulerDescUpdatedAt := schedulerFields[6].Descriptor()
	// scheduler.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	scheduler.DefaultUpdatedAt = schedulerDescUpdatedAt.Default.(func() time.Time)
	// scheduler.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	scheduler.UpdateDefaultUpdatedAt = schedulerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// schedulerDescID is the schema descriptor for id field.
	schedulerDescID := schedulerFields[0].Descriptor()
	// scheduler.DefaultID holds the default value on creation for the id field.
	scheduler.DefaultID = schedulerDescID.Default.(func() uuid.UUID)
	upscaleFields := schema.Upscale{}.Fields()
	_ = upscaleFields
	// upscaleDescSystemGenerated is the schema descriptor for system_generated field.
	upscaleDescSystemGenerated := upscaleFields[8].Descriptor()
	// upscale.DefaultSystemGenerated holds the default value on creation for the system_generated field.
	upscale.DefaultSystemGenerated = upscaleDescSystemGenerated.Default.(bool)
	// upscaleDescCreatedAt is the schema descriptor for created_at field.
	upscaleDescCreatedAt := upscaleFields[15].Descriptor()
	// upscale.DefaultCreatedAt holds the default value on creation for the created_at field.
	upscale.DefaultCreatedAt = upscaleDescCreatedAt.Default.(func() time.Time)
	// upscaleDescUpdatedAt is the schema descriptor for updated_at field.
	upscaleDescUpdatedAt := upscaleFields[16].Descriptor()
	// upscale.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	upscale.DefaultUpdatedAt = upscaleDescUpdatedAt.Default.(func() time.Time)
	// upscale.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	upscale.UpdateDefaultUpdatedAt = upscaleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// upscaleDescID is the schema descriptor for id field.
	upscaleDescID := upscaleFields[0].Descriptor()
	// upscale.DefaultID holds the default value on creation for the id field.
	upscale.DefaultID = upscaleDescID.Default.(func() uuid.UUID)
	upscalemodelFields := schema.UpscaleModel{}.Fields()
	_ = upscalemodelFields
	// upscalemodelDescIsActive is the schema descriptor for is_active field.
	upscalemodelDescIsActive := upscalemodelFields[2].Descriptor()
	// upscalemodel.DefaultIsActive holds the default value on creation for the is_active field.
	upscalemodel.DefaultIsActive = upscalemodelDescIsActive.Default.(bool)
	// upscalemodelDescIsDefault is the schema descriptor for is_default field.
	upscalemodelDescIsDefault := upscalemodelFields[3].Descriptor()
	// upscalemodel.DefaultIsDefault holds the default value on creation for the is_default field.
	upscalemodel.DefaultIsDefault = upscalemodelDescIsDefault.Default.(bool)
	// upscalemodelDescIsHidden is the schema descriptor for is_hidden field.
	upscalemodelDescIsHidden := upscalemodelFields[4].Descriptor()
	// upscalemodel.DefaultIsHidden holds the default value on creation for the is_hidden field.
	upscalemodel.DefaultIsHidden = upscalemodelDescIsHidden.Default.(bool)
	// upscalemodelDescCreatedAt is the schema descriptor for created_at field.
	upscalemodelDescCreatedAt := upscalemodelFields[5].Descriptor()
	// upscalemodel.DefaultCreatedAt holds the default value on creation for the created_at field.
	upscalemodel.DefaultCreatedAt = upscalemodelDescCreatedAt.Default.(func() time.Time)
	// upscalemodelDescUpdatedAt is the schema descriptor for updated_at field.
	upscalemodelDescUpdatedAt := upscalemodelFields[6].Descriptor()
	// upscalemodel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	upscalemodel.DefaultUpdatedAt = upscalemodelDescUpdatedAt.Default.(func() time.Time)
	// upscalemodel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	upscalemodel.UpdateDefaultUpdatedAt = upscalemodelDescUpdatedAt.UpdateDefault.(func() time.Time)
	// upscalemodelDescID is the schema descriptor for id field.
	upscalemodelDescID := upscalemodelFields[0].Descriptor()
	// upscalemodel.DefaultID holds the default value on creation for the id field.
	upscalemodel.DefaultID = upscalemodelDescID.Default.(func() uuid.UUID)
	upscaleoutputFields := schema.UpscaleOutput{}.Fields()
	_ = upscaleoutputFields
	// upscaleoutputDescCreatedAt is the schema descriptor for created_at field.
	upscaleoutputDescCreatedAt := upscaleoutputFields[6].Descriptor()
	// upscaleoutput.DefaultCreatedAt holds the default value on creation for the created_at field.
	upscaleoutput.DefaultCreatedAt = upscaleoutputDescCreatedAt.Default.(func() time.Time)
	// upscaleoutputDescUpdatedAt is the schema descriptor for updated_at field.
	upscaleoutputDescUpdatedAt := upscaleoutputFields[7].Descriptor()
	// upscaleoutput.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	upscaleoutput.DefaultUpdatedAt = upscaleoutputDescUpdatedAt.Default.(func() time.Time)
	// upscaleoutput.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	upscaleoutput.UpdateDefaultUpdatedAt = upscaleoutputDescUpdatedAt.UpdateDefault.(func() time.Time)
	// upscaleoutputDescID is the schema descriptor for id field.
	upscaleoutputDescID := upscaleoutputFields[0].Descriptor()
	// upscaleoutput.DefaultID holds the default value on creation for the id field.
	upscaleoutput.DefaultID = upscaleoutputDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLastSeenAt is the schema descriptor for last_seen_at field.
	userDescLastSeenAt := userFields[5].Descriptor()
	// user.DefaultLastSeenAt holds the default value on creation for the last_seen_at field.
	user.DefaultLastSeenAt = userDescLastSeenAt.Default.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	userroleFields := schema.UserRole{}.Fields()
	_ = userroleFields
	// userroleDescCreatedAt is the schema descriptor for created_at field.
	userroleDescCreatedAt := userroleFields[3].Descriptor()
	// userrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	userrole.DefaultCreatedAt = userroleDescCreatedAt.Default.(func() time.Time)
	// userroleDescUpdatedAt is the schema descriptor for updated_at field.
	userroleDescUpdatedAt := userroleFields[4].Descriptor()
	// userrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userrole.DefaultUpdatedAt = userroleDescUpdatedAt.Default.(func() time.Time)
	// userrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userrole.UpdateDefaultUpdatedAt = userroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userroleDescID is the schema descriptor for id field.
	userroleDescID := userroleFields[0].Descriptor()
	// userrole.DefaultID holds the default value on creation for the id field.
	userrole.DefaultID = userroleDescID.Default.(func() uuid.UUID)
}
