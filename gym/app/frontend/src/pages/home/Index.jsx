import React, { useEffect, useState } from 'react'
import Layout from '../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'
import { Carousel } from 'antd';

function Home() {


  return (
    <Layout>
      <div className='container-fluid h-75'>
        <Carousel arrows autoplay>
          <div>
            <h3>1</h3>
          </div>
          <div>
            <h3>2</h3>
          </div>
          <div>
            <h3>3</h3>
          </div>
          <div>
            <h3>4</h3>
          </div>
        </Carousel>
      </div>
    </Layout>
  );
}

export default Home;