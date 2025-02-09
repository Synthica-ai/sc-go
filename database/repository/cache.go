package repository

import (
	"github.com/stablecog/sc-go/database/ent/disposableemail"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/shared"
)

// Update the cache from the database
func (r *Repository) UpdateCache() error {
	generationModels, err := r.GetAllGenerationModels()
	if err != nil {
		log.Error("Failed to get generation_models", "err", err)
		return err
	}
	shared.GetCache().UpdateGenerationModels(generationModels)

	upscaleModels, err := r.GetAllUpscaleModels()
	if err != nil {
		log.Error("Failed to get upscale_models", "err", err)
		return err
	}
	shared.GetCache().UpdateUpscaleModels(upscaleModels)

	schedulers, err := r.GetAllSchedulers()
	if err != nil {
		log.Error("Failed to get schedulers", "err", err)
		return err
	}
	shared.GetCache().UpdateSchedulers(schedulers)

	admins, err := r.GetSuperAdminUserIDs()
	if err != nil {
		log.Error("Failed to get super admins", "err", err)
		return err
	}
	shared.GetCache().SetAdminUUIDs(admins)

	disposableEmailDomains, err := r.DB.DisposableEmail.Query().Select(disposableemail.FieldDomain).All(r.Ctx)
	if err != nil {
		log.Error("Failed to get disposable email domains", "err", err)
		return err
	}
	disposableEmailDomainsStr := make([]string, len(disposableEmailDomains))
	for i, domain := range disposableEmailDomains {
		disposableEmailDomainsStr[i] = domain.Domain
	}
	shared.GetCache().UpdateDisposableEmailDomains(disposableEmailDomainsStr)
	return nil
}
