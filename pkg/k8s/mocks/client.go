// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"

	appsv1 "k8s.io/api/apps/v1"

	corev1 "k8s.io/api/core/v1"

	mock "github.com/stretchr/testify/mock"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *Client) Delete(_a0 context.Context, _a1 *unstructured.Unstructured) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Client_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *unstructured.Unstructured
func (_e *Client_Expecter) Delete(_a0 interface{}, _a1 interface{}) *Client_Delete_Call {
	return &Client_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *Client_Delete_Call) Run(run func(_a0 context.Context, _a1 *unstructured.Unstructured)) *Client_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*unstructured.Unstructured))
	})
	return _c
}

func (_c *Client_Delete_Call) Return(_a0 error) *Client_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Delete_Call) RunAndReturn(run func(context.Context, *unstructured.Unstructured) error) *Client_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePVCsWithLabel provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) DeletePVCsWithLabel(_a0 context.Context, _a1 string, _a2 string, _a3 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_DeletePVCsWithLabel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePVCsWithLabel'
type Client_DeletePVCsWithLabel_Call struct {
	*mock.Call
}

// DeletePVCsWithLabel is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
//   - _a3 string
func (_e *Client_Expecter) DeletePVCsWithLabel(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *Client_DeletePVCsWithLabel_Call {
	return &Client_DeletePVCsWithLabel_Call{Call: _e.mock.On("DeletePVCsWithLabel", _a0, _a1, _a2, _a3)}
}

func (_c *Client_DeletePVCsWithLabel_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string, _a3 string)) *Client_DeletePVCsWithLabel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *Client_DeletePVCsWithLabel_Call) Return(_a0 error) *Client_DeletePVCsWithLabel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_DeletePVCsWithLabel_Call) RunAndReturn(run func(context.Context, string, string, string) error) *Client_DeletePVCsWithLabel_Call {
	_c.Call.Return(run)
	return _c
}

// DestinationRuleCRDExists provides a mock function with given fields: _a0
func (_m *Client) DestinationRuleCRDExists(_a0 context.Context) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DestinationRuleCRDExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DestinationRuleCRDExists'
type Client_DestinationRuleCRDExists_Call struct {
	*mock.Call
}

// DestinationRuleCRDExists is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Client_Expecter) DestinationRuleCRDExists(_a0 interface{}) *Client_DestinationRuleCRDExists_Call {
	return &Client_DestinationRuleCRDExists_Call{Call: _e.mock.On("DestinationRuleCRDExists", _a0)}
}

func (_c *Client_DestinationRuleCRDExists_Call) Run(run func(_a0 context.Context)) *Client_DestinationRuleCRDExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_DestinationRuleCRDExists_Call) Return(_a0 bool, _a1 error) *Client_DestinationRuleCRDExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DestinationRuleCRDExists_Call) RunAndReturn(run func(context.Context) (bool, error)) *Client_DestinationRuleCRDExists_Call {
	_c.Call.Return(run)
	return _c
}

// GetCRD provides a mock function with given fields: _a0, _a1
func (_m *Client) GetCRD(_a0 context.Context, _a1 string) (*v1.CustomResourceDefinition, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.CustomResourceDefinition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v1.CustomResourceDefinition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1.CustomResourceDefinition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.CustomResourceDefinition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetCRD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCRD'
type Client_GetCRD_Call struct {
	*mock.Call
}

// GetCRD is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *Client_Expecter) GetCRD(_a0 interface{}, _a1 interface{}) *Client_GetCRD_Call {
	return &Client_GetCRD_Call{Call: _e.mock.On("GetCRD", _a0, _a1)}
}

func (_c *Client_GetCRD_Call) Run(run func(_a0 context.Context, _a1 string)) *Client_GetCRD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_GetCRD_Call) Return(_a0 *v1.CustomResourceDefinition, _a1 error) *Client_GetCRD_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetCRD_Call) RunAndReturn(run func(context.Context, string) (*v1.CustomResourceDefinition, error)) *Client_GetCRD_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecret provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) GetSecret(_a0 context.Context, _a1 string, _a2 string) (*corev1.Secret, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *corev1.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*corev1.Secret, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *corev1.Secret); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecret'
type Client_GetSecret_Call struct {
	*mock.Call
}

// GetSecret is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *Client_Expecter) GetSecret(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Client_GetSecret_Call {
	return &Client_GetSecret_Call{Call: _e.mock.On("GetSecret", _a0, _a1, _a2)}
}

func (_c *Client_GetSecret_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *Client_GetSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Client_GetSecret_Call) Return(_a0 *corev1.Secret, _a1 error) *Client_GetSecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetSecret_Call) RunAndReturn(run func(context.Context, string, string) (*corev1.Secret, error)) *Client_GetSecret_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatefulSet provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) GetStatefulSet(_a0 context.Context, _a1 string, _a2 string) (*appsv1.StatefulSet, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *appsv1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*appsv1.StatefulSet, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *appsv1.StatefulSet); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetStatefulSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatefulSet'
type Client_GetStatefulSet_Call struct {
	*mock.Call
}

// GetStatefulSet is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *Client_Expecter) GetStatefulSet(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Client_GetStatefulSet_Call {
	return &Client_GetStatefulSet_Call{Call: _e.mock.On("GetStatefulSet", _a0, _a1, _a2)}
}

func (_c *Client_GetStatefulSet_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *Client_GetStatefulSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Client_GetStatefulSet_Call) Return(_a0 *appsv1.StatefulSet, _a1 error) *Client_GetStatefulSet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetStatefulSet_Call) RunAndReturn(run func(context.Context, string, string) (*appsv1.StatefulSet, error)) *Client_GetStatefulSet_Call {
	_c.Call.Return(run)
	return _c
}

// PatchApply provides a mock function with given fields: _a0, _a1
func (_m *Client) PatchApply(_a0 context.Context, _a1 *unstructured.Unstructured) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_PatchApply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchApply'
type Client_PatchApply_Call struct {
	*mock.Call
}

// PatchApply is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *unstructured.Unstructured
func (_e *Client_Expecter) PatchApply(_a0 interface{}, _a1 interface{}) *Client_PatchApply_Call {
	return &Client_PatchApply_Call{Call: _e.mock.On("PatchApply", _a0, _a1)}
}

func (_c *Client_PatchApply_Call) Run(run func(_a0 context.Context, _a1 *unstructured.Unstructured)) *Client_PatchApply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*unstructured.Unstructured))
	})
	return _c
}

func (_c *Client_PatchApply_Call) Return(_a0 error) *Client_PatchApply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_PatchApply_Call) RunAndReturn(run func(context.Context, *unstructured.Unstructured) error) *Client_PatchApply_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
