package env

import (
	log "github.com/newrelic/newrelic-diagnostics-cli/logger"
	"github.com/newrelic/newrelic-diagnostics-cli/tasks"
)

// BaseEnvCollectEnvVars - This struct defined the sample plugin which can be used as a starting point
type BaseEnvCollectEnvVars struct {
}

// Identifier - This returns the Category, Subcategory and Name of each task
func (t BaseEnvCollectEnvVars) Identifier() tasks.Identifier {
	return tasks.IdentifierFromString("Base/Env/CollectEnvVars")
}

// Explain - Returns the help text for each individual task
func (t BaseEnvCollectEnvVars) Explain() string {
	return "Collect environment variables"
}

// Dependencies - Returns the dependencies for ech task.
func (t BaseEnvCollectEnvVars) Dependencies() []string {
	return []string{}
}

// Execute - The core work within each task
func (t BaseEnvCollectEnvVars) Execute(options tasks.Options, upstream map[string]tasks.Result) tasks.Result {
	var result tasks.Result
	envVars, err := tasks.GetShellEnvVars()

	if err != nil {
		log.Debug(err.Error())
		result.Status = tasks.Warning
		result.Summary = "Unable to gather any Environment Variables from the current shell."
		// result.URL should be set to something, but what?
		return result
	}

	filteredEnvVars := envVars.WithDefaultFilter()

	result.Payload = filteredEnvVars
	result.Status = tasks.Info
	result.Summary = "Gathered Environment variables of current shell."

	return result
}
