# vmware-rest-proxy

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vmware-rest-proxy)](https://artifacthub.io/packages/search?repo=vmware-rest-proxy) ![Version: 0.1.3](https://img.shields.io/badge/Version-0.1.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square)

## Introduction

This helm chart installs the [VMware REST Proxy](https://github.com/dodevops/vmware-rest-proxy) for easy access
to commonly used vSphere information.

## Installation

Use

    helm install <name of release> vmware-rest-proy --repo https://dodevops.io/vmware-rest-proxy

to install this chart.

## Configuration

Set config.baseUrl to the URL of your vCenter server. See other `config.`-parameters for more configuration.

**Homepage:** <https://github.com/dodevops/vmware-rest-proxy>

## Source Code

* <https://github.com/dodevops/vmware-rest-proxy/tree/main/charts/vmware-rest-proxy>
* <https://github.com/dodevops/vmware-rest-proxy/blob/main/Dockerfile>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| config.baseUrl | string | `""` | base URL of the vCenter server |
| config.logLevel | string | `"INFO"` | Maximum log level to use (see (https://pkg.go.dev/github.com/sirupsen/logrus#readme-level-logging)) [INFO] |
| config.tlsSkipVerify | string | `"false"` | If set, will disable TLS verification for the API client |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"ghcr.io/dodevops/vmware-rest-proxy"` |  |
| image.tag | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.className | string | `""` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths[0].path | string | `"/"` |  |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| securityContext | object | `{}` |  |
| service.port | int | `8080` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |

