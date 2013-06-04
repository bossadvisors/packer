package vmware

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

const BuilderId = "mitchellh.vmware"

type Builder struct {
	config config
	runner multistep.Runner
}

type config struct {
	OutputDir string `mapstructure:"output_directory"`
}

func (b *Builder) Prepare(raw interface{}) (err error) {
	if b.config.OutputDir == "" {
		b.config.OutputDir = "vmware"
	}

	return nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook) packer.Artifact {
	steps := []multistep.Step{
		&stepPrepareOutputDir{},
		&stepCreateDisk{},
	}

	// Setup the state bag
	state := make(map[string]interface{})
	state["config"] = &b.config
	state["hook"] = hook
	state["ui"] = ui

	// Run!
	b.runner = &multistep.BasicRunner{Steps: steps}
	b.runner.Run(state)

	return nil
}

func (b *Builder) Cancel() {
}
