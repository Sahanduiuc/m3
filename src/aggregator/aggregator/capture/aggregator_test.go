// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package capture

import (
	"fmt"
	"testing"
	"time"

	"github.com/m3db/m3metrics/aggregation"
	"github.com/m3db/m3metrics/metadata"
	"github.com/m3db/m3metrics/metric"
	"github.com/m3db/m3metrics/metric/aggregated"
	"github.com/m3db/m3metrics/metric/id"
	"github.com/m3db/m3metrics/metric/unaggregated"
	"github.com/m3db/m3metrics/op"
	"github.com/m3db/m3metrics/op/applied"
	"github.com/m3db/m3metrics/policy"
	xtime "github.com/m3db/m3x/time"

	"github.com/stretchr/testify/require"
)

var (
	testCounter = unaggregated.MetricUnion{
		Type:       metric.CounterType,
		ID:         id.RawID("testCounter"),
		CounterVal: 1234,
	}
	testBatchTimer = unaggregated.MetricUnion{
		Type:          metric.TimerType,
		ID:            id.RawID("testCounter"),
		BatchTimerVal: []float64{1.0, 3.5, 2.2, 6.5, 4.8},
	}
	testGauge = unaggregated.MetricUnion{
		Type:     metric.GaugeType,
		ID:       id.RawID("testCounter"),
		GaugeVal: 123.456,
	}
	testForwarded = aggregated.Metric{
		Type:      metric.CounterType,
		ID:        []byte("testForwarded"),
		TimeNanos: 12345,
		Value:     908,
	}
	testInvalid = unaggregated.MetricUnion{
		Type: metric.UnknownType,
		ID:   id.RawID("invalid"),
	}
	testDefaultMetadatas = metadata.DefaultStagedMetadatas
	testForwardMetadata  = metadata.ForwardMetadata{
		AggregationID: aggregation.DefaultID,
		StoragePolicy: policy.NewStoragePolicy(time.Minute, xtime.Minute, 12*time.Hour),
		Pipeline: applied.NewPipeline([]applied.Union{
			{
				Type: op.RollupType,
				Rollup: applied.Rollup{
					ID:            []byte("foo"),
					AggregationID: aggregation.MustCompressTypes(aggregation.Count),
				},
			},
		}),
		SourceID:          []byte("testForwardSource"),
		NumForwardedTimes: 3,
	}
)

func TestAggregator(t *testing.T) {
	agg := NewAggregator()

	// Adding an invalid metric should result in an error.
	metadatas := testDefaultMetadatas
	require.Error(t, agg.AddUntimed(testInvalid, metadatas))

	// Add valid untimed metrics with policies.
	var expected SnapshotResult
	for _, mu := range []unaggregated.MetricUnion{testCounter, testBatchTimer, testGauge} {
		switch mu.Type {
		case metric.CounterType:
			expected.CountersWithMetadatas = append(
				expected.CountersWithMetadatas,
				unaggregated.CounterWithMetadatas{
					Counter:         mu.Counter(),
					StagedMetadatas: metadatas,
				})
		case metric.TimerType:
			expected.BatchTimersWithMetadatas = append(
				expected.BatchTimersWithMetadatas,
				unaggregated.BatchTimerWithMetadatas{
					BatchTimer:      mu.BatchTimer(),
					StagedMetadatas: metadatas,
				})
		case metric.GaugeType:
			expected.GaugesWithMetadatas = append(
				expected.GaugesWithMetadatas,
				unaggregated.GaugeWithMetadatas{
					Gauge:           mu.Gauge(),
					StagedMetadatas: metadatas,
				})
		default:
			require.Fail(t, fmt.Sprintf("unknown metric type %v", mu.Type))
		}
		require.NoError(t, agg.AddUntimed(mu, metadatas))
	}

	// Add valid forwarded metrics with metadata.
	expected.MetricsWithForwardMetadata = append(
		expected.MetricsWithForwardMetadata,
		aggregated.MetricWithForwardMetadata{
			Metric:          testForwarded,
			ForwardMetadata: testForwardMetadata,
		},
	)
	require.NoError(t, agg.AddForwarded(testForwarded, testForwardMetadata))

	require.Equal(t, 4, agg.NumMetricsAdded())

	res := agg.Snapshot()
	require.Equal(t, expected, res)
}
