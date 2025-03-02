/*
STACKIT Secrets Manager API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package secretsmanager

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

func Test_secretsmanager_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateAcl", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/acls"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Acl{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		createAclPayload := CreateAclPayload{}

		resp, reqErr := apiClient.CreateAcl(context.Background(), projectId, instanceId).CreateAclPayload(createAclPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Instance{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		createInstancePayload := CreateInstancePayload{}

		resp, reqErr := apiClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService CreateUser", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/users"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := User{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		createUserPayload := CreateUserPayload{}

		resp, reqErr := apiClient.CreateUser(context.Background(), projectId, instanceId).CreateUserPayload(createUserPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteAcl", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/acls/{aclId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		aclIdValue := "aclId"
		path = strings.Replace(path, "{"+"aclId"+"}", url.PathEscape(ParameterValueToString(aclIdValue, "aclId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		aclId := "aclId"

		reqErr := apiClient.DeleteAcl(context.Background(), projectId, instanceId, aclId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService DeleteInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"

		reqErr := apiClient.DeleteInstance(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService DeleteUser", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/users/{userId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		userIdValue := "userId"
		path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(ParameterValueToString(userIdValue, "userId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		userId := "userId"

		reqErr := apiClient.DeleteUser(context.Background(), projectId, instanceId, userId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService GetAcl", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/acls/{aclId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		aclIdValue := "aclId"
		path = strings.Replace(path, "{"+"aclId"+"}", url.PathEscape(ParameterValueToString(aclIdValue, "aclId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Acl{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		aclId := "aclId"

		resp, reqErr := apiClient.GetAcl(context.Background(), projectId, instanceId, aclId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetAcls", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/acls"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := AclList{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"

		resp, reqErr := apiClient.GetAcls(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetInstance", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := Instance{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"

		resp, reqErr := apiClient.GetInstance(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetInstances", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := InstanceList{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"

		resp, reqErr := apiClient.GetInstances(context.Background(), projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetUser", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/users/{userId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		userIdValue := "userId"
		path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(ParameterValueToString(userIdValue, "userId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := User{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		userId := "userId"

		resp, reqErr := apiClient.GetUser(context.Background(), projectId, instanceId, userId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetUsers", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/users"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			data := UserList{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"

		resp, reqErr := apiClient.GetUsers(context.Background(), projectId, instanceId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
		if resp == nil {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService UpdateAcl", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/acls/{aclId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		aclIdValue := "aclId"
		path = strings.Replace(path, "{"+"aclId"+"}", url.PathEscape(ParameterValueToString(aclIdValue, "aclId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		aclId := "aclId"
		updateAclPayload := UpdateAclPayload{}

		reqErr := apiClient.UpdateAcl(context.Background(), projectId, instanceId, aclId).UpdateAclPayload(updateAclPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

	t.Run("Test DefaultApiService UpdateUser", func(t *testing.T) {
		path := "/v1/projects/{projectId}/instances/{instanceId}/users/{userId}"
		projectIdValue := "projectId"
		path = strings.Replace(path, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		instanceIdValue := "instanceId"
		path = strings.Replace(path, "{"+"instanceId"+"}", url.PathEscape(ParameterValueToString(instanceIdValue, "instanceId")), -1)
		userIdValue := "userId"
		path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(ParameterValueToString(userIdValue, "userId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for secretsmanager_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := "projectId"
		instanceId := "instanceId"
		userId := "userId"
		updateUserPayload := UpdateUserPayload{}

		reqErr := apiClient.UpdateUser(context.Background(), projectId, instanceId, userId).UpdateUserPayload(updateUserPayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", err)
		}
	})

}
