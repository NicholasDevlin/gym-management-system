import React from 'react'
import Styles from './CardProduct.module.css'

function CardProduct({title}){
  return (
    <div className={Styles.card}>
      <div className={Styles.title}>{title}</div>
      <div className={Styles.icon}>
      
      </div>

      <div className={Styles.features}>
        <ul>
          <li><span>5</span> Edits</li>
          <li><span>1GB</span> Storage</li>
          <li><span>3</span> Pages</li>
          <li><span>1</span> Hour free support</li>
        </ul>
      </div>
      <a href="#" className={Styles.btn}>Check it out</a>
    </div>
  )
}

export default CardProduct;