// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.0.1-next (1420d48f2ccd197d05de90be67840e6a2b0584cd)
//	go go1.21.13
//	protoc (unknown)
//
// source: example/updatabletimer/v1/updatabletimer.proto
package updatabletimerv1xns

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/example/updatabletimer/v1"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	enumsv1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// ExampleOptions is used to configure example.updatabletimer.v1.Example xns activity registration
type ExampleOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewExampleOptions initializes a new ExampleOptions value
func NewExampleOptions() *ExampleOptions {
	return &ExampleOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *ExampleOptions) WithErrorConverter(errorConverter func(error) error) *ExampleOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *ExampleOptions) WithFilter(filter func(string) string) *ExampleOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *ExampleOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *ExampleOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// exampleOptions is a reference to the ExampleOptions initialized at registration
var exampleOptions *ExampleOptions

// RegisterExampleActivities registers example.updatabletimer.v1.Example cross-namespace activities
func RegisterExampleActivities(r worker.ActivityRegistry, c v1.ExampleClient, options ...*ExampleOptions) {
	if exampleOptions == nil && len(options) > 0 && options[0] != nil {
		exampleOptions = options[0]
	}
	a := &exampleActivities{c}
	if name := exampleOptions.filterActivity("example.updatabletimer.v1.Example.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.UpdatableTimerWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.UpdatableTimer, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.GetWakeUpTimeQueryName); name != "" {
		r.RegisterActivityWithOptions(a.GetWakeUpTime, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.UpdateWakeUpTimeSignalName); name != "" {
		r.RegisterActivityWithOptions(a.UpdateWakeUpTime, activity.RegisterOptions{Name: name})
	}
}

// UpdatableTimerWorkflowOptions are used to configure a(n) UpdatableTimer workflow execution
type UpdatableTimerWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewUpdatableTimerWorkflowOptions initializes a new UpdatableTimerWorkflowOptions value
func NewUpdatableTimerWorkflowOptions() *UpdatableTimerWorkflowOptions {
	return &UpdatableTimerWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *UpdatableTimerWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *UpdatableTimerWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *UpdatableTimerWorkflowOptions) WithDetached(d bool) *UpdatableTimerWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *UpdatableTimerWorkflowOptions) WithHeartbeatInterval(d time.Duration) *UpdatableTimerWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *UpdatableTimerWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *UpdatableTimerWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *UpdatableTimerWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *UpdatableTimerWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// UpdatableTimerRun provides a handle to a UpdatableTimer workflow execution
type UpdatableTimerRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) error

	// ID returns the workflow id
	ID() string

	// GetWakeUpTime retrieves the current timer expiration timestamp
	GetWakeUpTime(workflow.Context, ...*GetWakeUpTimeQueryOptions) (*v1.GetWakeUpTimeOutput, error)

	// GetWakeUpTime retrieves the current timer expiration timestamp
	GetWakeUpTimeAsync(workflow.Context, ...*GetWakeUpTimeQueryOptions) (GetWakeUpTimeQueryHandle, error)

	// UpdateWakeUpTime updates the timer expiration timestamp
	UpdateWakeUpTime(workflow.Context, *v1.UpdateWakeUpTimeInput, ...*UpdateWakeUpTimeSignalOptions) error

	// UpdateWakeUpTime updates the timer expiration timestamp
	UpdateWakeUpTimeAsync(workflow.Context, *v1.UpdateWakeUpTimeInput, ...*UpdateWakeUpTimeSignalOptions) (UpdateWakeUpTimeSignalHandle, error)
}

// updatableTimerRun provides a(n) UpdatableTimerRun implementation
type updatableTimerRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *updatableTimerRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelExampleWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *updatableTimerRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *updatableTimerRun) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *updatableTimerRun) ID() string {
	return r.id
}

// GetWakeUpTime retrieves the current timer expiration timestamp
func (r *updatableTimerRun) GetWakeUpTime(ctx workflow.Context, opts ...*GetWakeUpTimeQueryOptions) (*v1.GetWakeUpTimeOutput, error) {
	return GetWakeUpTime(ctx, r.ID(), "", opts...)
}

