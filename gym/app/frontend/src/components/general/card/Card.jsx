import React from 'react'
import Styles from './Card.module.css'

function Card({ title, body, button: Button }) {
  return (
    <div className={Styles.card}>
      <div className={Styles.title}>
        <h2>{title}</h2>
      </div>
      <div className={Styles.body}>
        <div className={Styles.content}>
          <div>{body}</div>
        </div>
        <div className={Styles.button}>
          <Button />
        </div>
      </div>
    </div>
  )
}

export default Card;