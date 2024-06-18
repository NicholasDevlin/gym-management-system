import { API_URLS } from "../apiConfig";

export async function GetUsers() {
  try {
    const response = await fetch(`${API_URLS.USER}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + localStorage.getItem('authToken')
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}\n${response.message}`);
    }
    const responseData = await response.json();
    return responseData.data;
  } catch (error) {
    return null;
  }
}

export async function GetRole() {
  try {
    const response = await fetch(`${API_URLS.ROLE}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + localStorage.getItem('authToken')
      },
    });

    const responseData = await response.json();
    return responseData.data;
  } catch (error) {
    return null;
  }
}