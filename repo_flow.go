package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type FlowRepository ClientService

func (t FlowRepository) Search(ctx ApiContext, criteria Criteria) (*FlowCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/flow", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(FlowCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t FlowRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/flow", criteria)

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

func (t FlowRepository) Upsert(ctx ApiContext, entity []Flow) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow": {
		Entity:  "flow",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t FlowRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow": {
		Entity:  "flow",
		Action:  "delete",
		Payload: payload,
	}})
}

type Flow struct {
	Id string `json:"id,omitempty"`

	EventName string `json:"eventName,omitempty"`

	Invalid bool `json:"invalid,omitempty"`

	Sequences []FlowSequence `json:"sequences,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Active bool `json:"active,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type FlowCollection struct {
	EntityCollection

	Data []Flow `json:"data"`
}