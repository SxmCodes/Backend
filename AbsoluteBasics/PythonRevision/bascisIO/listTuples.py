# learning about list and tuples.
fruits = ["apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew"]

print(fruits)

# Basic functions 
print(fruits[0])

fruits.append("nothing")
print(fruits)

fruits.extend("nothing2")
print(fruits)

# There is a huge difference between append and extend, append mai if we are adding "orange" to the list, then the whole word "orange" will be regarded as 1 character. But in extend, each letter is considered individual character.

# This is to insert any new character at any particular index.
fruits.insert(3, "orange")
print(fruits)

# This is to remove any particular characters from the list.
fruits.remove("cherry")
print(fruits)

# This is the concept of immutable and mutable in python, list can be changed and are changed when we call any function. 
# original string is not changed but original list changes. 



# Now here comes the immutable data-type in python, called tuple. 

# this is a valid typle. 
example1 = (1, 2, 3, 4, 5)
a = type(example1)
print(a)
print(example1)

# But this ain't 
example2 = (1, 2, 3)
print(type(a)) # this will be considered as a integer, not tuple. To make it tuple, we will add a comma (,) at teh end of `1`.
print(example2)

# some fo the functions of tuples. 
list1 = (34,56,100,3778,190)
number = list1.count(34)

# contatenation
contatinatedlist = example1 + example2 + list1
print(contatinatedlist)
# Repetition - in tuples the values are not repeated, to do so, just do this
simpletuple = (1,2,3)
repeatedtuple = simpletuple *3
print(repeatedtuple) # The values got repeated. 

# Membership
my_tuple = (2,3,4,5)
print(2 in my_tuple) # true
print(10 in my_tuple) # true
print(3 in my_tuple) # false

# Length 




         
