-- Insert roles
INSERT INTO Roles (RoleID, RoleName, RoleDesc)
VALUES
    (UUID(), 'admin', 'Administrator with full permissions'),
    (UUID(), 'seller', 'Seller with permissions to manage products'),
    (UUID(), 'user', 'Regular user with basic access');

-- Insert permissions
INSERT INTO Permissions (PermissionID, PermissionName, PermissionDesc)
VALUES
    (UUID(), 'manage_users', 'Manage users in the system'),
    (UUID(), 'view_products', 'View available products'),
    (UUID(), 'add_products', 'Add new products to the catalog'),
    (UUID(), 'edit_products', 'Edit existing products in the catalog'),
    (UUID(), 'delete_products', 'Delete products from the catalog');

-- Assign permissions to roles
-- Assign permissions to admin role
INSERT INTO RolePermissions (RoleID, PermissionID)
SELECT r.RoleID, p.PermissionID
FROM Roles r, Permissions p
WHERE r.RoleName = 'admin'
  AND p.PermissionName IN ('manage_users', 'view_products', 'add_products', 'edit_products', 'delete_products');

-- Assign permissions to seller role
INSERT INTO RolePermissions (RoleID, PermissionID)
SELECT r.RoleID, p.PermissionID
FROM Roles r, Permissions p
WHERE r.RoleName = 'seller'
  AND p.PermissionName IN ('view_products', 'add_products', 'edit_products', 'delete_products');

-- Assign permissions to user role
INSERT INTO RolePermissions (RoleID, PermissionID)
SELECT r.RoleID, p.PermissionID
FROM Roles r, Permissions p
WHERE r.RoleName = 'user'
  AND p.PermissionName = 'view_products';

-- Create an admin user and assign the 'admin' role
INSERT INTO Users (UserID, FirstName, LastName, Email, Password, Role, IsVerified, PhoneNo)
VALUES
    (UUID(), 'Admin', 'User', 'admin@frubay.com', '$2a$04$OQQ8anXh4kBcibADYjiL6OPgfUWY5MowDznH0eYkdq0rLgFC6uM36', -- Replace with hashed password
     (SELECT RoleID FROM Roles WHERE RoleName = 'admin'),
     TRUE, '+91 9557037766');
