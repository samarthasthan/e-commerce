-- name: CreateAccount :execresult
INSERT INTO Users (UserID, FirstName, LastName, Email, PhoneNo, Password, Role)
VALUES (?,?,?,?,?,?,?);
