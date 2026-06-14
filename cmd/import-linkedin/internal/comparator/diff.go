// Package comparator handles comparison between LinkedIn data and current config.
package comparator

import (
	"fmt"
	"reflect"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

// Diff represents the differences between LinkedIn data and current config.
type Diff struct {
	Added          []models.Change
	Modified       []models.Change
	Removed        []models.Change
	Experiences    EntityDiff[models.Experience]
	Education      EntityDiff[models.Education]
	Certifications EntityDiff[models.Certification]
}

// EntityDiff represents changes for a specific entity type.
type EntityDiff[T any] struct {
	Added    []T
	Modified []ChangePair[T]
	Removed  []T
}

// ChangePair represents an old and new version of an entity.
type ChangePair[T any] struct {
	Old T
	New T
}

// NewDiff creates a new empty Diff.
func NewDiff() *Diff {
	return &Diff{
		Added:    make([]models.Change, 0),
		Modified: make([]models.Change, 0),
		Removed:  make([]models.Change, 0),
		Experiences: EntityDiff[models.Experience]{
			Added:    make([]models.Experience, 0),
			Modified: make([]ChangePair[models.Experience], 0),
			Removed:  make([]models.Experience, 0),
		},
		Education: EntityDiff[models.Education]{
			Added:    make([]models.Education, 0),
			Modified: make([]ChangePair[models.Education], 0),
			Removed:  make([]models.Education, 0),
		},
		Certifications: EntityDiff[models.Certification]{
			Added:    make([]models.Certification, 0),
			Modified: make([]ChangePair[models.Certification], 0),
			Removed:  make([]models.Certification, 0),
		},
	}
}

// HasChanges returns true if there are any differences.
func (d *Diff) HasChanges() bool {
	return d.CountAdded() > 0 || d.CountModified() > 0 || d.CountRemoved() > 0
}

// CountAdded returns the total number of added items.
func (d *Diff) CountAdded() int {
	return len(d.Experiences.Added) + len(d.Education.Added) + len(d.Certifications.Added)
}

// CountModified returns the total number of modified items.
func (d *Diff) CountModified() int {
	return len(d.Experiences.Modified) + len(d.Education.Modified) + len(d.Certifications.Modified)
}

// CountRemoved returns the total number of removed items.
func (d *Diff) CountRemoved() int {
	return len(d.Experiences.Removed) + len(d.Education.Removed) + len(d.Certifications.Removed)
}

// CompareExperiences compares experiences from LinkedIn with current config.
func CompareExperiences(linkedin []models.Experience, current []models.Experience) EntityDiff[models.Experience] {
	result := EntityDiff[models.Experience]{
		Added:    make([]models.Experience, 0),
		Modified: make([]ChangePair[models.Experience], 0),
		Removed:  make([]models.Experience, 0),
	}

	currentMap := make(map[string]models.Experience)
	for _, exp := range current {
		currentMap[exp.ID()] = exp
	}

	linkedinMap := make(map[string]models.Experience)
	for _, exp := range linkedin {
		linkedinMap[exp.ID()] = exp
	}

	for id, linkedinExp := range linkedinMap {
		if currentExp, exists := currentMap[id]; exists {
			if !experiencesEqual(linkedinExp, currentExp) {
				result.Modified = append(result.Modified, ChangePair[models.Experience]{
					Old: currentExp,
					New: linkedinExp,
				})
			}
		} else {
			result.Added = append(result.Added, linkedinExp)
		}
	}

	for id, currentExp := range currentMap {
		if _, exists := linkedinMap[id]; !exists {
			result.Removed = append(result.Removed, currentExp)
		}
	}

	return result
}

// CompareEducation compares education from LinkedIn with current config.
func CompareEducation(linkedin []models.Education, current []models.Education) EntityDiff[models.Education] {
	result := EntityDiff[models.Education]{
		Added:    make([]models.Education, 0),
		Modified: make([]ChangePair[models.Education], 0),
		Removed:  make([]models.Education, 0),
	}

	currentMap := make(map[string]models.Education)
	for _, edu := range current {
		currentMap[edu.ID()] = edu
	}

	linkedinMap := make(map[string]models.Education)
	for _, edu := range linkedin {
		linkedinMap[edu.ID()] = edu
	}

	for id, linkedinEdu := range linkedinMap {
		if currentEdu, exists := currentMap[id]; exists {
			if !educationEqual(linkedinEdu, currentEdu) {
				result.Modified = append(result.Modified, ChangePair[models.Education]{
					Old: currentEdu,
					New: linkedinEdu,
				})
			}
		} else {
			result.Added = append(result.Added, linkedinEdu)
		}
	}

	for id, currentEdu := range currentMap {
		if _, exists := linkedinMap[id]; !exists {
			result.Removed = append(result.Removed, currentEdu)
		}
	}

	return result
}

// CompareCertifications compares certifications from LinkedIn with current config.
func CompareCertifications(linkedin []models.Certification, current []models.Certification) EntityDiff[models.Certification] {
	result := EntityDiff[models.Certification]{
		Added:    make([]models.Certification, 0),
		Modified: make([]ChangePair[models.Certification], 0),
		Removed:  make([]models.Certification, 0),
	}

	currentMap := make(map[string]models.Certification)
	for _, cert := range current {
		currentMap[cert.ID()] = cert
	}

	linkedinMap := make(map[string]models.Certification)
	for _, cert := range linkedin {
		linkedinMap[cert.ID()] = cert
	}

	for id, linkedinCert := range linkedinMap {
		if currentCert, exists := currentMap[id]; exists {
			if !certificationsEqual(linkedinCert, currentCert) {
				result.Modified = append(result.Modified, ChangePair[models.Certification]{
					Old: currentCert,
					New: linkedinCert,
				})
			}
		} else {
			result.Added = append(result.Added, linkedinCert)
		}
	}

	for id, currentCert := range currentMap {
		if _, exists := linkedinMap[id]; !exists {
			result.Removed = append(result.Removed, currentCert)
		}
	}

	return result
}

// CompareAll compares all entities and returns a complete diff.
func CompareAll(linkedinExp []models.Experience, currentExp []models.Experience,
	linkedinEdu []models.Education, currentEdu []models.Education,
	linkedinCert []models.Certification, currentCert []models.Certification) *Diff {

	diff := NewDiff()
	diff.Experiences = CompareExperiences(linkedinExp, currentExp)
	diff.Education = CompareEducation(linkedinEdu, currentEdu)
	diff.Certifications = CompareCertifications(linkedinCert, currentCert)

	return diff
}

func experiencesEqual(a, b models.Experience) bool {
	return a.Company == b.Company &&
		a.Title == b.Title &&
		a.StartDate == b.StartDate &&
		a.EndDate == b.EndDate &&
		a.Location == b.Location &&
		reflect.DeepEqual(a.Description, b.Description)
}

func educationEqual(a, b models.Education) bool {
	return a.Institution == b.Institution &&
		a.Degree == b.Degree &&
		a.Field == b.Field &&
		a.StartDate == b.StartDate &&
		a.EndDate == b.EndDate &&
		reflect.DeepEqual(a.Description, b.Description)
}

func certificationsEqual(a, b models.Certification) bool {
	return a.Name == b.Name &&
		a.Organization == b.Organization &&
		a.IssueDate == b.IssueDate &&
		a.ExpirationDate == b.ExpirationDate &&
		a.CredentialID == b.CredentialID &&
		a.CredentialURL == b.CredentialURL
}

// GetChangedFields returns a list of field names that differ between two experiences.
func GetChangedFields(old, new models.Experience) []string {
	var fields []string
	if old.Company != new.Company {
		fields = append(fields, "company")
	}
	if old.Title != new.Title {
		fields = append(fields, "title")
	}
	if old.StartDate != new.StartDate {
		fields = append(fields, "start_date")
	}
	if old.EndDate != new.EndDate {
		fields = append(fields, "end_date")
	}
	if old.Location != new.Location {
		fields = append(fields, "location")
	}
	if !reflect.DeepEqual(old.Description, new.Description) {
		fields = append(fields, "description")
	}
	return fields
}

// FormatID returns a human-readable identifier for display.
func FormatID(entity any) string {
	switch e := entity.(type) {
	case models.Experience:
		return fmt.Sprintf("%s @ %s", e.Title, e.Company)
	case *models.Experience:
		if e == nil {
			return ""
		}
		return fmt.Sprintf("%s @ %s", e.Title, e.Company)
	case models.Education:
		return fmt.Sprintf("%s at %s", e.Degree, e.Institution)
	case *models.Education:
		if e == nil {
			return ""
		}
		return fmt.Sprintf("%s at %s", e.Degree, e.Institution)
	case models.Certification:
		return fmt.Sprintf("%s (%s)", e.Name, e.Organization)
	case *models.Certification:
		if e == nil {
			return ""
		}
		return fmt.Sprintf("%s (%s)", e.Name, e.Organization)
	default:
		return "Unknown"
	}
}
