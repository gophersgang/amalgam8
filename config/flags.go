package config

import (
	"time"

	"strings"

	"github.com/amalgam8/registry/cluster"
	"github.com/codegangsta/cli"
)

// Flag names
const (
	LogLevelFlag  = "log_level"
	LogFormatFlag = "log_format"

	AuthModeFlag     = "auth_mode"
	JWTSecretFlag    = "jwt_secret"
	RequireHTTPSFlag = "require_https"

	RestAPIPortFlag     = "api_port"
	ReplicationPortFlag = "replication_port"

	ReplicationFlag = "replication"
	SyncTimeoutFlag = "sync_timeout"

	ClusterDirectoryFlag = "cluster_dir"
	ClusterSizeFlag      = "cluster_size"

	NamespaceCapacityFlag = "namespace_capacity"
	DefaultTTLFlag        = "default_ttl"
	MaxTTLFlag            = "max_ttl"
	MinTTLFlag            = "min_ttl"
)

// Flags represents the set of supported flags
var Flags = []cli.Flag{

	cli.StringFlag{
		Name:   LogLevelFlag,
		EnvVar: envVarFromFlag(LogLevelFlag),
		Value:  "debug",
		Usage:  "Logging level. Supported values are: 'debug', 'info', 'warn', 'error', 'fatal', 'panic'",
	},

	cli.StringFlag{
		Name:   LogFormatFlag,
		EnvVar: envVarFromFlag(LogFormatFlag),
		Value:  "text",
		Usage:  "Logging format. Supported values are: 'text', 'json', 'logstash'",
	},

	cli.StringSliceFlag{
		Name:   AuthModeFlag,
		EnvVar: envVarFromFlag(AuthModeFlag),
		Usage:  "Authentication modes. Supported values are: 'trusted', 'jwt'",
	},

	cli.StringFlag{
		Name:   JWTSecretFlag,
		EnvVar: envVarFromFlag(JWTSecretFlag),
		Usage:  "Secret key for JWT authentication",
	},

	cli.BoolFlag{
		Name:   RequireHTTPSFlag,
		EnvVar: envVarFromFlag(RequireHTTPSFlag),
		Usage:  "Require clients to use HTTPS for API calls",
	},

	cli.IntFlag{
		Name:   RestAPIPortFlag,
		EnvVar: envVarFromFlag(RestAPIPortFlag),
		Value:  8080,
		Usage:  "REST API port number",
	},

	cli.IntFlag{
		Name:   ReplicationPortFlag,
		EnvVar: envVarFromFlag(ReplicationPortFlag),
		Value:  6100,
		Usage:  "Replication port number",
	},

	cli.BoolFlag{
		Name:   ReplicationFlag,
		EnvVar: envVarFromFlag(ReplicationFlag),
		Usage:  "Enable replication",
	},

	cli.DurationFlag{
		Name:   SyncTimeoutFlag,
		EnvVar: envVarFromFlag(SyncTimeoutFlag),
		Value:  30 * time.Second,
		Usage:  "Registry timeout for establishing peer synchronization connection",
	},

	cli.StringFlag{
		Name:   ClusterDirectoryFlag,
		EnvVar: envVarFromFlag(ClusterDirectoryFlag),
		Value:  cluster.DefaultDirectory,
		Usage:  "Filesystem directory for cluster membership",
	},

	cli.IntFlag{
		Name:   ClusterSizeFlag,
		EnvVar: envVarFromFlag(ClusterSizeFlag),
		Value:  0,
		Usage:  "Cluster minimal healthy size",
	},

	cli.DurationFlag{
		Name:   DefaultTTLFlag,
		EnvVar: envVarFromFlag(DefaultTTLFlag),
		Value:  30 * time.Second,
		Usage:  "Registry default TTL",
	},

	cli.DurationFlag{
		Name:   MaxTTLFlag,
		EnvVar: envVarFromFlag(MaxTTLFlag),
		Value:  10 * time.Minute,
		Usage:  "Registry maximum TTL",
	},

	cli.DurationFlag{
		Name:   MinTTLFlag,
		EnvVar: envVarFromFlag(MinTTLFlag),
		Value:  10 * time.Second,
		Usage:  "Registry minimum TTL",
	},

	cli.IntFlag{
		Name:   NamespaceCapacityFlag,
		EnvVar: envVarFromFlag(NamespaceCapacityFlag),
		Value:  -1,
		Usage:  "Registry namespace capacity, value of -1 indicates no capacity limit",
	},
}

// envVarFromFlag returns the environment variable bound to the given flag
func envVarFromFlag(name string) string {
	return strings.ToUpper(name)
}
