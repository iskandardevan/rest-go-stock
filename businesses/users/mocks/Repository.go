package mocks

import (
	"context" 
	"rest-go-stock/businesses/users"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) GetByID (id uint, ctx context.Context) (users.Domain, error){
	ret := _m.Called(ctx, id)
	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context)users.Domain); ok{
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context)error); ok{
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GetByEmail (email string, ctx context.Context) (users.Domain, error){
	ret := _m.Called(ctx, email)
	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string, context.Context)users.Domain); ok{
		r0 = rf(email, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context)error); ok{
		r1 = rf(email, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}


func (_m *Repository) DeleteUserByID (id uint, ctx context.Context ) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) RegisterUser (ctx context.Context, domain *users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
    ret := _m.Called(ctx)

    var r0 []users.Domain
    if rf, ok := ret.Get(0).(func(context.Context) []users.Domain); ok {
        r0 = rf(ctx)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).([]users.Domain)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context) error); ok {
        r1 = rf(ctx)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *Repository) UpdateUserByID (id uint, ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context, users.Domain) users.Domain); ok {
		r0 = rf(id, ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context, users.Domain) error); ok {
		r1 = rf(id, ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdatePasswordByID (id uint, ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context, users.Domain) users.Domain); ok {
		r0 = rf(id, ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context, users.Domain) error); ok {
		r1 = rf(id, ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetEmail(ctx context.Context, email string) (users.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) CheckUser(email string, ctx context.Context) (users.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
