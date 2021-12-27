SELECT users.id, users.username, parents.username as parentsusername FROM USER as users LEFT JOIN user as parents 
ON parents.id = users.parent;