package go_haoke_admin_sdk

import (
	"net/http"

	"time"
)

type ImportExportProfileRepository ClientService

func (t ImportExportProfileRepository) Search(ctx ApiContext, criteria Criteria) (*ImportExportProfileCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/import-export-profile", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ImportExportProfileCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ImportExportProfileRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ImportExportProfileCollection, *http.Response, error) {
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

func (t ImportExportProfileRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/import-export-profile", criteria)

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

func (t ImportExportProfileRepository) Upsert(ctx ApiContext, entity []ImportExportProfile) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile": {
		Entity:  "import_export_profile",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ImportExportProfileRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile": {
		Entity:  "import_export_profile",
		Action:  "delete",
		Payload: payload,
	}})
}

type ImportExportProfile struct {
	Type string `json:"type,omitempty"`

	Enclosure string `json:"enclosure,omitempty"`

	Name string `json:"name,omitempty"`

	Label string `json:"label,omitempty"`

	SystemDefault bool `json:"systemDefault,omitempty"`

	FileType string `json:"fileType,omitempty"`

	Mapping interface{} `json:"mapping,omitempty"`

	ImportExportLogs []ImportExportLog `json:"importExportLogs,omitempty"`

	Translations []ImportExportProfileTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	SourceEntity string `json:"sourceEntity,omitempty"`

	Delimiter string `json:"delimiter,omitempty"`

	UpdateBy interface{} `json:"updateBy,omitempty"`

	Config interface{} `json:"config,omitempty"`
}

type ImportExportProfileCollection struct {
	EntityCollection

	Data []ImportExportProfile `json:"data"`
}