// GetWakeUpTime retrieves the current timer expiration timestamp
func (r *updatableTimerRun) GetWakeUpTimeAsync(ctx workflow.Context, opts ...*GetWakeUpTimeQueryOptions) (GetWakeUpTimeQueryHandle, error) {
	return GetWakeUpTimeAsync(ctx, r.ID(), "", opts...)
}

// UpdateWakeUpTime updates the timer expiration timestamp
func (r *updatableTimerRun) UpdateWakeUpTime(ctx workflow.Context, req *v1.UpdateWakeUpTimeInput, opts ...*UpdateWakeUpTimeSignalOptions) error {
	return UpdateWakeUpTime(ctx, r.ID(), "", req, opts...)
}

// UpdateWakeUpTime updates the timer expiration timestamp
func (r *updatableTimerRun) UpdateWakeUpTimeAsync(ctx workflow.Context, req *v1.UpdateWakeUpTimeInput, opts ...*UpdateWakeUpTimeSignalOptions) (UpdateWakeUpTimeSignalHandle, error) {
	return UpdateWakeUpTimeAsync(ctx, r.ID(), "", req, opts...)
}

// UpdatableTimer describes an updatable timer workflow
func UpdatableTimer(ctx workflow.Context, req *v1.UpdatableTimerInput, opts ...*UpdatableTimerWorkflowOptions) error {
	run, err := UpdatableTimerAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// UpdatableTimer describes an updatable timer workflow
func UpdatableTimerAsync(ctx workflow.Context, req *v1.UpdatableTimerInput, opts ...*UpdatableTimerWorkflowOptions) (UpdatableTimerRun, error) {
	activityName := exampleOptions.filterActivity(v1.UpdatableTimerWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.UpdatableTimerWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &UpdatableTimerWorkflowOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}
	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 86400000000000 // 1 day
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.UpdatableTimerIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.updatabletimer.v1.Example.UpdatableTimer\" workflow", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := uuid.NewRandom()
			if err != nil {
				workflow.GetLogger(ctx).Error("error generating workflow id", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
	}
	if wo.ID == "" {
		return nil, temporal.NewNonRetryableApplicationError("workflow id is required", "InvalidArgument", nil)
	}

	// marshal start workflow options protobuf message
	swo, err := xns.MarshalStartWorkflowOptions(wo)
	if err != nil {
		return nil, fmt.Errorf("error marshalling start workflow options: %w", err)
	}

	// marshal workflow request protobuf message
	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling workflow request: %w", err)
	}

	var parentClosePolicy temporalv1.ParentClosePolicy
	switch opt.ParentClosePolicy {
	case enumsv1.PARENT_CLOSE_POLICY_ABANDON:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON
	case enumsv1.PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case enumsv1.PARENT_CLOSE_POLICY_TERMINATE:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &updatableTimerRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			ParentClosePolicy:    parentClosePolicy,
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// GetWakeUpTimeQueryOptions are used to configure a(n) example.updatabletimer.v1.Example.GetWakeUpTime query execution
type GetWakeUpTimeQueryOptions struct {
	ActivityOptions   *workflow.ActivityOptions
	HeartbeatInterval time.Duration
}

// NewGetWakeUpTimeQueryOptions initializes a new GetWakeUpTimeQueryOptions value
func NewGetWakeUpTimeQueryOptions() *GetWakeUpTimeQueryOptions {
	return &GetWakeUpTimeQueryOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *GetWakeUpTimeQueryOptions) WithActivityOptions(ao workflow.ActivityOptions) *GetWakeUpTimeQueryOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *GetWakeUpTimeQueryOptions) WithHeartbeatInterval(d time.Duration) *GetWakeUpTimeQueryOptions {
	opts.HeartbeatInterval = d
	return opts
}

// GetWakeUpTimeQueryHandle provides a handle for a example.updatabletimer.v1.Example.GetWakeUpTime query activity
type GetWakeUpTimeQueryHandle interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*v1.GetWakeUpTimeOutput, error)
}

// getWakeUpTimeQueryHandle provides a(n) GetWakeUpTimeQueryHandle implementation
type getWakeUpTimeQueryHandle struct {
	cancel func()
	future workflow.Future
}

