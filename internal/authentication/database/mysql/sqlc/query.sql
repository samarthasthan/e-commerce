-- name: CreateAccount :exec
INSERT INTO Users (UserID, FirstName, LastName, Email, PhoneNo, Password, RoleID)
VALUES (?,?,?,?,?,?,(SELECT RoleID FROM Roles WHERE RoleName = ?));


-- name: CreateVerification :exec
INSERT INTO Verifications (VerificationId, UserID, OTP, ExpiresAt)
VALUES (?,?,?,?);

-- name: GetOTP :one
SELECT OTP, UserID, ExpiresAt FROM Verifications WHERE VerificationId = ?;

-- name: VerifyAccount :exec
UPDATE Users SET IsVerified = 1 WHERE UserID = ?;

-- name: DeleteVerification :exec
DELETE FROM Verifications WHERE VerificationId = ?;