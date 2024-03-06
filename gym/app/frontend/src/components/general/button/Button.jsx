import React from 'react'
import Styles from './Button.module.css'
function Button({ text, bgColor, onClick }) {
  const buttonStyle = {
    backgroundColor: bgColor
  }
  return (
    <button className={Styles.button} onClick={onClick} style={buttonStyle}>{text}</button>
  );
}

Button.defaultProps = {
  text: "submit",
  bgColor: "#d8cdb9"
}

export default Button