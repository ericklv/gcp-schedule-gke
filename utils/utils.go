package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type KBody struct {
	Cluster      string    `json:"cluster" validate:"required"`
	Zone         string    `json:"zone" validate:"required"`
	Project      string    `json:"project" validate:"required"`
	Namespace    string    `json:"namespace" validate:"required"`
	Nodepool     string    `json:"nodepool" validate:"required"`
	MachineSpecs *KMachine `json:"machineSpecs,omitempty"`
}

type KMachine struct {
	Type     string `json:"type" validate:"required"`
	Disk     string `json:"disk" validate:"required"`
	NumNodes string `json:"numNodes" validate:"required"`
}

type Request struct {
	Params
	KBody
}

type Params struct {
	Action string `param:"action"`
}

type Error_ struct {
	Key   string
	Error string
}

const g_cluster = "gcloud container clusters get-credentials"
const g_pool = "gcloud container node-pools"
const k_scale = "kubectl scale deploy -n"
const _z = "--zone"
const _p = "--project"
const _c = "--cluster"
const _r = "--replicas"
const _a = "--all"
const _y = "yes |"

const _m = "--machine-type"
const _d = "--disk-size"
const _n = "--num-nodes"

func Validator(t interface{}) error {
	validate := validator.New()
	err_ := validate.Struct(t)

	if err_ != nil {
		return err_
	}
	return nil
}

func CmdDown(b KBody) (string, error) {
	if err := Validator(b); err != nil {
		return "", err
	}

	g_cred := strings.Join([]string{g_cluster, b.Cluster, _z, b.Zone, _p, b.Project}, " ")
	scale_n := strings.Join([]string{k_scale, b.Namespace, _r, "0", _a}, " ")
	drop_pool := strings.Join([]string{_y, g_pool, "delete", b.Nodepool, _c, b.Cluster, _z, b.Zone, "--async"}, " ")

	return strings.Join([]string{g_cred, "&&", scale_n, "&&", drop_pool}, " "), nil
}

func CmdUp(b KBody) (string, error) {
	m_specs := b.MachineSpecs

	if m_specs == nil {
		return "", errors.New("machineSpecs are required")
	}

	if err := Validator(b); err != nil {
		return "", err
	}

	m_type := m_specs.Type
	m_disk := m_specs.Disk
	m_num_n := m_specs.NumNodes

	g_cred := strings.Join([]string{g_cluster, b.Cluster, _z, b.Zone, _p, b.Project}, " ")
	create_pool := strings.Join([]string{g_pool, "create", b.Nodepool, _c, b.Cluster, _z, b.Zone, _m, m_type, _d, m_disk, _n, m_num_n}, " ")
	scale_n := strings.Join([]string{k_scale, b.Namespace, _r, "1", _a}, " ")

	return strings.Join([]string{g_cred, "&&", create_pool, "&&", scale_n}, " "), nil
}

func S200(msg string) Response {
	return Response{200, msg}
}

func S4xx(msg string) Response {
	return Response{400, msg}
}

func S5xx(msg string) Response {
	return Response{500, msg}
}
