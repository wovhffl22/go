class Stack:
    # Initialize an object and its attribute, list.
    def __init__(self):
        self.item = []

    # push stores data x at the top of stack.
    def push(self, x):
        self.item.append(x)

    # top returns the latest data.
    def top(self):
        if len(self.item) != 0:
            return self.item[-1]
        else:
            return "Can't define it."

    # pop deletes the latest data.
    def pop(self):
        self.item.remove(self.item[-1])

    # empty returns either 1(true) or 0(false)
    def empty1(self):
        # Check if array is empty
        # bool(1) = true, bool(0) = false
        if len(self.item) == 0:
            return 1
        else:
            return 0

    def empty2(self):
        # int(true) = 1, int(false) = 0
        if len(self.item) == 0:
            return True
        else:
            return False

# Create an object
myStack = Stack()

# Approach to (list item) and add 1,2,3,4,5.
for i in range(1, 6):
    myStack.item.append(i)
    print("items added: ", myStack.item, "\n")

    # Delete the latest element.
    myStack.item.pop()
    print("the latest item deleted: ", myStack.item, "\n")

    # Return the latest element
    print("the latest item for now: ", myStack.top(), "\n")

    # Check if the list is empty - empty1
    isEmpty1 = myStack.empty1()
    print("Is my object empty now? ", myStack.item, isEmpty1, bool(isEmpty1), "\n")

    # Delete all the elements
    for i in range(len(myStack.item)):
        myStack.item.pop()
    print("deleting all the element...")

    # Check if the list is empty - empty2
    isEmpty2 = myStack.empty2()
    print("Is my object empty now? ", myStack.item, isEmpty2, int(isEmpty2), "\n")

    # Check what method and attribute the object holds - check the followings: item, push, pop, top, empty1, 2
    myMethod = dir(myStack)
    print(myMethod)
