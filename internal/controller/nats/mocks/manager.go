// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	config "sigs.k8s.io/controller-runtime/pkg/config"

	context "context"

	healthz "sigs.k8s.io/controller-runtime/pkg/healthz"

	http "net/http"

	logr "github.com/go-logr/logr"

	manager "sigs.k8s.io/controller-runtime/pkg/manager"

	meta "k8s.io/apimachinery/pkg/api/meta"

	mock "github.com/stretchr/testify/mock"

	record "k8s.io/client-go/tools/record"

	rest "k8s.io/client-go/rest"

	runtime "k8s.io/apimachinery/pkg/runtime"

	webhook "sigs.k8s.io/controller-runtime/pkg/webhook"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

type Manager_Expecter struct {
	mock *mock.Mock
}

func (_m *Manager) EXPECT() *Manager_Expecter {
	return &Manager_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *Manager) Add(_a0 manager.Runnable) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(manager.Runnable) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Manager_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 manager.Runnable
func (_e *Manager_Expecter) Add(_a0 interface{}) *Manager_Add_Call {
	return &Manager_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *Manager_Add_Call) Run(run func(_a0 manager.Runnable)) *Manager_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(manager.Runnable))
	})
	return _c
}

func (_c *Manager_Add_Call) Return(_a0 error) *Manager_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_Add_Call) RunAndReturn(run func(manager.Runnable) error) *Manager_Add_Call {
	_c.Call.Return(run)
	return _c
}

// AddHealthzCheck provides a mock function with given fields: name, check
func (_m *Manager) AddHealthzCheck(name string, check healthz.Checker) error {
	ret := _m.Called(name, check)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, healthz.Checker) error); ok {
		r0 = rf(name, check)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_AddHealthzCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddHealthzCheck'
type Manager_AddHealthzCheck_Call struct {
	*mock.Call
}

// AddHealthzCheck is a helper method to define mock.On call
//   - name string
//   - check healthz.Checker
func (_e *Manager_Expecter) AddHealthzCheck(name interface{}, check interface{}) *Manager_AddHealthzCheck_Call {
	return &Manager_AddHealthzCheck_Call{Call: _e.mock.On("AddHealthzCheck", name, check)}
}

func (_c *Manager_AddHealthzCheck_Call) Run(run func(name string, check healthz.Checker)) *Manager_AddHealthzCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(healthz.Checker))
	})
	return _c
}

func (_c *Manager_AddHealthzCheck_Call) Return(_a0 error) *Manager_AddHealthzCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_AddHealthzCheck_Call) RunAndReturn(run func(string, healthz.Checker) error) *Manager_AddHealthzCheck_Call {
	_c.Call.Return(run)
	return _c
}

// AddMetricsExtraHandler provides a mock function with given fields: path, handler
func (_m *Manager) AddMetricsExtraHandler(path string, handler http.Handler) error {
	ret := _m.Called(path, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, http.Handler) error); ok {
		r0 = rf(path, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_AddMetricsExtraHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMetricsExtraHandler'
type Manager_AddMetricsExtraHandler_Call struct {
	*mock.Call
}

// AddMetricsExtraHandler is a helper method to define mock.On call
//   - path string
//   - handler http.Handler
func (_e *Manager_Expecter) AddMetricsExtraHandler(path interface{}, handler interface{}) *Manager_AddMetricsExtraHandler_Call {
	return &Manager_AddMetricsExtraHandler_Call{Call: _e.mock.On("AddMetricsExtraHandler", path, handler)}
}

func (_c *Manager_AddMetricsExtraHandler_Call) Run(run func(path string, handler http.Handler)) *Manager_AddMetricsExtraHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.Handler))
	})
	return _c
}

func (_c *Manager_AddMetricsExtraHandler_Call) Return(_a0 error) *Manager_AddMetricsExtraHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_AddMetricsExtraHandler_Call) RunAndReturn(run func(string, http.Handler) error) *Manager_AddMetricsExtraHandler_Call {
	_c.Call.Return(run)
	return _c
}

