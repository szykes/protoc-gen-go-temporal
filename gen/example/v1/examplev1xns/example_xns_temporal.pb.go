// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.2-next (0b994cbcb4a60e84ee2566d444a20428b160f204)
//	go go1.22.1
//	protoc (unknown)
//
// source: example/v1/example.proto
package examplev1xns

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	v11 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	v13 "go.temporal.io/api/enums/v1"
	v12 "go.temporal.io/api/update/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// ExampleOptions is used to configure example.v1.Example xns activity registration
type ExampleOptions struct {
	// Filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	Filter func(string) string
}

// filter is used to filter xns activity registrations
func (opts *ExampleOptions) filter(name string) string {
	if opts == nil || opts.Filter == nil {
		return name
	}
	return opts.Filter(name)
}

// exampleOptions is a reference to the ExampleOptions initialized at registration
var exampleOptions *ExampleOptions

// RegisterExampleActivities registers example.v1.Example cross-namespace activities
func RegisterExampleActivities(r worker.ActivityRegistry, c v1.ExampleClient, opts ...ExampleOptions) {
	if len(opts) > 0 {
		exampleOptions = &opts[0]
	}
	a := &exampleActivities{c}
	if name := exampleOptions.filter("example.v1.Example.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filter(v1.CreateFooWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.CreateFoo, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filter("example.v1.Example.CreateFooWithSetFooProgress"); name != "" {
		r.RegisterActivityWithOptions(a.CreateFooWithSetFooProgress, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filter(v1.GetFooProgressQueryName); name != "" {
		r.RegisterActivityWithOptions(a.GetFooProgress, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filter(v1.SetFooProgressSignalName); name != "" {
		r.RegisterActivityWithOptions(a.SetFooProgress, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filter(v1.UpdateFooProgressUpdateName); name != "" {
		r.RegisterActivityWithOptions(a.UpdateFooProgress, activity.RegisterOptions{Name: name})
	}
}

// CreateFooWorkflowOptions are used to configure a(n) example.v1.Example.CreateFoo workflow execution
type CreateFooWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewCreateFooWorkflowOptions initializes a new CreateFooWorkflowOptions value
func NewCreateFooWorkflowOptions() *CreateFooWorkflowOptions {
	return &CreateFooWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *CreateFooWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *CreateFooWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *CreateFooWorkflowOptions) WithDetached(d bool) *CreateFooWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *CreateFooWorkflowOptions) WithHeartbeatInterval(d time.Duration) *CreateFooWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *CreateFooWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *CreateFooWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// CreateFooRun provides a handle to a example.v1.Example.CreateFoo workflow execution
type CreateFooRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*v1.CreateFooResponse, error)

	// ID returns the workflow id
	ID() string

	// GetFooProgress returns the status of a CreateFoo operation
	GetFooProgress(workflow.Context, ...*GetFooProgressQueryOptions) (*v1.GetFooProgressResponse, error)

	// GetFooProgress returns the status of a CreateFoo operation
	GetFooProgressAsync(workflow.Context, ...*GetFooProgressQueryOptions) (GetFooProgressQueryHandle, error)

	// SetFooProgress sets the current status of a CreateFoo operation
	SetFooProgress(workflow.Context, *v1.SetFooProgressRequest, ...*SetFooProgressSignalOptions) error

	// SetFooProgress sets the current status of a CreateFoo operation
	SetFooProgressAsync(workflow.Context, *v1.SetFooProgressRequest, ...*SetFooProgressSignalOptions) (SetFooProgressSignalHandle, error)

	// UpdateFooProgress sets the current status of a CreateFoo operation
	UpdateFooProgress(workflow.Context, *v1.SetFooProgressRequest, ...*UpdateFooProgressUpdateOptions) (*v1.GetFooProgressResponse, error)

	// UpdateFooProgress sets the current status of a CreateFoo operation
	UpdateFooProgressAsync(workflow.Context, *v1.SetFooProgressRequest, ...*UpdateFooProgressUpdateOptions) (UpdateFooProgressHandle, error)
}

// createFooRun provides a(n) CreateFooRun implementation
type createFooRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *createFooRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelExampleWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *createFooRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *createFooRun) Get(ctx workflow.Context) (*v1.CreateFooResponse, error) {
	var resp v1.CreateFooResponse
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ID returns the underlying workflow id
func (r *createFooRun) ID() string {
	return r.id
}

// GetFooProgress returns the status of a CreateFoo operation
func (r *createFooRun) GetFooProgress(ctx workflow.Context, opts ...*GetFooProgressQueryOptions) (*v1.GetFooProgressResponse, error) {
	return GetFooProgress(ctx, r.ID(), "", opts...)
}

// GetFooProgress returns the status of a CreateFoo operation
func (r *createFooRun) GetFooProgressAsync(ctx workflow.Context, opts ...*GetFooProgressQueryOptions) (GetFooProgressQueryHandle, error) {
	return GetFooProgressAsync(ctx, r.ID(), "", opts...)
}

// SetFooProgress sets the current status of a CreateFoo operation
func (r *createFooRun) SetFooProgress(ctx workflow.Context, req *v1.SetFooProgressRequest, opts ...*SetFooProgressSignalOptions) error {
	return SetFooProgress(ctx, r.ID(), "", req, opts...)
}

// SetFooProgress sets the current status of a CreateFoo operation
func (r *createFooRun) SetFooProgressAsync(ctx workflow.Context, req *v1.SetFooProgressRequest, opts ...*SetFooProgressSignalOptions) (SetFooProgressSignalHandle, error) {
	return SetFooProgressAsync(ctx, r.ID(), "", req, opts...)
}

// UpdateFooProgress sets the current status of a CreateFoo operation
func (r *createFooRun) UpdateFooProgress(ctx workflow.Context, req *v1.SetFooProgressRequest, opts ...*UpdateFooProgressUpdateOptions) (*v1.GetFooProgressResponse, error) {
	return UpdateFooProgress(ctx, r.ID(), "", req, opts...)
}

// UpdateFooProgress sets the current status of a CreateFoo operation
func (r *createFooRun) UpdateFooProgressAsync(ctx workflow.Context, req *v1.SetFooProgressRequest, opts ...*UpdateFooProgressUpdateOptions) (UpdateFooProgressHandle, error) {
	return UpdateFooProgressAsync(ctx, r.ID(), "", req, opts...)
}

// CreateFoo creates a new foo operation
func CreateFoo(ctx workflow.Context, req *v1.CreateFooRequest, opts ...*CreateFooWorkflowOptions) (*v1.CreateFooResponse, error) {
	run, err := CreateFooAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// CreateFoo creates a new foo operation
func CreateFooAsync(ctx workflow.Context, req *v1.CreateFooRequest, opts ...*CreateFooWorkflowOptions) (CreateFooRun, error) {
	activityName := exampleOptions.filter(v1.CreateFooWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.CreateFooWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &CreateFooWorkflowOptions{}
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
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 3600000000000 // 1 hour
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.CreateFooIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.v1.Example.CreateFoo\" workflow", "error", err)
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

	ctx, cancel := workflow.WithCancel(ctx)
	return &createFooRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// CreateFooWithSetFooProgress sends a(n) example.v1.Example.SetFooProgress signal to a example.v1.Example.CreateFoo workflow, starting it if necessary, and blocks until the workflow completes
func CreateFooWithSetFooProgress(ctx workflow.Context, req *v1.CreateFooRequest, signal *v1.SetFooProgressRequest, opts ...*CreateFooWorkflowOptions) (*v1.CreateFooResponse, error) {
	run, err := CreateFooWithSetFooProgressAsync(ctx, req, signal, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// CreateFooWithSetFooProgressAsync sends a(n) example.v1.Example.SetFooProgress signal to a(n) example.v1.Example.CreateFoo workflow, starting it if necessary, and returns a handle to the underlying activity
func CreateFooWithSetFooProgressAsync(ctx workflow.Context, req *v1.CreateFooRequest, signal *v1.SetFooProgressRequest, opts ...*CreateFooWorkflowOptions) (CreateFooRun, error) {
	activityName := exampleOptions.filter("example.v1.Example.CreateFooWithSetFooProgress")
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", "example.v1.Example.CreateFooWithSetFooProgress"),
			"Unimplemented",
			nil,
		)
	}

	opt := &CreateFooWorkflowOptions{}
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
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 3600000000000 // 1 hour
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.CreateFooIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.v1.Example.CreateFoo\" workflow", "error", err)
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

	// marshal signal request protobuf message
	wsignal, err := anypb.New(signal)
	if err != nil {
		return nil, fmt.Errorf("error marshalling signal request: %w", err)
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &createFooRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			Request:              wreq,
			Signal:               wsignal,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// GetFooProgressQueryOptions are used to configure a(n) example.v1.Example.GetFooProgress query execution
type GetFooProgressQueryOptions struct {
	ActivityOptions   *workflow.ActivityOptions
	HeartbeatInterval time.Duration
}

// NewGetFooProgressQueryOptions initializes a new GetFooProgressQueryOptions value
func NewGetFooProgressQueryOptions() *GetFooProgressQueryOptions {
	return &GetFooProgressQueryOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *GetFooProgressQueryOptions) WithActivityOptions(ao workflow.ActivityOptions) *GetFooProgressQueryOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *GetFooProgressQueryOptions) WithHeartbeatInterval(d time.Duration) *GetFooProgressQueryOptions {
	opts.HeartbeatInterval = d
	return opts
}

// GetFooProgressQueryHandle provides a handle for a example.v1.Example.GetFooProgress query activity
type GetFooProgressQueryHandle interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*v1.GetFooProgressResponse, error)
}

// getFooProgressQueryHandle provides a(n) GetFooProgressQueryHandle implementation
type getFooProgressQueryHandle struct {
	cancel func()
	future workflow.Future
}

// Cancel the underlying query activity
func (r *getFooProgressQueryHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *getFooProgressQueryHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying query result
func (r *getFooProgressQueryHandle) Get(ctx workflow.Context) (*v1.GetFooProgressResponse, error) {
	var resp v1.GetFooProgressResponse
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFooProgress returns the status of a CreateFoo operation
func GetFooProgress(ctx workflow.Context, workflowID string, runID string, opts ...*GetFooProgressQueryOptions) (*v1.GetFooProgressResponse, error) {
	handle, err := GetFooProgressAsync(ctx, workflowID, runID, opts...)
	if err != nil {
		return nil, err
	}
	return handle.Get(ctx)
}

// GetFooProgressAsync executes a(n) example.v1.Example.GetFooProgress query and returns a handle to the activity
func GetFooProgressAsync(ctx workflow.Context, workflowID string, runID string, opts ...*GetFooProgressQueryOptions) (GetFooProgressQueryHandle, error) {
	activityName := exampleOptions.filter(v1.GetFooProgressQueryName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.GetFooProgressQueryName),
			"Unimplemented",
			nil,
		)
	}

	opt := &GetFooProgressQueryOptions{}
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
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 60000000000 // 1 minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	ctx, cancel := workflow.WithCancel(ctx)
	return &getFooProgressQueryHandle{
		cancel: cancel,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.QueryRequest{
			HeartbeatInterval: durationpb.New(opt.HeartbeatInterval),
			WorkflowId:        workflowID,
			RunId:             runID,
		}),
	}, nil
}

// SetFooProgressSignalOptions are used to configure a(n) example.v1.Example.SetFooProgress signal execution
type SetFooProgressSignalOptions struct {
	ActivityOptions   *workflow.ActivityOptions
	HeartbeatInterval time.Duration
}

// NewSetFooProgressSignalOptions initializes a new SetFooProgressSignalOptions value
func NewSetFooProgressSignalOptions() *SetFooProgressSignalOptions {
	return &SetFooProgressSignalOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *SetFooProgressSignalOptions) WithActivityOptions(ao workflow.ActivityOptions) *SetFooProgressSignalOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *SetFooProgressSignalOptions) WithHeartbeatInterval(d time.Duration) *SetFooProgressSignalOptions {
	opts.HeartbeatInterval = d
	return opts
}

// SetFooProgressSignalHandle provides a handle for a example.v1.Example.SetFooProgress signal activity
type SetFooProgressSignalHandle interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error
	// Future returns the inner workflow.Future
	Future() workflow.Future
	// Get returns the inner workflow.Future
	Get(workflow.Context) error
}

