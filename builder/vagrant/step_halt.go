package vagrant

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepHalt struct {
	TeardownMethod string
	Provider       string
	GlobalID       string
}

func (s *StepHalt) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(VagrantDriver)
	ui := state.Get("ui").(packer.Ui)

	ui.Say(fmt.Sprintf("%sing Vagrant box...", s.TeardownMethod))

	var err error
	if s.TeardownMethod == "halt" {
		err = driver.Halt(s.GlobalID)
	} else if s.TeardownMethod == "suspend" {
		err = driver.Suspend(s.GlobalID)
	} else if s.TeardownMethod == "destroy" {
		err = driver.Destroy(s.GlobalID)
	} else {
		// Should never get here because of template validation
		state.Put("error", fmt.Errorf("Invalid teardown method selected; must be either halt, suspend, or destory."))
		return multistep.ActionHalt
	}
	if err != nil {
		state.Put("error", fmt.Errorf("Error halting Vagrant machine; please try to do this manually"))
		return multistep.ActionHalt
	}
	//continue
	return multistep.ActionContinue
}

func (s *StepHalt) Cleanup(state multistep.StateBag) {}
