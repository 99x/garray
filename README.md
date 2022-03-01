# G-array

Garray is a library written in Go (Golang). Which have collection of functions
to do the array operations.

## Installation

To install G-array package

`go get -u github.com/99x/garray`

import the code

`import "github.com/99x/garray"`


## Usage


The following test data array is used for all the examples of this article.


#### Sample data array

```
[
  {
    "guid": "58deff39-7366-4edf-8a7f-482bcd9e9bb7",
    "isActive": true,
    "age": 24,
    "name": "Maryanne Case",
    "gender": "female"
  },
  {
    "guid": "5d65f406-6007-4695-a9eb-73425edc4fea",
    "isActive": false,
    "age": 51,
    "name": "Christensen Acosta",
    "gender": "male"
  },
  {
    "guid": "45cc3685-bf2f-48a3-93eb-7351325eef43",
    "isActive": true,
    "age": 50,
    "name": "Shelton Brooks",
    "gender": "male"
  },
  {
    "guid": "99055545-b8c9-4d16-9cef-b0cf0527a3e9",
    "isActive": false,
    "age": 79,
    "name": "Bass Riley",
    "gender": "male"
  },
  {
    "guid": "de082e93-24a4-4219-81b8-6013097848d9",
    "isActive": true,
    "age": 52,
    "name": "Eddie Vaughan",
    "gender": "female"
  },
  {
    "guid": "24f4ca25-9485-443e-bc20-2c48f977bcef",
    "isActive": false,
    "age": 50,
    "name": "Krista Duncan",
    "gender": "female"
  }
]
```

This is the struct for the data deserialize

```
type User struct {
   Guid     string `json:"guid"`
   IsActive bool   `json:"isActive"`
   Age      int    `json:"age"`
   Name     string `json:"name"`
   Gender   string `json:"gender"`
}
```


### Find

The ***Find()*** function returns the first element in the provided array that satisfies the provided testing function. If no values satisfy the testing function, ***nill*** is returned.

Example:

```
func main() {
   found := *garray.Find(users, func(user User) bool {
      return user.Age > 50
   })
   log.Println(found)
}
```

As the result of the return function, it will return the following result

```
{
  "guid": "5d65f406-6007-4695-a9eb-73425edc4fea",
  "isActive": false,
  "age": 51,
  "name": "Christensen Acosta",
  "gender": "male"
}
```

### FindIndex

The** *FindIndex()* **function returns the index of the first element in the provided array that satisfies the provided testing function. If no values satisfy the testing function, ***-1*** is returned.

Example 1:

```
func main() {

   findIndex := garray.FindIndex(users, func(user User) bool {
      return user.Age > 50
   })
   log.Println(findIndex) // result: 1
}
```

Example 2:

```
func main() {

   findIndex := garray.FindIndex(users, func(user User) bool {
      return user.Age > 100
   })
   log.Println(findIndex) // result: -1
}
```

### **Search**

The ***Search()*** function returns the array of elements  from  the provided array that satisfies the provided testing function. If no values satisfy the testing function, an ***empty array ***is returned.

Example:

```
func main() {

   log.Println(users)
   log.Printf("Count : %d", len(users)) // output: 6
   searchData := garray.Search(users, func(user User, _ int) bool {
      return user.Age > 50
   })
   log.Println(searchData)
   log.Printf("Count : %d", len(searchData)) // output: 3
}
```

### Remove

The ***Remove()*** function returns the given array excluding the element of the given index

Example

```
func main() {
```

```
   log.Printf("Count: %d", len(users)) // Count: 6
   log.Printf("1st user's name: %s", users[0].Name) // 1st user's name: Maryanne Case
   log.Printf("2nd user's name: %s", users[1].Name) // 2nd user's name: Christensen Acosta

   users = garray.Remove(users, 0)
   log.Println("First element removed")

   log.Printf("Count: %d", len(users)) //  Count: 5
   log.Printf("1st user's name: %s", users[0].Name) // 1st user's name: Christensen Acosta
}
```

### Splice

The ***Splice()*** function is the same as ***Remove()*** function. Additionally, we can pass delete count as the  3rd parameter.

```
func main() {

   log.Printf("Count: %d", len(users)) // Count: 6
   log.Printf("1st user's name: %s", users[0].Name) // 1st user's name: Maryanne Case
   log.Printf("5th user's name: %s", users[4].Name) // 5th user's name: Eddie Vaughan

   users = garray.Splice(users, 1, 3) 
   log.Printf("Count: %d", len(users)) //  Count: 3
   log.Printf("1st user's name: %s", users[0].Name) // 1st user's name: Maryanne Case
   log.Printf("2nd user's name: %s", users[1].Name) // 2nd user's name: Eddie Vaughan
}
```

### Filter

The ***Filter()*** function returns an array  excluding all elements that pass the test implemented by the provided function.

```
func main() {

   ReadUsersData()
   log.Printf("Count: %d", len(users)) // Count: 6

   users = garray.Filter(users, func(user User, _ int) bool {
      return user.Age > 50
   })
   log.Println("Array filtered if age > 50")
   log.Printf("Count: %d", len(users)) // Count: 3
}
```

### Map

The ***Map()*** function returns an array populated with the results of calling a provided function on every element in the calling array.

```
func main() {

   ReadUsersData()
   log.Println(users)
   names := *garray.Map(users, func(users User, _ int) string {
      return users.Name
   })

   log.Println(names) // [Maryanne Case Christensen Acosta Shelton Brooks Bass Riley Eddie Vaughan Krista Duncan]
}
```
