/*
 * Model Registry REST API
 *
 * REST API for Model Registry to create and manage ML model metadata
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"

	model "github.com/kubeflow/model-registry/pkg/openapi"
)

// ModelRegistryServiceAPIRouter defines the required methods for binding the api requests to a responses for the ModelRegistryServiceAPI
// The ModelRegistryServiceAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ModelRegistryServiceAPIServicer to perform the required actions, then write the service results to the http response.
type ModelRegistryServiceAPIRouter interface {
	CreateEnvironmentInferenceService(http.ResponseWriter, *http.Request)
	CreateInferenceService(http.ResponseWriter, *http.Request)
	CreateInferenceServiceServe(http.ResponseWriter, *http.Request)
	CreateModelArtifact(http.ResponseWriter, *http.Request)
	CreateModelVersion(http.ResponseWriter, *http.Request)
	CreateModelVersionArtifact(http.ResponseWriter, *http.Request)
	CreateRegisteredModel(http.ResponseWriter, *http.Request)
	CreateRegisteredModelVersion(http.ResponseWriter, *http.Request)
	CreateServingEnvironment(http.ResponseWriter, *http.Request)
	FindInferenceService(http.ResponseWriter, *http.Request)
	FindModelArtifact(http.ResponseWriter, *http.Request)
	FindModelVersion(http.ResponseWriter, *http.Request)
	FindRegisteredModel(http.ResponseWriter, *http.Request)
	FindServingEnvironment(http.ResponseWriter, *http.Request)
	GetEnvironmentInferenceServices(http.ResponseWriter, *http.Request)
	GetInferenceService(http.ResponseWriter, *http.Request)
	GetInferenceServiceModel(http.ResponseWriter, *http.Request)
	GetInferenceServiceServes(http.ResponseWriter, *http.Request)
	GetInferenceServiceVersion(http.ResponseWriter, *http.Request)
	GetInferenceServices(http.ResponseWriter, *http.Request)
	GetModelArtifact(http.ResponseWriter, *http.Request)
	GetModelArtifacts(http.ResponseWriter, *http.Request)
	GetModelVersion(http.ResponseWriter, *http.Request)
	GetModelVersionArtifacts(http.ResponseWriter, *http.Request)
	GetModelVersions(http.ResponseWriter, *http.Request)
	GetRegisteredModel(http.ResponseWriter, *http.Request)
	GetRegisteredModelVersions(http.ResponseWriter, *http.Request)
	GetRegisteredModels(http.ResponseWriter, *http.Request)
	GetServingEnvironment(http.ResponseWriter, *http.Request)
	GetServingEnvironments(http.ResponseWriter, *http.Request)
	UpdateInferenceService(http.ResponseWriter, *http.Request)
	UpdateModelArtifact(http.ResponseWriter, *http.Request)
	UpdateModelVersion(http.ResponseWriter, *http.Request)
	UpdateRegisteredModel(http.ResponseWriter, *http.Request)
	UpdateServingEnvironment(http.ResponseWriter, *http.Request)
}

// ModelRegistryServiceAPIServicer defines the api actions for the ModelRegistryServiceAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ModelRegistryServiceAPIServicer interface {
	CreateEnvironmentInferenceService(context.Context, string, model.InferenceServiceCreate) (ImplResponse, error)
	CreateInferenceService(context.Context, model.InferenceServiceCreate) (ImplResponse, error)
	CreateInferenceServiceServe(context.Context, string, model.ServeModelCreate) (ImplResponse, error)
	CreateModelArtifact(context.Context, model.ModelArtifactCreate) (ImplResponse, error)
	CreateModelVersion(context.Context, model.ModelVersionCreate) (ImplResponse, error)
	CreateModelVersionArtifact(context.Context, string, model.Artifact) (ImplResponse, error)
	CreateRegisteredModel(context.Context, model.RegisteredModelCreate) (ImplResponse, error)
	CreateRegisteredModelVersion(context.Context, string, model.ModelVersion) (ImplResponse, error)
	CreateServingEnvironment(context.Context, model.ServingEnvironmentCreate) (ImplResponse, error)
	FindInferenceService(context.Context, string, string, string) (ImplResponse, error)
	FindModelArtifact(context.Context, string, string, string) (ImplResponse, error)
	FindModelVersion(context.Context, string, string, string) (ImplResponse, error)
	FindRegisteredModel(context.Context, string, string) (ImplResponse, error)
	FindServingEnvironment(context.Context, string, string) (ImplResponse, error)
	GetEnvironmentInferenceServices(context.Context, string, string, string, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetInferenceService(context.Context, string) (ImplResponse, error)
	GetInferenceServiceModel(context.Context, string) (ImplResponse, error)
	GetInferenceServiceServes(context.Context, string, string, string, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetInferenceServiceVersion(context.Context, string) (ImplResponse, error)
	GetInferenceServices(context.Context, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetModelArtifact(context.Context, string) (ImplResponse, error)
	GetModelArtifacts(context.Context, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetModelVersion(context.Context, string) (ImplResponse, error)
	GetModelVersionArtifacts(context.Context, string, string, string, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetModelVersions(context.Context, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetRegisteredModel(context.Context, string) (ImplResponse, error)
	GetRegisteredModelVersions(context.Context, string, string, string, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetRegisteredModels(context.Context, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	GetServingEnvironment(context.Context, string) (ImplResponse, error)
	GetServingEnvironments(context.Context, string, model.OrderByField, model.SortOrder, string) (ImplResponse, error)
	UpdateInferenceService(context.Context, string, model.InferenceServiceUpdate) (ImplResponse, error)
	UpdateModelArtifact(context.Context, string, model.ModelArtifactUpdate) (ImplResponse, error)
	UpdateModelVersion(context.Context, string, model.ModelVersion) (ImplResponse, error)
	UpdateRegisteredModel(context.Context, string, model.RegisteredModelUpdate) (ImplResponse, error)
	UpdateServingEnvironment(context.Context, string, model.ServingEnvironmentUpdate) (ImplResponse, error)
}