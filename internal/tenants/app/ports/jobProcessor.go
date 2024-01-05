package ports

import (
	"context"

	"github.com/riverqueue/river"
)

type JobProcessor interface {
	ScheduleNewJob(ctx context.Context, args river.JobArgs) error
}
