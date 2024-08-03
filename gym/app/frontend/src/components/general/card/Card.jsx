import React from 'react'
import Styles from './Card.module.css'

function Card({ title, body }) {
  return (
    <div className={Styles.card}>
      <div className={Styles.title}>
        {title}
      </div>
      <div className={Styles.body}>
        <div className={Styles.content}>
          {body}
        </div>
        <div className={Styles.button}>
        </div>
      </div>
    </div>
  )
}

export default Card;