// Cancel the underlying query activity
func (r *getWakeUpTimeQueryHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *getWakeUpTimeQueryHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying query result
func (r *getWakeUpTimeQueryHandle) Get(ctx workflow.Context) (*v1.GetWakeUpTimeOutput, error) {
	var resp v1.GetWakeUpTimeOutput
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWakeUpTime retrieves the current timer expiration timestamp
func GetWakeUpTime(ctx workflow.Context, workflowID string, runID string, opts ...*GetWakeUpTimeQueryOptions) (*v1.GetWakeUpTimeOutput, error) {
	handle, err := GetWakeUpTimeAsync(ctx, workflowID, runID, opts...)
	if err != nil {
		return nil, err
	}
	return handle.Get(ctx)
}

// GetWakeUpTimeAsync executes a(n) example.updatabletimer.v1.Example.GetWakeUpTime query and returns a handle to the activity
func GetWakeUpTimeAsync(ctx workflow.Context, workflowID string, runID string, opts ...*GetWakeUpTimeQueryOptions) (GetWakeUpTimeQueryHandle, error) {
	activityName := exampleOptions.filterActivity(v1.GetWakeUpTimeQueryName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.GetWakeUpTimeQueryName),
			"Unimplemented",
			nil,
		)
	}

	opt := &GetWakeUpTimeQueryOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 60000000000 // 1 minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	ctx, cancel := workflow.WithCancel(ctx)
	return &getWakeUpTimeQueryHandle{
		cancel: cancel,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.QueryRequest{
			HeartbeatInterval: durationpb.New(opt.HeartbeatInterval),
			WorkflowId:        workflowID,
			RunId:             runID,
		}),
	}, nil
}

// UpdateWakeUpTimeSignalOptions are used to configure a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal execution
type UpdateWakeUpTimeSignalOptions struct {
	ActivityOptions   *workflow.ActivityOptions
	HeartbeatInterval time.Duration
}

// NewUpdateWakeUpTimeSignalOptions initializes a new UpdateWakeUpTimeSignalOptions value
func NewUpdateWakeUpTimeSignalOptions() *UpdateWakeUpTimeSignalOptions {
	return &UpdateWakeUpTimeSignalOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *UpdateWakeUpTimeSignalOptions) WithActivityOptions(ao workflow.ActivityOptions) *UpdateWakeUpTimeSignalOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *UpdateWakeUpTimeSignalOptions) WithHeartbeatInterval(d time.Duration) *UpdateWakeUpTimeSignalOptions {
	opts.HeartbeatInterval = d
	return opts
}

// UpdateWakeUpTimeSignalHandle provides a handle for a example.updatabletimer.v1.Example.UpdateWakeUpTime signal activity
type UpdateWakeUpTimeSignalHandle interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error
	// Future returns the inner workflow.Future
	Future() workflow.Future
	// Get returns the inner workflow.Future
	Get(workflow.Context) error
}

// updateWakeUpTimeSignalHandle provides a(n) UpdateWakeUpTimeQueryHandle implementation
type updateWakeUpTimeSignalHandle struct {
	cancel func()
	future workflow.Future
}

// Cancel the underlying signal activity
func (r *updateWakeUpTimeSignalHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *updateWakeUpTimeSignalHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion
func (r *updateWakeUpTimeSignalHandle) Get(ctx workflow.Context) error {
	return r.future.Get(ctx, nil)
}

// UpdateWakeUpTime updates the timer expiration timestamp
func UpdateWakeUpTime(ctx workflow.Context, workflowID string, runID string, req *v1.UpdateWakeUpTimeInput, opts ...*UpdateWakeUpTimeSignalOptions) error {
	handle, err := UpdateWakeUpTimeAsync(ctx, workflowID, runID, req, opts...)
	if err != nil {
		return err
	}
	return handle.Get(ctx)
}

// UpdateWakeUpTimeAsync executes a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal
func UpdateWakeUpTimeAsync(ctx workflow.Context, workflowID string, runID string, req *v1.UpdateWakeUpTimeInput, opts ...*UpdateWakeUpTimeSignalOptions) (UpdateWakeUpTimeSignalHandle, error) {
	activityName := exampleOptions.filterActivity(v1.UpdateWakeUpTimeSignalName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.UpdateWakeUpTimeSignalName),
			"Unimplemented",
			nil,
		)
	}

	opt := &UpdateWakeUpTimeSignalOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 60000000000 // 1 minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// marshal workflow request
	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling workflow request: %w", err)
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &updateWakeUpTimeSignalHandle{
		cancel: cancel,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.SignalRequest{
			HeartbeatInterval: durationpb.New(opt.HeartbeatInterval),
			WorkflowId:        workflowID,
			RunId:             runID,
			Request:           wreq,
		}),
	}, nil
}