// AddReadyzCheck provides a mock function with given fields: name, check
func (_m *Manager) AddReadyzCheck(name string, check healthz.Checker) error {
	ret := _m.Called(name, check)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, healthz.Checker) error); ok {
		r0 = rf(name, check)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_AddReadyzCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddReadyzCheck'
type Manager_AddReadyzCheck_Call struct {
	*mock.Call
}

// AddReadyzCheck is a helper method to define mock.On call
//   - name string
//   - check healthz.Checker
func (_e *Manager_Expecter) AddReadyzCheck(name interface{}, check interface{}) *Manager_AddReadyzCheck_Call {
	return &Manager_AddReadyzCheck_Call{Call: _e.mock.On("AddReadyzCheck", name, check)}
}

func (_c *Manager_AddReadyzCheck_Call) Run(run func(name string, check healthz.Checker)) *Manager_AddReadyzCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(healthz.Checker))
	})
	return _c
}

func (_c *Manager_AddReadyzCheck_Call) Return(_a0 error) *Manager_AddReadyzCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_AddReadyzCheck_Call) RunAndReturn(run func(string, healthz.Checker) error) *Manager_AddReadyzCheck_Call {
	_c.Call.Return(run)
	return _c
}

// Elected provides a mock function with given fields:
func (_m *Manager) Elected() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Manager_Elected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Elected'
type Manager_Elected_Call struct {
	*mock.Call
}

// Elected is a helper method to define mock.On call
func (_e *Manager_Expecter) Elected() *Manager_Elected_Call {
	return &Manager_Elected_Call{Call: _e.mock.On("Elected")}
}

func (_c *Manager_Elected_Call) Run(run func()) *Manager_Elected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_Elected_Call) Return(_a0 <-chan struct{}) *Manager_Elected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_Elected_Call) RunAndReturn(run func() <-chan struct{}) *Manager_Elected_Call {
	_c.Call.Return(run)
	return _c
}

// GetAPIReader provides a mock function with given fields:
func (_m *Manager) GetAPIReader() client.Reader {
	ret := _m.Called()

	var r0 client.Reader
	if rf, ok := ret.Get(0).(func() client.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Reader)
		}
	}

	return r0
}

// Manager_GetAPIReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAPIReader'
type Manager_GetAPIReader_Call struct {
	*mock.Call
}

// GetAPIReader is a helper method to define mock.On call
func (_e *Manager_Expecter) GetAPIReader() *Manager_GetAPIReader_Call {
	return &Manager_GetAPIReader_Call{Call: _e.mock.On("GetAPIReader")}
}

func (_c *Manager_GetAPIReader_Call) Run(run func()) *Manager_GetAPIReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetAPIReader_Call) Return(_a0 client.Reader) *Manager_GetAPIReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetAPIReader_Call) RunAndReturn(run func() client.Reader) *Manager_GetAPIReader_Call {
	_c.Call.Return(run)
	return _c
}

// GetCache provides a mock function with given fields:
func (_m *Manager) GetCache() cache.Cache {
	ret := _m.Called()

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func() cache.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}

// Manager_GetCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCache'
type Manager_GetCache_Call struct {
	*mock.Call
}

// GetCache is a helper method to define mock.On call
func (_e *Manager_Expecter) GetCache() *Manager_GetCache_Call {
	return &Manager_GetCache_Call{Call: _e.mock.On("GetCache")}
}

func (_c *Manager_GetCache_Call) Run(run func()) *Manager_GetCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetCache_Call) Return(_a0 cache.Cache) *Manager_GetCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetCache_Call) RunAndReturn(run func() cache.Cache) *Manager_GetCache_Call {
	_c.Call.Return(run)
	return _c
}

// GetClient provides a mock function with given fields:
func (_m *Manager) GetClient() client.Client {
	ret := _m.Called()

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}

// Manager_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type Manager_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
func (_e *Manager_Expecter) GetClient() *Manager_GetClient_Call {
	return &Manager_GetClient_Call{Call: _e.mock.On("GetClient")}
}

func (_c *Manager_GetClient_Call) Run(run func()) *Manager_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetClient_Call) Return(_a0 client.Client) *Manager_GetClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetClient_Call) RunAndReturn(run func() client.Client) *Manager_GetClient_Call {
	_c.Call.Return(run)
	return _c
}

// GetConfig provides a mock function with given fields:
func (_m *Manager) GetConfig() *rest.Config {
	ret := _m.Called()

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func() *rest.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	return r0
}

// Manager_GetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConfig'
type Manager_GetConfig_Call struct {
	*mock.Call
}

// GetConfig is a helper method to define mock.On call
func (_e *Manager_Expecter) GetConfig() *Manager_GetConfig_Call {
	return &Manager_GetConfig_Call{Call: _e.mock.On("GetConfig")}
}

