package go_haoke_admin_sdk

import (
	"net/http"

	"time"
)

type ProductConfiguratorSettingRepository ClientService

func (t ProductConfiguratorSettingRepository) Search(ctx ApiContext, criteria Criteria) (*ProductConfiguratorSettingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-configurator-setting", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductConfiguratorSettingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductConfiguratorSettingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductConfiguratorSettingCollection, *http.Response, error) {
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

func (t ProductConfiguratorSettingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-configurator-setting", criteria)

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

func (t ProductConfiguratorSettingRepository) Upsert(ctx ApiContext, entity []ProductConfiguratorSetting) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_configurator_setting": {
		Entity:  "product_configurator_setting",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductConfiguratorSettingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_configurator_setting": {
		Entity:  "product_configurator_setting",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductConfiguratorSetting struct {
	Id string `json:"id,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Media *Media `json:"media,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	OptionId string `json:"optionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	ProductId string `json:"productId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Option *PropertyGroupOption `json:"option,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Position float64 `json:"position,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductConfiguratorSettingCollection struct {
	EntityCollection

	Data []ProductConfiguratorSetting `json:"data"`
}
