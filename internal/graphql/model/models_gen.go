// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddonList struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Endpoint string `json:"endpoint"`
}

type AddonStatusInput struct {
	Selector     *MeshType `json:"selector"`
	Context      *string   `json:"context"`
	TargetStatus Status    `json:"targetStatus"`
}

type Container struct {
	ControlPlaneMemberName string           `json:"controlPlaneMemberName"`
	ContainerName          string           `json:"containerName"`
	Image                  string           `json:"image"`
	Status                 *ContainerStatus `json:"status"`
	Ports                  []*ContainerPort `json:"ports"`
	Resources              interface{}      `json:"resources"`
}

type ContainerPort struct {
	Name          *string `json:"name"`
	ContainerPort int     `json:"containerPort"`
	Protocol      string  `json:"protocol"`
}

type ContainerStatus struct {
	ContainerStatusName string      `json:"containerStatusName"`
	Image               string      `json:"image"`
	State               interface{} `json:"state"`
	LastState           interface{} `json:"lastState"`
	Ready               bool        `json:"ready"`
	RestartCount        interface{} `json:"restartCount"`
	Started             bool        `json:"started"`
	ImageID             interface{} `json:"imageID"`
	ContainerID         interface{} `json:"containerID"`
}

type ControlPlane struct {
	Name    string                `json:"name"`
	Members []*ControlPlaneMember `json:"members"`
}

type ControlPlaneMember struct {
	Name       string       `json:"name"`
	Component  string       `json:"component"`
	Version    string       `json:"version"`
	Namespace  string       `json:"namespace"`
	DataPlanes []*Container `json:"data_planes"`
}

