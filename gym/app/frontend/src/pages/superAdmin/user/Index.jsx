import React, { useState, useEffect, useCallback } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './User.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";
import { Table } from "antd";
import column, { activeOptions } from "./Config.jsx";
import TextField from "../../../components/general/input/inputTextField/TextField.jsx";
import Select from "../../../components/general/input/select/Select.jsx";
import { GetRole } from "../../../controller/UserController.js";

function Transaction() {
  const alert = useAlert()
  const [userData, setUserData] = useState([]);
  const [roleOptions, setRoleOptions] = useState([]);
  const [filter, setFilter] = useState({
    active: null,
    name: ""
  });

  useEffect(() => {
    getUsers();
    fetchUsersOptions();
  }, []);

  const fetchUsersOptions = useCallback(async () => {
    const result = await GetRole();
    if (result) {
      setRoleOptions(result.map((data) => ({
        value: data.id,
        label: data.role,
      })));
    }
  }, []);

  async function getUsers() {
    try {
      const queryParams = new URLSearchParams();
      Object.keys(filter).forEach(key => {
        if (filter[key] !== null && filter[key] !== "") {
          queryParams.append(key, filter[key]);
        }
      });
      const response = await fetch(`${API_URLS.USER}?${queryParams.toString()}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setUserData(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  const handleInputChange = (e) => {
    const { id, value } = e.target || e || {};
    setFilter((prevData) => ({
      ...prevData,
      [id]: value,
    }));
  };

  async function deleteUser(uuid) {
    try {
      const response = await fetch(`${API_URLS.USER}/${uuid}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        const newData = userData.filter((item) => item.uuid !== uuid);
        setUserData(newData);
        alert.success("Delete Successful!");
      } else {
        alert.error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  async function updateUser(record) {
    try {
      const response = await fetch(`${API_URLS.USER}/${record.uuid}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
        body: JSON.stringify(record)
      });

      const responseData = await response.json();
      if (responseData.success) {
        getUsers();
        alert.success("Update Successful!");
      } else {
        alert.error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  return (
    <Layout>
      <div className={Styles.container}>
        <div className="d-flex w-100">
          <div className="me-3">
            <TextField label="Name" id="name" onChange={handleInputChange} />
          </div>
          <div className="me-3">
            <Select id="active" options={activeOptions} value={filter.active} onSelect={(value) => {
              setFilter((prevData) => ({
                ...prevData,
                active: value,
              }));
            }} label="Active" />
          </div>
          <div className="me-3 d-flex align-items-center mt-2">
            <Button text="Filter" onClick={() => getUsers()} />
          </div>
        </div>
      </div>
      <div>
        <Table
          columns={column(deleteUser, roleOptions, updateUser)}
          dataSource={userData}
          className="h-100 m-3"
          scroll={{
            x: 1300,
            y: 500
          }}
        />
      </div>
    </Layout>
  );
}

export default Transaction;