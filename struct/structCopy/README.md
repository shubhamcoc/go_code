# Copy of struct in golang
Difference between deep and shallow copy of struct.

# Output of the structCopy code
```
  In case of Shallow Copy of complex struct like person

  Before changing element in object b
  a is: {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[line1 line2] Number:{Number1:123456789}} SomeDetail:map[key1:value1]}
  b is: {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[line1 line2] Number:{Number1:123456789}} SomeDetail:map[key1:value1]}

  After changing element (i.e.) line1 = test4 in object b
  a is {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[test4 line2] Number:{Number1:123456789}} SomeDetail:map[key1:value1]}
  b is {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[test4 line2] Number:{Number1:123456789}} SomeDetail:map[key1:value1]}

  In case of Deep Copy

  After changing element (i.e.) test4 = test5, key1 = key2 and value1 = value2 in object c
  a is {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[test4 line2] Number:{Number1:123456789}} SomeDetail:map[key1:value1]}
  c is {Name:test Age:24 Alias:[alias1 alias2] Address:{Line1:[test5 line2] Number:{Number1:123456789}} SomeDetail:map[key2:value2]}

  In case of Shallow Copy of struct like person1

  Before changing element in object per2
  per1 is: {name:person1 age:26 address:{line1:line64 city:somecity}}
  per2 is: {name:person1 age:26 address:{line1:line64 city:somecity}}

  After changing element (i.e.) line64 = test54 in object per2
  per1 is {name:person1 age:26 address:{line1:line64 city:somecity}}
  per2 is {name:person1 age:26 address:{line1:line54 city:somecity}}

  Before changing element in object per3
  per1 is: {name:person1 age:26 address:{line1:line64 city:somecity}}
  per2 is: {name:person1 age:26 address:{line1:line54 city:somecity}}
  per3 is: {name:person1 age:26 address:{line1:line64 city:somecity}}

  After changing element (i.e.) line64 = test44 in object per3
  per1 is {name:person1 age:26 address:{line1:line64 city:somecity}}
  per2 is {name:person1 age:26 address:{line1:line54 city:somecity}}
  per3 is {name:person1 age:26 address:{line1:line44 city:somecity}}
```

play.golang.org link :- https://play.golang.org/p/jkCJD3w7942
