package dotenv_test

import (
	"errors"
	"testing"

	"github.com/soumayg9673/dotenv"
)

func TestAddRqdKey(t *testing.T) {
	testCases := []struct {
		Name  string
		Key   string
		Value bool
		Err   error
	}{
		{
			Name:  "key with value required",
			Key:   "PACKAGE",
			Value: true,
			Err:   nil,
		},
		{
			Name:  "key without value",
			Key:   "PACKAGE",
			Value: false,
			Err:   nil,
		},
		{
			Name:  "empty key with value required",
			Key:   "",
			Value: true,
			Err:   dotenv.ErrEmptyKey,
		},
		{
			Name:  "empty key with no value required",
			Key:   "",
			Value: false,
			Err:   dotenv.ErrEmptyKey,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if err := dotenv.AddRqdKey(tc.Key, tc.Value); err != nil {
				if !errors.Is(tc.Err, err) { // failed
					t.Errorf("expected error %v but got %v", tc.Err, err)
				} else { // pass
					return
				}
			}

			m := dotenv.GetAllRqd()
			if _, ok := m[tc.Key]; !ok {
				t.Error("key should be listed on required list")
			}

			// clean up required env keys
			dotenv.DeleteRqdKey(tc.Key)
		})
	}
}

func TestAddMulRqdKeys(t *testing.T) {
	testCases := []struct {
		Name   string
		MapRqd map[string]bool
		Err    error
	}{
		{
			Name: "map of keys with value required",
			MapRqd: map[string]bool{
				"PACKAGE": true,
				"PUBLIC":  true,
			},
			Err: nil,
		},
		{
			Name: "map of keys with no value required",
			MapRqd: map[string]bool{
				"PACKAGE": false,
				"PUBLIC":  false,
			},
			Err: nil,
		},
		{
			Name: "map of empty key with value required",
			MapRqd: map[string]bool{
				"": true,
			},
			Err: dotenv.ErrEmptyKey,
		},
		{
			Name: "map of empty key with no value required",
			MapRqd: map[string]bool{
				"": false,
			},
			Err: dotenv.ErrEmptyKey,
		},
		{
			Name: "map of empty/on-empty keys with value, no-value required",
			MapRqd: map[string]bool{
				"":        false,
				"PACKAGE": true,
			},
			Err: dotenv.ErrEmptyKey,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if err := dotenv.AddMulRqdKeys(tc.MapRqd); err != nil {
				if !errors.Is(tc.Err, err) { // failed
					t.Errorf("expected error %v but got %v", tc.Err, err)
				} else { // pass
					return
				}
			}

			m := dotenv.GetAllRqd()

			for k := range tc.MapRqd {
				if _, ok := m[k]; !ok {
					t.Error("key should be listed on required list")
				}
				// clean up required env keys
				dotenv.DeleteRqdKey(k)
			}
		})
	}
}
