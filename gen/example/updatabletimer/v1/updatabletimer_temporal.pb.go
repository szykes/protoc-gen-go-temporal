// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.2-next (0b994cbcb4a60e84ee2566d444a20428b160f204)
//	go go1.22.1
//	protoc (unknown)
//
// source: example/updatabletimer/v1/updatabletimer.proto
package updatabletimerv1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpers "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	scheme "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	gohomedir "github.com/mitchellh/go-homedir"
	v2 "github.com/urfave/cli/v2"
	client "go.temporal.io/sdk/client"
	testsuite "go.temporal.io/sdk/testsuite"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	protojson "google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"os"
	"sort"
	"time"
)

// ExampleTaskQueue is the default task-queue for a example.updatabletimer.v1.Example worker
const ExampleTaskQueue = "updatable-timer"

// example.updatabletimer.v1.Example workflow names
const (
	UpdatableTimerWorkflowName = "UpdatableTimer"
)

// example.updatabletimer.v1.Example workflow id expressions
var (
	UpdatableTimerIdexpression = expression.MustParseExpression("updatable-timer/${! name.or(uuid_v4()) }")
)

// example.updatabletimer.v1.Example query names
const (
	GetWakeUpTimeQueryName = "example.updatabletimer.v1.Example.GetWakeUpTime"
)

// example.updatabletimer.v1.Example signal names
const (
	UpdateWakeUpTimeSignalName = "example.updatabletimer.v1.Example.UpdateWakeUpTime"
)

// ExampleClient describes a client for a(n) example.updatabletimer.v1.Example worker
type ExampleClient interface {
	// UpdatableTimer describes an updatable timer workflow
	UpdatableTimer(ctx context.Context, req *UpdatableTimerInput, opts ...*UpdatableTimerOptions) error

	// UpdatableTimerAsync starts a(n) UpdatableTimer workflow and returns a handle to the workflow run
	UpdatableTimerAsync(ctx context.Context, req *UpdatableTimerInput, opts ...*UpdatableTimerOptions) (UpdatableTimerRun, error)

	// GetUpdatableTimer retrieves a handle to an existing UpdatableTimer workflow execution
	GetUpdatableTimer(ctx context.Context, workflowID string, runID string) UpdatableTimerRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error

	// GetWakeUpTime retrieves the current timer expiration timestamp
	GetWakeUpTime(ctx context.Context, workflowID string, runID string) (*GetWakeUpTimeOutput, error)

	// UpdateWakeUpTime updates the timer expiration timestamp
	UpdateWakeUpTime(ctx context.Context, workflowID string, runID string, signal *UpdateWakeUpTimeInput) error
}

// exampleClient implements a temporal client for a example.updatabletimer.v1.Example service
type exampleClient struct {
	client client.Client
	log    *slog.Logger
}

// NewExampleClient initializes a new example.updatabletimer.v1.Example client
func NewExampleClient(c client.Client, options ...*exampleClientOptions) ExampleClient {
	var cfg *exampleClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewExampleClientOptions()
	}
	return &exampleClient{
		client: c,
		log:    cfg.getLogger(),
	}
}

// NewExampleClientWithOptions initializes a new Example client with the given options
func NewExampleClientWithOptions(c client.Client, opts client.Options, options ...*exampleClientOptions) (ExampleClient, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	var cfg *exampleClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewExampleClientOptions()
	}
	return &exampleClient{
		client: c,
		log:    cfg.getLogger(),
	}, nil
}

// exampleClientOptions describes optional runtime configuration for a ExampleClient
type exampleClientOptions struct {
	log *slog.Logger
}

// NewExampleClientOptions initializes a new exampleClientOptions value
func NewExampleClientOptions() *exampleClientOptions {
	return &exampleClientOptions{}
}

// WithLogger can be used to override the default logger
func (opts *exampleClientOptions) WithLogger(l *slog.Logger) *exampleClientOptions {
	if l != nil {
		opts.log = l
	}
	return opts
}

// getLogger returns the configured logger, or the default logger
func (opts *exampleClientOptions) getLogger() *slog.Logger {
	if opts != nil && opts.log != nil {
		return opts.log
	}
	return slog.Default()
}

