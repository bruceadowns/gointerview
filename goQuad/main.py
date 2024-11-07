import math

a = float(input("What is a? "))
b = float(input("What is b? "))
c = float(input("What is c? "))
d = math.pow(b, 2) - 4 * a * c
x = (-b + math.sqrt(d)) / (2 * a)
print(x)
x = (-b - math.sqrt(d)) / (2 * a)
print(x)
