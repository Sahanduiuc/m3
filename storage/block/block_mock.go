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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/storage/block/types.go

package block

import (
	gomock "github.com/golang/mock/gomock"
	context "github.com/m3db/m3db/context"
	encoding "github.com/m3db/m3db/encoding"
	ts "github.com/m3db/m3db/ts"
	io "github.com/m3db/m3db/x/io"
	pool "github.com/m3db/m3x/pool"
	time "time"
)

// Mock of FilteredBlocksMetadataIter interface
type MockFilteredBlocksMetadataIter struct {
	ctrl     *gomock.Controller
	recorder *_MockFilteredBlocksMetadataIterRecorder
}

// Recorder for MockFilteredBlocksMetadataIter (not exported)
type _MockFilteredBlocksMetadataIterRecorder struct {
	mock *MockFilteredBlocksMetadataIter
}

func NewMockFilteredBlocksMetadataIter(ctrl *gomock.Controller) *MockFilteredBlocksMetadataIter {
	mock := &MockFilteredBlocksMetadataIter{ctrl: ctrl}
	mock.recorder = &_MockFilteredBlocksMetadataIterRecorder{mock}
	return mock
}

func (_m *MockFilteredBlocksMetadataIter) EXPECT() *_MockFilteredBlocksMetadataIterRecorder {
	return _m.recorder
}

func (_m *MockFilteredBlocksMetadataIter) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockFilteredBlocksMetadataIterRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockFilteredBlocksMetadataIter) Current() (ts.ID, Metadata) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.ID)
	ret1, _ := ret[1].(Metadata)
	return ret0, ret1
}

func (_mr *_MockFilteredBlocksMetadataIterRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

// Mock of FetchBlockResult interface
type MockFetchBlockResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockResultRecorder
}

// Recorder for MockFetchBlockResult (not exported)
type _MockFetchBlockResultRecorder struct {
	mock *MockFetchBlockResult
}

func NewMockFetchBlockResult(ctrl *gomock.Controller) *MockFetchBlockResult {
	mock := &MockFetchBlockResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlockResult) EXPECT() *_MockFetchBlockResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockResult) Start() time.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockFetchBlockResult) Readers() []io.SegmentReader {
	ret := _m.ctrl.Call(_m, "Readers")
	ret0, _ := ret[0].([]io.SegmentReader)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Readers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Readers")
}

func (_m *MockFetchBlockResult) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

// Mock of FetchBlockMetadataResults interface
type MockFetchBlockMetadataResults struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockMetadataResultsRecorder
}

// Recorder for MockFetchBlockMetadataResults (not exported)
type _MockFetchBlockMetadataResultsRecorder struct {
	mock *MockFetchBlockMetadataResults
}

func NewMockFetchBlockMetadataResults(ctrl *gomock.Controller) *MockFetchBlockMetadataResults {
	mock := &MockFetchBlockMetadataResults{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockMetadataResultsRecorder{mock}
	return mock
}

func (_m *MockFetchBlockMetadataResults) EXPECT() *_MockFetchBlockMetadataResultsRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockMetadataResults) Add(res FetchBlockMetadataResult) {
	_m.ctrl.Call(_m, "Add", res)
}

func (_mr *_MockFetchBlockMetadataResultsRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0)
}

func (_m *MockFetchBlockMetadataResults) Results() []FetchBlockMetadataResult {
	ret := _m.ctrl.Call(_m, "Results")
	ret0, _ := ret[0].([]FetchBlockMetadataResult)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultsRecorder) Results() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Results")
}

func (_m *MockFetchBlockMetadataResults) Sort() {
	_m.ctrl.Call(_m, "Sort")
}

func (_mr *_MockFetchBlockMetadataResultsRecorder) Sort() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sort")
}