// UpdatableTimer describes an updatable timer workflow
func (c *exampleClient) UpdatableTimer(ctx context.Context, req *UpdatableTimerInput, options ...*UpdatableTimerOptions) error {
	run, err := c.UpdatableTimerAsync(ctx, req, options...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// UpdatableTimer describes an updatable timer workflow
func (c *exampleClient) UpdatableTimerAsync(ctx context.Context, req *UpdatableTimerInput, options ...*UpdatableTimerOptions) (UpdatableTimerRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(UpdatableTimerIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"UpdatableTimer\" workflow: %w", err)
		}
		opts.ID = id
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, UpdatableTimerWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &updatableTimerRun{
		client: c,
		run:    run,
	}, nil
}

// GetUpdatableTimer fetches an existing UpdatableTimer execution
func (c *exampleClient) GetUpdatableTimer(ctx context.Context, workflowID string, runID string) UpdatableTimerRun {
	return &updatableTimerRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *exampleClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return c.client.CancelWorkflow(ctx, workflowID, runID)
}

// TerminateWorkflow terminates an existing workflow execution
func (c *exampleClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.client.TerminateWorkflow(ctx, workflowID, runID, reason, details...)
}

// GetWakeUpTime retrieves the current timer expiration timestamp
func (c *exampleClient) GetWakeUpTime(ctx context.Context, workflowID string, runID string) (*GetWakeUpTimeOutput, error) {
	var resp GetWakeUpTimeOutput
	if val, err := c.client.QueryWorkflow(ctx, workflowID, runID, GetWakeUpTimeQueryName); err != nil {
		return nil, err
	} else if err = val.Get(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateWakeUpTime updates the timer expiration timestamp
func (c *exampleClient) UpdateWakeUpTime(ctx context.Context, workflowID string, runID string, signal *UpdateWakeUpTimeInput) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, UpdateWakeUpTimeSignalName, signal)
}

// UpdatableTimerOptions provides configuration for a UpdatableTimer workflow operation
type UpdatableTimerOptions struct {
	opts *client.StartWorkflowOptions
}

// NewUpdatableTimerOptions initializes a new UpdatableTimerOptions value
func NewUpdatableTimerOptions() *UpdatableTimerOptions {
	return &UpdatableTimerOptions{}
}

// WithStartWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *UpdatableTimerOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *UpdatableTimerOptions {
	opts.opts = &options
	return opts
}

// UpdatableTimerRun describes a(n) UpdatableTimer workflow run
type UpdatableTimerRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error

	// GetWakeUpTime retrieves the current timer expiration timestamp
	GetWakeUpTime(ctx context.Context) (*GetWakeUpTimeOutput, error)

	// UpdateWakeUpTime updates the timer expiration timestamp
	UpdateWakeUpTime(ctx context.Context, req *UpdateWakeUpTimeInput) error
}

// updatableTimerRun provides an internal implementation of a(n) UpdatableTimerRunRun
type updatableTimerRun struct {
	client *exampleClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *updatableTimerRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *updatableTimerRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *updatableTimerRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *updatableTimerRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *updatableTimerRun) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *updatableTimerRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// GetWakeUpTime retrieves the current timer expiration timestamp
func (r *updatableTimerRun) GetWakeUpTime(ctx context.Context) (*GetWakeUpTimeOutput, error) {
	return r.client.GetWakeUpTime(ctx, r.ID(), "")
}

// UpdateWakeUpTime updates the timer expiration timestamp
func (r *updatableTimerRun) UpdateWakeUpTime(ctx context.Context, req *UpdateWakeUpTimeInput) error {
	return r.client.UpdateWakeUpTime(ctx, r.ID(), "", req)
}

// Reference to generated workflow functions
var (
	// UpdatableTimer describes an updatable timer workflow
	UpdatableTimerFunction func(workflow.Context, *UpdatableTimerInput) error
)

// ExampleWorkflows provides methods for initializing new example.updatabletimer.v1.Example workflow values
type ExampleWorkflows interface {
	// UpdatableTimer describes an updatable timer workflow
	UpdatableTimer(ctx workflow.Context, input *UpdatableTimerWorkflowInput) (UpdatableTimerWorkflow, error)
}

// RegisterExampleWorkflows registers example.updatabletimer.v1.Example workflows with the given worker
func RegisterExampleWorkflows(r worker.WorkflowRegistry, workflows ExampleWorkflows) {
	RegisterUpdatableTimerWorkflow(r, workflows.UpdatableTimer)
}