type DataPlane struct {
	Name    string       `json:"name"`
	Proxies []*Container `json:"proxies"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type K8sContext struct {
	ID string `json:"id"`
}

type KctlDescribeDetails struct {
	Describe *string `json:"describe"`
	Ctxid    *string `json:"ctxid"`
}

type MesheryResult struct {
	MesheryID          *string                `json:"meshery_id"`
	Name               *string                `json:"name"`
	Mesh               *string                `json:"mesh"`
	PerformanceProfile *string                `json:"performance_profile"`
	TestID             *string                `json:"test_id"`
	RunnerResults      map[string]interface{} `json:"runner_results"`
	ServerMetrics      *string                `json:"server_metrics"`
	ServerBoardConfig  *string                `json:"server_board_config"`
	TestStartTime      *string                `json:"test_start_time"`
	UserID             *string                `json:"user_id"`
	UpdatedAt          *string                `json:"updated_at"`
	CreatedAt          *string                `json:"created_at"`
}

type NameSpace struct {
	Namespace string `json:"namespace"`
}

type OAMCapability struct {
	OamDefinition interface{} `json:"oam_definition"`
	ID            *string     `json:"id"`
	OamRefSchema  *string     `json:"oam_ref_schema"`
	Host          *string     `json:"host"`
	Restricted    *bool       `json:"restricted"`
	Metadata      interface{} `json:"metadata"`
}

type OperatorControllerStatus struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  Status `json:"status"`
	Error   *Error `json:"error"`
}

type OperatorStatus struct {
	Status      Status                      `json:"status"`
	Version     string                      `json:"version"`
	Controllers []*OperatorControllerStatus `json:"controllers"`
	Error       *Error                      `json:"error"`
	Context     *string                     `json:"context"`
}

type OperatorStatusInput struct {
	TargetStatus Status  `json:"targetStatus"`
	Context      *string `json:"context"`
}

type PageFilter struct {
	Page     string  `json:"page"`
	PageSize string  `json:"pageSize"`
	Order    *string `json:"order"`
	Search   *string `json:"search"`
	From     *string `json:"from"`
	To       *string `json:"to"`
}

type PatternLocation struct {
	Branch *string `json:"branch"`
	Host   *string `json:"host"`
	Path   *string `json:"path"`
	Type   *string `json:"type"`
}

type PatternPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Patterns   []*PatternResult `json:"patterns"`
}

type PatternResult struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	UserID      string           `json:"user_id"`
	Location    *PatternLocation `json:"location"`
	PatternFile string           `json:"pattern_file"`
	CanSupport  bool             `json:"canSupport"`
	Errmsg      *string          `json:"errmsg"`
	CreatedAt   *string          `json:"created_at"`
	UpdatedAt   *string          `json:"updated_at"`
}

type PerfPageProfiles struct {
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalCount int            `json:"total_count"`
	Profiles   []*PerfProfile `json:"profiles"`
}

type PerfPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Results    []*MesheryResult `json:"results"`
}

type PerfProfile struct {
	ConcurrentRequest int       `json:"concurrent_request"`
	CreatedAt         *string   `json:"created_at"`
	Duration          string    `json:"duration"`
	Endpoints         []*string `json:"endpoints"`
	ID                string    `json:"id"`
	LastRun           *string   `json:"last_run"`
	LoadGenerators    []*string `json:"load_generators"`
	Name              *string   `json:"name"`
	QPS               *int      `json:"qps"`
	TotalResults      *int      `json:"total_results"`
	UpdatedAt         *string   `json:"updated_at"`
	UserID            string    `json:"user_id"`
	RequestHeaders    *string   `json:"request_headers"`
	RequestCookies    *string   `json:"request_cookies"`
	RequestBody       *string   `json:"request_body"`
	ContentType       *string   `json:"content_type"`
	ServiceMesh       *string   `json:"service_mesh"`
}

type ReSyncActions struct {
	ClearDb string `json:"clearDB"`
	ReSync  string `json:"ReSync"`
}

type ServiceMeshFilter struct {
	Type         *MeshType `json:"type"`
	K8sServerIDs []string  `json:"k8sServerIDs"`
}

type MeshType string

const (
	MeshTypeAllMesh            MeshType = "ALL_MESH"
	MeshTypeInvalidMesh        MeshType = "INVALID_MESH"
	MeshTypeAppMesh            MeshType = "APP_MESH"
	MeshTypeCitrixServiceMesh  MeshType = "CITRIX_SERVICE_MESH"
	MeshTypeConsul             MeshType = "CONSUL"
	MeshTypeIstio              MeshType = "ISTIO"
	MeshTypeKuma               MeshType = "KUMA"
	MeshTypeLinkerd            MeshType = "LINKERD"
	MeshTypeTraefikMesh        MeshType = "TRAEFIK_MESH"
	MeshTypeOctarine           MeshType = "OCTARINE"
	MeshTypeNetworkServiceMesh MeshType = "NETWORK_SERVICE_MESH"
	MeshTypeTanzu              MeshType = "TANZU"
	MeshTypeOpenServiceMesh    MeshType = "OPEN_SERVICE_MESH"
	MeshTypeNginxServiceMesh   MeshType = "NGINX_SERVICE_MESH"
	MeshTypeCiliumServiceMesh  MeshType = "CILIUM_SERVICE_MESH"
)

var AllMeshType = []MeshType{
	MeshTypeAllMesh,
	MeshTypeInvalidMesh,
	MeshTypeAppMesh,
	MeshTypeCitrixServiceMesh,
	MeshTypeConsul,
	MeshTypeIstio,
	MeshTypeKuma,
	MeshTypeLinkerd,
	MeshTypeTraefikMesh,
	MeshTypeOctarine,
	MeshTypeNetworkServiceMesh,
	MeshTypeTanzu,
	MeshTypeOpenServiceMesh,
	MeshTypeNginxServiceMesh,
	MeshTypeCiliumServiceMesh,
}

func (e MeshType) IsValid() bool {
	switch e {
	case MeshTypeAllMesh, MeshTypeInvalidMesh, MeshTypeAppMesh, MeshTypeCitrixServiceMesh, MeshTypeConsul, MeshTypeIstio, MeshTypeKuma, MeshTypeLinkerd, MeshTypeTraefikMesh, MeshTypeOctarine, MeshTypeNetworkServiceMesh, MeshTypeTanzu, MeshTypeOpenServiceMesh, MeshTypeNginxServiceMesh, MeshTypeCiliumServiceMesh:
		return true
	}
	return false
}

func (e MeshType) String() string {
	return string(e)
}

func (e *MeshType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshType", str)
	}
	return nil
}

func (e MeshType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusEnabled    Status = "ENABLED"
	StatusConnected  Status = "CONNECTED"
	StatusDisabled   Status = "DISABLED"
	StatusProcessing Status = "PROCESSING"
	StatusUnknown    Status = "UNKNOWN"
)

var AllStatus = []Status{
	StatusEnabled,
	StatusConnected,
	StatusDisabled,
	StatusProcessing,
	StatusUnknown,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusEnabled, StatusConnected, StatusDisabled, StatusProcessing, StatusUnknown:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
