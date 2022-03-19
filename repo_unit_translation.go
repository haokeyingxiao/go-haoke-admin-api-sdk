package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type UnitTranslationRepository ClientService

func (t UnitTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*UnitTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/unit-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UnitTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UnitTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*UnitTranslationCollection, *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
}

func (t UnitTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/unit-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UnitTranslationRepository) Upsert(ctx ApiContext, entity []UnitTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit_translation": {
		Entity:  "unit_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UnitTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit_translation": {
		Entity:  "unit_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type UnitTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	ShortCode string `json:"shortCode,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Unit *Unit `json:"unit,omitempty"`
}

type UnitTranslationCollection struct {
	EntityCollection

	Data []UnitTranslation `json:"data"`
}
