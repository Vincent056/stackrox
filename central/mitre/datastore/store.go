package datastore

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/mitre"
)

var (
	log = logging.LoggerForModule()
)

// MitreAttackReadOnlyDataStore provides functionality to read MITRE ATT&CK vectors.
// A vector represents MITRE tactics (why) and its techniques/sub-techniques (how).
//go:generate mockgen-wrapper
type MitreAttackReadOnlyDataStore interface {
	GetAll() []*storage.MitreAttackVector
	Get(id string) (*storage.MitreAttackVector, error)
}

type mitreAttackStoreImpl struct {
	// mitreAttackVectors stores MITRE ATT&CK vectors keyed by tactic ID.
	mitreAttackVectors map[string]*storage.MitreAttackVector
}

func newMitreAttackStore() *mitreAttackStoreImpl {
	s := &mitreAttackStoreImpl{
		mitreAttackVectors: make(map[string]*storage.MitreAttackVector),
	}
	// If ATT&CK data cannot be loaded, fail open.
	if err := s.loadBundledData(); err != nil {
		log.Errorf("MITRE ATT&CK data for system policies unavailable: %v", err)
	}
	return s
}

func (s *mitreAttackStoreImpl) GetAll() []*storage.MitreAttackVector {
	resp := make([]*storage.MitreAttackVector, 0, len(s.mitreAttackVectors))
	for _, vector := range s.mitreAttackVectors {
		resp = append(resp, vector)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].GetTactic().GetName() < resp[j].GetTactic().GetName()
	})

	return resp
}

func (s *mitreAttackStoreImpl) Get(id string) (*storage.MitreAttackVector, error) {
	if id == "" {
		return nil, errors.Wrap(errorhelpers.ErrInvalidArgs, "MITRE ATT&CK tactic ID must be provided")
	}

	v := s.mitreAttackVectors[id]
	if v == nil {
		return nil, errors.Wrapf(errorhelpers.ErrNotFound, "MITRE ATT&CK vector for tactic %q not found. Please check the tactic ID and retry.", id)
	}
	return v, nil
}

func (s *mitreAttackStoreImpl) loadBundledData() error {
	attackBundle, err := mitre.GetMitreBundle()
	if err != nil {
		return errors.Wrap(err, "loading default MITRE ATT&CK data")
	}

	// Flatten vectors from all matrices since we populate all enterprise.
	vectors := flattenMatrices(attackBundle.GetMatrices()...)
	for _, vector := range vectors {
		s.mitreAttackVectors[vector.GetTactic().GetId()] = vector
	}
	return nil
}

func flattenMatrices(matrices ...*storage.MitreAttackMatrix) []*storage.MitreAttackVector {
	tactics := make(map[string]*storage.MitreTactic)
	techniques := make(map[string]*storage.MitreTechnique)
	tacticsTechniques := make(map[string]map[string]struct{})
	for _, matrix := range matrices {
		for _, vector := range matrix.GetVectors() {
			tacticID := vector.GetTactic().GetId()
			if tactics[tacticID] == nil {
				tactics[tacticID] = vector.GetTactic()
			}

			if tacticsTechniques[tacticID] == nil {
				tacticsTechniques[tacticID] = make(map[string]struct{})
			}

			for _, technique := range vector.GetTechniques() {
				if techniques[technique.GetId()] == nil {
					techniques[technique.GetId()] = technique
				}

				if _, ok := tacticsTechniques[tacticID][technique.GetId()]; ok {
					techniques[technique.GetId()] = technique
				}
				tacticsTechniques[tacticID][technique.GetId()] = struct{}{}
			}
		}
	}

	vectors := make([]*storage.MitreAttackVector, 0, len(tactics))
	for tacticID, techniqueIDs := range tacticsTechniques {
		techniquesForTactics := make([]*storage.MitreTechnique, 0, len(techniqueIDs))
		for techniqueID := range techniqueIDs {
			techniquesForTactics = append(techniquesForTactics, techniques[techniqueID])
		}

		vectors = append(vectors, &storage.MitreAttackVector{
			Tactic:     tactics[tacticID],
			Techniques: techniquesForTactics,
		})
	}
	return vectors
}