// setFooProgressSignalHandle provides a(n) SetFooProgressQueryHandle implementation
type setFooProgressSignalHandle struct {
	cancel func()
	future workflow.Future
}

// Cancel the underlying signal activity
func (r *setFooProgressSignalHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *setFooProgressSignalHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion
func (r *setFooProgressSignalHandle) Get(ctx workflow.Context) error {
	return r.future.Get(ctx, nil)
}

// SetFooProgress sets the current status of a CreateFoo operation
func SetFooProgress(ctx workflow.Context, workflowID string, runID string, req *v1.SetFooProgressRequest, opts ...*SetFooProgressSignalOptions) error {
	handle, err := SetFooProgressAsync(ctx, workflowID, runID, req, opts...)
	if err != nil {
		return err
	}
	return handle.Get(ctx)
}

// SetFooProgressAsync executes a(n) example.v1.Example.SetFooProgress signal
func SetFooProgressAsync(ctx workflow.Context, workflowID string, runID string, req *v1.SetFooProgressRequest, opts ...*SetFooProgressSignalOptions) (SetFooProgressSignalHandle, error) {
	activityName := exampleOptions.filter(v1.SetFooProgressSignalName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.SetFooProgressSignalName),
			"Unimplemented",
			nil,
		)
	}

	opt := &SetFooProgressSignalOptions{}
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
	return &setFooProgressSignalHandle{
		cancel: cancel,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.SignalRequest{
			HeartbeatInterval: durationpb.New(opt.HeartbeatInterval),
			WorkflowId:        workflowID,
			RunId:             runID,
			Request:           wreq,
		}),
	}, nil
}