func (_c *Manager_GetConfig_Call) Run(run func()) *Manager_GetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetConfig_Call) Return(_a0 *rest.Config) *Manager_GetConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetConfig_Call) RunAndReturn(run func() *rest.Config) *Manager_GetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetControllerOptions provides a mock function with given fields:
func (_m *Manager) GetControllerOptions() config.Controller {
	ret := _m.Called()

	var r0 config.Controller
	if rf, ok := ret.Get(0).(func() config.Controller); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Controller)
	}

	return r0
}

// Manager_GetControllerOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetControllerOptions'
type Manager_GetControllerOptions_Call struct {
	*mock.Call
}

// GetControllerOptions is a helper method to define mock.On call
func (_e *Manager_Expecter) GetControllerOptions() *Manager_GetControllerOptions_Call {
	return &Manager_GetControllerOptions_Call{Call: _e.mock.On("GetControllerOptions")}
}

func (_c *Manager_GetControllerOptions_Call) Run(run func()) *Manager_GetControllerOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetControllerOptions_Call) Return(_a0 config.Controller) *Manager_GetControllerOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetControllerOptions_Call) RunAndReturn(run func() config.Controller) *Manager_GetControllerOptions_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventRecorderFor provides a mock function with given fields: name
func (_m *Manager) GetEventRecorderFor(name string) record.EventRecorder {
	ret := _m.Called(name)

	var r0 record.EventRecorder
	if rf, ok := ret.Get(0).(func(string) record.EventRecorder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(record.EventRecorder)
		}
	}

	return r0
}

// Manager_GetEventRecorderFor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventRecorderFor'
type Manager_GetEventRecorderFor_Call struct {
	*mock.Call
}

// GetEventRecorderFor is a helper method to define mock.On call
//   - name string
func (_e *Manager_Expecter) GetEventRecorderFor(name interface{}) *Manager_GetEventRecorderFor_Call {
	return &Manager_GetEventRecorderFor_Call{Call: _e.mock.On("GetEventRecorderFor", name)}
}

func (_c *Manager_GetEventRecorderFor_Call) Run(run func(name string)) *Manager_GetEventRecorderFor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Manager_GetEventRecorderFor_Call) Return(_a0 record.EventRecorder) *Manager_GetEventRecorderFor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetEventRecorderFor_Call) RunAndReturn(run func(string) record.EventRecorder) *Manager_GetEventRecorderFor_Call {
	_c.Call.Return(run)
	return _c
}

// GetFieldIndexer provides a mock function with given fields:
func (_m *Manager) GetFieldIndexer() client.FieldIndexer {
	ret := _m.Called()

	var r0 client.FieldIndexer
	if rf, ok := ret.Get(0).(func() client.FieldIndexer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.FieldIndexer)
		}
	}

	return r0
}

// Manager_GetFieldIndexer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFieldIndexer'
type Manager_GetFieldIndexer_Call struct {
	*mock.Call
}

// GetFieldIndexer is a helper method to define mock.On call
func (_e *Manager_Expecter) GetFieldIndexer() *Manager_GetFieldIndexer_Call {
	return &Manager_GetFieldIndexer_Call{Call: _e.mock.On("GetFieldIndexer")}
}

func (_c *Manager_GetFieldIndexer_Call) Run(run func()) *Manager_GetFieldIndexer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetFieldIndexer_Call) Return(_a0 client.FieldIndexer) *Manager_GetFieldIndexer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetFieldIndexer_Call) RunAndReturn(run func() client.FieldIndexer) *Manager_GetFieldIndexer_Call {
	_c.Call.Return(run)
	return _c
}

// GetHTTPClient provides a mock function with given fields:
func (_m *Manager) GetHTTPClient() *http.Client {
	ret := _m.Called()

	var r0 *http.Client
	if rf, ok := ret.Get(0).(func() *http.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Client)
		}
	}

	return r0
}

// Manager_GetHTTPClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHTTPClient'
type Manager_GetHTTPClient_Call struct {
	*mock.Call
}

// GetHTTPClient is a helper method to define mock.On call
func (_e *Manager_Expecter) GetHTTPClient() *Manager_GetHTTPClient_Call {
	return &Manager_GetHTTPClient_Call{Call: _e.mock.On("GetHTTPClient")}
}

func (_c *Manager_GetHTTPClient_Call) Run(run func()) *Manager_GetHTTPClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetHTTPClient_Call) Return(_a0 *http.Client) *Manager_GetHTTPClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetHTTPClient_Call) RunAndReturn(run func() *http.Client) *Manager_GetHTTPClient_Call {
	_c.Call.Return(run)
	return _c
}

