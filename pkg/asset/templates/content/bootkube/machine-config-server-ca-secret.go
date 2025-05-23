package bootkube

import (
	"context"
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	machineConfigServerCASecretFileName = "machine-config-server-ca-secret.yaml.template" // #nosec G101
)

var _ asset.WritableAsset = (*MachineConfigServerCASecret)(nil)

// MachineConfigServerCASecret is the constant to represent contents of machine-config-server-ca-secret.yaml.template file.
type MachineConfigServerCASecret struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset.
func (t *MachineConfigServerCASecret) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *MachineConfigServerCASecret) Name() string {
	return "MachineConfigServerCASecret"
}

// Generate generates the actual files by this asset.
func (t *MachineConfigServerCASecret) Generate(_ context.Context, parents asset.Parents) error {
	fileName := machineConfigServerCASecretFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     data,
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *MachineConfigServerCASecret) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *MachineConfigServerCASecret) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, machineConfigServerCASecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
