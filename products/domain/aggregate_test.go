package domain

import (
	"errors"
	"testing"
	"time"

	//"github.com/kr/pretty"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/mocks"
)

func TestAggregateHandleCommand(t *testing.T) {
	TimeNow = func() time.Time {
		return time.Date(2017, time.November, 20, 42, 0, 0, 0, time.Local)
	}

	id := eh.NewUUID()
	cases := map[string]struct {
		agg             *AggregateProduct
		cmd             eh.Command
		expectedfEvents []eh.Event
		expectedErr     error
	}{
		"unknown command": {
			&Aggregate{
				AggregateBase: events.NewAggregateBase(AggregateType, id),
				created:       true,
			},
			&mocks.Command{
				ID:      id,
				Content: "testcontent",
			},
			nil,
			errors.New("could not handle command: Command"),
		},
	}
}
