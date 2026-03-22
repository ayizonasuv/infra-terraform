package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/infra-terraform/infra-terraform/constants"
	"github.com/infra-terraform/infra-terraform/logger"
	"github.com/infra-terraform/infra-terraform/terraform"
)

func GetProjectRoot() string {
	return os.Getenv(constants.ENV_PROJECT_ROOT)
}

func GetProjectName() string {
	return filepath.Base(GetProjectRoot())
}

func GetTfWorkDir() string {
	return filepath.Join(GetProjectRoot(), terraform.TF_WORK_DIR)
}

func GetConfigFile() string {
	return filepath.Join(GetProjectRoot(), terraform.TF_CONFIG_FILE)
}

func GetBackendConfig() map[string]string {
	backendConfig := make(map[string]string)

	backendConfig[terraform.TF_BACKEND_CONFIG_KEY_ENGINE] = os.Getenv(constants.ENV_BACKEND_ENGINE)
	backendConfig[terraform.TF_BACKEND_CONFIG_KEY_BUCKET] = os.Getenv(constants.ENV_BACKEND_BUCKET)
	backendConfig[terraform.TF_BACKEND_CONFIG_KEY_DYNAMODB_TABLE] = os.Getenv(constants.ENV_BACKEND_DYNAMODB_TABLE)
	backendConfig[terraform.TF_BACKEND_CONFIG_KEY_ACCESS_KEY_ID] = os.Getenv(constants.ENV_BACKEND_ACCESS_KEY_ID)
	backendConfig[terraform.TF_BACKEND_CONFIG_KEY_SECRET_ACCESS_KEY] = os.Getenv(constants.ENV_BACKEND_SECRET_ACCESS_KEY)

	return backendConfig
}

func IsRemoteBackend() bool {
	return strings.ToLower(os.Getenv(constants.ENV_BACKEND_ENGINE)) == "remote"
}

func GetRemoteBackendConfig() (string, string) {
	return os.Getenv(constants.ENV_BACKEND_BUCKET), os.Getenv(constants.ENV_BACKEND_DYNAMODB_TABLE)
}

func GetLogger() *logger.Logger {
	return logger.NewLogger(GetProjectName())
}