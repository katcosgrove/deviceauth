// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import migrate "github.com/mendersoftware/go-lib-micro/mongo/migrate"
import mock "github.com/stretchr/testify/mock"

// Migration is an autogenerated mock type for the Migration type
type Migration struct {
	mock.Mock
}

// Up provides a mock function with given fields: from
func (_m *Migration) Up(from migrate.Version) error {
	ret := _m.Called(from)

	var r0 error
	if rf, ok := ret.Get(0).(func(migrate.Version) error); ok {
		r0 = rf(from)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *Migration) Version() migrate.Version {
	ret := _m.Called()

	var r0 migrate.Version
	if rf, ok := ret.Get(0).(func() migrate.Version); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(migrate.Version)
	}

	return r0
}

var _ migrate.Migration = (*Migration)(nil)
