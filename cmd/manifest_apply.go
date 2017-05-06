package cmd

import (
	"os"

	"github.com/bpicode/fritzctl/assert"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/bpicode/fritzctl/manifest"
	"github.com/mitchellh/cli"
)

type manifestApplyCommand struct {
}

func (cmd *manifestApplyCommand) Help() string {
	return "Apply a given manifest against the state of the FRITZ!Box."
}

func (cmd *manifestApplyCommand) Synopsis() string {
	return "apply a given manifest"
}

func (cmd *manifestApplyCommand) Run(args []string) int {
	assert.StringSliceHasAtLeast(args, 1, "insufficient input: path to input manifest expected.")
	target := cmd.parseManifest(args[0])
	api := fritz.HomeAutomation(clientLogin())
	src := cmd.obtainSourcePlan(api)
	err := manifest.AhaApiApplier(api).Apply(src, target)
	assert.NoError(err, "plan (dry-run) of manifest was not successful:", err)
	return 0
}

// ManifestApply is a factory creating commands for applying manifest files.
func ManifestApply() (cli.Command, error) {
	p := manifestApplyCommand{}
	return &p, nil
}

func (cmd *manifestApplyCommand) parseManifest(filename string) *manifest.Plan {
	file, err := os.Open(filename)
	assert.NoError(err, "cannot open manifest file:", err)
	defer file.Close()
	p, err := manifest.Parse(file)
	assert.NoError(err, "cannot parse manifest file:", err)
	return p
}

func (cmd *manifestApplyCommand) obtainSourcePlan(api fritz.HomeAutomationApi) *manifest.Plan {
	l, err := api.ListDevices()
	assert.NoError(err, "cannot obtain device data:", err)
	return manifest.ConvertDevicelist(l)
}
