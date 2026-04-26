// Package registry provides a procedure registry that scans SQL procedure directories
// and builds a mapping of procedure names to their file paths for fast lookups.
//
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.6-35b-a3b
// Percentage written by Agent: 60%

package registry

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ProcedureRegistry maps procedure names (without sp_ prefix) to their SQL file paths
type ProcedureRegistry map[string]string

// Registry manages the procedure registry with lazy initialization
type Registry struct {
	baseDir     string
	procedures  []string // list of procedure directories to search
	registry    ProcedureRegistry
	initialized bool
	once        sync.Once
	err         error
}

// NewRegistry creates a new Registry instance
// baseDir should point to the SQL Procedures directory (e.g., .../sql/Procedures)
// procedureDirs are subdirectories within baseDir to scan (e.g., ["company"])
func NewRegistry() *Registry {
	cwd := "/workspaces/cc520-company-dbms-project"
	sqlPath := filepath.Join(cwd, "src", "company", "sql", "Procedures")
	procedureDirs := []string{"company"}

	return &Registry{
		baseDir:     sqlPath,
		procedures:  procedureDirs,
		registry:    make(ProcedureRegistry),
		initialized: false,
	}
}

// init initializes the registry if not already done
func (r *Registry) init() error {
	r.once.Do(func() {
		r.err = r.buildRegistry()
		r.initialized = true
	})
	return r.err
}

// buildRegistry scans procedure directories and builds the registry
func (r *Registry) buildRegistry() error {
	for _, procDir := range r.procedures {
		procPath := filepath.Join(r.baseDir, procDir)

		entries, err := os.ReadDir(procPath)
		if err != nil {
			return fmt.Errorf("failed to read procedure directory %s: %w", procPath, err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			// Only process .sql files
			if !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}

			// Skip the wildcard SELECT file
			if entry.Name() == "Org.sql" {
				continue
			}

			// Extract procedure name from filename
			filename := entry.Name()

			// Remove .sql extension
			name := strings.TrimSuffix(filename, ".sql")

			// Add sp_ prefix for registry key
			procName := fmt.Sprintf("sp_%s", name)

			// Store the full path to the SQL file
			sqlPath := filepath.Join(procPath, filename)
			r.registry[procName] = sqlPath
		}
	}
	return nil
}

// Get returns the SQL file path for a given procedure name
func (r *Registry) Get(procedureName string) (string, error) {
	if err := r.init(); err != nil {
		return "", err
	}

	// Correct procedure name if sp_ prefix is not in it
	if !strings.Contains(procedureName, "sp_") {
		procedureName = fmt.Sprintf("sp_%s", procedureName)
	}

	_, ok := r.registry[procedureName]
	if !ok {
		return "", fmt.Errorf("procedure not found: %s", procedureName)
	}

	return procedureName, nil
}

// GetAll returns a copy of the entire registry mapping
func (r *Registry) GetAll() ProcedureRegistry {
	if err := r.init(); err != nil {
		return nil
	}

	result := make(ProcedureRegistry, len(r.registry))
	for k, v := range r.registry {
		result[k] = v
	}
	return result
}

// Initialized is a getter to return if the registry was initialized or not
func (r *Registry) Initialized() bool {
	return r.initialized
}

// ClearRegistry clears the registry mapping and resets the initialized state
// for testing
func (r *Registry) ClearRegistry() {
	for k := range r.registry {
		delete(r.registry, k)
	}
	r.initialized = false
}

// Count returns the current number of registered stored procedures for testing
func (r *Registry) Count() int {
	return len(r.registry)
}