// RegisterUpdatableTimerWorkflow registers a example.updatabletimer.v1.Example.UpdatableTimer workflow with the given worker
func RegisterUpdatableTimerWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *UpdatableTimerWorkflowInput) (UpdatableTimerWorkflow, error)) {
	UpdatableTimerFunction = buildUpdatableTimer(wf)
	r.RegisterWorkflowWithOptions(UpdatableTimerFunction, workflow.RegisterOptions{Name: UpdatableTimerWorkflowName})
}

// buildUpdatableTimer converts a UpdatableTimer workflow struct into a valid workflow function
func buildUpdatableTimer(ctor func(workflow.Context, *UpdatableTimerWorkflowInput) (UpdatableTimerWorkflow, error)) func(workflow.Context, *UpdatableTimerInput) error {
	return func(ctx workflow.Context, req *UpdatableTimerInput) error {
		input := &UpdatableTimerWorkflowInput{
			Req: req,
			UpdateWakeUpTime: &UpdateWakeUpTimeSignal{
				Channel: workflow.GetSignalChannel(ctx, UpdateWakeUpTimeSignalName),
			},
		}
		wf, err := ctor(ctx, input)
		if err != nil {
			return err
		}
		if initializable, ok := wf.(helpers.Initializable); ok {
			if err := initializable.Initialize(ctx); err != nil {
				return err
			}
		}
		if err := workflow.SetQueryHandler(ctx, GetWakeUpTimeQueryName, wf.GetWakeUpTime); err != nil {
			return err
		}
		return wf.Execute(ctx)
	}
}

// UpdatableTimerWorkflowInput describes the input to a(n) UpdatableTimer workflow constructor
type UpdatableTimerWorkflowInput struct {
	Req              *UpdatableTimerInput
	UpdateWakeUpTime *UpdateWakeUpTimeSignal
}

// UpdatableTimer describes an updatable timer workflow
//
// workflow details: (id: "updatable-timer/${! name.or(uuid_v4()) }")
type UpdatableTimerWorkflow interface {
	// Execute defines the entrypoint to a(n) UpdatableTimer workflow
	Execute(ctx workflow.Context) error

	// GetWakeUpTime retrieves the current timer expiration timestamp
	GetWakeUpTime() (*GetWakeUpTimeOutput, error)
}

// UpdatableTimer describes an updatable timer workflow
func UpdatableTimerChild(ctx workflow.Context, req *UpdatableTimerInput, options ...*UpdatableTimerChildOptions) error {
	childRun, err := UpdatableTimerChildAsync(ctx, req, options...)
	if err != nil {
		return err
	}
	return childRun.Get(ctx)
}

// UpdatableTimer describes an updatable timer workflow
func UpdatableTimerChildAsync(ctx workflow.Context, req *UpdatableTimerInput, options ...*UpdatableTimerChildOptions) (*UpdatableTimerChildRun, error) {
	var opts *workflow.ChildWorkflowOptions
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	} else {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(UpdatableTimerIdexpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &UpdatableTimerChildRun{Future: workflow.ExecuteChildWorkflow(ctx, UpdatableTimerWorkflowName, req)}, nil
}

// UpdatableTimerChildOptions provides configuration for a UpdatableTimer workflow operation
type UpdatableTimerChildOptions struct {
	opts *workflow.ChildWorkflowOptions
}

// NewUpdatableTimerChildOptions initializes a new UpdatableTimerChildOptions value
func NewUpdatableTimerChildOptions() *UpdatableTimerChildOptions {
	return &UpdatableTimerChildOptions{}
}

// WithChildWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *UpdatableTimerChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *UpdatableTimerChildOptions {
	opts.opts = &options
	return opts
}