func (_m *MockFetchBlockMetadataResults) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockFetchBlockMetadataResultsRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockFetchBlockMetadataResults) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockFetchBlockMetadataResultsRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of FetchBlocksMetadataResults interface
type MockFetchBlocksMetadataResults struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlocksMetadataResultsRecorder
}

// Recorder for MockFetchBlocksMetadataResults (not exported)
type _MockFetchBlocksMetadataResultsRecorder struct {
	mock *MockFetchBlocksMetadataResults
}

func NewMockFetchBlocksMetadataResults(ctrl *gomock.Controller) *MockFetchBlocksMetadataResults {
	mock := &MockFetchBlocksMetadataResults{ctrl: ctrl}
	mock.recorder = &_MockFetchBlocksMetadataResultsRecorder{mock}
	return mock
}

func (_m *MockFetchBlocksMetadataResults) EXPECT() *_MockFetchBlocksMetadataResultsRecorder {
	return _m.recorder
}

func (_m *MockFetchBlocksMetadataResults) Add(res FetchBlocksMetadataResult) {
	_m.ctrl.Call(_m, "Add", res)
}

func (_mr *_MockFetchBlocksMetadataResultsRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0)
}

func (_m *MockFetchBlocksMetadataResults) Results() []FetchBlocksMetadataResult {
	ret := _m.ctrl.Call(_m, "Results")
	ret0, _ := ret[0].([]FetchBlocksMetadataResult)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultsRecorder) Results() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Results")
}

func (_m *MockFetchBlocksMetadataResults) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockFetchBlocksMetadataResultsRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockFetchBlocksMetadataResults) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockFetchBlocksMetadataResultsRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of DatabaseBlock interface
type MockDatabaseBlock struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockRecorder
}

// Recorder for MockDatabaseBlock (not exported)
type _MockDatabaseBlockRecorder struct {
	mock *MockDatabaseBlock
}

func NewMockDatabaseBlock(ctrl *gomock.Controller) *MockDatabaseBlock {
	mock := &MockDatabaseBlock{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlock) EXPECT() *_MockDatabaseBlockRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlock) StartTime() time.Time {
	ret := _m.ctrl.Call(_m, "StartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) StartTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTime")
}

func (_m *MockDatabaseBlock) Checksum() *uint32 {
	ret := _m.ctrl.Call(_m, "Checksum")
	ret0, _ := ret[0].(*uint32)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) Checksum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Checksum")
}

func (_m *MockDatabaseBlock) Stream(blocker context.Context) (io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Stream", blocker)
	ret0, _ := ret[0].(io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseBlockRecorder) Stream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stream", arg0)
}

func (_m *MockDatabaseBlock) Reset(startTime time.Time, segment ts.Segment) {
	_m.ctrl.Call(_m, "Reset", startTime, segment)
}

func (_mr *_MockDatabaseBlockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1)
}

func (_m *MockDatabaseBlock) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseBlockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of DatabaseSeriesBlocks interface
type MockDatabaseSeriesBlocks struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesBlocksRecorder
}

// Recorder for MockDatabaseSeriesBlocks (not exported)
type _MockDatabaseSeriesBlocksRecorder struct {
	mock *MockDatabaseSeriesBlocks
}

func NewMockDatabaseSeriesBlocks(ctrl *gomock.Controller) *MockDatabaseSeriesBlocks {
	mock := &MockDatabaseSeriesBlocks{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesBlocksRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeriesBlocks) EXPECT() *_MockDatabaseSeriesBlocksRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeriesBlocks) Options() Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Options")
}

func (_m *MockDatabaseSeriesBlocks) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockDatabaseSeriesBlocks) AddBlock(block DatabaseBlock) {
	_m.ctrl.Call(_m, "AddBlock", block)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddBlock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBlock", arg0)
}

func (_m *MockDatabaseSeriesBlocks) AddSeries(other DatabaseSeriesBlocks) {
	_m.ctrl.Call(_m, "AddSeries", other)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddSeries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSeries", arg0)
}

