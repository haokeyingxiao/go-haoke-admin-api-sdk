package go_shopware_admin_sdk

import (
	"net/http"
)

type SalesChannelShippingMethodRepository ClientService

func (t SalesChannelShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelShippingMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-shipping-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelShippingMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-shipping-method", criteria)

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

func (t SalesChannelShippingMethodRepository) Upsert(ctx ApiContext, entity []SalesChannelShippingMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_shipping_method": {
		Entity:  "sales_channel_shipping_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_shipping_method": {
		Entity:  "sales_channel_shipping_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelShippingMethod struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`
}

type SalesChannelShippingMethodCollection struct {
	EntityCollection

	Data []SalesChannelShippingMethod `json:"data"`
}