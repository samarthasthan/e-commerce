-- Insert roles
INSERT INTO Roles (RoleID, RoleName, RoleDesc)
VALUES
    ('admin-role-id', 'admin', 'Administrator with full permissions'),
    ('seller-role-id', 'seller', 'Seller with permissions to manage products'),
    ('user-role-id', 'user', 'Regular user with basic access');

-- Insert permissions
INSERT INTO Permissions (PermissionID, PermissionName, PermissionDesc, ActionType)
VALUES
    ('manage_users-permission-id', 'manage_users', 'Manage users in the system', 'write'),
    ('view_products-permission-id', 'view_products', 'View available products', 'read'),
    ('add_products-permission-id', 'add_products', 'Add new products to the catalog', 'write'),
    ('edit_products-permission-id', 'edit_products', 'Edit existing products in the catalog', 'write'),
    ('delete_products-permission-id', 'delete_products', 'Delete products from the catalog', 'delete');

-- Assign permissions to roles
-- Assign permissions to admin role
INSERT INTO RolePermissions (RoleID, PermissionID)
VALUES
    ('admin-role-id', 'manage_users-permission-id'),
    ('admin-role-id', 'view_products-permission-id'),
    ('admin-role-id', 'add_products-permission-id'),
    ('admin-role-id', 'edit_products-permission-id'),
    ('admin-role-id', 'delete_products-permission-id');

-- Assign permissions to seller role
INSERT INTO RolePermissions (RoleID, PermissionID)
VALUES
    ('seller-role-id', 'view_products-permission-id'),
    ('seller-role-id', 'add_products-permission-id'),
    ('seller-role-id', 'edit_products-permission-id'),
    ('seller-role-id', 'delete_products-permission-id');

-- Assign permissions to user role
INSERT INTO RolePermissions (RoleID, PermissionID)
VALUES
    ('user-role-id', 'view_products-permission-id');

-- Create an admin user and assign the 'admin' role
INSERT INTO Users (UserID, FirstName, LastName, Email, Password, RoleID, IsVerified, PhoneNo)
VALUES
    ('admin-user-id', 'Admin', 'User', 'admin@frubay.com', '$2a$04$OQQ8anXh4kBcibADYjiL6OPgfUWY5MowDznH0eYkdq0rLgFC6uM36', -- Replace with hashed password
     'admin-role-id',
     TRUE, '+91 9557030000');