func (_m *MockDatabaseSeriesBlocks) MinTime() time.Time {
	ret := _m.ctrl.Call(_m, "MinTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) MinTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MinTime")
}

func (_m *MockDatabaseSeriesBlocks) MaxTime() time.Time {
	ret := _m.ctrl.Call(_m, "MaxTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) MaxTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxTime")
}

func (_m *MockDatabaseSeriesBlocks) BlockAt(t time.Time) (DatabaseBlock, bool) {
	ret := _m.ctrl.Call(_m, "BlockAt", t)
	ret0, _ := ret[0].(DatabaseBlock)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) BlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) AllBlocks() map[time.Time]DatabaseBlock {
	ret := _m.ctrl.Call(_m, "AllBlocks")
	ret0, _ := ret[0].(map[time.Time]DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AllBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllBlocks")
}

func (_m *MockDatabaseSeriesBlocks) RemoveBlockAt(t time.Time) {
	_m.ctrl.Call(_m, "RemoveBlockAt", t)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) RemoveBlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveBlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) RemoveAll() {
	_m.ctrl.Call(_m, "RemoveAll")
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) RemoveAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveAll")
}

func (_m *MockDatabaseSeriesBlocks) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of DatabaseBlockPool interface
type MockDatabaseBlockPool struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockPoolRecorder
}

// Recorder for MockDatabaseBlockPool (not exported)
type _MockDatabaseBlockPoolRecorder struct {
	mock *MockDatabaseBlockPool
}

func NewMockDatabaseBlockPool(ctrl *gomock.Controller) *MockDatabaseBlockPool {
	mock := &MockDatabaseBlockPool{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockPoolRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlockPool) EXPECT() *_MockDatabaseBlockPoolRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlockPool) Init(alloc DatabaseBlockAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockDatabaseBlockPool) Get() DatabaseBlock {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseBlockPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockDatabaseBlockPool) Put(block DatabaseBlock) {
	_m.ctrl.Call(_m, "Put", block)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of FetchBlockMetadataResultsPool interface
type MockFetchBlockMetadataResultsPool struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockMetadataResultsPoolRecorder
}

// Recorder for MockFetchBlockMetadataResultsPool (not exported)
type _MockFetchBlockMetadataResultsPoolRecorder struct {
	mock *MockFetchBlockMetadataResultsPool
}

func NewMockFetchBlockMetadataResultsPool(ctrl *gomock.Controller) *MockFetchBlockMetadataResultsPool {
	mock := &MockFetchBlockMetadataResultsPool{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockMetadataResultsPoolRecorder{mock}
	return mock
}

func (_m *MockFetchBlockMetadataResultsPool) EXPECT() *_MockFetchBlockMetadataResultsPoolRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockMetadataResultsPool) Get() FetchBlockMetadataResults {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(FetchBlockMetadataResults)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultsPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockFetchBlockMetadataResultsPool) Put(res FetchBlockMetadataResults) {
	_m.ctrl.Call(_m, "Put", res)
}

func (_mr *_MockFetchBlockMetadataResultsPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of FetchBlocksMetadataResultsPool interface
type MockFetchBlocksMetadataResultsPool struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlocksMetadataResultsPoolRecorder
}

// Recorder for MockFetchBlocksMetadataResultsPool (not exported)
type _MockFetchBlocksMetadataResultsPoolRecorder struct {
	mock *MockFetchBlocksMetadataResultsPool
}

func NewMockFetchBlocksMetadataResultsPool(ctrl *gomock.Controller) *MockFetchBlocksMetadataResultsPool {
	mock := &MockFetchBlocksMetadataResultsPool{ctrl: ctrl}
	mock.recorder = &_MockFetchBlocksMetadataResultsPoolRecorder{mock}
	return mock
}

func (_m *MockFetchBlocksMetadataResultsPool) EXPECT() *_MockFetchBlocksMetadataResultsPoolRecorder {
	return _m.recorder
}

