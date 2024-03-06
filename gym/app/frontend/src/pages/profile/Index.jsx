import React, { useState } from "react";
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

function Profile() {
  const userData = useUserData();
  const [profileData, setprofileData] = useState({
    name: "",
    birthDate: "",
    gender: "",
    phoneNumber: "",
  });

  const handleInputChange = (e) => {
    const { id, value } = e.target;

    setprofileData((prevData) => ({
      ...prevData,
      [id]: value,
    }));
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
        console.log("Save successful. Response:", responseData);
      } else {
        console.error("Save unsuccessful. Response:", responseData);
      }
    } catch (error) {
      console.error("Error:", error);
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
        <TextInput id={"name"} label={"Name"} onChange={handleInputChange} />
        <PhoneNumberInput id={"phoneNumber"} label={"Phone Number"} onChange={handleInputChange} />
        <DatetimePicker label={"Birthdate"} id={"birthDate"} onChange={handleInputDateChange} />
        <GenderPicker onChange={handleInputChange} id={"gender"} />
        <div className="d-flex justify-content-end">
          <Button onClick={saveUserProfile} text={"Save"} />
        </div>
      </div>
    </Layout>
  );
}

export default Profile;
