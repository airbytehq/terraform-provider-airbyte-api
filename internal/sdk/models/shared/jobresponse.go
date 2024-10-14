// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// JobResponse - Provides details of a single job.
type JobResponse struct {
	BytesSynced  *int64 `json:"bytesSynced,omitempty"`
	ConnectionID string `json:"connectionId"`
	// Duration of a sync in ISO_8601 format
	Duration *string `json:"duration,omitempty"`
	JobID    int64   `json:"jobId"`
	// Enum that describes the different types of jobs that the platform runs.
	JobType       JobTypeEnum   `json:"jobType"`
	LastUpdatedAt *string       `json:"lastUpdatedAt,omitempty"`
	RowsSynced    *int64        `json:"rowsSynced,omitempty"`
	StartTime     string        `json:"startTime"`
	Status        JobStatusEnum `json:"status"`
}

func (o *JobResponse) GetBytesSynced() *int64 {
	if o == nil {
		return nil
	}
	return o.BytesSynced
}

func (o *JobResponse) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *JobResponse) GetDuration() *string {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *JobResponse) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *JobResponse) GetJobType() JobTypeEnum {
	if o == nil {
		return JobTypeEnum("")
	}
	return o.JobType
}

func (o *JobResponse) GetLastUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedAt
}

func (o *JobResponse) GetRowsSynced() *int64 {
	if o == nil {
		return nil
	}
	return o.RowsSynced
}

func (o *JobResponse) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

func (o *JobResponse) GetStatus() JobStatusEnum {
	if o == nil {
		return JobStatusEnum("")
	}
	return o.Status
}
