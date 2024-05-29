import React, { createContext, useContext, useEffect, useState } from 'react';
import { decodeToken } from './JwtUtils.js';

const UserDataContext = createContext();

export const useUserData = () => useContext(UserDataContext);

export const UserDataProvider = ({ children }) => {
  const [userData, setUserData] = useState(null);
  const [key, setKey] = useState(0);

  const fetchData = async () => {
    try {
      const authToken = localStorage.getItem('authToken');
      if (authToken) {
        const decodedUserData = await decodeToken(authToken);
        setUserData(decodedUserData);
      }
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  useEffect(() => {
    setKey(prevKey => prevKey + 1);
  }, [userData]);

  const handleLogout = () => {
    setUserData(null);
    localStorage.removeItem('authToken');
  };

  return (
    <UserDataContext.Provider value={{ userData, fetchData, handleLogout }}>
      <div key={key} className='w-100 h-100'>{children}</div>
    </UserDataContext.Provider>
  );
};
