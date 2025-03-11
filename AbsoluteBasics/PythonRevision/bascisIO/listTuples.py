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