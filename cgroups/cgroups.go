package cgroups

import (
	sub "github.com/iamwwc/wwcdocker/cgroups/subsystems"
)

func CreateCgroup(id string, pid int) error {
	for _, sub := range sub.Subsystems {
		if err := sub.Apply(id,pid); err != nil {
			return err
		}
	}
	return nil
}

func RemoveFromCgroup(id string) error {
	for _, sub := range sub.Subsystems {
		if err := sub.Remove(id); err != nil {
			return err
		}
	}
	return nil
}

func CreateAndSetLimit(id string, pid int, config *sub.ResourceConfig) error {
	if err := CreateCgroup(id, pid); err != nil {
		return err
	}

	if err := SetResourceLimit(id,config); err != nil {
		return err
	}
	return nil
}

func SetResourceLimit(id string, config *sub.ResourceConfig) error {
	for _, sub := range sub.Subsystems {
		if err := sub.SetLimit(id, config); err != nil {
			return err
		}
	}
	return nil
}

