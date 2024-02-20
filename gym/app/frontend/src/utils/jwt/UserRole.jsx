// UserRoleProvider.js
import React, { createContext, useContext, useEffect, useState } from 'react';
import { decodeToken } from './JwtUtils.js';

const UserRoleContext = createContext();

export const useUserRole = () => useContext(UserRoleContext);

export const UserRoleProvider = ({ children }) => {
  const [userRole, setUserRole] = useState(null);

  useEffect(() => {
    const authToken = localStorage.getItem('authToken');

    if (authToken) {
      const decodedToken = decodeToken(authToken);

      if (decodedToken && decodedToken.role) {
        setUserRole(decodedToken.role);
      }
    }
  }, []);

  return (
    <UserRoleContext.Provider value={{ userRole, setUserRole }}>
      {children}
    </UserRoleContext.Provider>
  );
};
