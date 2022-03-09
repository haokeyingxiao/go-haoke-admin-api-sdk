package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppActionButtonRepository ClientService

func (t AppActionButtonRepository) Search(ctx ApiContext, criteria Criteria) (*AppActionButtonCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-action-button", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppActionButtonCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppActionButtonRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-action-button", criteria)

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

func (t AppActionButtonRepository) Upsert(ctx ApiContext, entity []AppActionButton) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button": {
		Entity:  "app_action_button",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppActionButtonRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button": {
		Entity:  "app_action_button",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppActionButton struct {
	Id string `json:"id,omitempty"`

	Entity string `json:"entity,omitempty"`

	Action string `json:"action,omitempty"`

	Translations []AppActionButtonTranslation `json:"translations,omitempty"`

	App *App `json:"app,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	View string `json:"view,omitempty"`

	Url string `json:"url,omitempty"`

	OpenNewTab bool `json:"openNewTab,omitempty"`

	Label string `json:"label,omitempty"`

	AppId string `json:"appId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AppActionButtonCollection struct {
	EntityCollection

	Data []AppActionButton `json:"data"`
}