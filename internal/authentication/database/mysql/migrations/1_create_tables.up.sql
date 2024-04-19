CREATE TABLE Roles(
    RoleID varchar(255) PRIMARY KEY,
    RoleName varchar(255) NOT NULL,
    RoleDesc varchar(255),
    CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UpdatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    DeleteAt timestamp
);

CREATE TABLE Users(
    UserID varchar(255) PRIMARY KEY,
    FirstName varchar(255) NOT NULL,
    LastName varchar(255) NOT NULL,
    Email varchar(255) NOT NULL,
    IsVerified boolean DEFAULT FALSE,
    PhoneNo varchar(255),
    Password varchar(255) NOT NULL,
    Role varchar(255),
    Blocked boolean DEFAULT FALSE,
    CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UpdatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    DeleteAt timestamp,
    FOREIGN KEY (Role) REFERENCES Roles(RoleID)
);

CREATE TABLE Resets(
    ResetID varchar(255) PRIMARY KEY,
    User varchar(255),
    ResetToken varchar(255) NOT NULL,
    ExpiresAt timestamp NOT NULL,
    IsUsed boolean DEFAULT FALSE,
    CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UpdatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    DeleteAt timestamp,
    FOREIGN KEY (User) REFERENCES Users(UserID)
);

CREATE TABLE Verifications(
    VerificationId varchar(255) PRIMARY KEY,
    User varchar(255),
    VerifyToken varchar(255) NOT NULL,
    ExpiresAt timestamp NOT NULL,
    IsUsed boolean DEFAULT FALSE,
    CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UpdatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    DeleteAt timestamp,
    FOREIGN KEY (User) REFERENCES Users(UserID)
);