// UpdatableTimerChildRun describes a child UpdatableTimer workflow run
type UpdatableTimerChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *UpdatableTimerChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *UpdatableTimerChildRun) Select(sel workflow.Selector, fn func(*UpdatableTimerChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *UpdatableTimerChildRun) SelectStart(sel workflow.Selector, fn func(*UpdatableTimerChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *UpdatableTimerChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// UpdateWakeUpTime sends a(n) "example.updatabletimer.v1.Example.UpdateWakeUpTime" signal request to the child workflow
func (r *UpdatableTimerChildRun) UpdateWakeUpTime(ctx workflow.Context, input *UpdateWakeUpTimeInput) error {
	return r.UpdateWakeUpTimeAsync(ctx, input).Get(ctx, nil)
}

// UpdateWakeUpTimeAsync sends a(n) "example.updatabletimer.v1.Example.UpdateWakeUpTime" signal request to the child workflow
func (r *UpdatableTimerChildRun) UpdateWakeUpTimeAsync(ctx workflow.Context, input *UpdateWakeUpTimeInput) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, UpdateWakeUpTimeSignalName, input)
}

// UpdateWakeUpTimeSignal describes a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal
type UpdateWakeUpTimeSignal struct {
	Channel workflow.ReceiveChannel
}

// NewUpdateWakeUpTimeSignal initializes a new example.updatabletimer.v1.Example.UpdateWakeUpTime signal wrapper
func NewUpdateWakeUpTimeSignal(ctx workflow.Context) *UpdateWakeUpTimeSignal {
	return &UpdateWakeUpTimeSignal{Channel: workflow.GetSignalChannel(ctx, UpdateWakeUpTimeSignalName)}
}

// Receive blocks until a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal is received
func (s *UpdateWakeUpTimeSignal) Receive(ctx workflow.Context) (*UpdateWakeUpTimeInput, bool) {
	var resp UpdateWakeUpTimeInput
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a example.updatabletimer.v1.Example.UpdateWakeUpTime signal without blocking
func (s *UpdateWakeUpTimeSignal) ReceiveAsync() *UpdateWakeUpTimeInput {
	var resp UpdateWakeUpTimeInput
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// ReceiveWithTimeout blocks until a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal is received or timeout expires.
// Returns more value of false when Channel is closed.
// Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.
// resp will be nil if ok is false.
func (s *UpdateWakeUpTimeSignal) ReceiveWithTimeout(ctx workflow.Context, timeout time.Duration) (resp *UpdateWakeUpTimeInput, ok bool, more bool) {
	resp = &UpdateWakeUpTimeInput{}
	if ok, more = s.Channel.ReceiveWithTimeout(ctx, timeout, &resp); !ok {
		return nil, false, more
	}
	return
}

// Select checks for a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal without blocking
func (s *UpdateWakeUpTimeSignal) Select(sel workflow.Selector, fn func(*UpdateWakeUpTimeInput)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// UpdateWakeUpTime updates the timer expiration timestamp
func UpdateWakeUpTimeExternal(ctx workflow.Context, workflowID string, runID string, req *UpdateWakeUpTimeInput) error {
	return UpdateWakeUpTimeExternalAsync(ctx, workflowID, runID, req).Get(ctx, nil)
}

// UpdateWakeUpTime updates the timer expiration timestamp
func UpdateWakeUpTimeExternalAsync(ctx workflow.Context, workflowID string, runID string, req *UpdateWakeUpTimeInput) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, UpdateWakeUpTimeSignalName, req)
}

// ExampleActivities describes available worker activities
type ExampleActivities interface{}

// RegisterExampleActivities registers activities with a worker
func RegisterExampleActivities(r worker.ActivityRegistry, activities ExampleActivities) {}

// TestClient provides a testsuite-compatible Client
type TestExampleClient struct {
	env       *testsuite.TestWorkflowEnvironment
	workflows ExampleWorkflows
}

var _ ExampleClient = &TestExampleClient{}

// NewTestExampleClient initializes a new TestExampleClient value
func NewTestExampleClient(env *testsuite.TestWorkflowEnvironment, workflows ExampleWorkflows, activities ExampleActivities) *TestExampleClient {
	if workflows != nil {
		RegisterExampleWorkflows(env, workflows)
	}
	if activities != nil {
		RegisterExampleActivities(env, activities)
	}
	return &TestExampleClient{env, workflows}
}

// UpdatableTimer executes a(n) UpdatableTimer workflow in the test environment
func (c *TestExampleClient) UpdatableTimer(ctx context.Context, req *UpdatableTimerInput, opts ...*UpdatableTimerOptions) error {
	run, err := c.UpdatableTimerAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// UpdatableTimerAsync executes a(n) UpdatableTimer workflow in the test environment
func (c *TestExampleClient) UpdatableTimerAsync(ctx context.Context, req *UpdatableTimerInput, options ...*UpdatableTimerOptions) (UpdatableTimerRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(UpdatableTimerIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"UpdatableTimer\" workflow: %w", err)
		}
		opts.ID = id
	}
	return &testUpdatableTimerRun{client: c, env: c.env, opts: opts, req: req, workflows: c.workflows}, nil
}

// GetUpdatableTimer is a noop
func (c *TestExampleClient) GetUpdatableTimer(ctx context.Context, workflowID string, runID string) UpdatableTimerRun {
	return &testUpdatableTimerRun{env: c.env, workflows: c.workflows}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *TestExampleClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	c.env.CancelWorkflow()
	return nil
}

// TerminateWorkflow terminates an existing workflow execution
func (c *TestExampleClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.CancelWorkflow(ctx, workflowID, runID)
}

// GetWakeUpTime executes a example.updatabletimer.v1.Example.GetWakeUpTime query
func (c *TestExampleClient) GetWakeUpTime(ctx context.Context, workflowID string, runID string) (*GetWakeUpTimeOutput, error) {
	val, err := c.env.QueryWorkflow(GetWakeUpTimeQueryName)
	if err != nil {
		return nil, err
	} else if !val.HasValue() {
		return nil, nil
	} else {
		var result GetWakeUpTimeOutput
		if err := val.Get(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}
}

// UpdateWakeUpTime executes a example.updatabletimer.v1.Example.UpdateWakeUpTime signal
func (c *TestExampleClient) UpdateWakeUpTime(ctx context.Context, workflowID string, runID string, req *UpdateWakeUpTimeInput) error {
	c.env.SignalWorkflow(UpdateWakeUpTimeSignalName, req)
	return nil
}

var _ UpdatableTimerRun = &testUpdatableTimerRun{}

// testUpdatableTimerRun provides convenience methods for interacting with a(n) UpdatableTimer workflow in the test environment
type testUpdatableTimerRun struct {
	client    *TestExampleClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *UpdatableTimerInput
	workflows ExampleWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testUpdatableTimerRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test UpdatableTimer workflow result
func (r *testUpdatableTimerRun) Get(context.Context) error {
	r.env.ExecuteWorkflow(UpdatableTimerWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return err
	}
	return nil
}

// ID returns a test UpdatableTimer workflow run's workflow ID
func (r *testUpdatableTimerRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testUpdatableTimerRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testUpdatableTimerRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testUpdatableTimerRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// GetWakeUpTime executes a example.updatabletimer.v1.Example.GetWakeUpTime query against a test UpdatableTimer workflow
func (r *testUpdatableTimerRun) GetWakeUpTime(ctx context.Context) (*GetWakeUpTimeOutput, error) {
	return r.client.GetWakeUpTime(ctx, r.ID(), r.RunID())
}

// UpdateWakeUpTime executes a example.updatabletimer.v1.Example.UpdateWakeUpTime signal against a test UpdatableTimer workflow
func (r *testUpdatableTimerRun) UpdateWakeUpTime(ctx context.Context, req *UpdateWakeUpTimeInput) error {
	return r.client.UpdateWakeUpTime(ctx, r.ID(), r.RunID(), req)
}

// ExampleCliOptions describes runtime configuration for example.updatabletimer.v1.Example cli
type ExampleCliOptions struct {
	after            func(*v2.Context) error
	before           func(*v2.Context) error
	clientForCommand func(*v2.Context) (client.Client, error)
	worker           func(*v2.Context, client.Client) (worker.Worker, error)
}

// NewExampleCliOptions initializes a new ExampleCliOptions value
func NewExampleCliOptions() *ExampleCliOptions {
	return &ExampleCliOptions{}
}

// WithAfter injects a custom After hook to be run after any command invocation
func (opts *ExampleCliOptions) WithAfter(fn func(*v2.Context) error) *ExampleCliOptions {
	opts.after = fn
	return opts
}

// WithBefore injects a custom Before hook to be run prior to any command invocation
func (opts *ExampleCliOptions) WithBefore(fn func(*v2.Context) error) *ExampleCliOptions {
	opts.before = fn
	return opts
}

// WithClient provides a Temporal client factory for use by commands
func (opts *ExampleCliOptions) WithClient(fn func(*v2.Context) (client.Client, error)) *ExampleCliOptions {
	opts.clientForCommand = fn
	return opts
}

// WithWorker provides an method for initializing a worker
func (opts *ExampleCliOptions) WithWorker(fn func(*v2.Context, client.Client) (worker.Worker, error)) *ExampleCliOptions {
	opts.worker = fn
	return opts
}

// NewExampleCli initializes a cli for a(n) example.updatabletimer.v1.Example service
func NewExampleCli(options ...*ExampleCliOptions) (*v2.App, error) {
	commands, err := newExampleCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.App{
		Name:     "example",
		Commands: commands,
	}, nil
}

// NewExampleCliCommand initializes a cli command for a example.updatabletimer.v1.Example service with subcommands for each query, signal, update, and workflow
func NewExampleCliCommand(options ...*ExampleCliOptions) (*v2.Command, error) {
	subcommands, err := newExampleCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.Command{
		Name:        "example",
		Subcommands: subcommands,
	}, nil
}

// newExampleCommands initializes (sub)commands for a example.updatabletimer.v1.Example cli or command
func newExampleCommands(options ...*ExampleCliOptions) ([]*v2.Command, error) {
	opts := &ExampleCliOptions{}
	if len(options) > 0 {
		opts = options[0]
	}
	if opts.clientForCommand == nil {
		opts.clientForCommand = func(*v2.Context) (client.Client, error) {
			return client.Dial(client.Options{})
		}
	}
	commands := []*v2.Command{
		{
			Name:                   "get-wake-up-time",
			Usage:                  "GetWakeUpTime retrieves the current timer expiration timestamp",
			Category:               "QUERIES",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.StringFlag{
					Name:     "workflow-id",
					Usage:    "workflow id",
					Required: true,
					Aliases:  []string{"w"},
				},
				&v2.StringFlag{
					Name:    "run-id",
					Usage:   "run id",
					Aliases: []string{"r"},
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewExampleClient(c)
				if resp, err := client.GetWakeUpTime(cmd.Context, cmd.String("workflow-id"), cmd.String("run-id")); err != nil {
					return fmt.Errorf("error executing %q query: %w", GetWakeUpTimeQueryName, err)
				} else {
					b, err := protojson.Marshal(resp)
					if err != nil {
						return fmt.Errorf("error serializing response json: %w", err)
					}
					var out bytes.Buffer
					if err := json.Indent(&out, b, "", "  "); err != nil {
						return fmt.Errorf("error formatting json: %w", err)
					}
					fmt.Println(out.String())
					return nil
				}
			},
		},
		{
			Name:                   "update-wake-up-time",
			Usage:                  "UpdateWakeUpTime updates the timer expiration timestamp",
			Category:               "SIGNALS",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.StringFlag{
					Name:     "workflow-id",
					Usage:    "workflow id",
					Required: true,
					Aliases:  []string{"w"},
				},
				&v2.StringFlag{
					Name:    "run-id",
					Usage:   "run id",
					Aliases: []string{"r"},
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "wake-up-time",
					Usage:    "set the value of the operation's \"WakeUpTime\" parameter (e.g. \"2017-01-15T01:30:15.01Z\")",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewExampleClient(c)
				req, err := UnmarshalCliFlagsToUpdateWakeUpTimeInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				if err := client.UpdateWakeUpTime(cmd.Context, cmd.String("workflow-id"), cmd.String("run-id"), req); err != nil {
					return fmt.Errorf("error sending %q signal: %w", UpdateWakeUpTimeSignalName, err)
				}
				fmt.Println("success")
				return nil
			},
		},
		{
			Name:                   "updatable-timer",
			Usage:                  "UpdatableTimer describes an updatable timer workflow",
			Category:               "WORKFLOWS",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.BoolFlag{
					Name:    "detach",
					Usage:   "run workflow in the background and print workflow and execution id",
					Aliases: []string{"d"},
				},
				&v2.StringFlag{
					Name:    "task-queue",
					Usage:   "task queue name",
					Aliases: []string{"t"},
					EnvVars: []string{"TEMPORAL_TASK_QUEUE_NAME", "TEMPORAL_TASK_QUEUE", "TASK_QUEUE_NAME", "TASK_QUEUE"},
					Value:   "updatable-timer",
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "initial-wake-up-time",
					Usage:    "set the value of the operation's \"InitialWakeUpTime\" parameter (e.g. \"2017-01-15T01:30:15.01Z\")",
					Category: "INPUT",
				},
				&v2.StringFlag{
					Name:     "name",
					Usage:    "set the value of the operation's \"Name\" parameter",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				tc, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer tc.Close()
				c := NewExampleClient(tc)
				req, err := UnmarshalCliFlagsToUpdatableTimerInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.UpdatableTimerAsync(cmd.Context, req, NewUpdatableTimerOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", UpdatableTimerWorkflowName, err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow id: %s\n", run.ID())
					fmt.Printf("run id: %s\n", run.RunID())
					return nil
				}
				if err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					return nil
				}
			},
		},
	}
	if opts.worker != nil {
		commands = append(commands, []*v2.Command{
			{
				Name:                   "worker",
				Usage:                  "runs a example.updatabletimer.v1.Example worker process",
				UseShortOptionHandling: true,
				Before:                 opts.before,
				After:                  opts.after,
				Action: func(cmd *v2.Context) error {
					c, err := opts.clientForCommand(cmd)
					if err != nil {
						return fmt.Errorf("error initializing client for command: %w", err)
					}
					defer c.Close()
					w, err := opts.worker(cmd, c)
					if opts.worker != nil {
						if err != nil {
							return fmt.Errorf("error initializing worker: %w", err)
						}
					}
					if err := w.Start(); err != nil {
						return fmt.Errorf("error starting worker: %w", err)
					}
					defer w.Stop()
					<-cmd.Context.Done()
					return nil
				},
			},
		}...)
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
	return commands, nil
}

// UnmarshalCliFlagsToUpdateWakeUpTimeInput unmarshals a UpdateWakeUpTimeInput from command line flags
func UnmarshalCliFlagsToUpdateWakeUpTimeInput(cmd *v2.Context) (*UpdateWakeUpTimeInput, error) {
	var result UpdateWakeUpTimeInput
	var hasValues bool
	if cmd.IsSet("input-file") {
		inputFile, err := gohomedir.Expand(cmd.String("input-file"))
		if err != nil {
			inputFile = cmd.String("input-file")
		}
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error reading input-file: %w", err)
		}
		if err := protojson.Unmarshal(b, &result); err != nil {
			return nil, fmt.Errorf("error parsing input-file json: %w", err)
		}
		hasValues = true
	}
	if cmd.IsSet("wake-up-time") {
		hasValues = true
		v, err := time.Parse(time.RFC3339Nano, cmd.String("wake-up-time"))
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling \"wake-up-time\" timestamp flag: %w", err)
		}
		result.WakeUpTime = timestamppb.New(v)
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// UnmarshalCliFlagsToUpdatableTimerInput unmarshals a UpdatableTimerInput from command line flags
func UnmarshalCliFlagsToUpdatableTimerInput(cmd *v2.Context) (*UpdatableTimerInput, error) {
	var result UpdatableTimerInput
	var hasValues bool
	if cmd.IsSet("input-file") {
		inputFile, err := gohomedir.Expand(cmd.String("input-file"))
		if err != nil {
			inputFile = cmd.String("input-file")
		}
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error reading input-file: %w", err)
		}
		if err := protojson.Unmarshal(b, &result); err != nil {
			return nil, fmt.Errorf("error parsing input-file json: %w", err)
		}
		hasValues = true
	}
	if cmd.IsSet("initial-wake-up-time") {
		hasValues = true
		v, err := time.Parse(time.RFC3339Nano, cmd.String("initial-wake-up-time"))
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling \"initial-wake-up-time\" timestamp flag: %w", err)
		}
		result.InitialWakeUpTime = timestamppb.New(v)
	}
	if cmd.IsSet("name") {
		hasValues = true
		result.Name = cmd.String("name")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// WithExampleSchemeTypes registers all Example protobuf types with the given scheme
func WithExampleSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_example_updatabletimer_v1_updatabletimer_proto.Messages().ByName("GetWakeUpTimeOutput"))
		s.RegisterType(File_example_updatabletimer_v1_updatabletimer_proto.Messages().ByName("UpdateWakeUpTimeInput"))
		s.RegisterType(File_example_updatabletimer_v1_updatabletimer_proto.Messages().ByName("UpdatableTimerInput"))
	}
}
