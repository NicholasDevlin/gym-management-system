import React, { useEffect, useState } from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";

function Home() {
  const alert = useAlert();
  const [totalMember, setTotalMember] = useState("Loading...")
  const [activeMember, setActiveMember] = useState("Loading...")
  const [totalTransaction, setTotalTransaction] = useState("Loading...")
  const [todayExpired, setTodayExpired] = useState([])

  useEffect(() => {
    getTransactionThisMonth();
    getTotalMember();
    getActiveMember();
    getExpireMemberToday();
  }, []);

  async function getTransactionThisMonth() {
    try {
      const response = await fetch(`${API_URLS.TRANSACTION}-this-month`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setTotalTransaction(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  async function getTotalMember() {
    try {
      const queryParams = new URLSearchParams({ role: "user" });
      const response = await fetch(`${API_URLS.USER}-total?${queryParams}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setTotalMember(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  async function getActiveMember() {
    try {
      const queryParams = new URLSearchParams({ role: "user", active: true });
      const response = await fetch(`${API_URLS.USER}-total?${queryParams}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setActiveMember(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  async function getExpireMemberToday() {
    try {
      const queryParams = new URLSearchParams({ role: "user", lastDayActive: true });
      const response = await fetch(`${API_URLS.USER}?${queryParams}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setTodayExpired(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  return (
    <Layout>
      <div className='container-fluid h-75'>
        <div className='d-flex justify-content-between'>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='w-100 text-center'><h3>{totalMember}</h3></div>
              <div className='text-center'><h4>Total Member</h4></div>
            </div>
          </div>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='w-100 text-center'><h3>{activeMember}</h3></div>
              <div className='text-center'><h4>Active Member</h4></div>
            </div>
          </div>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='w-100 text-center'><h3>{totalTransaction}</h3></div>
              <div className='text-center'><h4>Total Transaction this month</h4></div>
            </div>
          </div>
        </div>
        <div className='m-3 row'>
          <div className={`${Styles.memberList} col-sm-12 col-md-12 col-lg-6 card p-3`}>
            <h4 className='w-100'>
              List of members whose membership expires today:
            </h4>
            <div className='card-body overflow-auto h-75'>
              <ul>
                {todayExpired.map((value) => {
                  return (<li>
                    <h5>- {value.name} ({value.phoneNumber})</h5>
                  </li>)
                })}
              </ul>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;