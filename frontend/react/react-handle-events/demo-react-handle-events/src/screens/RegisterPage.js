import { useState } from "react";

import "./RegisterPage.css";

const RegisterPage = (props) => {
  const [isEmailError, setIsEmailError] = useState(true);
  const [isAddressError, setIsAddressError] = useState(true);
  const [isLanguageError, setIsLanguageError] = useState(true);
  const { userData, setUserData, setPage } = props;
  const handleInputChange = (event) => {
    let eventName = event.target.name;
    let eventValue = event.target.value;
    setUserData(previousValues => {
      return {
        ...previousValues,
        [eventName]: eventValue
      }
    });
    if(eventName === "userName") {
      setIsEmailError(eventValue.length < 8);
    };
    if(eventName === "userEmail") {
      setIsEmailError(eventValue.length < 8);
    };
    if(eventName === "userAddress") {
      setIsAddressError(eventValue.length < 8);
    };
    if(eventName === "userLanguage") {
      setIsLanguageError(!eventValue);
    };
  };

  const handleSubmit = (event) => {
    setPage('movieList');
    event.preventDefault();
  };

  return (
    <form className="register-container" onSubmit={(e) => handleSubmit(e)}>
      <div className="register-input-container">
        <label className="register-input">Name :</label>
        <input
          name="userName"
          type="text"
          onChange={handleInputChange}
          value={userData.userName ? userData.userName : ""}
          className="register-input"
        ></input>
      </div>
      <div className="register-input-container">
        <label className="register-input">Email :</label>
        <input
          name="userEmail"
          type="email"
          onChange={handleInputChange}
          value={userData.userEmail ? userData.userEmail : ""}
          className="register-input"
        ></input>
      </div>
      <div className={'register-input-container'}>
          <label className="register-input">Address : </label>
          <textarea
            name="userAddress"
            rows="4"
            cols="40"
            className="register-input"
            value={userData.userAddress ? userData.userAddress : ""}
            onChange={handleInputChange}
          />
      </div>
      <div className="register-input-container">
        <label className="register-input">Language :</label>
        <select
            name="userLanguage"
            className="register-input"
            onChange={handleInputChange}
            value={userData.userLanguage ? userData.userLanguage : ""}
        >
          <option value=""></option>
          <option value="en">English</option>
          <option value="id">Indonesian</option>
        </select>
      </div>
      <button type="submit" disabled={isEmailError || isAddressError || isLanguageError}>REGISTER</button>
    </form>
  );
};

export default RegisterPage;
