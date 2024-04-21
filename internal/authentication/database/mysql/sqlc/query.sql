-- name: CreateAccount :execresult
INSERT INTO Users (UserID, FirstName, LastName, Email, PhoneNo, Password, Role)
VALUES (?,?,?,?,?,?,?);

-- name: GetPassword :one
SELECT Password from Users
WHERE Email = ?;

-- name: GetRole :one
SELECT RoleID from Roles
WHERE RoleName = ?;
