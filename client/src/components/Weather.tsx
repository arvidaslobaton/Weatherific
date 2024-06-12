import React, { useEffect, useState } from "react";
import { getWeather } from "../api";

interface WeatherData {
  name: string;
  temp: number;
  weather: string;
  feels_like: number;
  humidity: number;
  wind_speed: number;
}

const Weather: React.FC = () => {
  const [data, setData] = useState<WeatherData | null>(null);
  const [location, setLocation] = useState<string>("");
  const [inputValue, setInputValue] = useState<string>("");

  useEffect(() => {
    const fetchWeather = async () => {
      if (!location) return;

      try {
        const weatherData = await getWeather(location);
        if (weatherData) {
          setData(weatherData);
        }
      } catch (error) {
        console.error("Error fetching the weather data:", error);
      }
    };

    fetchWeather();
  }, [location]);

  const handleKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      setLocation(inputValue);
    }
  };

  return (
    <div className="app">
      <div className="search">
        <input
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="Enter location"
        />
      </div>
      {data && (
        <div className="container">
          <div className="top">
            <div className="location">
              <p>{data.name}</p>
            </div>
            <div className="temp">
              <h1>{data.temp}°C</h1>
            </div>
            <div className="description">
              <p>{data.weather}</p>
            </div>
          </div>
          <div className="bottom">
            <div className="feels">
              <p className="bold">{data.feels_like}°C</p>
              <p>Feels</p>
            </div>
            <div className="humidity">
              <p className="bold">{data.humidity}%</p>
              <p>Humidity</p>
            </div>
            <div className="wind">
              <p className="bold">{data.wind_speed} mph</p>
              <p>Wind Speed</p>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Weather;
