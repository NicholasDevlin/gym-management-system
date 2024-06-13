import React from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'

function Home() {
  return (
    <Layout>
      <div className='container-fluid h-100'>
        <div className='d-flex justify-content-between'>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Total Member</h4></div>
            </div>
          </div>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Active Member</h4></div>
            </div>
          </div>
          <div className={`${Styles.card} card m-3 p-3 min-w-25 w-25`}>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Total Transaction this month</h4></div>
            </div>
          </div>
        </div>
        <div className='m-3 w-50 h-75'>
          <div className={`${Styles.memberList} card p-3 mt-3`}>
            <h4 className='w-100'>
              List of members whose membership expires today:
            </h4>
            <div className='card-body overflow-auto'>
              <ul>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
                <li>1</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;