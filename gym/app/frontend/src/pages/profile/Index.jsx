import React, { useState, useEffect } from "react";
import Layout from "../../layout/MainLayout/Layout.jsx";
import Styles from "./Profile.module.css";
import TextInput from "../../components/general/input/inputTextField/TextField.jsx";
import PhoneNumberInput from "../../components/general/input/phoneNumberInput/PhoneNumberInput.jsx";
import DatetimePicker from "../../components/general/input/datetimePicker/DatetimePicker.jsx";
import GenderPicker from "../../components/general/input/genderPicker/GenderPicker.jsx";
import Button from "../../components/general/button/Button.jsx";
import { Icon } from "@iconify/react";
import { API_URLS } from "../../apiConfig.js";
import { useUserData } from "../../utils/jwt/UserData.jsx";
import { useAlert } from "react-alert";

function Profile() {
  const alert = useAlert()
  const userData = useUserData();
  const [profileData, setprofileData] = useState({
    name: "",
    birthDate: "",
    gender: "",
    phoneNumber: "",
  });

  const handleInputChange = (e) => {
    const { id, name, value } = e.target || {};

    if (id) {
      setprofileData((prevData) => ({
        ...prevData,
        [id]: value,
      }));
    } else if (name) {
      setprofileData((prevData) => ({
        ...prevData,
        [name]: value,
      }));
    }
  };

  const handleInputDateChange = (date) => {
    const { id } = date;
    const value = date.value;

    setprofileData((prevData) => ({
      ...prevData,
      [id]: value,
    }));
  };

  async function saveUserProfile() {
    try {
      const response = await fetch(`${API_URLS.USER}/${userData.uuid}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
        body: JSON.stringify(profileData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();

      if (responseData.success) {
        alert.success("Save successful");
      } else {
        alert.error("Save unsuccessful. Response:", responseData);
      }
    } catch (error) {
      alert.error("Error:", error);
    }
  }

  useEffect(() => {
    getUserProfile();
  }, []);

  async function getUserProfile() {
    try {
      const response = await fetch(`${API_URLS.USER}/${userData.uuid}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();
      if (responseData.success) {
        setprofileData(responseData.data);
      } else {
        alert.error("Get data unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  return (
    <Layout>
      <div className="d-flex justify-content-center mt-4">
        <div className={Styles.imageContainer}>
          <img
            alt="profile"
            className={`rounded-circle ${Styles.image}`}
            src="https://w7.pngwing.com/pngs/79/184/png-transparent-mannequin-head-dummy-model-face-male-fashion-bold-thumbnail.png"
          />
          <div className={Styles.iconContainer}>
            <label htmlFor="inputFile">
              <Icon
                className={Styles.icon}
                icon="mdi:pencil-outline"
                width="25"
                height="25"
              />
            </label>
            <input id="inputFile" type="file" className="d-none" />
          </div>
        </div>
      </div>
      <div className="container">
        <TextInput id={"name"} label={"Name"} value={profileData.name || ''} onChange={handleInputChange} />
        <PhoneNumberInput id={"phoneNumber"} value={profileData.phoneNumber || ''} label={"Phone Number"} onChange={handleInputChange} />
        <DatetimePicker label={"Birthdate"} value={new Date(profileData.birthDate)} id={"birthDate"} onChange={handleInputDateChange} />
        <GenderPicker onChange={handleInputChange} value={profileData.gender} id={"gender"} />
        <div className="d-flex justify-content-end">
          <Button onClick={saveUserProfile} text={"Save"} />
        </div>
      </div>
    </Layout>
  );
}

export default Profile;
