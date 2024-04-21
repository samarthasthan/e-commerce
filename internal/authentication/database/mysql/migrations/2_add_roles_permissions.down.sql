-- Delete the admin user
DELETE FROM Users
WHERE Email = 'admin@frubay.com';

-- Delete role-permission mappings for each role
-- For the admin role
DELETE FROM RolePermissions
WHERE RoleID = (SELECT RoleID FROM Roles WHERE RoleName = 'admin');

-- For the seller role
DELETE FROM RolePermissions
WHERE RoleID = (SELECT RoleID FROM Roles WHERE RoleName = 'seller');

-- For the user role
DELETE FROM RolePermissions
WHERE RoleID = (SELECT RoleID FROM Roles WHERE RoleName = 'user');

-- Delete permissions
DELETE FROM Permissions
WHERE PermissionName IN ('manage_users', 'view_products', 'add_products', 'edit_products', 'delete_products');

-- Delete roles
DELETE FROM Roles
WHERE RoleName IN ('admin', 'seller', 'user');
