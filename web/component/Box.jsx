import React from 'react'

function Box({title, unit, value}) {
  return (
    <div className='bg-amber-600 text-white p-4 rounded-xl'>
        <h1>{title}</h1>
        <p className='text-2xl py-2'>{unit}</p>
        <p>{value}</p>
    </div>
  )
}

export default Box