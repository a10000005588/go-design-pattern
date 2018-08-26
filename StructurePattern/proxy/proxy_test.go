package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(133424)
	for i := 0; i < 10000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    UserList{},
	}
	// someDatabase[0] someDatabase[1] someDatabase[2]
	knownIDs := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		// proxy will check whether this request for user has been searched or not.
		// If it is the first time to search one of the user, the info come from the DB
		// if not , the info came from the cache.

		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}

		if proxy.DidDidLastSearchUsedCached {
			t.Error("No user can be returned from an empty cache")
		}
	})
	// search knownIDs[0] again.
	// this time the user data must came from "Cache" rather than "DB"
	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object htat is stored on it")
		}
		// this time the last search must came from the cache.
		if !proxy.DidDidLastSearchUsedCached {
			t.Error("The user should have been returned from the cache")
		}
	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}
	user2, err := proxy.FindUser(knownIDs[1])
	if err != nil {
		t.Error("The user wasn't stored on the proxy cache yet")
	}
	user3, err := proxy.FindUser(knownIDs[2])
	if err != nil {
		t.Error("The user wasn't stored on the proxy cache yet")
	}

	// because oldest data in cache will be removed.
	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("User that should be gone that was found")
		}
	}
	// because the size of cache was defined be 2.
	if len(proxy.StackCache) != 2 {
		t.Error("After inserting 3 users the cahce should not grow more than to two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("A non expected user was found on the cache")
		}
	}
}
