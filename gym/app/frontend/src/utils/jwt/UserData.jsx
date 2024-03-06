import React, { createContext, useContext, useEffect, useState } from 'react';
import { decodeToken } from './JwtUtils.js';

const UserDataContext = createContext();

export const useUserData = () => useContext(UserDataContext);

export const UserDataProvider = ({ children }) => {
  const [userData, setUserData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const authToken = localStorage.getItem('authToken');
        if (authToken) {
          const userData = await decodeToken(authToken); 
          setUserData(userData);
        }
      } catch (error) {
        console.error('Error fetching user data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <UserDataContext.Provider value={userData}>
      {children}
    </UserDataContext.Provider>
  );
};