// UpdateFooProgressUpdateOptions are used to configure a(n) example.v1.Example.UpdateFooProgress update execution
type UpdateFooProgressUpdateOptions struct {
	ActivityOptions       *workflow.ActivityOptions
	HeartbeatInterval     time.Duration
	UpdateWorkflowOptions *client.UpdateWorkflowWithOptionsRequest
}

// NewUpdateFooProgressUpdateOptions initializes a new UpdateFooProgressUpdateOptions value
func NewUpdateFooProgressUpdateOptions() *UpdateFooProgressUpdateOptions {
	return &UpdateFooProgressUpdateOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *UpdateFooProgressUpdateOptions) WithActivityOptions(ao workflow.ActivityOptions) *UpdateFooProgressUpdateOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *UpdateFooProgressUpdateOptions) WithHeartbeatInterval(d time.Duration) *UpdateFooProgressUpdateOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithUpdateWorkflowOptions can be used to customize the update workflow options
func (opts *UpdateFooProgressUpdateOptions) WithUpdateWorkflowOptions(uwo client.UpdateWorkflowWithOptionsRequest) *UpdateFooProgressUpdateOptions {
	opts.UpdateWorkflowOptions = &uwo
	return opts
}

// UpdateFooProgressHandle provides a handle to a example.v1.Example.UpdateFooProgress workflow update
type UpdateFooProgressHandle interface {
	// Cancel cancels the update activity
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get blocks on update completion and returns the result
	Get(workflow.Context) (*v1.GetFooProgressResponse, error)

	// ID returns the update id
	ID() string
}

