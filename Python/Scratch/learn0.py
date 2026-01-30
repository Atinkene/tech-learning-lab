fruits = ["apple", "banana", "cherry", "kiwi", "mango"]

fruits.append("orange")
fruits.pop(1)
fruits.extend(["grape", "pineapple"])

newlist = [x for x in fruits if "o" in x]
print(fruits)
print(newlist)

match fruits:
    case ["apple", "banana", *rest]:
        print("Starts with apple and banana")
    case ["kiwi", *rest]:
        print("Starts with kiwi")
    case _:
        print("No match")  

for fruit in fruits:
    if fruit.endswith("a"):
        print(fruit)