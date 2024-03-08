import React from 'react'
import Styles from './CardProduct.module.css'
import { Link } from 'react-router-dom'

function CardProduct({ title, duration, price, description }) {
  return (
    <div className={Styles.card}>
      <div className={Styles.title}>{title}</div>
      <div className={Styles.icon}>

      </div>
      <div className={Styles.features}>
        <ul>
          <li>{duration}</li>
          <li>{price}</li>
        </ul>
      </div>
      <div className={`'row' ${Styles.description}`}>
        <p>{description}</p>
      </div>
      <Link href="#" className={Styles.btn}>Check it out</Link>
    </div>
  )
}

export default CardProduct;