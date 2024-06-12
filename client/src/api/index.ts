import axios from "axios";

interface WeatherResponse {
  cod: string | number;
  name: string;
  temp: number;
  weather: string;
  feels_like: number;
  humidity: number;
  wind_speed: number;
}

export const getWeather = async (
  location: string
): Promise<WeatherResponse | null> => {
  try {
    const response = await axios.post<WeatherResponse>(
      "http://localhost:8081/getWeather",
      {
        name: location,
      }
    );
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Error fetching the weather data:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    return null;
  }
};
