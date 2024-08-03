import React from 'react'
import Styles from './Button.module.css'
export default function Button({ text, bgColor, onClick, className }) {
  const buttonStyle = {
    backgroundColor: bgColor
  }
  return (
    <button className={`${Styles.button} ${className}`} onClick={onClick} style={buttonStyle}>{text}</button>
  );
}

export function DangerButton({text, onClick}) {
  return (
    <button className={Styles.dangerButton} onClick={onClick} >{text}</button>
  );
}

Button.defaultProps = {
  text: "submit",
  bgColor: "#d8cdb9"
}