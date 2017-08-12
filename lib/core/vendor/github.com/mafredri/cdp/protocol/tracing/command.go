// Code generated by cdpgen. DO NOT EDIT.

package tracing

// StartArgs represents the arguments for Start in the Tracing domain.
type StartArgs struct {
	// Categories is deprecated.
	//
	// Deprecated: Category/tag filter
	Categories *string `json:"categories,omitempty"`
	// Options is deprecated.
	//
	// Deprecated: Tracing options
	Options                      *string  `json:"options,omitempty"`
	BufferUsageReportingInterval *float64 `json:"bufferUsageReportingInterval,omitempty"` // If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	// TransferMode Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to ReportEvents).
	//
	// Values: "ReportEvents", "ReturnAsStream".
	TransferMode *string      `json:"transferMode,omitempty"`
	TraceConfig  *TraceConfig `json:"traceConfig,omitempty"` // No description.
}

// NewStartArgs initializes StartArgs with the required arguments.
func NewStartArgs() *StartArgs {
	args := new(StartArgs)

	return args
}

// SetCategories sets the Categories optional argument.
//
// Deprecated: Category/tag filter
func (a *StartArgs) SetCategories(categories string) *StartArgs {
	a.Categories = &categories
	return a
}

// SetOptions sets the Options optional argument.
//
// Deprecated: Tracing options
func (a *StartArgs) SetOptions(options string) *StartArgs {
	a.Options = &options
	return a
}

// SetBufferUsageReportingInterval sets the BufferUsageReportingInterval optional argument. If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
func (a *StartArgs) SetBufferUsageReportingInterval(bufferUsageReportingInterval float64) *StartArgs {
	a.BufferUsageReportingInterval = &bufferUsageReportingInterval
	return a
}

// SetTransferMode sets the TransferMode optional argument. Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to ReportEvents).
//
// Values: "ReportEvents", "ReturnAsStream".
func (a *StartArgs) SetTransferMode(transferMode string) *StartArgs {
	a.TransferMode = &transferMode
	return a
}

// SetTraceConfig sets the TraceConfig optional argument.
func (a *StartArgs) SetTraceConfig(traceConfig TraceConfig) *StartArgs {
	a.TraceConfig = &traceConfig
	return a
}

// GetCategoriesReply represents the return values for GetCategories in the Tracing domain.
type GetCategoriesReply struct {
	Categories []string `json:"categories"` // A list of supported tracing categories.
}

// RequestMemoryDumpReply represents the return values for RequestMemoryDump in the Tracing domain.
type RequestMemoryDumpReply struct {
	DumpGUID string `json:"dumpGuid"` // GUID of the resulting global memory dump.
	Success  bool   `json:"success"`  // True iff the global memory dump succeeded.
}

// RecordClockSyncMarkerArgs represents the arguments for RecordClockSyncMarker in the Tracing domain.
type RecordClockSyncMarkerArgs struct {
	SyncID string `json:"syncId"` // The ID of this clock sync marker
}

// NewRecordClockSyncMarkerArgs initializes RecordClockSyncMarkerArgs with the required arguments.
func NewRecordClockSyncMarkerArgs(syncID string) *RecordClockSyncMarkerArgs {
	args := new(RecordClockSyncMarkerArgs)
	args.SyncID = syncID
	return args
}
