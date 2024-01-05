import React from 'react'
import Styles from './Button.module.css'
function Button({ text, bgColor }) {
  const buttonStyle = {
    backgroundColor: bgColor
  }
  return (
    <button className={Styles.button} style={buttonStyle}>{text}</button>
  );
}

Button.defaultProps = {
  text: "submit",
  // bgColor: "#252B48"
  bgColor: "#d8cdb9"
}

export default Button