// CancelExampleWorkflow cancels an existing workflow
func CancelExampleWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelExampleWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelExampleWorkflowAsync cancels an existing workflow
func CancelExampleWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := exampleOptions.filterActivity("example.updatabletimer.v1.Example.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for example.updatabletimer.v1.Example.CancelWorkflow",
			"Unimplemented",
			nil,
		))
		return f
	}
	ao := workflow.GetActivityOptions(ctx)
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.StartToCloseTimeout = time.Minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	return workflow.ExecuteActivity(ctx, activityName, workflowID, runID)
}

// exampleActivities provides activities that can be used to interact with a(n) Example service's workflow, queries, signals, and updates across namespaces
type exampleActivities struct {
	client v1.ExampleClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *exampleActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// UpdatableTimer executes a(n) UpdatableTimer workflow via an activity
func (a *exampleActivities) UpdatableTimer(ctx context.Context, input *xnsv1.WorkflowRequest) (err error) {
	// unmarshal workflow request
	var req v1.UpdatableTimerInput
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return exampleOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/updatabletimer/v1.UpdatableTimerInput", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.UpdatableTimerRun
	run, err = a.client.UpdatableTimerAsync(ctx, &req, v1.NewUpdatableTimerOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return exampleOptions.convertError(err)
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = run.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 30
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		// send heartbeats periodically
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())

		// return retryable error on worker close
		case <-activity.GetWorkerStopChannel(ctx):
			return temporal.NewApplicationError("worker is stopping", "WorkerStopped")

		// catch parent activity context cancellation. in most cases, this should indicate a
		// server-sent cancellation, but there's a non-zero possibility that this cancellation
		// is received due to the worker stopping, prior to detecting the closing of the worker
		// stop channel. to give us an opportunity to detect a cancellation stemming from the
		// worker closing, we again check to see if the worker stop channel is closed before
		// propagating the cancellation
		case <-ctx.Done():
			select {
			case <-activity.GetWorkerStopChannel(ctx):
				return temporal.NewApplicationError("worker is stopping", "WorkerStopped")
			default:
				parentClosePolicy := input.GetParentClosePolicy()
				if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL || parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE {
					disconnectedCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL {
						err = run.Cancel(disconnectedCtx)
					} else {
						err = run.Terminate(disconnectedCtx, "xns activity cancellation received", "error", ctx.Err())
					}
					if err != nil {
						return exampleOptions.convertError(err)
					}
				}
				return exampleOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return exampleOptions.convertError(err)
		}
	}
}

// GetWakeUpTime executes a(n) example.updatabletimer.v1.Example.GetWakeUpTime query via an activity
func (a *exampleActivities) GetWakeUpTime(ctx context.Context, input *xnsv1.QueryRequest) (resp *v1.GetWakeUpTimeOutput, err error) {
	// execute signal in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = a.client.GetWakeUpTime(ctx, input.GetWorkflowId(), input.GetRunId())
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 10
	}

	// heartbeat activity while waiting for signal to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx)
		case <-ctx.Done():
			return nil, exampleOptions.convertError(ctx.Err())
		case <-doneCh:
			return resp, exampleOptions.convertError(err)
		}
	}
}

// UpdateWakeUpTime executes a(n) example.updatabletimer.v1.Example.UpdateWakeUpTime signal via an activity
func (a *exampleActivities) UpdateWakeUpTime(ctx context.Context, input *xnsv1.SignalRequest) (err error) {
	// unmarshal signal request
	var req v1.UpdateWakeUpTimeInput
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return exampleOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling signal request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/updatabletimer/v1.UpdateWakeUpTimeInput", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}
	// execute signal in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = a.client.UpdateWakeUpTime(ctx, input.GetWorkflowId(), input.GetRunId(), &req)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 10
	}

	// heartbeat activity while waiting for signal to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx)
		case <-ctx.Done():
			exampleOptions.convertError(ctx.Err())
		case <-doneCh:
			return exampleOptions.convertError(err)
		}
	}
}
