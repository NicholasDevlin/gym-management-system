import React, { useEffect, useState } from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";
import { Badge, Calendar } from 'antd';
import Card from '../../../components/general/card/Card.jsx';
import { useUserData } from '../../../utils/jwt/UserData.jsx';

function Home() {
  const alert = useAlert();
  const [absensi, setAbsensi] = useState();
  const [transaction, setTransaction] = useState();
  const [user, setUser] = useState();
  const [activeMember, setActiveMember] = useState([]);
  const { userData } = useUserData();

  useEffect(() => {
    getAbsensi();
    getUser();
    getMemberTransaction();
    getActiveMember();
  }, [])

  async function getAbsensi() {
    try {
      const response = await fetch(`${API_URLS.ABSENSI}/my-absensi`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setAbsensi(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  async function getUser() {
    try {
      const response = await fetch(`${API_URLS.USER}/${userData.uuid}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setUser(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  async function getActiveMember() {
    try {
      const response = await fetch(`${API_URLS.USER}`, {
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
      alert.error(`${error}`);
    }
  }

  function getTransactionByUserUUID(data, userUUID) {
    for (const transactionDetail of data.transactionDetail) {
      for (const transactionMemberDetail of transactionDetail.transactionMemberDetail) {
        if (transactionMemberDetail.userUUID === userUUID) {
          return transactionDetail;
        }
      }
    }
    return null;
  }

  async function getMemberTransaction() {
    try {
      const queryParams = new URLSearchParams({ memberUUID: userData.uuid, isComplete: true });
      const response = await fetch(`${API_URLS.TRANSACTION}-member?${queryParams.toString()}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });
      const responseData = await response.json();
      if (responseData.success) {
        setTransaction(responseData.data);
        CardBody();
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  const getListData = (value) => {
    let listData;
    if (!absensi) {
      return [];
    }
    const date = new Date(value);
    const foundObject = absensi.find(obj => {
      const objDate = new Date(obj.date);
      return objDate.getFullYear() === date.getFullYear() &&
        objDate.getMonth() === date.getMonth() &&
        objDate.getDate() === date.getDate();
    });
    if (foundObject) {
      listData = [
        {
          type: 'success',
          content: '',
        },
      ];
    }
    return listData || [];
  };

  const dateCellRender = (value) => {
    const listData = getListData(value);
    return (
      <ul className="events">
        {listData.map((item) => (
          <li key={item.content}>
            <Badge status={item.type} text={item.content} />
          </li>
        ))}
      </ul>
    );
  };
  const cellRender = (current, info) => {
    if (info.type === 'date') return dateCellRender(current);
    if (info.type === 'month') return null;
    return info.originNode;
  };

  const CardTitle = (<>
    <h5 className='ps-2 pt-2'>{`Hi welcome back, ${user && user.name}`}</h5>
  </>);

  const CardBody = () => {
    let memberTransaction = null;
    if (transaction) {
      memberTransaction = getTransactionByUserUUID(transaction, userData.uuid);
    }
    let options = { day: 'numeric', month: 'long', year: 'numeric' };
    return (<>
      <div className='d-flex mt-2 justify-content-between'>
        <div>
          <h6>
            Active Membership Plan
          </h6>
          <h4>
            {memberTransaction != null ? memberTransaction.membershipPlan.name : "-"}
          </h4>
        </div>
        <div>
          <h6 className='text-end'>
            Active until
          </h6>
          <h4 className='text-end'>
            {user ? new Date(user.subscriptionDueDate).toLocaleDateString("id-ID", options) : "-"}
          </h4>
        </div>
      </div>
    </>);
  }

  return (
    <Layout>
      <div className='container-fluid h-75'>
        <div className={Styles.container}>
          <Card title={CardTitle} body={CardBody()} />
        </div>
        <div className={Styles.container}>
          <div className='w-100 row mx-3'>
            <div className={`col-md-6 col-sm-12 h-100 card ${Styles.Card}`}>
              <Calendar
                fullscreen={false}
                cellRender={cellRender}
              />
            </div>
            <div className={`col-md-6 col-sm-12 card ${Styles.Card}`}>
              <h5 className='text-dark'>Other Active Member</h5>
              <div className='card-body overflow-auto h-75'>
                <ul>
                  {activeMember.map((value) => (
                    <li className='text-dark my-1'>- {value.name}</li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;