// updateFooProgressHandle provides a(n) UpdateFooProgressHandle implementation
type updateFooProgressHandle struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow update
func (r *updateFooProgressHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *updateFooProgressHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying update result
func (r *updateFooProgressHandle) Get(ctx workflow.Context) (*v1.GetFooProgressResponse, error) {
	var resp v1.GetFooProgressResponse
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ID returns the underlying workflow id
func (r *updateFooProgressHandle) ID() string {
	return r.id
}

// UpdateFooProgress sets the current status of a CreateFoo operation
func UpdateFooProgress(ctx workflow.Context, workflowID string, runID string, req *v1.SetFooProgressRequest, opts ...*UpdateFooProgressUpdateOptions) (*v1.GetFooProgressResponse, error) {
	run, err := UpdateFooProgressAsync(ctx, workflowID, runID, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// UpdateFooProgressAsync executes a(n) example.v1.Example.UpdateFooProgress update and blocks until error or response received
func UpdateFooProgressAsync(ctx workflow.Context, workflowID string, runID string, req *v1.SetFooProgressRequest, opts ...*UpdateFooProgressUpdateOptions) (UpdateFooProgressHandle, error) {
	activityName := exampleOptions.filter(v1.UpdateFooProgressUpdateName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.UpdateFooProgressUpdateName),
			"Unimplemented",
			nil,
		)
	}

	opt := &UpdateFooProgressUpdateOptions{}
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
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 300000000000 // 5 minutes
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	uo := client.UpdateWorkflowWithOptionsRequest{}
	if opt.UpdateWorkflowOptions != nil {
		uo = *opt.UpdateWorkflowOptions
	}
	uo.WorkflowID = workflowID
	uo.RunID = runID
	if uo.UpdateID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := expression.EvalExpression(v1.UpdateFooProgressIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.v1.Example.UpdateFooProgress\" update", "error", err)
				return nil
			}
			return id
		}).Get(&uo.UpdateID); err != nil {
			return nil, err
		}
	}
	if uo.UpdateID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := uuid.NewRandom()
			if err != nil {
				workflow.GetLogger(ctx).Error("error generating update id", "error", err)
				return nil
			}
			return id
		}).Get(&uo.UpdateID); err != nil {
			return nil, err
		}
	}
	if uo.UpdateID == "" {
		return nil, temporal.NewNonRetryableApplicationError("update id is required", "InvalidArgument", nil)
	}

	uopb, err := xns.MarshalUpdateWorkflowOptions(uo)
	if err != nil {
		return nil, fmt.Errorf("error marshalling update workflow options: %w", err)
	}

	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling update request: %w", err)
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &updateFooProgressHandle{
		cancel: cancel,
		id:     uo.UpdateID,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.UpdateRequest{
			HeartbeatInterval:     durationpb.New(opt.HeartbeatInterval),
			Request:               wreq,
			UpdateWorkflowOptions: uopb,
		}),
	}, nil
}

