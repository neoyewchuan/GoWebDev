You should never store a password without encrypting it.

We will use ````bcrypt``` to encrypt our password.

# step 1

You will need to go get this package:
```
go get golang.org/x/crypto/bcrypt
```

# step 2
change your user struct's password field from string to the data type []byte.

# step 3
Use bcrypt to encrypt your password before storing it 