func (_m *MockFetchBlocksMetadataResultsPool) Get() FetchBlocksMetadataResults {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(FetchBlocksMetadataResults)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultsPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockFetchBlocksMetadataResultsPool) Put(res FetchBlocksMetadataResults) {
	_m.ctrl.Call(_m, "Put", res)
}

func (_mr *_MockFetchBlocksMetadataResultsPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsRecorder
}

// Recorder for MockOptions (not exported)
type _MockOptionsRecorder struct {
	mock *MockOptions
}

func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &_MockOptionsRecorder{mock}
	return mock
}

func (_m *MockOptions) EXPECT() *_MockOptionsRecorder {
	return _m.recorder
}

func (_m *MockOptions) SetDatabaseBlockAllocSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockAllocSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockAllocSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockAllocSize", arg0)
}

func (_m *MockOptions) DatabaseBlockAllocSize() int {
	ret := _m.ctrl.Call(_m, "DatabaseBlockAllocSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockAllocSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockAllocSize")
}

func (_m *MockOptions) SetDatabaseBlockPool(value DatabaseBlockPool) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockPool", arg0)
}

func (_m *MockOptions) DatabaseBlockPool() DatabaseBlockPool {
	ret := _m.ctrl.Call(_m, "DatabaseBlockPool")
	ret0, _ := ret[0].(DatabaseBlockPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockPool")
}

func (_m *MockOptions) SetContextPool(value context.Pool) Options {
	ret := _m.ctrl.Call(_m, "SetContextPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetContextPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetContextPool", arg0)
}

func (_m *MockOptions) ContextPool() context.Pool {
	ret := _m.ctrl.Call(_m, "ContextPool")
	ret0, _ := ret[0].(context.Pool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ContextPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContextPool")
}

func (_m *MockOptions) SetEncoderPool(value encoding.EncoderPool) Options {
	ret := _m.ctrl.Call(_m, "SetEncoderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetEncoderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEncoderPool", arg0)
}

func (_m *MockOptions) EncoderPool() encoding.EncoderPool {
	ret := _m.ctrl.Call(_m, "EncoderPool")
	ret0, _ := ret[0].(encoding.EncoderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncoderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncoderPool")
}

func (_m *MockOptions) SetReaderIteratorPool(value encoding.ReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReaderIteratorPool", arg0)
}

func (_m *MockOptions) ReaderIteratorPool() encoding.ReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "ReaderIteratorPool")
	ret0, _ := ret[0].(encoding.ReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorPool")
}

func (_m *MockOptions) SetMultiReaderIteratorPool(value encoding.MultiReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetMultiReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetMultiReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMultiReaderIteratorPool", arg0)
}

func (_m *MockOptions) MultiReaderIteratorPool() encoding.MultiReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "MultiReaderIteratorPool")
	ret0, _ := ret[0].(encoding.MultiReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) MultiReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiReaderIteratorPool")
}

func (_m *MockOptions) SetSegmentReaderPool(value io.SegmentReaderPool) Options {
	ret := _m.ctrl.Call(_m, "SetSegmentReaderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetSegmentReaderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSegmentReaderPool", arg0)
}

func (_m *MockOptions) SegmentReaderPool() io.SegmentReaderPool {
	ret := _m.ctrl.Call(_m, "SegmentReaderPool")
	ret0, _ := ret[0].(io.SegmentReaderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) SegmentReaderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SegmentReaderPool")
}

func (_m *MockOptions) SetBytesPool(value pool.CheckedBytesPool) Options {
	ret := _m.ctrl.Call(_m, "SetBytesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetBytesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBytesPool", arg0)
}

func (_m *MockOptions) BytesPool() pool.CheckedBytesPool {
	ret := _m.ctrl.Call(_m, "BytesPool")
	ret0, _ := ret[0].(pool.CheckedBytesPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) BytesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BytesPool")
}
