message='Hi \U0001f600'
width=20
padding=(width-len(message))//2
print(' '*padding,end='')
print(message)
print('-'*width)