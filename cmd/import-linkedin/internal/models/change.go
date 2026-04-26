// Package models defines the data structures for LinkedIn import entities.
package models

// ChangeType represents the type of change detected.
type ChangeType string

const (
	// ChangeTypeAdded indicates a new entity was found.
	ChangeTypeAdded ChangeType = "added"
	// ChangeTypeModified indicates an existing entity was modified.
	ChangeTypeModified ChangeType = "modified"
	// ChangeTypeRemoved indicates an entity was removed.
	ChangeTypeRemoved ChangeType = "removed"
)

// Change represents a difference between LinkedIn data and current config.
type Change struct {
	EntityType    string   `yaml:"entity_type" json:"entity_type"`
	EntityID      string   `yaml:"entity_id" json:"entity_id"`
	OldValue      any      `yaml:"old_value,omitempty" json:"old_value,omitempty"`
	NewValue      any      `yaml:"new_value" json:"new_value"`
	ChangeType    string   `yaml:"change_type" json:"change_type"`
	FieldsChanged []string `yaml:"fields_changed,omitempty" json:"fields_changed,omitempty"`
}

// IsAccepted returns true if the change should be applied.
// This is set during the interactive confirmation phase.
func (c Change) IsAccepted() bool {
	// Changes are pending until explicitly accepted or rejected
	// This method is a placeholder for future state tracking
	return true
}

// Summary returns a human-readable summary of the change.
func (c Change) Summary() string {
	switch ChangeType(c.ChangeType) {
	case ChangeTypeAdded:
		return "Added new " + c.EntityType + ": " + c.EntityID
	case ChangeTypeModified:
		return "Modified " + c.EntityType + ": " + c.EntityID
	case ChangeTypeRemoved:
		return "Removed " + c.EntityType + ": " + c.EntityID
	default:
		return "Unknown change for " + c.EntityType + ": " + c.EntityID
	}
}