// CancelExampleWorkflow cancels an existing workflow
func CancelExampleWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelExampleWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelExampleWorkflowAsync cancels an existing workflow
func CancelExampleWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := exampleOptions.filter("example.v1.Example.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for example.v1.Example.CancelWorkflow",
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

// CreateFoo executes a(n) example.v1.Example.CreateFoo workflow via an activity
func (a *exampleActivities) CreateFoo(ctx context.Context, input *v11.WorkflowRequest) (resp *v1.CreateFooResponse, err error) {
	// unmarshal workflow request
	var req v1.CreateFooRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/v1.CreateFooRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		)
	}

	// initialize workflow execution
	var run v1.CreateFooRun
	run, err = a.client.CreateFooAsync(ctx, &req, v1.NewCreateFooOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return nil, err
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil, nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = run.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())
		case <-ctx.Done():
			if err := run.Cancel(ctx); err != nil {
				return nil, err
			}
			return nil, workflow.ErrCanceled
		case <-doneCh:
			return resp, err
		}
	}
}

// CreateFooWithSetFooProgress sends a(n) example.v1.Example.SetFooProgress signal to a(n) example.v1.Example.CreateFoo workflow via an activity
func (a *exampleActivities) CreateFooWithSetFooProgress(ctx context.Context, input *v11.WorkflowRequest) (resp *v1.CreateFooResponse, err error) {
	// unmarshal workflow request
	var req v1.CreateFooRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/v1.CreateFooRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		)
	}

	// unmarshal signal request
	var signal v1.SetFooProgressRequest
	if err := input.Signal.UnmarshalTo(&signal); err != nil {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling signal request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/v1.SetFooProgressRequest", input.Signal.GetTypeUrl()),
			"InvalidArgument",
			err,
		)
	}

	// initialize workflow execution
	var run v1.CreateFooRun
	run, err = a.client.CreateFooWithSetFooProgressAsync(ctx, &req, &signal, v1.NewCreateFooOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return nil, err
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil, nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = run.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())
		case <-ctx.Done():
			if err := run.Cancel(ctx); err != nil {
				return nil, err
			}
			return nil, workflow.ErrCanceled
		case <-doneCh:
			return resp, err
		}
	}
}

