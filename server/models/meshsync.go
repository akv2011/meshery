package models

import (
	"github.com/layer5io/meshsync/pkg/model"
)

type MeshSyncResourcesAPIResponse struct {
	Page       int                        `json:"page"`
	PageSize   int                        `json:"page_size"`
	TotalCount int64                      `json:"total_count"`
	Resources  []model.KubernetesResource `json:"resources"`
}

type MeshSyncResourcesSummaryAPIResponse struct {
	Kinds []struct {
		Kind  string
		Count int64
	} `json:"kinds"`
	Namespaces []string `json:"namespaces"`
}
