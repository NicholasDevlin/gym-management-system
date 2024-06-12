import React from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'

function Home() {
  return (
    <Layout>
      <div className='container-fluid h-100'>
        <div className='d-flex justify-content-between'>
          <div className='card m-3 p-3 min-w-25 w-25'>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Total Member</h4></div>
            </div>
          </div>
          <div className='card m-3 p-3 min-w-25 w-25'>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Active Member</h4></div>
            </div>
          </div>
          <div className='card m-3 p-3 min-w-25 w-25'>
            <div className='d-flex justify-content-center w-100 align-content-around h-100 flex-wrap'>
              <div className='text-center'><h3>Loading...</h3></div>
              <div className='text-center'><h4>Total Transaction this month</h4></div>
            </div>
          </div>
        </div>
        <div className='m-3 w-50 h-100'>
          <div className='card mt-3'>
            <h4>
              List of members whose membership expires today:
            </h4>
            <ul>
              <li>1</li>
              <li>1</li>
              <li>1</li>
            </ul>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;