// GetFooProgress executes a(n) example.v1.Example.GetFooProgress query via an activity
func (a *exampleActivities) GetFooProgress(ctx context.Context, input *v11.QueryRequest) (resp *v1.GetFooProgressResponse, err error) {
	// execute signal in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = a.client.GetFooProgress(ctx, input.GetWorkflowId(), input.GetRunId())
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
			return nil, ctx.Err()
		case <-doneCh:
			return resp, err
		}
	}
}

// SetFooProgress executes a(n) example.v1.Example.SetFooProgress signal via an activity
func (a *exampleActivities) SetFooProgress(ctx context.Context, input *v11.SignalRequest) (err error) {
	// unmarshal signal request
	var req v1.SetFooProgressRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling signal request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/v1.SetFooProgressRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		)
	}
	// execute signal in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = a.client.SetFooProgress(ctx, input.GetWorkflowId(), input.GetRunId(), &req)
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
			return ctx.Err()
		case <-doneCh:
			return err
		}
	}
}

// UpdateFooProgress executes a(n) example.v1.Example.UpdateFooProgress update via an activity
func (a *exampleActivities) UpdateFooProgress(ctx context.Context, input *v11.UpdateRequest) (resp *v1.GetFooProgressResponse, err error) {
	var handle v1.UpdateFooProgressHandle
	if activity.HasHeartbeatDetails(ctx) {
		// extract update id from heartbeat details
		var updateID string
		if err := activity.GetHeartbeatDetails(ctx, &updateID); err != nil {
			return nil, err
		}

		// retrieve handle for existing update
		handle, err = a.client.GetUpdateFooProgress(ctx, client.GetWorkflowUpdateHandleOptions{
			WorkflowID: input.GetUpdateWorkflowOptions().GetWorkflowId(),
			RunID:      input.GetUpdateWorkflowOptions().GetRunId(),
			UpdateID:   updateID,
		})
		if err != nil {
			return nil, err
		}
	} else {
		// unmarshal update request
		var req v1.SetFooProgressRequest
		if err := input.Request.UnmarshalTo(&req); err != nil {
			return nil, temporal.NewNonRetryableApplicationError(
				fmt.Sprintf("error unmarshalling update request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/v1.SetFooProgressRequest", input.Request.GetTypeUrl()),
				"InvalidArgument",
				err,
			)
		}

		uo := xns.UnmarshalUpdateWorkflowOptions(input.GetUpdateWorkflowOptions())
		uo.WaitPolicy = &v12.WaitPolicy{
			LifecycleStage: v13.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED,
		}

		// initialize update execution
		handle, err = a.client.UpdateFooProgressAsync(
			ctx,
			input.GetUpdateWorkflowOptions().GetWorkflowId(),
			input.GetUpdateWorkflowOptions().GetRunId(),
			&req,
			v1.NewUpdateFooProgressOptions().WithUpdateWorkflowOptions(uo),
		)
		if err != nil {
			return nil, err
		}
		activity.RecordHeartbeat(ctx, handle.UpdateID())
	}

	// wait for update to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = handle.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow update to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, handle.UpdateID())
		case <-ctx.Done():
			return nil, workflow.ErrCanceled
		case <-doneCh:
			return resp, err
		}
	}
}
