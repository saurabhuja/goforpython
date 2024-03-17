
#average of all numbers between 1 to 100 which is divided by either 3 or 5.
count,total=0,0
for i in range(1,101):
    if i%3==0 or i%5==0:
        count=count+1
        total=total+i
average=total/count
print(average)