// GetLogger provides a mock function with given fields:
func (_m *Manager) GetLogger() logr.Logger {
	ret := _m.Called()

	var r0 logr.Logger
	if rf, ok := ret.Get(0).(func() logr.Logger); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(logr.Logger)
	}

	return r0
}

// Manager_GetLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogger'
type Manager_GetLogger_Call struct {
	*mock.Call
}

// GetLogger is a helper method to define mock.On call
func (_e *Manager_Expecter) GetLogger() *Manager_GetLogger_Call {
	return &Manager_GetLogger_Call{Call: _e.mock.On("GetLogger")}
}

func (_c *Manager_GetLogger_Call) Run(run func()) *Manager_GetLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetLogger_Call) Return(_a0 logr.Logger) *Manager_GetLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetLogger_Call) RunAndReturn(run func() logr.Logger) *Manager_GetLogger_Call {
	_c.Call.Return(run)
	return _c
}

// GetRESTMapper provides a mock function with given fields:
func (_m *Manager) GetRESTMapper() meta.RESTMapper {
	ret := _m.Called()

	var r0 meta.RESTMapper
	if rf, ok := ret.Get(0).(func() meta.RESTMapper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meta.RESTMapper)
		}
	}

	return r0
}

// Manager_GetRESTMapper_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRESTMapper'
type Manager_GetRESTMapper_Call struct {
	*mock.Call
}

// GetRESTMapper is a helper method to define mock.On call
func (_e *Manager_Expecter) GetRESTMapper() *Manager_GetRESTMapper_Call {
	return &Manager_GetRESTMapper_Call{Call: _e.mock.On("GetRESTMapper")}
}

func (_c *Manager_GetRESTMapper_Call) Run(run func()) *Manager_GetRESTMapper_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetRESTMapper_Call) Return(_a0 meta.RESTMapper) *Manager_GetRESTMapper_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetRESTMapper_Call) RunAndReturn(run func() meta.RESTMapper) *Manager_GetRESTMapper_Call {
	_c.Call.Return(run)
	return _c
}

// GetScheme provides a mock function with given fields:
func (_m *Manager) GetScheme() *runtime.Scheme {
	ret := _m.Called()

	var r0 *runtime.Scheme
	if rf, ok := ret.Get(0).(func() *runtime.Scheme); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.Scheme)
		}
	}

	return r0
}

// Manager_GetScheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetScheme'
type Manager_GetScheme_Call struct {
	*mock.Call
}

// GetScheme is a helper method to define mock.On call
func (_e *Manager_Expecter) GetScheme() *Manager_GetScheme_Call {
	return &Manager_GetScheme_Call{Call: _e.mock.On("GetScheme")}
}

func (_c *Manager_GetScheme_Call) Run(run func()) *Manager_GetScheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetScheme_Call) Return(_a0 *runtime.Scheme) *Manager_GetScheme_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetScheme_Call) RunAndReturn(run func() *runtime.Scheme) *Manager_GetScheme_Call {
	_c.Call.Return(run)
	return _c
}

// GetWebhookServer provides a mock function with given fields:
func (_m *Manager) GetWebhookServer() webhook.Server {
	ret := _m.Called()

	var r0 webhook.Server
	if rf, ok := ret.Get(0).(func() webhook.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(webhook.Server)
		}
	}

	return r0
}

// Manager_GetWebhookServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWebhookServer'
type Manager_GetWebhookServer_Call struct {
	*mock.Call
}

// GetWebhookServer is a helper method to define mock.On call
func (_e *Manager_Expecter) GetWebhookServer() *Manager_GetWebhookServer_Call {
	return &Manager_GetWebhookServer_Call{Call: _e.mock.On("GetWebhookServer")}
}

func (_c *Manager_GetWebhookServer_Call) Run(run func()) *Manager_GetWebhookServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Manager_GetWebhookServer_Call) Return(_a0 webhook.Server) *Manager_GetWebhookServer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GetWebhookServer_Call) RunAndReturn(run func() webhook.Server) *Manager_GetWebhookServer_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Manager) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Manager_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Manager_Expecter) Start(ctx interface{}) *Manager_Start_Call {
	return &Manager_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Manager_Start_Call) Run(run func(ctx context.Context)) *Manager_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Manager_Start_Call) Return(_a0 error) *Manager_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_Start_Call) RunAndReturn(run func(context.Context) error) *Manager_Start_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
