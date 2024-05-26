package go_haoke_admin_sdk

import (
	"net/http"

	"time"
)

type PropertyGroupTranslationRepository ClientService

func (t PropertyGroupTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PropertyGroupTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/property-group-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PropertyGroupTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PropertyGroupTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PropertyGroupTranslationCollection, *http.Response, error) {
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

func (t PropertyGroupTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/property-group-translation", criteria)

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

func (t PropertyGroupTranslationRepository) Upsert(ctx ApiContext, entity []PropertyGroupTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_translation": {
		Entity:  "property_group_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PropertyGroupTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_translation": {
		Entity:  "property_group_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PropertyGroupTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	Description string `json:"description,omitempty"`

	Position float64 `json:"position,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	PropertyGroup *PropertyGroup `json:"propertyGroup,omitempty"`

	Name string `json:"name,omitempty"`

	PropertyGroupId string `json:"propertyGroupId,omitempty"`
}

type PropertyGroupTranslationCollection struct {
	EntityCollection

	Data []PropertyGroupTranslation `json:"data"`
}
