package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct 
	ID int32
}

type UserList []User

// Proxy will cover the object you want, such as DB and Cached.
type UserListProxy struct {
	SomeDatabase               UserList
	StackCache                 UserList
	StackCapacity              int
	DidDidLastSearchUsedCached bool
}

// if struct type is array, you should use *(struct array) to get the instance.
func (t *UserList) FindUser(id int32) (User, error) {

	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}

	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

// Using Proxy to deal with data retrieving from cache or DB.
// 先找cache 沒找到資料再去找DB 找完DB 再把結果存到 cache內
func (u *UserListProxy) FindUser(id int32) (User, error) {

	// StackCache is also UserList , the FindUser method in UserList is just return data in User[]
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		// if last search used cacahse, then mark 'true'
		u.DidDidLastSearchUsedCached = true
		return user, nil
	}
	// if StackCache return err , it means no data in cache, then do search DB
	user, err = u.SomeDatabase.FindUser(id)

	if err != nil {
		return User{}, err
	}
	// if there are someone retriving data from db, then store the data in cache in order to handle the same request.
	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.DidDidLastSearchUsedCached = false
	return user, nil

}

// Proxy can call addUserToStack
func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

// UserList type can addUser
func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
