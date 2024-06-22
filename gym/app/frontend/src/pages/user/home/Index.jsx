import React, { useEffect, useState } from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";
import { Badge, Calendar } from 'antd';
import Card from '../../../components/general/card/Card.jsx';

function Home() {
  const alert = useAlert();
  const [absensi, setAbsensi] = useState();

  useEffect(() => {
    getAbsensi();
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

  return (
    <Layout>
      <div className='container-fluid h-75'>
        <div className={Styles.container}>
          <Card title={"test"} body={"loremmmm "} />
        </div>
        <div className={Styles.container}>
          <div className='w-100 mx-3'>
            <div className={`w-50 h-100 ${Styles.Calendar}`}>
              <Calendar
                fullscreen={false}
                cellRender={cellRender}
              />
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;