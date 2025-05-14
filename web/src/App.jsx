import { useEffect, useState } from 'react'
import './App.css'
import { formatDateTime, formatTime } from './utils/Date'
import Box from '../component/Box'

function App() {
  const [city, setCity] = useState("")
  const [weatherData, setWeatherData] = useState(null)

  // useEffect(() => {
  const fetchData = async () => {
    try {
      const res = await fetch(`http://localhost:3000/v1/weather?city=${city}`)
      const data = await res.json()
      console.log(data)
      setWeatherData(data.data)
    } catch (error) {
      console.log("Error fetching weather data", error)
    }
  }
  // fetchData()
  // }, [city])

  const handleSubmit = (e) => {
    e.preventDefault()
    if (city.trim()) fetchData()
  }


  return (
    <>
      <div className='max-w-[1100px] grid grid-cols-[35%_65%] mx-auto mt-[60px]'>
        <div className='bg-white px-10 py-8 h-[600px]'>
          <form onSubmit={handleSubmit} className='px-4'>
            <input value={city} onChange={(e) => setCity(e.target.value)} type="search"
              className="bg-slate-200 rounded-[18px] w-full px-4 py-1.5  outline-none" placeholder='Enter city name' />
          </form>
          <div>
            {weatherData ? (
              <div className='flex flex-col justify-center items-center'>
                <img src={weatherData.icon} className='w-[170px]' alt="" />
                <div className=''>
                  <h1 className='text-6xl'>{weatherData.temp_c}°C</h1>
                  <p className='text-center'>{weatherData.text}</p>
                </div>
                <div className='border-b-2 h-2 border-gray-200 mt-4 w-8/12 '></div>
                <p className='text-center mt-7'>{formatDateTime(weatherData.localtime)}</p>
                <p className='text-center text-xl'>{formatTime(weatherData.localtime)} <span className='block'>Day</span></p>
                <div className='mt-10'>
                  <h1 className='text-4xl'>{weatherData.name}</h1>
                </div>
              </div>
            ) : <p className='text-center h-[500px] flex justify-center items-center text-sm'>Enter a city to get the weather data.</p>}
          </div>
        </div>
        <div className='bg-slate-100 px-10 py-8'>
          {weatherData ? (
            <div className=''>
              <h1 className='text-2xl'>Today</h1>
              <div className='grid grid-cols-3 mt-4 gap-x-6 gap-y-10'>
                <Box title="Wind" unit={"6km/h"} value={"East"} />
                <Box title="Humidity" unit={"69%"} value={""} />
                <Box title="UV Index" unit={"1"} value={""} />
                <Box title="Pressure" unit={"1000mb"} value={""} />
                <Box title="Real Feel" unit={"31°C"} value={""} />
                <Box title="Precipitation" unit={"mm"} value={""} />
                <Box title="Cloud cover" unit={"%"} value={""} />
                <Box title="Wind gust" unit={"kph"} value={""} />
                <Box title="Visibility" unit={"6km"} value={""} />
              </div>
            </div>
          ) : <p className='text-center h-[500px] flex justify-center items-center text-sm'>Enter a city to get the weather data.</p>}
        </div>
      </div>
    </>
  )
}

export default App
