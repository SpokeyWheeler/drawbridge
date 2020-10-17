# Permissions

There are 3 ordering options.

1 allows you to set wide-ranging permissions.

2 allows you to revoke some specific permissions.

3 allows you to grant narrow permissions.

For example: 

1. GRANT ALL ON hrdata.* TO fred;
2. REVOKE ALL on hrdata.salaries FROM fred;
3. GRANT SELECT on hrdata.salaries(id, paygrade, ssn) TO fred;