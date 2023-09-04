package identity_test

import (
	"errors"
	dbtesting "gopher/infra/db/db_testing"
	"gopher/services/identity"
	"testing"
)

func TestFindByUsername(t *testing.T) {
	testCases := []dbtesting.Expectations{
		{
			TestName:       "store.FindByUsername should return a user",
			ExpectedError:  nil,
			ExpectedResult: &identity.User{},
		},
		{
			TestName:       "store.FindByUsername should return error",
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := identity.NewStore(fakeDb)
			_, err := service.FindByUsername("some username")
			if tc.ExpectedError == nil && err != nil {
				t.Fatalf("Expected to have no error but got %q", err)
			} else if tc.ExpectedError != nil && err == nil {
				t.Fatalf("Unexpected error: %q", err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []dbtesting.Expectations{
		{
			TestName:       "store.Create should return a user",
			ExpectedError:  nil,
			ExpectedResult: &identity.User{},
		},
		{
			TestName:       "store.Create should return error",
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := identity.NewStore(fakeDb)
			_, err := service.Create(&identity.CreateIdentityDTO{})
			if tc.ExpectedError == nil && err != nil {
				t.Fatalf("Expected to have no error but got %q", err)
			} else if tc.ExpectedError != nil && err == nil {
				t.Fatalf("Unexpected error: %q", err)
			}
		})
	}
}
